package cli

// Prepaid hourly billing (Gyro Cloud Istanbul tr-ist-1, Gyro Containers): gyro billing balance|usage|estimate|statement.
// Endpoints: GET /cloud/billing/balance, GET /cloud/billing/usage, POST /cloud/billing/estimate,
// GET /cloud/billing/statement/{yyyy-mm}. Scope: billing_read.

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
	"github.com/Gyro-Hosting/gyro-cli/internal/output"
)

var monthRe = regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])$`)

func addPrepaidBilling(a *App, root *cobra.Command) {
	// spec 2.11.0 carries these operations: the hand-written commands below replace the generated ones (topup-url stays generated)
	markHandled("getPrepaidBalance", "listPrepaidUsage", "getPrepaidStatement", "estimatePrepaidCost")
	g := ensureGroup(root, []string{"billing"})
	g.GroupID = "acct"
	g.Short = i18n.S("Prepaid hourly usage (Istanbul cloud, containers): balance, usage, estimate, statement",
		"Ön ödemeli saatlik kullanım (İstanbul bulut, kapsayıcı): bakiye, kullanım, tahmin, döküm")
	delete(g.Annotations, "autoShort")

	get := func(path string, q url.Values) (*api.Response, error) {
		ctx, cancel := a.Ctx()
		defer cancel()
		c, err := a.Client()
		if err != nil {
			return nil, err
		}
		return c.Do(ctx, api.Request{Method: "GET", Path: path, Query: q})
	}
	save := func(data json.RawMessage, file string) error {
		var f struct {
			Filename string `json:"filename"`
			B64      string `json:"content_base64"`
		}
		if err := json.Unmarshal(data, &f); err != nil || f.B64 == "" {
			return fmt.Errorf("%s", i18n.S("the answer carries no file", "yanıtta dosya yok"))
		}
		b, err := base64.StdEncoding.DecodeString(f.B64)
		if err != nil {
			return err
		}
		if file == "" || file == "." {
			file = f.Filename
		}
		if file == "-" {
			_, err = a.Out.Write(b)
			return err
		}
		if err := os.WriteFile(file, b, 0o600); err != nil {
			return err
		}
		a.Info("%s: %s (%d bytes)", i18n.S("Saved", "Kaydedildi"), file, len(b))
		return nil
	}

	/* -------------------------------------------------------------- balance */
	bal := &cobra.Command{
		Use:     "balance",
		Short:   i18n.S("Balance, hourly cost and estimated time left", "Bakiye, saatlik ücret ve tahmini kalan süre"),
		Example: "  gyro billing balance\n  gyro billing balance -o json --query hours_left",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := get("/cloud/billing/balance", nil)
			if err != nil {
				return err
			}
			if a.Machine() || a.Opt.Query != "" || a.Printer().Format == "text" {
				return a.PrintRaw(resp.Data, nil, "")
			}
			var b struct {
				Balance    map[string]any `json:"balance"`
				Hourly     map[string]any `json:"hourly_cost"`
				Daily      map[string]any `json:"daily_cost"`
				Upcoming   map[string]any `json:"upcoming_invoices_7d"`
				HoursLeft  *float64       `json:"hours_left"`
				RunsOut    *string        `json:"runs_out_at"`
				Level      string         `json:"level"`
				Services   map[string]int `json:"services"`
				SuspHourly map[string]any `json:"suspended_hourly_cost"`
				EurRate    float64        `json:"eur_rate"`
				TopupURL   string         `json:"topup_url"`
			}
			if err := json.Unmarshal(resp.Data, &b); err != nil || b.Balance == nil {
				return a.PrintRaw(resp.Data, nil, "")
			}
			w := a.Out
			fmt.Fprintf(w, "%-28s %s\n", i18n.S("Balance", "Bakiye")+":", output.Scalar(b.Balance))
			fmt.Fprintf(w, "%-28s %s  (%s %s)\n", i18n.S("Hourly cost", "Saatlik ücret")+":", output.Scalar(b.Hourly), i18n.S("daily", "günlük"), output.Scalar(b.Daily))
			if b.HoursLeft != nil {
				h := *b.HoursLeft
				d := int(h) / 24
				left := fmt.Sprintf("%d %s %d %s", d, i18n.S("d", "gün"), int(h)-d*24, i18n.S("h", "saat"))
				until := ""
				if b.RunsOut != nil {
					until = " (" + strings.Replace((*b.RunsOut)[:16], "T", " ", 1) + ")"
				}
				fmt.Fprintf(w, "%-28s %s%s [%s]\n", i18n.S("Estimated time left", "Tahmini kalan süre")+":", left, until, b.Level)
			} else {
				fmt.Fprintf(w, "%-28s %s\n", i18n.S("Estimated time left", "Tahmini kalan süre")+":", i18n.S("no running usage", "çalışan kullanım yok"))
			}
			if amt, _ := b.Upcoming["amount"].(float64); amt > 0 {
				fmt.Fprintf(w, "%-28s %s\n", i18n.S("Invoices due in 7 days", "7 günde vadesi gelen fatura")+":", output.Scalar(b.Upcoming))
			}
			fmt.Fprintf(w, "%-28s %d %s, %d %s\n", i18n.S("Prepaid services", "Ön ödemeli hizmetler")+":", b.Services["active"], i18n.S("running", "çalışıyor"), b.Services["suspended"], i18n.S("suspended", "askıda"))
			if b.Services["suspended"] > 0 {
				fmt.Fprintf(w, "%-28s %s\n", i18n.S("Needed to resume (1 h)", "Açılması için (1 saat)")+":", output.Scalar(b.SuspHourly))
			}
			fmt.Fprintf(w, "%-28s %s\n", i18n.S("Top up", "Bakiye yükle")+":", b.TopupURL)
			return nil
		},
	}
	g.AddCommand(bal)

	/* -------------------------------------------------------------- usage */
	var from, to, status, csvFile string
	var service int
	var daily, monthly, all bool
	usage := &cobra.Command{
		Use:   "usage",
		Short: i18n.S("Hour-by-hour usage (or daily/monthly totals)", "Saat saat kullanım (ya da günlük/aylık toplam)"),
		Example: "  gyro billing usage --from 2026-11-01 --to 2026-11-30\n  gyro billing usage --service 17501 --daily\n" +
			"  gyro billing usage --monthly -o json\n  gyro billing usage --csv kullanim.csv",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			q := url.Values{}
			for k, v := range map[string]string{"from": from, "to": to, "status": status} {
				if v != "" {
					q.Set(k, v)
				}
			}
			if service > 0 {
				q.Set("service_id", strconv.Itoa(service))
			}
			if csvFile != "" {
				q.Set("format", "csv")
				q.Set("lang", i18n.S("en", "tr"))
				resp, err := get("/cloud/billing/usage", q)
				if err != nil {
					return err
				}
				return save(resp.Data, csvFile)
			}
			gran := "hour"
			if daily {
				gran = "day"
			}
			if monthly {
				gran = "month"
			}
			q.Set("granularity", gran)
			q.Set("limit", "100")
			var rows []any
			var head map[string]any
			for page := 1; ; page++ {
				q.Set("page", strconv.Itoa(page))
				resp, err := get("/cloud/billing/usage", q)
				if err != nil {
					return err
				}
				var d map[string]any
				if err := json.Unmarshal(resp.Data, &d); err != nil {
					return err
				}
				if head == nil {
					head = d
				}
				if items, ok := d["items"].([]any); ok {
					rows = append(rows, items...)
				}
				if !all || resp.Meta == nil || resp.Meta.NextPage <= page {
					break
				}
			}
			if rows == nil {
				rows = []any{}
			}
			if a.Machine() || a.Opt.Query != "" {
				head["items"] = rows
				b, _ := json.Marshal(head)
				return a.PrintRaw(b, nil, "")
			}
			var cols []output.Column
			if gran == "hour" {
				cols = []output.Column{{Header: "hour", Path: "hour"}, {Header: "service", Path: "service_id"}, {Header: "item", Path: "description"},
					{Header: "status", Path: "status"}, {Header: "power", Path: "power"}, {Header: "eur_h", Path: "unit_price_eur"}, {Header: "rate", Path: "eur_rate"}, {Header: "amount", Path: "amount"}}
			} else {
				cols = []output.Column{{Header: "period", Path: "period"}, {Header: "hours", Path: "hours"}, {Header: "amount", Path: "amount"}, {Header: "amount_eur", Path: "amount_eur"}}
			}
			if err := a.Printer().Print(rows, cols, ""); err != nil {
				return err
			}
			if t, ok := head["totals"].([]any); ok && !a.Opt.Quiet {
				for _, x := range t {
					fmt.Fprintf(a.Out, "%s: %s\n", i18n.S("Total", "Toplam"), output.Scalar(x))
				}
			}
			if gran == "hour" && !all && len(rows) == 100 {
				a.Info("%s", i18n.S("First 100 hours shown; use --all, --daily or --csv.", "İlk 100 saat gösterildi; --all, --daily ya da --csv kullanın."))
			}
			return nil
		},
	}
	f := usage.Flags()
	f.StringVar(&from, "from", "", i18n.S("first day YYYY-MM-DD (default: 1st of this month)", "ilk gün YYYY-AA-GG (varsayılan: bu ayın 1'i)"))
	f.StringVar(&to, "to", "", i18n.S("last day YYYY-MM-DD (default: today)", "son gün YYYY-AA-GG (varsayılan: bugün)"))
	f.IntVar(&service, "service", 0, i18n.S("only this service id", "yalnız bu hizmet no"))
	f.StringVar(&status, "status", "", "ucretli|tavan|kismi|odenmedi|ucretsiz")
	f.BoolVar(&daily, "daily", false, i18n.S("daily totals", "günlük toplamlar"))
	f.BoolVar(&monthly, "monthly", false, i18n.S("monthly totals", "aylık toplamlar"))
	f.BoolVar(&all, "all", false, i18n.S("fetch every page", "tüm sayfaları getir"))
	f.StringVar(&csvFile, "csv", "", i18n.S("save the hours as CSV (file, '.' = server name, '-' = stdout)", "saatleri CSV kaydet (dosya, '.' = sunucunun adı, '-' = ekran)"))
	g.AddCommand(usage)

	/* -------------------------------------------------------------- estimate */
	var plan, region string
	var count, productID int
	var items []string
	est := &cobra.Command{
		Use:   "estimate",
		Short: i18n.S("Hourly, daily and monthly (capped) cost of plans; how long your balance lasts", "Planların saatlik, günlük ve aylık (tavanlı) maliyeti; bakiyenizin ne kadar yeteceği"),
		Example: "  gyro billing estimate --plan standart-2-4 --region tr-ist-1 --count 3\n" +
			"  gyro billing estimate --item standart-2-4:2 --item standart-4-8:1",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			var list []map[string]any
			for _, it := range items {
				p, n := it, 1
				if i := strings.LastIndex(it, ":"); i > 0 {
					p = it[:i]
					if v, err := strconv.Atoi(it[i+1:]); err == nil {
						n = v
					}
				}
				list = append(list, map[string]any{"plan": p, "region": region, "count": n})
			}
			if plan != "" || productID > 0 {
				m := map[string]any{"region": region, "count": count}
				if plan != "" {
					m["plan"] = plan
				}
				if productID > 0 {
					m["product_id"] = productID
				}
				list = append(list, m)
			}
			if len(list) == 0 {
				return fmt.Errorf("%s", i18n.S("give --plan or --item", "--plan ya da --item verin"))
			}
			ctx, cancel := a.Ctx()
			defer cancel()
			c, err := a.Client()
			if err != nil {
				return err
			}
			resp, err := c.Do(ctx, api.Request{Method: "POST", Path: "/cloud/billing/estimate", Body: map[string]any{"items": list}})
			if err != nil {
				return err
			}
			if a.Machine() || a.Opt.Query != "" {
				return a.PrintRaw(resp.Data, nil, "")
			}
			var e map[string]any
			_ = json.Unmarshal(resp.Data, &e)
			rows, _ := e["items"].([]any)
			cols := []output.Column{{Header: "plan", Path: "plan"}, {Header: "count", Path: "count"}, {Header: "hourly", Path: "hourly"}, {Header: "daily", Path: "daily"}, {Header: "monthly_cap", Path: "monthly_cap"}}
			if err := a.Printer().Print(rows, cols, ""); err != nil {
				return err
			}
			fmt.Fprintf(a.Out, "%s: %s · %s: %s · %s: %s\n", i18n.S("Hourly", "Saatlik"), output.Scalar(e["hourly"]), i18n.S("Daily", "Günlük"), output.Scalar(e["daily"]),
				i18n.S("Monthly cap", "Aylık tavan"), output.Scalar(e["monthly_cap"]))
			if hc, ok := e["hours_covered"].(float64); ok {
				fmt.Fprintf(a.Out, "%s: %s · %s %.1f %s\n", i18n.S("Balance", "Bakiye"), output.Scalar(e["balance"]), i18n.S("lasts about", "yaklaşık"), hc, i18n.S("hours (with what already runs)", "saat yeter (çalışanlarla birlikte)"))
			}
			if ok, _ := e["can_order"].(bool); !ok {
				fmt.Fprintf(a.Out, "%s %s\n", i18n.S("To order, the balance must cover the first 24 hours:", "Sipariş için bakiye ilk 24 saati karşılamalı:"), output.Scalar(e["order_requires"]))
			}
			return nil
		},
	}
	ef := est.Flags()
	ef.StringVar(&plan, "plan", "", i18n.S("plan name, e.g. standart-2-4", "plan adı, ör. standart-2-4"))
	ef.IntVar(&productID, "product-id", 0, i18n.S("or the product id", "ya da ürün no"))
	ef.StringVar(&region, "region", "tr-ist-1", i18n.S("region (prepaid: tr-ist-1)", "bölge (ön ödemeli: tr-ist-1)"))
	ef.IntVar(&count, "count", 1, i18n.S("how many", "adet"))
	ef.StringArrayVar(&items, "item", nil, i18n.S("plan:count, repeatable", "plan:adet, tekrarlanabilir"))
	g.AddCommand(est)

	/* -------------------------------------------------------------- statement */
	var pdfFile, stCSV string
	st := &cobra.Command{
		Use:     "statement YYYY-MM",
		Short:   i18n.S("Monthly usage statement (not an invoice); --pdf / --csv to download", "Aylık kullanım dökümü (fatura değildir); indirmek için --pdf / --csv"),
		Example: "  gyro billing statement 2026-11\n  gyro billing statement 2026-11 --pdf .\n  gyro billing statement 2026-11 --csv kasim.csv",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !monthRe.MatchString(args[0]) {
				return fmt.Errorf("%s", i18n.S("month must be YYYY-MM", "ay YYYY-AA olmalı"))
			}
			q := url.Values{}
			q.Set("lang", i18n.S("en", "tr"))
			file := ""
			switch {
			case pdfFile != "":
				q.Set("format", "pdf")
				file = pdfFile
			case stCSV != "":
				q.Set("format", "csv")
				file = stCSV
			}
			resp, err := get("/cloud/billing/statement/"+args[0], q)
			if err != nil {
				return err
			}
			if file != "" {
				return save(resp.Data, file)
			}
			if a.Machine() || a.Opt.Query != "" {
				return a.PrintRaw(resp.Data, nil, "")
			}
			var s map[string]any
			_ = json.Unmarshal(resp.Data, &s)
			rows, _ := s["by_service"].([]any)
			cols := []output.Column{{Header: "service", Path: "service_id"}, {Header: "description", Path: "description"}, {Header: "hours", Path: "hours"}, {Header: "amount", Path: "amount"}}
			if err := a.Printer().Print(rows, cols, ""); err != nil {
				return err
			}
			if t, ok := s["totals"].([]any); ok {
				for _, x := range t {
					fmt.Fprintf(a.Out, "%s: %s\n", i18n.S("Total", "Toplam"), output.Scalar(x))
				}
			}
			fmt.Fprintln(a.Out, i18n.S("Usage statement, not an invoice.", "Kullanım dökümüdür, fatura değildir."))
			return nil
		},
	}
	st.Flags().StringVar(&pdfFile, "pdf", "", i18n.S("save as PDF (file, '.' = server name)", "PDF kaydet (dosya, '.' = sunucunun adı)"))
	st.Flags().StringVar(&stCSV, "csv", "", i18n.S("save as CSV (file, '.' = server name)", "CSV kaydet (dosya, '.' = sunucunun adı)"))
	g.AddCommand(st)
}
