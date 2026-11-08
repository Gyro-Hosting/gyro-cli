package cli

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"golang.org/x/term"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/config"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
	"github.com/Gyro-Hosting/gyro-cli/internal/output"
	"github.com/Gyro-Hosting/gyro-cli/internal/spec"
)

/* ------------------------------------------------------------------ configure */

func addConfigure(a *App, root *cobra.Command) {
	var keyring, tokenStdin, noVerify bool
	var apiURL, output_, lang string
	cfg := &cobra.Command{
		Use:     "configure",
		GroupID: "cli",
		Short:   i18n.S("Store an API key and preferences in a named profile", "API anahtarını ve tercihleri adlı bir profile kaydet"),
		Long: i18n.S(`Interactive setup, like 'aws configure'. Profiles live in ~/.config/gyro/config.toml (mode 0600).
Use --profile NAME for more than one account; --keyring keeps the key in the OS keychain instead of the file.
Non-interactive: echo "$KEY" | gyro configure --token-stdin --profile ci`,
			`'aws configure' gibi etkileşimli kurulum. Profiller ~/.config/gyro/config.toml dosyasındadır (izin 0600).
Birden çok hesap için --profile AD; --keyring anahtarı dosya yerine işletim sistemi anahtar zincirinde tutar.
Etkileşimsiz: echo "$ANAHTAR" | gyro configure --token-stdin --profile ci`),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := a.Config()
			if err != nil {
				return err
			}
			name := a.ProfileName()
			p := c.Get(name)
			r := bufio.NewReader(a.In)
			interactive := a.IsTTY() && !tokenStdin
			token := ""
			if tokenStdin {
				line, _ := r.ReadString('\n')
				token = strings.TrimSpace(line)
			}
			if apiURL != "" {
				p.APIURL = apiURL
			}
			if output_ != "" {
				p.Output = output_
			}
			if lang != "" {
				p.Lang = i18n.Normalize(lang)
			}
			if interactive {
				fmt.Fprintln(a.Err, i18n.F("Configuring profile %q (%s)", "%q profili yapılandırılıyor (%s)", name, c.FilePath()))
				old, _, _ := c.TokenFor(name)
				label := i18n.S("API key (wck_...)", "API anahtarı (wck_...)")
				if old != "" {
					label += " [" + api.MaskToken(old) + "]"
				}
				fmt.Fprintf(a.Err, "%s: ", label)
				if f, ok := a.In.(*os.File); ok && term.IsTerminal(int(f.Fd())) {
					b, err := term.ReadPassword(int(f.Fd()))
					fmt.Fprintln(a.Err)
					if err != nil {
						return err
					}
					token = strings.TrimSpace(string(b))
				} else {
					line, _ := r.ReadString('\n')
					token = strings.TrimSpace(line)
				}
				if token == "" {
					token = old
				}
				def := p.APIURL
				if def == "" {
					def = api.DefaultURL
				}
				p.APIURL = a.Prompt(r, i18n.S("API URL", "API adresi"), def)
				if p.APIURL == api.DefaultURL {
					p.APIURL = ""
				}
				p.Output = a.Prompt(r, i18n.S("Default output (table/json/yaml/text)", "Varsayılan çıktı (table/json/yaml/text)"), firstNonEmpty(p.Output, "table"))
				p.Lang = i18n.Normalize(a.Prompt(r, i18n.S("Language (tr/en)", "Dil (tr/en)"), firstNonEmpty(p.Lang, i18n.Lang())))
				if !cmd.Flags().Changed("keyring") {
					ans := a.Prompt(r, i18n.S("Store the key in the OS keychain? (y/N)", "Anahtar işletim sistemi anahtar zincirinde saklansın mı? (e/H)"), "")
					keyring = strings.HasPrefix(strings.ToLower(ans), "y") || strings.HasPrefix(strings.ToLower(ans), "e")
				}
			}
			if token != "" {
				if !strings.HasPrefix(token, "wck_") && !strings.HasPrefix(token, "wak_") {
					fmt.Fprintln(a.Err, i18n.S("warning: client keys normally start with wck_", "uyarı: müşteri anahtarları normalde wck_ ile başlar"))
				}
				if !noVerify {
					cl := a.NewClient(firstNonEmpty(p.APIURL, a.Env("GYRO_API_URL"), api.DefaultURL), token)
					ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
					var who struct {
						Name        string   `json:"name"`
						OwnerID     int64    `json:"owner_id"`
						Permissions []string `json:"permissions"`
					}
					_, err := cl.Get(ctx, "/whoami", nil, &who)
					cancel()
					if err != nil {
						return fmt.Errorf("%s: %w", i18n.S("key check failed (use --no-verify to save anyway)", "anahtar doğrulanamadı (yine de kaydetmek için --no-verify)"), err)
					}
					fmt.Fprintln(a.Err, i18n.F("Key OK: %q, account #%d, %d permission(s).", "Anahtar geçerli: %q, hesap #%d, %d yetki.", who.Name, who.OwnerID, len(who.Permissions)))
				}
				if err := c.SetToken(name, token, keyring); err != nil {
					return err
				}
			}
			if c.DefaultProfile == "" {
				c.DefaultProfile = name
			}
			if err := c.Save(); err != nil {
				return err
			}
			fmt.Fprintln(a.Err, i18n.F("Saved profile %q to %s", "%q profili %s dosyasına kaydedildi", name, c.FilePath()))
			return nil
		},
	}
	cfg.Flags().BoolVar(&keyring, "keyring", false, i18n.S("store the key in the OS keychain", "anahtarı işletim sistemi anahtar zincirinde sakla"))
	cfg.Flags().BoolVar(&tokenStdin, "token-stdin", false, i18n.S("read the key from stdin (non-interactive)", "anahtarı stdin'den oku (etkileşimsiz)"))
	cfg.Flags().BoolVar(&noVerify, "no-verify", false, i18n.S("do not call /whoami to check the key", "anahtarı /whoami ile doğrulama"))
	cfg.Flags().StringVar(&apiURL, "set-api-url", "", i18n.S("API URL to store in the profile", "profile yazılacak API adresi"))
	cfg.Flags().StringVar(&output_, "set-output", "", i18n.S("default output to store", "kaydedilecek varsayılan çıktı"))
	cfg.Flags().StringVar(&lang, "set-lang", "", i18n.S("language to store (tr/en)", "kaydedilecek dil (tr/en)"))

	list := &cobra.Command{
		Use: "list", Short: i18n.S("List profiles", "Profilleri listele"), Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := a.Config()
			if err != nil {
				return err
			}
			var rows []any
			for _, n := range c.ProfileNames() {
				p := c.Profiles[n]
				store := "file"
				if p.TokenStorage == "keyring" {
					store = "keyring"
				} else if p.Token == "" {
					store = "-"
				}
				def := ""
				if n == c.ActiveName(a.Opt.Profile) {
					def = "*"
				}
				rows = append(rows, map[string]any{"active": def, "name": n, "api_url": firstNonEmpty(p.APIURL, api.DefaultURL), "token": store, "output": p.Output, "lang": p.Lang})
			}
			if rows == nil {
				rows = []any{}
			}
			cols := []output.Column{{Header: "active", Path: "active"}, {Header: "name", Path: "name"}, {Header: "api_url", Path: "api_url"},
				{Header: "token", Path: "token"}, {Header: "output", Path: "output"}, {Header: "lang", Path: "lang"}}
			return a.Printer().Print(rows, cols, "name")
		},
	}
	get := &cobra.Command{
		Use: "get <KEY>", Short: i18n.S("Print a profile setting (api_url, output, lang, ssh_user, ssh_identity)", "Profil ayarını yaz (api_url, output, lang, ssh_user, ssh_identity)"),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			p := a.Profile()
			v, ok := profileField(p, args[0])
			if !ok {
				return &ExitErr{ExitUsage, fmt.Errorf("%s: %s", i18n.S("unknown key", "bilinmeyen anahtar"), args[0])}
			}
			fmt.Fprintln(a.Out, *v)
			return nil
		},
	}
	set := &cobra.Command{
		Use: "set <KEY> <VALUE>", Short: i18n.S("Change a profile setting (the API key: use 'gyro configure --token-stdin')", "Profil ayarını değiştir (API anahtarı için 'gyro configure --token-stdin')"),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := a.Config()
			if err != nil {
				return err
			}
			p := c.Get(a.ProfileName())
			v, ok := profileField(p, args[0])
			if !ok {
				return &ExitErr{ExitUsage, fmt.Errorf("%s: %s", i18n.S("unknown key", "bilinmeyen anahtar"), args[0])}
			}
			*v = args[1]
			return c.Save()
		},
	}
	del := &cobra.Command{
		Use: "delete <PROFILE>", Short: i18n.S("Remove a profile", "Profili sil"), Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := a.Config()
			if err != nil {
				return err
			}
			if _, ok := c.Profiles[args[0]]; !ok {
				return &ExitErr{ExitNotFound, fmt.Errorf("%s: %s", i18n.S("no such profile", "böyle bir profil yok"), args[0])}
			}
			if err := a.Confirm(i18n.F("Delete profile %q?", "%q profili silinsin mi?", args[0])); err != nil {
				return err
			}
			c.Delete(args[0])
			return c.Save()
		},
	}
	cfg.AddCommand(list, get, set, del)
	root.AddCommand(cfg)
}

func profileField(pp *config.Profile, key string) (*string, bool) {
	switch strings.ReplaceAll(key, "-", "_") {
	case "api_url":
		return &pp.APIURL, true
	case "output":
		return &pp.Output, true
	case "lang":
		return &pp.Lang, true
	case "ssh_user":
		return &pp.SSHUser, true
	case "ssh_identity":
		return &pp.SSHIdentity, true
	}
	return nil, false
}

func firstNonEmpty(v ...string) string {
	for _, s := range v {
		if s != "" {
			return s
		}
	}
	return ""
}

/* ------------------------------------------------------------------ context */

func addContext(a *App, root *cobra.Command) {
	ctxCmd := &cobra.Command{
		Use:     "context",
		Aliases: []string{"ctx"},
		GroupID: "cli",
		Short:   i18n.S("Show the active profile, API URL and key (and check the key)", "Etkin profili, API adresini ve anahtarı göster (anahtarı doğrula)"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			tok, src, terr := a.Token()
			info := map[string]any{
				"profile": a.ProfileName(), "api_url": a.APIURL(), "token": api.MaskToken(tok), "token_source": src,
				"config_file": func() string {
					c, _ := a.Config()
					if c != nil {
						return c.FilePath()
					}
					return ""
				}(),
				"language": i18n.Lang(), "cli_version": Version, "api_spec_version": spec.APIVersion,
			}
			if terr != nil {
				info["token_error"] = terr.Error()
			}
			if tok != "" {
				ctx, cancel := a.Ctx()
				defer cancel()
				c, err := a.Client()
				if err == nil {
					var who map[string]any
					if _, err := c.Get(ctx, "/whoami", nil, &who); err != nil {
						info["key_status"] = err.Error()
					} else {
						info["key_status"] = "ok"
						info["key_name"] = who["name"]
						info["account_id"] = who["owner_id"]
						info["permissions"] = who["permissions"]
					}
				}
			}
			return a.Printer().Print(toGeneric(info), nil, "profile")
		},
	}
	use := &cobra.Command{
		Use: "use <PROFILE>", Short: i18n.S("Make a profile the default", "Profili varsayılan yap"), Args: cobra.ExactArgs(1),
		ValidArgsFunction: func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
			c, _ := a.Config()
			if c == nil {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			return c.ProfileNames(), cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := a.Config()
			if err != nil {
				return err
			}
			if _, ok := c.Profiles[args[0]]; !ok {
				return &ExitErr{ExitNotFound, fmt.Errorf("%s: %s", i18n.S("no such profile (run gyro configure --profile NAME)", "böyle bir profil yok (gyro configure --profile AD)"), args[0])}
			}
			c.DefaultProfile = args[0]
			if err := c.Save(); err != nil {
				return err
			}
			a.Info("%s", i18n.F("Default profile is now %q.", "Varsayılan profil artık %q.", args[0]))
			return nil
		},
	}
	ctxCmd.AddCommand(use)
	root.AddCommand(ctxCmd)
}

/* ------------------------------------------------------------------ version / update check */

// UpdateURL is the GitHub "latest release" API endpoint (GYRO_UPDATE_URL overrides; "off" disables).
var UpdateURL = "https://api.github.com/repos/Gyro-Hosting/gyro-cli/releases/latest"

func addVersion(a *App, root *cobra.Command) {
	var check bool
	cmd := &cobra.Command{
		Use: "version", GroupID: "cli",
		Short: i18n.S("Print the version (and check for a newer release with --check)", "Sürümü yaz (--check ile yeni sürüm denetimi)"),
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			info := map[string]any{"version": Version, "commit": Commit, "built": Date, "go": runtime.Version(),
				"platform": runtime.GOOS + "/" + runtime.GOARCH, "api_spec": spec.APIVersion}
			if check {
				latest, url_, err := latestRelease(a)
				switch {
				case err != nil:
					info["update_check"] = err.Error()
				case latest == "":
					info["update_check"] = i18n.S("no published release yet", "henüz yayınlanmış sürüm yok")
				default:
					info["latest"] = latest
					if newer(latest, Version) {
						info["update_check"] = i18n.F("update available: %s (%s)", "güncelleme var: %s (%s)", latest, url_)
						info["upgrade"] = upgradeHint()
					} else {
						info["update_check"] = i18n.S("up to date", "güncel")
					}
				}
			}
			if a.Machine() || a.Opt.Query != "" {
				return a.Printer().Print(toGeneric(info), nil, "version")
			}
			fmt.Fprintf(a.Out, "gyro %s (%s, %s) %s %s, API %s\n", Version, Commit, Date, runtime.Version(), info["platform"], spec.APIVersion)
			if v, ok := info["update_check"]; ok {
				fmt.Fprintln(a.Out, v)
			}
			if v, ok := info["upgrade"]; ok {
				fmt.Fprintln(a.Out, v)
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&check, "check", false, i18n.S("check GitHub for a newer release", "GitHub'da yeni sürüm var mı bak"))
	root.AddCommand(cmd)
}

func latestRelease(a *App) (tag, htmlURL string, err error) {
	u := UpdateURL
	if v := a.Env("GYRO_UPDATE_URL"); v != "" {
		u = v
	}
	if u == "off" {
		return "", "", fmt.Errorf("%s", i18n.S("update check disabled", "güncelleme denetimi kapalı"))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, "GET", u, nil)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "gyro-cli/"+Version)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return "", "", nil
	}
	if resp.StatusCode != 200 {
		return "", "", fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	var r struct {
		TagName string `json:"tag_name"`
		HTMLURL string `json:"html_url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return "", "", err
	}
	return r.TagName, r.HTMLURL, nil
}

// newer compares v1.2.3 style versions.
func newer(latest, cur string) bool {
	if cur == "dev" || cur == "" {
		return false
	}
	p := func(s string) [3]int {
		var out [3]int
		s = strings.TrimPrefix(s, "v")
		s = strings.SplitN(s, "-", 2)[0]
		for i, part := range strings.SplitN(s, ".", 3) {
			fmt.Sscanf(part, "%d", &out[i])
		}
		return out
	}
	l, c := p(latest), p(cur)
	for i := 0; i < 3; i++ {
		if l[i] != c[i] {
			return l[i] > c[i]
		}
	}
	return false
}

func upgradeHint() string {
	exe, _ := os.Executable()
	switch {
	case strings.Contains(exe, "Cellar") || strings.Contains(exe, "homebrew"):
		return "brew upgrade gyro"
	case strings.Contains(strings.ToLower(exe), "scoop"):
		return "scoop update gyro"
	case runtime.GOOS == "windows":
		return i18n.S("download the new zip from the release page", "yeni zip'i sürüm sayfasından indirin")
	}
	return "curl -fsSL https://www.gyrohosting.com/cli/install.sh | sh   " + i18n.S("(or your package manager)", "(ya da paket yöneticiniz)")
}

/* ------------------------------------------------------------------ completion */

func addCompletion(a *App, root *cobra.Command) {
	cmd := &cobra.Command{
		Use:       "completion bash|zsh|fish|powershell",
		GroupID:   "cli",
		Short:     i18n.S("Generate a shell completion script", "Kabuk tamamlama betiği üret"),
		ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Long: i18n.S(`Load completions:
  bash:  source <(gyro completion bash)        # or: gyro completion bash > /etc/bash_completion.d/gyro
  zsh:   gyro completion zsh > "${fpath[1]}/_gyro"
  fish:  gyro completion fish > ~/.config/fish/completions/gyro.fish
  pwsh:  gyro completion powershell | Out-String | Invoke-Expression
Server, SSH key and other ids complete from the API.`,
			`Tamamlamayı yükleyin:
  bash:  source <(gyro completion bash)        # ya da: gyro completion bash > /etc/bash_completion.d/gyro
  zsh:   gyro completion zsh > "${fpath[1]}/_gyro"
  fish:  gyro completion fish > ~/.config/fish/completions/gyro.fish
  pwsh:  gyro completion powershell | Out-String | Invoke-Expression
Sunucu, SSH anahtarı ve diğer kimlikler API'den tamamlanır.`),
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "bash":
				return root.GenBashCompletionV2(a.Out, true)
			case "zsh":
				return root.GenZshCompletion(a.Out)
			case "fish":
				return root.GenFishCompletion(a.Out, true)
			default:
				return root.GenPowerShellCompletionWithDesc(a.Out)
			}
		},
	}
	root.AddCommand(cmd)
	root.CompletionOptions.DisableDefaultCmd = true
}

/* ------------------------------------------------------------------ docs (hidden) */

func addDocs(a *App, root *cobra.Command) {
	var dir, format string
	cmd := &cobra.Command{
		Use: "gen-docs", Hidden: true, Short: "Generate markdown / man pages for all commands",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return err
			}
			root.DisableAutoGenTag = true
			switch format {
			case "man":
				date := time.Date(2026, 11, 9, 0, 0, 0, 0, time.UTC)
				return doc.GenManTree(root, &doc.GenManHeader{Title: "GYRO", Section: "1", Source: "Gyro Hosting", Manual: "gyro manual", Date: &date}, dir)
			default:
				return doc.GenMarkdownTreeCustom(root, dir, func(string) string { return "" }, func(s string) string { return s })
			}
		},
	}
	cmd.Flags().StringVar(&dir, "dir", "docs/commands", "output directory")
	cmd.Flags().StringVar(&format, "format", "markdown", "markdown|man")
	root.AddCommand(cmd)
}

/* ------------------------------------------------------------------ raw API access */

func addAPICommand(a *App, root *cobra.Command) {
	g := &cobra.Command{
		Use: "api", GroupID: "cli",
		Short: i18n.S("Call any API endpoint directly (like 'gh api')", "Herhangi bir API ucunu doğrudan çağır ('gh api' gibi)"),
		Args:  cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error { return cmd.Help() },
	}
	var data, idem string
	var params, set []string
	var all bool
	req := &cobra.Command{
		Use:   "request <METHOD> <PATH>",
		Short: i18n.S("Send a raw request to a path below the API URL", "API adresinin altındaki bir yola ham istek gönder"),
		Example: "  gyro api request GET /cloud/servers -p status=active\n" +
			"  gyro api request POST /cloud/ssh-keys --set name=ci --set public_key=@~/.ssh/id_ed25519.pub",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			method := strings.ToUpper(args[0])
			body, err := a.ReadData(data)
			if err != nil {
				return err
			}
			extra, err := ParseKV(set)
			if err != nil {
				return err
			}
			for k, v := range extra {
				if s, ok := v.(string); ok && strings.HasPrefix(s, "@") {
					if s2, err := a.ReadValue(s); err == nil {
						v = strings.TrimSpace(s2)
					}
				}
				if body == nil {
					body = map[string]any{}
				}
				body[k] = v
			}
			q := url.Values{}
			for _, p := range params {
				k, v, _ := strings.Cut(p, "=")
				q.Add(k, v)
			}
			if method != "GET" {
				if err := a.Confirm(i18n.F("Send %s %s?", "%s %s gönderilsin mi?", method, args[1])); err != nil {
					return err
				}
			}
			ctx, cancel := a.Ctx()
			defer cancel()
			c, err := a.Client()
			if err != nil {
				return err
			}
			if all && method == "GET" {
				items, err := c.All(ctx, args[1], q, "", 0)
				if err != nil {
					return err
				}
				return a.PrintList(items, nil, "")
			}
			r := api.Request{Method: method, Path: args[1], Query: q, IdempotencyKey: idem}
			if body != nil {
				r.Body = body
			}
			resp, err := c.Do(ctx, r)
			if err != nil {
				return err
			}
			if err := a.PrintRaw(resp.Data, nil, ""); err != nil {
				return err
			}
			warnSecrets(a, resp.Data)
			return nil
		},
	}
	req.Flags().StringVar(&data, "data", "", i18n.S("JSON body, @file or @-", "JSON gövde, @dosya ya da @-"))
	req.Flags().StringArrayVar(&set, "set", nil, i18n.S("body field key=value (value @file reads a file)", "gövde alanı anahtar=değer (@dosya dosyadan okur)"))
	req.Flags().StringArrayVarP(&params, "param", "P", nil, i18n.S("query parameter key=value", "sorgu parametresi anahtar=değer"))
	req.Flags().StringVar(&idem, "idempotency-key", "", i18n.S("send an Idempotency-Key header", "Idempotency-Key başlığı gönder"))
	req.Flags().BoolVar(&all, "all", false, i18n.S("follow pagination (GET)", "sayfalamayı izle (GET)"))

	ops := &cobra.Command{
		Use: "operations", Aliases: []string{"ops"},
		Short: i18n.S("List every API operation and its CLI command", "Tüm API işlemlerini ve CLI komutlarını listele"),
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			var rows []any
			for _, o := range spec.Ops {
				command := "gyro " + strings.Join(o.Command, " ")
				if c, _, err := root.Find(o.Command); err != nil || c == nil || c.Annotations["operationId"] != o.ID {
					command = handwrittenFor(root, o) // covered by a hand-written command
				}
				rows = append(rows, map[string]any{"operation": o.ID, "method": o.Method, "path": o.Path, "scope": o.Scope, "command": command, "summary": summary(&o)})
			}
			sort.SliceStable(rows, func(i, j int) bool {
				return rows[i].(map[string]any)["path"].(string) < rows[j].(map[string]any)["path"].(string)
			})
			cols := []output.Column{{Header: "method", Path: "method"}, {Header: "path", Path: "path"}, {Header: "command", Path: "command"}, {Header: "scope", Path: "scope"}}
			return a.Printer().Print(toGeneric(rows), cols, "operation")
		},
	}
	g.AddCommand(req, ops)
	root.AddCommand(g)
}

// handwrittenFor names the hand-written command that covers an operation.
var handwrittenMap = map[string]string{
	"listServers": "gyro server list", "getServer": "gyro server get", "createServer": "gyro server create",
	"deleteServer": "gyro server delete", "updateServer": "gyro server update | rename", "getMetrics": "gyro server metrics",
	"createSshKey": "gyro ssh-key create", "listProducts": "gyro plan list", "listInvoices": "gyro invoice list",
	"getBalance": "gyro balance", "whoami": "gyro whoami",
}

func handwrittenFor(root *cobra.Command, o spec.Op) string {
	if s, ok := handwrittenMap[o.ID]; ok {
		return s
	}
	if o.ID == "serverAction" && len(o.Command) > 0 {
		switch o.Command[len(o.Command)-1] {
		case "poweron":
			return "gyro server power on"
		case "poweroff":
			return "gyro server power off"
		case "shutdown", "reboot", "reset":
			return "gyro server power " + o.Command[len(o.Command)-1]
		case "rebuild":
			return "gyro server rebuild"
		case "reset-password":
			return "gyro server reset-password"
		case "enable-rescue":
			return "gyro server rescue enable"
		case "disable-rescue":
			return "gyro server rescue disable"
		}
	}
	return "gyro api request " + o.Method + " " + o.Path
}
