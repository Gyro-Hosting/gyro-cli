package cli

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
	"github.com/Gyro-Hosting/gyro-cli/internal/output"
)

/* ------------------------------------------------------------------ SSH keys */

func addSSHKeys(a *App, root *cobra.Command) {
	g := ensureGroup(root, []string{"ssh-key"})
	g.Short = i18n.S("SSH public keys of your account", "Hesabınızdaki SSH açık anahtarları")
	delete(g.Annotations, "autoShort")
	markHandled("createSshKey")
	var name, file string
	add := &cobra.Command{
		Use: "create [PUBLIC-KEY-FILE]", Aliases: []string{"add", "import"},
		Short: i18n.S("Upload an SSH public key", "SSH açık anahtarı yükle"),
		Example: "  gyro ssh-key create                      # ~/.ssh/id_ed25519.pub\n" +
			"  gyro ssh-key create ~/.ssh/work.pub --name work\n  cat key.pub | gyro ssh-key create - --name ci",
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				file = args[0]
			}
			var pub string
			switch {
			case file == "-":
				b, err := readAll(a.In)
				if err != nil {
					return err
				}
				pub = string(b)
			case file != "":
				b, err := os.ReadFile(expandHome(file))
				if err != nil {
					return err
				}
				pub = string(b)
			default:
				home, _ := os.UserHomeDir()
				for _, n := range []string{"id_ed25519.pub", "id_ecdsa.pub", "id_rsa.pub"} {
					if b, err := os.ReadFile(filepath.Join(home, ".ssh", n)); err == nil {
						pub, file = string(b), filepath.Join(home, ".ssh", n)
						break
					}
				}
				if pub == "" {
					return &ExitErr{ExitUsage, fmt.Errorf("%s", i18n.S("no public key found in ~/.ssh; give a file", "~/.ssh içinde açık anahtar yok; dosya verin"))}
				}
				a.Info("%s", i18n.F("Using %s", "%s kullanılıyor", file))
			}
			pub = strings.TrimSpace(pub)
			if name == "" {
				if f := strings.Fields(pub); len(f) >= 3 {
					name = f[2]
				} else if file != "" && file != "-" {
					name = strings.TrimSuffix(filepath.Base(file), ".pub")
				} else {
					name = "gyro-cli"
				}
			}
			ctx, cancel := a.Ctx()
			defer cancel()
			c, err := a.Client()
			if err != nil {
				return err
			}
			resp, err := c.Do(ctx, api.Request{Method: "POST", Path: "/cloud/ssh-keys", Body: map[string]any{"name": name, "public_key": pub}})
			if err != nil {
				return err
			}
			return a.PrintRaw(resp.Data, sshKeyCols, "id")
		},
	}
	add.Flags().StringVar(&name, "name", "", i18n.S("key name (default: the key comment)", "anahtar adı (varsayılan: anahtar yorumu)"))
	g.AddCommand(add)
}

var sshKeyCols = []output.Column{{Header: "id", Path: "id"}, {Header: "name", Path: "name"}, {Header: "fingerprint", Path: "fingerprint"}, {Header: "created", Path: "created_at"}}

/* ------------------------------------------------------------------ catalog */

func addCatalog(a *App, root *cobra.Command) {
	markHandled("listProducts")
	plan := &cobra.Command{
		Use: "plan", Aliases: []string{"plans"},
		Short: i18n.S("Server plans and prices (catalog)", "Sunucu planları ve fiyatlar (katalog)"),
		Args:  cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error { return cmd.Help() },
	}
	var loc, typ, cycle string
	var all bool
	list := &cobra.Command{
		Use: "list", Aliases: []string{"ls"},
		Short:   i18n.S("List orderable plans with prices", "Sipariş verilebilen planları fiyatlarıyla listele"),
		Example: "  gyro plan list --location falkenstein\n  gyro plan list --cycle annually -o json",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := a.Ctx()
			defer cancel()
			c, err := a.Client()
			if err != nil {
				return err
			}
			code, err := a.resolveLocation(ctx, loc)
			if err != nil {
				return err
			}
			q := url.Values{}
			if code != "" {
				q.Set("location", code)
			}
			if typ != "" {
				q.Set("server_type", typ)
			}
			items, err := c.All(ctx, "/cloud/products", q, "", 0)
			if err != nil {
				return err
			}
			var rows []any
			for _, it := range items {
				v, _ := output.Decode(it)
				m, ok := v.(map[string]any)
				if !ok {
					continue
				}
				if !all && m["available"] == false {
					continue
				}
				m["plan"] = PlanSlug(fmt.Sprint(m["name"]))
				if prices, ok := m["prices"].([]any); ok {
					for _, p := range prices {
						pm, _ := p.(map[string]any)
						if pm != nil && pm["cycle"] == cycle {
							m["price"] = map[string]any{"amount": pm["amount"], "currency": pm["currency"]}
						}
					}
				}
				rows = append(rows, m)
			}
			sort.SliceStable(rows, func(i, j int) bool {
				mi, mj := rows[i].(map[string]any), rows[j].(map[string]any)
				pi, _ := output.Lookup(mi, "price.amount").(float64)
				pj, _ := output.Lookup(mj, "price.amount").(float64)
				if pi != pj {
					return pi < pj
				}
				return fmt.Sprint(mi["location"]) < fmt.Sprint(mj["location"])
			})
			cols := []output.Column{{Header: "id", Path: "id"}, {Header: "plan", Path: "plan"}, {Header: "type", Path: "server_type"},
				{Header: "location", Path: "location"}, {Header: "vcpu", Path: "cores"}, {Header: "ram_gb", Path: "memory_gb"},
				{Header: "disk_gb", Path: "disk_gb"}, {Header: "cpu", Path: "cpu_type"}, {Header: "price/" + cycle, Path: "price"}}
			if all {
				cols = append(cols, output.Column{Header: "available", Path: "available"})
			}
			return a.Printer().Print(rows, cols, "id")
		},
	}
	list.Flags().StringVar(&loc, "location", "", i18n.S("location code or city", "lokasyon kodu ya da şehir"))
	list.Flags().StringVar(&typ, "server-type", "", i18n.S("server type, e.g. cx23", "sunucu tipi, ör. cx23"))
	list.Flags().StringVar(&cycle, "cycle", "monthly", "monthly|quarterly|semiannually|annually")
	list.Flags().BoolVar(&all, "all", false, i18n.S("include out-of-stock plans", "stokta olmayanları da göster"))
	_ = list.RegisterFlagCompletionFunc("location", completeField(a, "/cloud/locations", "code"))
	plan.AddCommand(list)
	root.AddCommand(plan)
}

/* ------------------------------------------------------------------ billing */

func addBilling(a *App, root *cobra.Command) {
	markHandled("listInvoices", "getBalance", "whoami")
	inv := ensureGroup(root, []string{"invoice"})
	inv.Short = i18n.S("Invoices", "Faturalar")
	delete(inv.Annotations, "autoShort")
	inv.GroupID = "acct"
	var unpaid, all bool
	var status string
	list := &cobra.Command{
		Use: "list", Aliases: []string{"ls"},
		Short:   i18n.S("List invoices", "Faturaları listele"),
		Example: "  gyro invoice list --unpaid\n  gyro invoice list --all -o json --query \"[?state=='paid'].total.amount\"",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := a.Ctx()
			defer cancel()
			c, err := a.Client()
			if err != nil {
				return err
			}
			q := url.Values{}
			if status != "" {
				q.Set("status", status)
			}
			items, err := c.All(ctx, "/invoices", q, "", pagesFor(all || unpaid))
			if err != nil {
				return err
			}
			var rows []any
			for _, it := range items {
				v, _ := output.Decode(it)
				m, _ := v.(map[string]any)
				if m == nil {
					continue
				}
				if unpaid {
					st := fmt.Sprint(m["state"])
					if st != "unpaid" && st != "overdue" {
						continue
					}
				}
				rows = append(rows, m)
			}
			if rows == nil {
				rows = []any{}
			}
			cols := []output.Column{{Header: "id", Path: "invoice_id"}, {Header: "number", Path: "number"}, {Header: "state", Path: "state"},
				{Header: "total", Path: "total"}, {Header: "due_amount", Path: "amount_due"}, {Header: "created", Path: "created_at"},
				{Header: "due", Path: "due_date"}, {Header: "item", Path: "first_item"}}
			if err := a.Printer().Print(rows, cols, "invoice_id"); err != nil {
				return err
			}
			if unpaid && len(rows) == 0 && !a.Machine() {
				a.Info("%s", i18n.S("No unpaid invoices.", "Ödenmemiş fatura yok."))
			}
			return nil
		},
	}
	list.Flags().BoolVar(&unpaid, "unpaid", false, i18n.S("only unpaid and overdue invoices", "yalnız ödenmemiş ve gecikmiş faturalar"))
	list.Flags().StringVar(&status, "status", "", i18n.S("API status filter", "API durum süzgeci"))
	list.Flags().BoolVar(&all, "all", false, i18n.S("fetch every page", "tüm sayfaları getir"))
	inv.AddCommand(list)

	bal := ensureGroup(root, []string{"balance"})
	bal.GroupID = "acct"
	bal.Short = i18n.S("Account balance (bare command) and transactions", "Hesap bakiyesi (tek başına) ve hareketler")
	delete(bal.Annotations, "autoShort")
	bal.Annotations = map[string]string{"generated": "1"}
	bal.RunE = func(cmd *cobra.Command, args []string) error {
		ctx, cancel := a.Ctx()
		defer cancel()
		c, err := a.Client()
		if err != nil {
			return err
		}
		resp, err := c.Do(ctx, api.Request{Method: "GET", Path: "/balance"})
		if err != nil {
			return err
		}
		if a.Machine() || a.Opt.Query != "" || a.Printer().Format == "text" {
			return a.PrintRaw(resp.Data, nil, "")
		}
		var b struct {
			Balance    map[string]any `json:"balance"`
			LowBalance bool           `json:"low_balance"`
			AutoPay    struct {
				FromBalance bool `json:"from_balance"`
			} `json:"auto_pay"`
		}
		if err := json.Unmarshal(resp.Data, &b); err != nil || b.Balance == nil {
			return a.PrintRaw(resp.Data, nil, "")
		}
		fmt.Fprintf(a.Out, "%s: %s\n", i18n.S("Balance", "Bakiye"), output.Scalar(b.Balance))
		if b.LowBalance {
			fmt.Fprintln(a.Out, i18n.S("Warning: balance is low.", "Uyarı: bakiye düşük."))
		}
		fmt.Fprintf(a.Out, "%s: %s\n", i18n.S("Pay invoices automatically from balance", "Faturaları bakiyeden otomatik öde"), output.Scalar(b.AutoPay.FromBalance))
		return nil
	}

	who := &cobra.Command{
		Use:     "whoami",
		GroupID: "acct",
		Short:   i18n.S("Show the API key in use and its permissions", "Kullanılan API anahtarını ve yetkilerini göster"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := a.Ctx()
			defer cancel()
			c, err := a.Client()
			if err != nil {
				return err
			}
			resp, err := c.Do(ctx, api.Request{Method: "GET", Path: "/whoami"})
			if err != nil {
				return err
			}
			return a.PrintRaw(resp.Data, nil, "id")
		},
	}
	root.AddCommand(who)
}
