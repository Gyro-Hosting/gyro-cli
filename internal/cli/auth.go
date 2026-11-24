package cli

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
	"github.com/Gyro-Hosting/gyro-cli/internal/output"
)

/* ------------------------------------------------------------------ auth (browser sign-in, API 2.31.0) */

type deviceCode struct {
	DeviceCode              string `json:"device_code"`
	UserCode                string `json:"user_code"`
	VerificationURI         string `json:"verification_uri"`
	VerificationURIComplete string `json:"verification_uri_complete"`
	ExpiresIn               int    `json:"expires_in"`
	Interval                int    `json:"interval"`
}

type authKey struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	TokenPrefix string   `json:"token_prefix"`
	Source      string   `json:"source"`
	ExpiresAt   *string  `json:"expires_at"`
	Permissions int      `json:"permissions"`
	Scopes      []string `json:"scopes"`
	IPs         []string `json:"ips"`
}

type authAccount struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type authSession struct {
	Key            authKey      `json:"key"`
	Account        authAccount  `json:"account"`
	AccountContext *authAccount `json:"account_context"` // API 2.46.0: account selected with X-Account-Id
}

type deviceToken struct {
	AccessToken    string       `json:"access_token"`
	Key            authKey      `json:"key"`
	Account        authAccount  `json:"account"`
	AccountContext *authAccount `json:"account_context"` // API 2.46.0: default account picked on the approval page
}

// headless reports an environment where opening a browser makes no sense (SSH session, no display).
func headless(env func(string) string) bool {
	if env("SSH_CONNECTION") != "" || env("SSH_TTY") != "" {
		return true
	}
	if runtime.GOOS == "linux" || runtime.GOOS == "freebsd" {
		return env("DISPLAY") == "" && env("WAYLAND_DISPLAY") == ""
	}
	return false
}

// openBrowser is replaced in tests.
var openBrowser = func(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	cmd.Stdout, cmd.Stderr = io.Discard, io.Discard
	return cmd.Start()
}

func hostname() string {
	h, _ := os.Hostname()
	if i := strings.IndexByte(h, '.'); i > 0 {
		h = h[:i]
	}
	return h
}

func expiryText(p *string) string {
	if p == nil || *p == "" {
		return i18n.S("never", "süresiz")
	}
	if t, err := time.Parse(time.RFC3339, *p); err == nil {
		return t.Local().Format("2006-01-02 15:04")
	}
	return *p
}

// deviceLogin runs the browser sign-in and stores the new key in the profile. Returns the key info.
func (a *App) deviceLogin(ctx context.Context, apiURL, template string, scopes []string, noBrowser, useKeyring bool) (*deviceToken, error) {
	cl := a.NewClient(apiURL, "")
	cl.Lang = i18n.Lang()
	if a.Opt.Debug {
		cl.Debug = a.Err
	}
	body := map[string]any{"client_name": "gyro-cli", "client_version": Version, "hostname": hostname(),
		"os": runtime.GOOS + "/" + runtime.GOARCH, "lang": i18n.Lang()}
	if template != "" {
		body["template"] = template
	}
	if len(scopes) > 0 {
		body["scopes"] = scopes
	}
	resp, err := cl.Do(ctx, api.Request{Method: "POST", Path: "/auth/device/code", Body: body})
	if err != nil {
		return nil, err
	}
	var dc deviceCode
	if err := json.Unmarshal(resp.Data, &dc); err != nil || dc.DeviceCode == "" {
		return nil, errors.New(i18n.S("unexpected answer from the sign-in service", "giriş servisinden beklenmeyen yanıt"))
	}
	fmt.Fprintln(a.Err, i18n.S("To sign in, open this address in a browser and approve the request:", "Giriş için bu adresi tarayıcıda açıp isteği onaylayın:"))
	fmt.Fprintf(a.Err, "\n  %s\n\n  %s %s\n\n", dc.VerificationURIComplete, i18n.S("Code:", "Kod:"), dc.UserCode)
	if !noBrowser && !headless(a.Env) {
		if openBrowser(dc.VerificationURIComplete) == nil {
			fmt.Fprintln(a.Err, i18n.S("(Your browser was opened. Check that it shows the same code.)", "(Tarayıcınız açıldı. Aynı kodu gösterdiğini kontrol edin.)"))
		}
	}
	interval := time.Duration(max(dc.Interval, 5)) * time.Second
	deadline := time.Now().Add(time.Duration(max(dc.ExpiresIn, 60)) * time.Second)
	sp := a.NewSpinner(i18n.S("Waiting for approval", "Onay bekleniyor"))
	defer sp.Done("")
	for {
		for waited := time.Duration(0); waited < interval; waited += 250 * time.Millisecond {
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}
			a.Sleep(250 * time.Millisecond)
			sp.Tick("")
		}
		if time.Now().After(deadline.Add(30 * time.Second)) {
			return nil, &ExitErr{ExitTimeout, errors.New(i18n.S("the code expired; run `gyro auth login` again", "kodun süresi doldu; `gyro auth login` komutunu yeniden çalıştırın"))}
		}
		r, err := cl.Do(ctx, api.Request{Method: "POST", Path: "/auth/device/token", Body: map[string]any{"device_code": dc.DeviceCode}})
		if err == nil {
			var tok deviceToken
			if json.Unmarshal(r.Data, &tok) != nil || tok.AccessToken == "" {
				return nil, errors.New(i18n.S("unexpected answer from the sign-in service", "giriş servisinden beklenmeyen yanıt"))
			}
			return &tok, nil
		}
		e, ok := api.AsError(err)
		if !ok {
			return nil, err
		}
		switch e.Code {
		case "authorization_pending":
			continue
		case "slow_down":
			if v, ok := e.Details["interval"].(float64); ok && v > 0 {
				interval = time.Duration(v) * time.Second
			} else {
				interval += 5 * time.Second
			}
			continue
		case "rate_limited", "network_error":
			interval += 5 * time.Second
			continue
		case "access_denied":
			return nil, &ExitErr{ExitAuth, errors.New(i18n.S("the sign-in was denied in the browser", "giriş tarayıcıda reddedildi") + messageSuffix(e))}
		case "expired_token":
			return nil, &ExitErr{ExitTimeout, errors.New(i18n.S("the code expired; run `gyro auth login` again", "kodun süresi doldu; `gyro auth login` komutunu yeniden çalıştırın"))}
		default:
			return nil, err
		}
	}
}

func messageSuffix(e *api.Error) string {
	if e.Message == "" || strings.HasPrefix(e.Message, "The sign-in request was denied") {
		return ""
	}
	return ": " + e.Message
}

// storeLogin writes the new key into the profile (same file/keychain rules as `gyro configure`).
func (a *App) storeLogin(name, apiURL string, tok *deviceToken, useKeyring bool) error {
	c, err := a.Config()
	if err != nil {
		return err
	}
	p := c.Get(name)
	if apiURL != "" && apiURL != api.DefaultURL {
		p.APIURL = apiURL
	}
	if err := c.SetToken(name, tok.AccessToken, useKeyring); err != nil {
		return err
	}
	// API 2.46.0: the account picked on the approval page becomes the profile default (own account = none)
	p.Account = 0
	if tok.AccountContext != nil && tok.AccountContext.ID > 0 && tok.AccountContext.ID != tok.Account.ID {
		p.Account = tok.AccountContext.ID
	}
	if c.DefaultProfile == "" {
		c.DefaultProfile = name
	}
	if err := c.Save(); err != nil {
		return err
	}
	a.client = nil
	fmt.Fprintln(a.Err, i18n.F("Signed in to account #%d %s (%s).", "#%d %s hesabına giriş yapıldı (%s).", tok.Account.ID, tok.Account.Name, tok.Account.Email))
	fmt.Fprintln(a.Err, i18n.F("Key %q: %d permission(s)%s, expires %s.", "%q anahtarı: %d izin%s, bitiş %s.", tok.Key.Name, tok.Key.Permissions,
		scopeList(tok.Key.Scopes), expiryText(tok.Key.ExpiresAt)))
	if p.Account > 0 {
		fmt.Fprintln(a.Err, i18n.F("Default account: #%d %s (as a panel user; override with --account self).", "Varsayılan hesap: #%d %s (panel kullanıcısı olarak; --account self ile geçersiz kılın).", p.Account, tok.AccountContext.Name))
	}
	fmt.Fprintln(a.Err, i18n.F("Saved to profile %q (%s).", "%q profiline kaydedildi (%s).", name, c.FilePath()))
	return nil
}

func scopeList(s []string) string {
	if len(s) == 0 {
		return ""
	}
	return " [" + strings.Join(s, ", ") + "]"
}

func addAuth(a *App, root *cobra.Command) {
	markHandled("createDeviceCode", "pollDeviceToken", "getAuthSession", "deleteAuthSession")
	auth := &cobra.Command{
		Use:     "auth",
		GroupID: "cli",
		Short:   i18n.S("Sign in with the browser, show or end the CLI session", "Tarayıcıyla giriş yap, CLI oturumunu göster ya da kapat"),
	}
	var noBrowser, useKeyring bool
	var template, scopes string
	login := &cobra.Command{
		Use:   "login",
		Short: i18n.S("Sign in with the browser (creates a time-limited API key for this profile)", "Tarayıcıyla giriş (bu profil için süreli bir API anahtarı oluşturur)"),
		Long: i18n.S(`Opens the Gyro panel in your browser; you approve the request there (permissions and expiry are chosen on that page)
and the new key is stored in the profile, exactly like 'gyro configure'. No key is copied by hand.
Over SSH or without a display the address is only printed: open it on any device where you are signed in.
In CI keep using a key: echo "$KEY" | gyro configure --token-stdin`,
			`Gyro panelini tarayıcıda açar; isteği orada onaylarsınız (izinler ve süre o sayfada seçilir) ve yeni anahtar
'gyro configure' ile aynı şekilde profile kaydedilir. Elle anahtar kopyalanmaz.
SSH bağlantısında ya da ekransız ortamda adres yalnız yazdırılır: giriş yapmış olduğunuz herhangi bir cihazda açın.
CI'da anahtar kullanmaya devam edin: echo "$ANAHTAR" | gyro configure --token-stdin`),
		Example: "  gyro auth login\n  gyro auth login --profile work --no-browser\n  gyro auth login --template readonly",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := a.Ctx()
			defer cancel()
			var sc []string
			for _, s := range strings.Split(scopes, ",") {
				if s = strings.TrimSpace(s); s != "" {
					sc = append(sc, s)
				}
			}
			tok, err := a.deviceLogin(ctx, a.APIURL(), template, sc, noBrowser, useKeyring)
			if err != nil {
				return err
			}
			return a.storeLogin(a.ProfileName(), a.Opt.APIURL, tok, useKeyring)
		},
	}
	login.Flags().BoolVar(&noBrowser, "no-browser", false, i18n.S("only print the address, do not open a browser", "yalnız adresi yaz, tarayıcı açma"))
	login.Flags().BoolVar(&useKeyring, "keyring", false, i18n.S("store the key in the OS keychain", "anahtarı işletim sistemi anahtar zincirinde sakla"))
	login.Flags().StringVar(&template, "template", "", i18n.S("preselect a permission template: readonly, dns, cloud, containers, storage, billing, full", "izin şablonunu önceden seç: readonly, dns, cloud, containers, storage, billing, full"))
	login.Flags().StringVar(&scopes, "scopes", "", i18n.S("preselect scopes (comma separated, e.g. cloud_read,services_read); you can still change them in the browser", "kapsamları önceden seç (virgülle, ör. cloud_read,services_read); tarayıcıda yine değiştirebilirsiniz"))
	_ = login.RegisterFlagCompletionFunc("template", fixedCompletion("readonly", "dns", "cloud", "containers", "storage", "billing", "full"))

	status := &cobra.Command{
		Use:   "status",
		Short: i18n.S("Show the account, key, permissions and expiry of the current profile", "Geçerli profilin hesabını, anahtarını, izinlerini ve bitişini göster"),
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cl, err := a.Client()
			if err != nil {
				return err
			}
			ctx, cancel := a.Ctx()
			defer cancel()
			var s authSession
			if _, err := cl.Get(ctx, "/auth/session", nil, &s); err != nil {
				return authHint(err)
			}
			_, src, _ := a.Token()
			row := map[string]any{"profile": a.ProfileName(), "token_source": src, "account_id": s.Account.ID, "account": s.Account.Name, "email": s.Account.Email,
				"key_id": s.Key.ID, "key_name": s.Key.Name, "key_prefix": s.Key.TokenPrefix, "source": s.Key.Source, "expires_at": expiryText(s.Key.ExpiresAt),
				"permissions": s.Key.Permissions, "scopes": strings.Join(s.Key.Scopes, ","), "ips": strings.Join(s.Key.IPs, ","), "acting_in": ""}
			if s.AccountContext != nil {
				row["acting_in"] = fmt.Sprintf("#%d %s", s.AccountContext.ID, s.AccountContext.Name)
			}
			if a.Machine() {
				return a.Printer().Print(s, nil, "")
			}
			cols := []output.Column{{Header: "profile", Path: "profile"}, {Header: "account_id", Path: "account_id"}, {Header: "account", Path: "account"},
				{Header: "email", Path: "email"}, {Header: "key_name", Path: "key_name"}, {Header: "key_prefix", Path: "key_prefix"}, {Header: "source", Path: "source"},
				{Header: "expires_at", Path: "expires_at"}, {Header: "permissions", Path: "permissions"}, {Header: "scopes", Path: "scopes"}, {Header: "ips", Path: "ips"}, {Header: "token_source", Path: "token_source"},
				{Header: "acting_in", Path: "acting_in"}}
			return a.Printer().Print(row, cols, "key_id")
		},
	}

	var revoke, keepKey bool
	logout := &cobra.Command{
		Use:   "logout",
		Short: i18n.S("Remove the key from this profile and revoke it on the server", "Anahtarı bu profilden sil ve sunucuda iptal et"),
		Long: i18n.S(`Keys created by 'gyro auth login' are revoked on the server and removed from the profile.
A key you pasted with 'gyro configure' is only removed locally (it may be used elsewhere); add --revoke to revoke it too.`,
			`'gyro auth login' ile oluşturulan anahtar sunucuda iptal edilir ve profilden silinir.
'gyro configure' ile yapıştırdığınız anahtar yalnız yerelden silinir (başka yerde kullanılıyor olabilir); onu da iptal etmek için --revoke.`),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if a.Env("GYRO_TOKEN") != "" {
				return &ExitErr{ExitUsage, errors.New(i18n.S("the key comes from GYRO_TOKEN; unset it (and revoke the key in the panel if needed)", "anahtar GYRO_TOKEN ortam değişkeninden geliyor; değişkeni kaldırın (gerekirse anahtarı panelden iptal edin)"))}
			}
			c, err := a.Config()
			if err != nil {
				return err
			}
			name := a.ProfileName()
			tok, _, err := c.TokenFor(name)
			if err != nil {
				return err
			}
			if tok == "" {
				fmt.Fprintln(a.Err, i18n.F("Profile %q has no key; nothing to do.", "%q profilinde anahtar yok; yapılacak bir şey yok.", name))
				return nil
			}
			ctx, cancel := a.Ctx()
			defer cancel()
			cl := a.NewClient(a.APIURL(), tok)
			cl.Lang = i18n.Lang()
			var s authSession
			_, serr := cl.Get(ctx, "/auth/session", nil, &s)
			if serr == nil && !keepKey && (s.Key.Source == "device" || revoke) {
				if _, err := cl.Do(ctx, api.Request{Method: "DELETE", Path: "/auth/session"}); err != nil {
					return fmt.Errorf("%s: %w", i18n.S("could not revoke the key on the server (profile unchanged)", "anahtar sunucuda iptal edilemedi (profil değişmedi)"), err)
				}
				fmt.Fprintln(a.Err, i18n.F("Key %q revoked on the server.", "%q anahtarı sunucuda iptal edildi.", s.Key.Name))
			} else if serr == nil {
				fmt.Fprintln(a.Err, i18n.F("Key %q was not revoked (it was not created by gyro auth login); revoke it in the panel or use --revoke.",
					"%q anahtarı iptal edilmedi (gyro auth login ile oluşturulmamış); panelden iptal edin ya da --revoke kullanın.", s.Key.Name))
			}
			if err := c.SetToken(name, "", false); err != nil {
				return err
			}
			if err := c.Save(); err != nil {
				return err
			}
			fmt.Fprintln(a.Err, i18n.F("Key removed from profile %q.", "Anahtar %q profilinden silindi.", name))
			return nil
		},
	}
	logout.Flags().BoolVar(&revoke, "revoke", false, i18n.S("also revoke a key that was not created by gyro auth login", "gyro auth login ile oluşturulmamış anahtarı da iptal et"))
	logout.Flags().BoolVar(&keepKey, "keep-key", false, i18n.S("only remove the key locally, never revoke it", "anahtarı yalnız yerelden sil, iptal etme"))

	auth.AddCommand(login, status, logout)
	root.AddCommand(auth)
}

// authHint adds the next step to expired/revoked key errors.
func authHint(err error) error {
	if e, ok := api.AsError(err); ok && (e.Code == "key_expired" || e.Code == "key_revoked" || e.Code == "invalid_token") {
		return &ExitErr{ExitAuth, fmt.Errorf("%w (%s)", err, i18n.S("sign in again: gyro auth login", "yeniden giriş yapın: gyro auth login"))}
	}
	return err
}
