package cli

// gyro limits: every account limit with the current usage (GET /cloud/account/limits, API 2.26.0) and how to raise it.
// Limits are raised with a support ticket: gyro ticket create --department-id 3 --subject "Limit increase: <key>" --message "...".

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
)

// LimitTicketCmd is the ticket command printed for limit refusals when the API does not send one.
const LimitTicketCmd = `gyro ticket create --department-id 3 --subject "Limit increase: %s" --message "Requested limit: <N>. Reason: <what you are building>."`

func limitsHelp() string {
	return i18n.S(`Account limits: every limit of your account with the current usage (new-account limit of the prepaid cloud,
Gyro Cloud Istanbul servers, templates, additional IPs and disks per server, API order caps, keys, webhooks,
firewall rules and security groups, storage buckets and keys, DNS zones, container plans, Gyro Pipeline,
uptime monitors, projects).

Limits are raised on request: open a support ticket.
  gyro ticket create --department-id 3 --subject "Limit increase: <key>" --message "Requested limit: <N>. Reason: ..."
Plan limits (buckets, zones, container services ...) also grow with a bigger plan or an add-on.
When a command is refused because of a limit, the error shows the limit, the usage and the ready ticket command.`,
		`Hesap sınırları: hesabınızın tüm sınırları ve şu anki kullanım (ön ödemeli bulutta yeni hesap sınırı, Gyro Bulut İstanbul
sunucuları, şablonlar, sunucu başına ek IP ve disk, API sipariş sınırları, anahtarlar, webhook'lar, güvenlik duvarı
kuralları ve grupları, depolama kovaları ve anahtarları, DNS zone'ları, kapsayıcı paketleri, Gyro Pipeline,
uptime izleyicileri, projeler).

Sınırlar talep üzerine artırılır: destek talebi açın.
  gyro ticket create --department-id 3 --subject "Limit increase: <anahtar>" --message "İstenen sınır: <N>. Gerekçe: ..."
Paket sınırları (kova, zone, kapsayıcı servisi ...) paketi yükselterek ya da ek hizmetle de büyür.
Bir komut sınır yüzünden reddedilirse hata çıktısı sınırı, kullanımı ve hazır talep komutunu gösterir.`)
}

func addLimits(a *App, root *cobra.Command) {
	markHandled("getAccountLimits")
	var fresh bool
	cmd := &cobra.Command{
		Use:     "limits",
		Aliases: []string{"limit", "quota", "quotas"},
		GroupID: "acct",
		Short:   i18n.S("Account limits and usage; how to raise them (support ticket)", "Hesap sınırları ve kullanım; nasıl artırılır (destek talebi)"),
		Long:    limitsHelp(),
		Example: "  gyro limits\n  gyro limits -o json --query \"limits[?key=='compute.max_vms']\"\n  gyro ticket create --department-id 3 --subject \"Limit increase: new_account.servers\" --message \"Requested limit: 10. Reason: ...\"",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := a.Ctx()
			defer cancel()
			c, err := a.Client()
			if err != nil {
				return err
			}
			q := url.Values{}
			if fresh {
				q.Set("fresh", "1")
			}
			resp, err := c.Do(ctx, api.Request{Method: "GET", Path: "/cloud/account/limits", Query: q})
			if err != nil {
				return err
			}
			if a.Machine() || a.Opt.Query != "" || a.Printer().Format == "text" {
				return a.PrintRaw(resp.Data, nil, "")
			}
			var d struct {
				NewAccount *struct {
					Applies    bool                      `json:"applies"`
					Unlimited  bool                      `json:"unlimited"`
					PeriodDays int                       `json:"period_days"`
					EndsAt     string                    `json:"ends_at"`
					Servers    struct{ Limit, Used int } `json:"servers"`
					Envs       struct{ Limit, Used int } `json:"container_environments"`
				} `json:"new_account"`
				Limits []struct {
					Key, Scope, Unit, How, Desc, DescTR string
					Limit, Used                         *int
					Custom                              bool
					MaxSize                             *struct {
						Vcpu     int `json:"vcpu"`
						MemoryGB int `json:"memory_gb"`
						DiskGB   int `json:"disk_gb"`
					} `json:"max_size"`
					Services []struct {
						ID          int    `json:"service_id"`
						Name        string `json:"name"`
						Limit, Used *int
					} `json:"services"`
				} `json:"limits"`
			}
			if err := json.Unmarshal(fixLimitKeys(resp.Data), &d); err != nil {
				return a.PrintRaw(resp.Data, nil, "")
			}
			w := a.Out
			num := func(p *int) string {
				if p == nil {
					return "-"
				}
				return fmt.Sprint(*p)
			}
			if n := d.NewAccount; n != nil {
				state := i18n.S("not active (account older than ", "geçerli değil (hesap ") + fmt.Sprint(n.PeriodDays) + i18n.S(" days)", " günden eski)")
				if n.Unlimited {
					state = i18n.S("unlimited", "sınırsız")
				} else if n.Applies {
					state = i18n.S("active until ", "geçerli, bitiş ") + n.EndsAt
				}
				fmt.Fprintf(w, "%s: %s\n  %s %d/%d   %s %d/%d\n\n", i18n.S("New-account limit (prepaid cloud)", "Yeni hesap sınırı (ön ödemeli bulut)"), state,
					i18n.S("Istanbul servers", "İstanbul sunucuları"), n.Servers.Used, n.Servers.Limit,
					i18n.S("container environments", "kapsayıcı ortamları"), n.Envs.Used, n.Envs.Limit)
			}
			if !a.Opt.NoHeader {
				fmt.Fprintf(w, "%-28s %-8s %8s %8s  %s\n", i18n.S("LIMIT", "SINIR"), i18n.S("SCOPE", "KAPSAM"), i18n.S("MAX", "EN ÇOK"), i18n.S("USED", "KULLANIM"), i18n.S("HOW TO RAISE", "NASIL ARTAR"))
			}
			for _, l := range d.Limits {
				max := num(l.Limit)
				if l.Limit == nil && l.Scope == "service" {
					max = i18n.S("plan", "paket")
				}
				if l.MaxSize != nil {
					max = fmt.Sprintf("%d/%d/%d", l.MaxSize.Vcpu, l.MaxSize.MemoryGB, l.MaxSize.DiskGB)
				}
				if l.Custom {
					max += "*"
				}
				fmt.Fprintf(w, "%-28s %-8s %8s %8s  %s\n", l.Key, l.Scope, max, num(l.Used), howText(l.How))
				for _, s := range l.Services {
					if a.Opt.Wide || len(l.Services) <= 3 {
						fmt.Fprintf(w, "  #%-25d %-8s %8s %8s\n", s.ID, "", num(s.Limit), num(s.Used))
					}
				}
			}
			fmt.Fprintf(w, "\n%s\n  %s\n", i18n.S("Raise a limit with a support ticket (* = changed for your account):", "Sınırı destek talebiyle artırın (* = hesabınıza özel):"),
				fmt.Sprintf(LimitTicketCmd, "<"+i18n.S("limit", "sınır")+">"))
			return nil
		},
	}
	cmd.Flags().BoolVar(&fresh, "fresh", false, i18n.S("read again (answers are cached for 60 s)", "yeniden oku (yanıt 60 sn önbellekli)"))
	root.AddCommand(cmd)
}

// fixLimitKeys maps the API field names onto the short struct field names above.
func fixLimitKeys(b json.RawMessage) json.RawMessage {
	r := strings.NewReplacer(`"how_to_increase":`, `"How":`, `"description_tr":`, `"DescTR":`, `"description":`, `"Desc":`)
	return json.RawMessage(r.Replace(string(b)))
}

func howText(h string) string {
	switch h {
	case "open_support_ticket":
		return i18n.S("support ticket", "destek talebi")
	case "upgrade_plan_or_open_support_ticket":
		return i18n.S("bigger plan / support ticket", "büyük paket / destek talebi")
	case "upgrade_plan":
		return i18n.S("bigger plan or add-on", "büyük paket ya da ek hizmet")
	case "wait":
		return i18n.S("rate limit, wait", "hız sınırı, bekleyin")
	case "fixed":
		return i18n.S("fixed", "sabit")
	}
	return h
}

// limitHint turns details.how_to_increase (API 2.26.0) into one or two lines for the error output.
func limitHint(e *api.Error) string {
	how, _ := e.Details["how_to_increase"].(string)
	if how == "" {
		return ""
	}
	key, _ := e.Details["limit_key"].(string)
	if how == "wait" {
		return i18n.S("This is a short rate limit: wait and retry.", "Bu kısa süreli bir hız sınırı: bekleyip tekrar deneyin.")
	}
	cli := ""
	if t, ok := e.Details["support_ticket"].(map[string]any); ok {
		cli, _ = t["cli"].(string)
	}
	if cli == "" {
		k := key
		if k == "" {
			k = e.Code
		}
		cli = fmt.Sprintf(LimitTicketCmd, k)
	}
	pre := ""
	if how == "upgrade_plan" {
		pre = i18n.S("Choose a bigger plan or an add-on; for more than the largest plan, ", "Daha büyük paket ya da ek hizmet seçin; en büyük paketin üstü için ")
	}
	return pre + i18n.S("To raise the limit: ", "Sınırı artırmak için: ") + cli + "\n  " +
		i18n.S("All limits and usage: gyro limits", "Tüm sınırlar ve kullanım: gyro limits")
}

// isLimitDetail reports detail keys that limitHint prints in a readable form.
func isLimitDetail(k string) bool {
	switch k {
	case "support_ticket", "hint", "hint_tr", "limits_endpoint", "how_to_increase", "limit_exceeded":
		return true
	}
	return false
}

// quoteLimitError turns quote.limit_exceeded of a dry-run answer into the same error the real order would return.
func quoteLimitError(m map[string]any) *api.Error {
	d := map[string]any{}
	for k, v := range m {
		if k != "code" && k != "message" {
			d[k] = v
		}
	}
	msg, _ := m["message"].(string)
	return &api.Error{Status: 422, Code: "limit_exceeded", Message: msg, Details: d}
}
