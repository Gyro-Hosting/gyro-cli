package cli

// Account context (API 2.46.0, X-Account-Id): work in an account whose owner added you as a panel user,
// with your own key. Precedence: --account > GYRO_ACCOUNT > profile "account" (gyro account use / gyro auth login).

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
	"github.com/Gyro-Hosting/gyro-cli/internal/output"
)

type accountAccess struct {
	Services     string `json:"services"`
	Domains      string `json:"domains"`
	Invoices     string `json:"invoices"`
	Tickets      string `json:"tickets"`
	Orders       bool   `json:"orders"`
	SSO          bool   `json:"sso"`
	Passwords    bool   `json:"passwords"`
	AccountAdmin bool   `json:"account_admin"`
}

type accountItem struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	Email       string        `json:"email"`
	IsSelf      bool          `json:"is_self"`
	Role        string        `json:"role"`
	Permissions []string      `json:"permissions"`
	Access      accountAccess `json:"access"`
}

// parseAccount normalises an --account / GYRO_ACCOUNT value: "" | "self" | "0" = own account ("").
func parseAccount(v, from string) (string, error) {
	v = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(v), "#"))
	if v == "" || strings.EqualFold(v, "self") || v == "0" {
		return "", nil
	}
	if n, err := strconv.ParseInt(v, 10, 64); err != nil || n < 1 {
		return "", &ExitErr{ExitUsage, fmt.Errorf("%s: %s", from, i18n.F("invalid account id %q (a number from `gyro account list`, or self)", "geçersiz hesap id %q (`gyro account list` içindeki sayı ya da self)", v))}
	}
	return v, nil
}

// AccountID is the account context sent as X-Account-Id ("" = own account).
func (a *App) AccountID() (string, error) {
	if a.Opt.Account != "" {
		return parseAccount(a.Opt.Account, "--account")
	}
	if e := a.Env("GYRO_ACCOUNT"); e != "" {
		return parseAccount(e, "GYRO_ACCOUNT")
	}
	if p := a.Profile(); p.Account > 0 {
		return strconv.FormatInt(p.Account, 10), nil
	}
	return "", nil
}

func addAccount(a *App, root *cobra.Command) {
	markHandled("listAccounts")
	grp := ensureGroup(root, []string{"account"})
	if grp.Annotations["autoShort"] != "" {
		grp.Short = i18n.S("Accounts you can act in, account settings", "İşlem yapabildiğiniz hesaplar, hesap ayarları")
		delete(grp.Annotations, "autoShort")
	}
	grp.GroupID = "acct"

	list := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   i18n.S("Accounts you can act in with --account (your own + accounts you are a panel user of)", "--account ile işlem yapabildiğiniz hesaplar (kendi hesabınız + panel kullanıcısı olduğunuz hesaplar)"),
		Long: i18n.S(`Lists your own account and every account whose owner added you under Account > Users, with your permissions there.
Act in one of them with --account <id>, GYRO_ACCOUNT=<id> or make it the profile default with 'gyro account use <id>'.
Permissions are exactly the panel permissions; owner-only operations (API keys, webhooks, users, balance, cards) are refused.`,
			`Kendi hesabınızı ve sizi Hesap > Kullanıcılar altında ekleyen her hesabı oradaki izinlerinizle listeler.
Birinde işlem yapmak için --account <id>, GYRO_ACCOUNT=<id> ya da profil varsayılanı için 'gyro account use <id>'.
İzinler paneldeki izinlerin aynısıdır; yalnız hesap sahibine açık işlemler (API anahtarları, webhook, kullanıcılar, bakiye, kartlar) reddedilir.`),
		Example:     "  gyro account list\n  gyro server list --account 12345\n  gyro account use 12345",
		Args:        cobra.NoArgs,
		Annotations: map[string]string{"operationId": "listAccounts"},
		RunE: func(cmd *cobra.Command, args []string) error {
			items, raw, err := a.listAccounts()
			if err != nil {
				return err
			}
			if a.Machine() || a.Opt.Query != "" || a.Printer().Format == "text" {
				return a.PrintRaw(raw, nil, "")
			}
			cur, _ := a.AccountID()
			rows := make([]any, 0, len(items))
			for _, it := range items {
				mark := ""
				if (cur == "" && it.IsSelf) || cur == strconv.FormatInt(it.ID, 10) {
					mark = "*"
				}
				rows = append(rows, map[string]any{"current": mark, "id": it.ID, "name": it.Name, "email": it.Email, "role": it.Role,
					"services": it.Access.Services, "domains": it.Access.Domains, "invoices": it.Access.Invoices, "tickets": it.Access.Tickets,
					"orders": it.Access.Orders, "passwords": it.Access.Passwords, "permissions": strings.Join(it.Permissions, ",")})
			}
			cols := []output.Column{{Header: "current", Path: "current"}, {Header: "id", Path: "id"}, {Header: "name", Path: "name"}, {Header: "email", Path: "email"},
				{Header: "role", Path: "role"}, {Header: "services", Path: "services"}, {Header: "domains", Path: "domains"}, {Header: "invoices", Path: "invoices"},
				{Header: "tickets", Path: "tickets"}, {Header: "orders", Path: "orders"}, {Header: "passwords", Path: "passwords"},
				{Header: "permissions", Path: "permissions"}}
			return a.Printer().Print(rows, cols, "id")
		},
	}

	use := &cobra.Command{
		Use:     "use <id|self>",
		Short:   i18n.S("Make an account the default of this profile (self = your own account)", "Bir hesabı bu profilin varsayılanı yap (self = kendi hesabınız)"),
		Example: "  gyro account use 12345\n  gyro account use self",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseAccount(args[0], "account")
			if err != nil {
				return err
			}
			var name string
			if id != "" {
				items, _, err := a.listAccounts()
				if err != nil {
					return err
				}
				for _, it := range items {
					if strconv.FormatInt(it.ID, 10) == id {
						if it.IsSelf {
							id = ""
						}
						name = it.Name
					}
				}
				if id != "" && name == "" {
					return &ExitErr{ExitAuth, errors.New(i18n.F("account #%s is not in `gyro account list` (you are not an active panel user there)", "#%s hesabı `gyro account list` içinde yok (orada etkin panel kullanıcısı değilsiniz)", id))}
				}
			}
			c, err := a.Config()
			if err != nil {
				return err
			}
			p := c.Get(a.ProfileName())
			p.Account = 0
			if id != "" {
				p.Account, _ = strconv.ParseInt(id, 10, 64)
			}
			if err := c.Save(); err != nil {
				return err
			}
			a.client = nil
			if id == "" {
				a.Info(i18n.F("Profile %q now works in your own account.", "%q profili artık kendi hesabınızda çalışır.", a.ProfileName()))
			} else {
				a.Info(i18n.F("Profile %q now works in account #%s %s (override with --account self).", "%q profili artık #%s %s hesabında çalışır (--account self ile geçersiz kılın).", a.ProfileName(), id, name))
			}
			return nil
		},
	}
	grp.AddCommand(list, use)
}

// listAccounts calls GET /accounts WITHOUT an account context (the list is always the key holder's).
func (a *App) listAccounts() ([]accountItem, []byte, error) {
	cl, err := a.Client()
	if err != nil {
		return nil, nil, err
	}
	ctx, cancel := a.Ctx()
	defer cancel()
	saved := cl.AccountID
	cl.AccountID = ""
	defer func() { cl.AccountID = saved }()
	resp, err := cl.Do(ctx, api.Request{Method: "GET", Path: "/accounts", Query: url.Values{"limit": {"100"}}})
	if err != nil {
		return nil, nil, err
	}
	var items []accountItem
	if err := json.Unmarshal(resp.Data, &items); err != nil {
		return nil, nil, err
	}
	return items, resp.Data, nil
}
