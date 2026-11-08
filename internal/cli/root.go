// Package cli builds the gyro command tree: generated commands for every API operation
// (generated.go, from internal/spec) plus hand-written commands with nicer UX (server.go, ...).
package cli

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/config"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
	"github.com/Gyro-Hosting/gyro-cli/internal/output"
	"github.com/Gyro-Hosting/gyro-cli/internal/spec"
)

// Build information (set with -ldflags by GoReleaser).
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

// Exit codes (documented in README).
const (
	ExitOK        = 0
	ExitError     = 1
	ExitUsage     = 2
	ExitAuth      = 3 // 401/403
	ExitNotFound  = 4
	ExitConflict  = 5 // 409
	ExitPayment   = 6 // 402
	ExitAborted   = 7 // user said no / confirmation missing
	ExitRateLimit = 8
	ExitTimeout   = 9 // --wait timed out
)

// ExitErr carries an exit code.
type ExitErr struct {
	Code int
	Err  error
}

func (e *ExitErr) Error() string { return e.Err.Error() }
func (e *ExitErr) Unwrap() error { return e.Err }

// Options are the global flags.
type Options struct {
	Profile  string
	APIURL   string
	Output   string
	Query    string
	NoHeader bool
	Quiet    bool
	NoColor  bool
	Lang     string
	Yes      bool
	Debug    bool
	Wide     bool
	Timeout  time.Duration
}

// App is shared by all commands.
type App struct {
	Opt    Options
	In     io.Reader
	Out    io.Writer
	Err    io.Writer
	Env    func(string) string
	IsTTY  func() bool // stdin is an interactive terminal (prompts allowed)
	cfg    *config.File
	client *api.Client
	// NewClient can be replaced in tests.
	NewClient func(baseURL, token string) *api.Client
	// Sleep for --wait polling (tests set it to a no-op).
	Sleep func(time.Duration)
	// Exec runs external programs (ssh); tests replace it.
	Exec func(name string, args []string) error
}

// NewApp returns an App wired to the real process.
func NewApp() *App {
	return &App{
		In: os.Stdin, Out: os.Stdout, Err: os.Stderr, Env: os.Getenv,
		IsTTY: func() bool { return term.IsTerminal(int(os.Stdin.Fd())) },
		NewClient: func(baseURL, token string) *api.Client {
			return api.New(baseURL, token, fmt.Sprintf("gyro-cli/%s (%s/%s)", Version, runtime.GOOS, runtime.GOARCH))
		},
		Sleep: time.Sleep,
		Exec:  execProgram,
	}
}

// Config loads the config file once.
func (a *App) Config() (*config.File, error) {
	if a.cfg != nil {
		return a.cfg, nil
	}
	c, err := config.Load()
	if err != nil {
		return nil, err
	}
	a.cfg = c
	return c, nil
}

// ProfileName is the active profile.
func (a *App) ProfileName() string {
	c, err := a.Config()
	if err != nil {
		return "default"
	}
	return c.ActiveName(a.Opt.Profile)
}

// Profile returns the active profile (may be empty).
func (a *App) Profile() *config.Profile {
	c, err := a.Config()
	if err != nil {
		return &config.Profile{}
	}
	if p, ok := c.Profiles[a.ProfileName()]; ok {
		return p
	}
	return &config.Profile{}
}

// APIURL resolves the base URL: flag > GYRO_API_URL > profile > default.
func (a *App) APIURL() string {
	switch {
	case a.Opt.APIURL != "":
		return a.Opt.APIURL
	case a.Env("GYRO_API_URL") != "":
		return a.Env("GYRO_API_URL")
	case a.Profile().APIURL != "":
		return a.Profile().APIURL
	}
	return api.DefaultURL
}

// Token resolves the API key: GYRO_TOKEN > profile (file or keychain).
func (a *App) Token() (token, source string, err error) {
	if t := a.Env("GYRO_TOKEN"); t != "" {
		return t, "env GYRO_TOKEN", nil
	}
	c, err := a.Config()
	if err != nil {
		return "", "", err
	}
	t, src, err := c.TokenFor(a.ProfileName())
	if src != "" {
		src = src + " (profile " + a.ProfileName() + ")"
	}
	return t, src, err
}

// Client returns an authenticated API client.
func (a *App) Client() (*api.Client, error) {
	if a.client != nil {
		return a.client, nil
	}
	tok, _, err := a.Token()
	if err != nil {
		return nil, &ExitErr{ExitAuth, err}
	}
	if tok == "" {
		return nil, &ExitErr{ExitAuth, errors.New(i18n.S(
			"no API key: run `gyro configure` or set GYRO_TOKEN (create keys in Client area → Account → API)",
			"API anahtarı yok: `gyro configure` çalıştırın ya da GYRO_TOKEN ayarlayın (anahtar: Müşteri paneli → Hesabım → API)"))}
	}
	c := a.NewClient(a.APIURL(), tok)
	c.Lang = i18n.Lang()
	if a.Opt.Debug {
		c.Debug = a.Err
	}
	a.client = c
	return c, nil
}

// Printer returns the output printer for the current flags.
func (a *App) Printer() *output.Printer {
	format := a.Opt.Output
	if format == "" {
		format = a.Env("GYRO_OUTPUT")
	}
	if format == "" {
		format = a.Profile().Output
	}
	if format == "" {
		format = "table"
	}
	return &output.Printer{W: a.Out, Format: format, Query: a.Opt.Query, NoHeader: a.Opt.NoHeader,
		Quiet: a.Opt.Quiet, Color: output.ColorEnabled(a.Out, a.Opt.NoColor), Wide: a.Opt.Wide}
}

// Machine reports json/yaml output (progress and hints then go to stderr only).
func (a *App) Machine() bool {
	p := a.Printer()
	return p.Format == "json" || p.Format == "yaml" || p.Quiet
}

// Ctx returns a context cancelled on Ctrl-C and after --timeout.
func (a *App) Ctx() (context.Context, context.CancelFunc) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	if a.Opt.Timeout > 0 {
		c2, cancel := context.WithTimeout(ctx, a.Opt.Timeout)
		return c2, func() { cancel(); stop() }
	}
	return ctx, stop
}

// Info prints a human message to stderr (suppressed in quiet mode).
func (a *App) Info(format string, args ...any) {
	if a.Opt.Quiet {
		return
	}
	fmt.Fprintf(a.Err, format+"\n", args...)
}

// NewRoot builds the whole command tree.
func NewRoot(a *App) *cobra.Command {
	root := &cobra.Command{
		Use:           "gyro",
		Short:         i18n.S("Gyro Hosting command-line interface", "Gyro Hosting komut satırı aracı"),
		Long:          rootLong(),
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if a.Opt.Lang != "" {
				i18n.Set(a.Opt.Lang)
			} else if l := a.Env("GYRO_LANG"); l == "" {
				if p := a.Profile(); p.Lang != "" {
					i18n.Set(p.Lang)
				}
			}
			switch a.Opt.Output {
			case "", "table", "json", "yaml", "text":
			default:
				return &ExitErr{ExitUsage, fmt.Errorf("--output: %s (table|json|yaml|text)", a.Opt.Output)}
			}
			if c, err := a.Config(); err == nil && c.InsecurePerms() {
				fmt.Fprintln(a.Err, i18n.F("warning: %s is readable by other users; run chmod 600", "uyarı: %s başka kullanıcılarca okunabilir; chmod 600 yapın", c.FilePath()))
			}
			return nil
		},
	}
	root.SetVersionTemplate("gyro {{.Version}}\n")
	f := root.PersistentFlags()
	f.StringVarP(&a.Opt.Profile, "profile", "p", "", i18n.S("profile to use (env GYRO_PROFILE)", "kullanılacak profil (ortam GYRO_PROFILE)"))
	f.StringVar(&a.Opt.APIURL, "api-url", "", i18n.S("API base URL (env GYRO_API_URL)", "API taban adresi (ortam GYRO_API_URL)"))
	f.StringVarP(&a.Opt.Output, "output", "o", "", i18n.S("output format: table, json, yaml, text", "çıktı biçimi: table, json, yaml, text"))
	f.StringVar(&a.Opt.Query, "query", "", i18n.S("JMESPath query applied to the result", "sonuca uygulanacak JMESPath sorgusu"))
	f.BoolVar(&a.Opt.NoHeader, "no-header", false, i18n.S("omit the table header", "tablo başlığını yazma"))
	f.BoolVarP(&a.Opt.Quiet, "quiet", "q", false, i18n.S("print only ids", "yalnız kimlikleri yaz"))
	f.BoolVar(&a.Opt.NoColor, "no-color", false, i18n.S("disable colours (also NO_COLOR)", "renkleri kapat (NO_COLOR da olur)"))
	f.StringVar(&a.Opt.Lang, "lang", "", i18n.S("message language: tr or en (default from locale)", "mesaj dili: tr ya da en (varsayılan yerel ayar)"))
	f.BoolVarP(&a.Opt.Yes, "yes", "y", false, i18n.S("do not ask for confirmation", "onay sorma"))
	f.BoolVar(&a.Opt.Debug, "debug", false, i18n.S("log HTTP requests to stderr (secrets masked)", "HTTP isteklerini stderr'e yaz (sırlar maskeli)"))
	f.BoolVar(&a.Opt.Wide, "wide", false, i18n.S("show all table columns", "tüm tablo sütunlarını göster"))
	f.DurationVar(&a.Opt.Timeout, "timeout", 0, i18n.S("overall timeout, e.g. 30s, 10m", "toplam zaman aşımı, ör. 30s, 10m"))
	_ = root.RegisterFlagCompletionFunc("output", fixedCompletion("table", "json", "yaml", "text"))
	_ = root.RegisterFlagCompletionFunc("lang", fixedCompletion("tr", "en"))
	_ = root.RegisterFlagCompletionFunc("profile", func(cmd *cobra.Command, args []string, s string) ([]string, cobra.ShellCompDirective) {
		c, err := a.Config()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return c.ProfileNames(), cobra.ShellCompDirectiveNoFileComp
	})

	root.AddGroup(&cobra.Group{ID: "res", Title: i18n.S("Resources:", "Kaynaklar:")},
		&cobra.Group{ID: "acct", Title: i18n.S("Account and billing:", "Hesap ve fatura:")},
		&cobra.Group{ID: "cli", Title: i18n.S("Configuration and tools:", "Yapılandırma ve araçlar:")})

	// hand-written commands first; generated ones fill the gaps
	addConfigure(a, root)
	addContext(a, root)
	addVersion(a, root)
	addCompletion(a, root)
	addDocs(a, root)
	addAPICommand(a, root)
	addServer(a, root)
	addSSHKeys(a, root)
	addCatalog(a, root)
	addBilling(a, root)
	addProductLines(a, root)
	addDNS(a, root)
	addGenerated(a, root)
	addCoreShortcuts(a, root)
	fillGroupShorts(root)
	acct := map[string]bool{"me": true, "address": true, "affiliate": true, "order": true, "ticket": true, "key": true,
		"webhook": true, "reseller": true, "reference": true, "activity": true, "sms": true, "ping": true}
	for _, c := range root.Commands() {
		if c.GroupID == "" && !c.Hidden && c.Name() != "help" {
			c.GroupID = "res"
			if acct[c.Name()] {
				c.GroupID = "acct"
			}
		}
	}
	root.SetHelpCommandGroupID("cli")
	root.SetCompletionCommandGroupID("cli")
	return root
}

func rootLong() string {
	return i18n.S(`gyro manages your Gyro Hosting services from the terminal: cloud servers, SSH keys, invoices,
balance and everything else the Gyro Public API offers (API spec `+spec.APIVersion+`).

Get started:
  gyro configure                 # store an API key (Client area → Account → API)
  gyro server list
  gyro server create --plan standart-2-4 --location falkenstein --image ubuntu-24.04 --ssh-key mykey --wait
  gyro server ssh web-1

Money and destructive actions ask for confirmation (use --yes in scripts, --dry-run for a quote).`,
		`gyro, Gyro Hosting hizmetlerinizi terminalden yönetir: bulut sunucular, SSH anahtarları, faturalar,
bakiye ve Gyro Public API'nin sunduğu her şey (API sürümü `+spec.APIVersion+`).

Başlarken:
  gyro configure                 # API anahtarını kaydedin (Müşteri paneli → Hesabım → API)
  gyro server list
  gyro server create --plan standart-2-4 --location falkenstein --image ubuntu-24.04 --ssh-key anahtarim --wait
  gyro server ssh web-1

Para harcayan ve geri alınamayan işlemler onay sorar (betiklerde --yes, fiyat için --dry-run).`)
}

func fixedCompletion(vals ...string) func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	return func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
		return vals, cobra.ShellCompDirectiveNoFileComp
	}
}

// Execute runs the CLI and returns the process exit code.
func Execute(a *App, args []string) int {
	root := NewRoot(a)
	root.SetArgs(args)
	root.SetIn(a.In)
	root.SetOut(a.Out)
	root.SetErr(a.Err)
	err := root.Execute()
	if err == nil {
		return ExitOK
	}
	return reportError(a, err)
}

func reportError(a *App, err error) int {
	code := ExitError
	var ee *ExitErr
	if errors.As(err, &ee) {
		code = ee.Code
	}
	prefix := i18n.S("Error", "Hata")
	if e, ok := api.AsError(err); ok {
		if code == ExitError {
			switch {
			case e.Status == 401 || e.Status == 403:
				code = ExitAuth
			case e.Status == 404:
				code = ExitNotFound
			case e.Status == 409:
				code = ExitConflict
			case e.Status == 402:
				code = ExitPayment
			case e.Status == 429:
				code = ExitRateLimit
			}
		}
		fmt.Fprintf(a.Err, "%s: %s\n", prefix, e.Message)
		meta := []string{}
		if e.Code != "" {
			meta = append(meta, i18n.S("code", "kod")+": "+e.Code)
		}
		if e.Status != 0 {
			meta = append(meta, fmt.Sprintf("HTTP %d", e.Status))
		}
		if e.RequestID != "" {
			meta = append(meta, i18n.S("request id", "istek no")+": "+e.RequestID)
		}
		fmt.Fprintf(a.Err, "  (%s)\n", strings.Join(meta, ", "))
		if len(e.Details) > 0 {
			for k, v := range e.Details {
				fmt.Fprintf(a.Err, "  %s: %v\n", k, v)
			}
		}
		if h := hint(e); h != "" {
			fmt.Fprintf(a.Err, "  %s\n", h)
		}
		return code
	}
	msg := err.Error()
	if strings.HasPrefix(msg, "unknown command") || strings.HasPrefix(msg, "unknown flag") ||
		strings.HasPrefix(msg, "accepts ") || strings.HasPrefix(msg, "requires ") || strings.Contains(msg, "flag needs an argument") ||
		strings.HasPrefix(msg, "invalid argument") || strings.Contains(msg, "required flag") {
		code = ExitUsage
		msg += "\n" + i18n.S("Run 'gyro --help' for usage.", "Kullanım için 'gyro --help'.")
	}
	fmt.Fprintf(a.Err, "%s: %s\n", prefix, msg)
	return code
}

func hint(e *api.Error) string {
	switch e.Code {
	case "missing_token", "invalid_token", "key_revoked":
		return i18n.S("Check your key with `gyro context`, or create a new one in Client area → Account → API.",
			"Anahtarı `gyro context` ile kontrol edin ya da Müşteri paneli → Hesabım → API'den yenisini oluşturun.")
	case "insufficient_scope":
		return i18n.S("The key lacks this permission; edit its scopes in Client area → Account → API.",
			"Anahtarın bu yetkisi yok; kapsamları Müşteri paneli → Hesabım → API'den düzenleyin.")
	case "ip_not_allowed":
		return i18n.S("Your IP is not in the key's allow list.", "IP adresiniz anahtarın izin listesinde değil.")
	case "insufficient_balance":
		return i18n.S("Top up your balance (`gyro balance`) and try again.", "Bakiyenizi yükleyip tekrar deneyin (`gyro balance`).")
	case "rate_limited":
		return i18n.S("Too many requests; the CLI already retried. Wait a minute.", "Çok fazla istek; CLI zaten tekrar denedi. Bir dakika bekleyin.")
	case "server_locked":
		return i18n.S("Another action is running on this server; add --wait or retry shortly.", "Sunucuda başka bir işlem sürüyor; --wait ekleyin ya da biraz sonra deneyin.")
	case "confirm_required":
		return i18n.S("Re-run with --yes to confirm.", "Onaylamak için --yes ile tekrar çalıştırın.")
	}
	return ""
}

// Run is the process entry point: language pre-scan, build, execute.
func Run(args []string) int {
	a := NewApp()
	preScan(a, args)
	return Execute(a, args)
}
