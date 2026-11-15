package cli

// GYR-12 (API 2.36.0): scheduled jobs of a container environment.
//   gyro container schedule create <ENV> <NAME> --cron "0 3 * * *" --service api [--timezone UTC] [--retries 2] -- npm run report
//   gyro container schedule create <ENV> <NAME> --every 15m --url https://example.com/cron --header "Authorization: Bearer ..."
//   gyro container schedule update <ENV> <NAME|ID> [same flags; only the given ones change]
//   gyro container schedule list <ENV>
// get / delete / pause / resume / run-now / preview come from the spec.

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
	"github.com/Gyro-Hosting/gyro-cli/internal/output"
)

const (
	opCreateContainerSchedule = "createContainerSchedule"
	opUpdateContainerSchedule = "updateContainerSchedule"
	opListContainerSchedules  = "listContainerSchedules"
)

type schedFlags struct {
	cron, at, tz, svc, target, url, method, body, desc, cpu string
	every, limit, httpTO                                    time.Duration
	headers                                                 []string
	retries, memory                                         int
	paused, noHeaders                                       bool
}

func (f *schedFlags) bind(c *cobra.Command, update bool) {
	fs := c.Flags()
	fs.StringVar(&f.cron, "cron", "", i18n.S(`cron expression: minute hour day month weekday, e.g. "0 3 * * *" (daily 03:00), "*/15 * * * *", "0 9 * * 1-5"; @daily / @hourly work too`,
		`cron ifadesi: dakika saat gün ay haftanın-günü, ör. "0 3 * * *" (her gün 03:00), "*/15 * * * *", "0 9 * * 1-5"; @daily / @hourly da olur`))
	fs.DurationVar(&f.every, "every", 0, i18n.S("run every N minutes instead of a cron expression, e.g. 15m, 2h (at least 1m; Serverless 5m)", "cron yerine her N dakikada, ör. 15m, 2h (en az 1m; Sunucusuz 5m)"))
	fs.StringVar(&f.at, "at", "", i18n.S("run once at this time: 2026-12-01T03:00 (in --timezone) or RFC 3339", "bu zamanda bir kez çalıştır: 2026-12-01T03:00 (--timezone'a göre) ya da RFC 3339"))
	fs.StringVar(&f.tz, "timezone", "", i18n.S("IANA time zone of the schedule (default Europe/Istanbul), e.g. UTC", "zamanlamanın IANA saat dilimi (varsayılan Europe/Istanbul), ör. UTC"))
	fs.StringVar(&f.svc, "service", "", i18n.S("run the command after \"--\" with this service's image, variables and secrets", "\"--\" sonrasındaki komutu bu servisin imajı, değişkenleri ve gizli değerleriyle çalıştır"))
	fs.StringVar(&f.url, "url", "", i18n.S("call this URL instead (http:// or https://)", "bunun yerine bu adresi çağır (http:// ya da https://)"))
	fs.StringVar(&f.method, "method", "", i18n.S("HTTP method (default GET)", "HTTP yöntemi (varsayılan GET)"))
	fs.StringArrayVar(&f.headers, "header", nil, i18n.S(`HTTP header "Name: value" (repeatable; values are kept secret)`, `HTTP başlığı "Ad: değer" (tekrarlanabilir; değerler gizli saklanır)`))
	fs.StringVar(&f.body, "body", "", i18n.S("HTTP body (text, or @file)", "HTTP gövdesi (metin ya da @dosya)"))
	fs.DurationVar(&f.httpTO, "http-timeout", 0, i18n.S("time limit of the HTTP call (1s-120s, default 30s)", "HTTP çağrısının zaman sınırı (1s-120s, varsayılan 30s)"))
	fs.DurationVar(&f.limit, "time-limit", 0, i18n.S("stop a run after this time (10s-6h, default 30m)", "bir çalışmayı bu süre sonunda durdur (10s-6h, varsayılan 30m)"))
	fs.StringVar(&f.cpu, "cpu", "", i18n.S("vCPU of a run (default: same as the service)", "çalışmanın vCPU'su (varsayılan: servisle aynı)"))
	fs.IntVar(&f.memory, "memory", 0, i18n.S("memory in MB of a run (default: same as the service)", "çalışmanın belleği MB (varsayılan: servisle aynı)"))
	fs.IntVar(&f.retries, "retries", 0, i18n.S("retries after a failed run, 0-3 (30 s × attempt apart)", "başarısız çalışmadan sonra tekrar deneme, 0-3 (30 sn × deneme arayla)"))
	fs.StringVar(&f.desc, "description", "", i18n.S("description", "açıklama"))
	if update {
		fs.BoolVar(&f.noHeaders, "clear-headers", false, i18n.S("remove the stored HTTP headers", "saklanan HTTP başlıklarını sil"))
	} else {
		fs.BoolVar(&f.paused, "paused", false, i18n.S("create it paused", "duraklatılmış olarak oluştur"))
	}
}

// body builds the request; on update only the flags that were set go out.
func (f *schedFlags) request(c *cobra.Command, args []string, update bool) (map[string]any, error) {
	ch := func(n string) bool { return !update || c.Flags().Changed(n) }
	b := map[string]any{}
	n := 0
	for _, x := range []string{"cron", "every", "at"} {
		if c.Flags().Changed(x) {
			n++
		}
	}
	if n > 1 {
		return nil, errors.New(i18n.S("give only one of --cron, --every, --at", "--cron, --every, --at'tan yalnız birini verin"))
	}
	if !update && n == 0 {
		return nil, errors.New(i18n.S("when should it run? --cron \"0 3 * * *\", --every 15m or --at 2026-12-01T03:00", "ne zaman çalışsın? --cron \"0 3 * * *\", --every 15m ya da --at 2026-12-01T03:00"))
	}
	switch {
	case c.Flags().Changed("cron"):
		b["cron"] = f.cron
	case c.Flags().Changed("every"):
		if f.every < time.Minute || f.every%time.Minute != 0 {
			return nil, errors.New(i18n.S("--every: whole minutes, at least 1m", "--every: tam dakika, en az 1m"))
		}
		b["rate_minutes"] = int(f.every / time.Minute)
	case c.Flags().Changed("at"):
		b["at"] = f.at
	}
	if ch("timezone") && f.tz != "" {
		b["timezone"] = f.tz
	}
	if ch("description") && (update || f.desc != "") {
		b["description"] = f.desc
	}
	if ch("retries") && (update || f.retries != 0) {
		b["retries"] = f.retries
	}
	dash := c.ArgsLenAtDash()
	if f.svc != "" && f.url != "" {
		return nil, errors.New(i18n.S("give either --service (command) or --url (HTTP call)", "ya --service (komut) ya --url (HTTP çağrısı) verin"))
	}
	if f.url != "" || (update && (c.Flags().Changed("header") || c.Flags().Changed("body") || c.Flags().Changed("method") || c.Flags().Changed("http-timeout") || f.noHeaders)) {
		if f.url != "" {
			b["target"], b["url"] = "http", f.url
		}
		if f.method != "" {
			b["method"] = strings.ToUpper(f.method)
		}
		if len(f.headers) > 0 || f.noHeaders {
			h := map[string]string{}
			for _, x := range f.headers {
				k, v, ok := strings.Cut(x, ":")
				if !ok || strings.TrimSpace(k) == "" {
					return nil, fmt.Errorf(i18n.S("--header %q: write \"Name: value\"", "--header %q: \"Ad: değer\" biçiminde yazın"), x)
				}
				h[strings.TrimSpace(k)] = strings.TrimSpace(v)
			}
			b["headers"] = h
		}
		if c.Flags().Changed("body") {
			b["body"] = f.body // "@file" is read by the caller (App.ReadValue)
		}
		if f.httpTO > 0 {
			b["http_timeout_seconds"] = int(f.httpTO.Seconds())
		}
		if dash >= 0 && len(args) > dash {
			return nil, errors.New(i18n.S("an HTTP call takes no command after \"--\"", "HTTP çağrısı \"--\" sonrasına komut almaz"))
		}
		return b, nil
	}
	if f.svc != "" {
		b["target"], b["service"] = "service", f.svc
	} else if !update {
		return nil, errors.New(i18n.S("what should run? --service api -- <command>, or --url https://...", "ne çalışsın? --service api -- <komut> ya da --url https://..."))
	}
	if dash >= 0 && len(args) > dash {
		b["command"] = args[dash:]
	}
	if f.limit > 0 {
		b["timeout_seconds"] = int(f.limit.Seconds())
	}
	if f.cpu != "" {
		var x float64
		if _, err := fmt.Sscanf(f.cpu, "%g", &x); err != nil || x <= 0 {
			return nil, errors.New(i18n.S("--cpu: vCPU as a number, e.g. 0.5", "--cpu: vCPU sayı olarak, ör. 0.5"))
		}
		b["cpu_m"] = int(x*1000 + 0.5)
	}
	if f.memory > 0 {
		b["memory_mi"] = f.memory
	}
	return b, nil
}

func schedWhen(s map[string]any) string {
	switch fmt.Sprint(s["kind"]) {
	case "cron":
		return fmt.Sprint(s["cron"])
	case "rate":
		return fmt.Sprintf(i18n.S("every %v min", "her %v dk"), s["rate_minutes"])
	case "once":
		return i18n.S("once ", "tek sefer ") + fmt.Sprint(s["at"])
	}
	return ""
}

func schedTarget(s map[string]any) string {
	if fmt.Sprint(s["target"]) == "http" {
		return fmt.Sprint(s["method"]) + " " + fmt.Sprint(s["url"])
	}
	var l []string
	for _, x := range anyList(s["command"]) {
		l = append(l, fmt.Sprint(x))
	}
	return fmt.Sprint(s["service"]) + ": " + jobCmdText(l)
}

func addContainerSchedules(a *App, root *cobra.Command) {
	markHandled(opCreateContainerSchedule, opUpdateContainerSchedule, opListContainerSchedules)
	grp := ensureGroup(root, []string{"container", "schedule"})
	grp.Short = i18n.S("Scheduled jobs (cron) of a container environment", "Kapsayıcı ortamının zamanlanmış görevleri (cron)")

	do := func(cmd *cobra.Command, method, path string, body map[string]any) error {
		ctx, cancel := a.Ctx()
		defer cancel()
		c, err := a.Client()
		if err != nil {
			return err
		}
		if v, ok := body["body"].(string); ok && strings.HasPrefix(v, "@") {
			if body["body"], err = a.ReadValue(v); err != nil {
				return &ExitErr{Code: ExitUsage, Err: err}
			}
		}
		req := api.Request{Method: method, Path: path, Body: body}
		if method == "POST" {
			req.IdempotencyKey = api.NewID()
		}
		resp, err := c.Do(ctx, req)
		if err != nil {
			return err
		}
		var w struct {
			Schedule map[string]any `json:"schedule"`
		}
		if json.Unmarshal(resp.Data, &w) == nil && w.Schedule != nil && !a.Machine() && a.Opt.Query == "" {
			a.Info(i18n.S("schedule %v (%v): %s, next run %v", "zamanlama %v (%v): %s, sonraki çalışma %v"), w.Schedule["name"], w.Schedule["id"], schedWhen(w.Schedule), w.Schedule["next_run_local"])
			return nil
		}
		return a.PrintRaw(resp.Data, nil, "schedule.id")
	}

	var cf schedFlags
	create := &cobra.Command{
		Use:   "create <ENVIRONMENT_ID> <NAME> (--cron EXPR | --every DURATION | --at TIME) (--service NAME [-- COMMAND...] | --url URL) [flags]",
		Short: i18n.S("Create a scheduled job (cron, every N minutes or once)", "Zamanlanmış görev oluştur (cron, her N dakikada ya da tek sefer)"),
		Long: i18n.S(`Runs a command of a service (as a one-off job: the image of the running version, its variables and secrets) or
calls a URL on a schedule. Times follow --timezone (default Europe/Istanbul). A run never starts while the previous
run of the same schedule is still going. Failed runs can be retried (--retries 0-3); after the last attempt you get
a schedule_failed notification (and schedule_missed when a run could not start in time). Free on node plans; on
Serverless each run is billed like a task. Limits: shortest interval 1 minute (Serverless 5), schedules per
environment 5-10 by plan (see "gyro limits").`,
			`Bir servisin komutunu (tek seferlik görev olarak: çalışan sürümün imajı, değişkenleri ve gizli değerleri) ya da bir
adresi takvime göre çalıştırır. Saatler --timezone'a göre (varsayılan Europe/Istanbul). Aynı zamanlamanın önceki
çalışması sürerken yenisi başlamaz. Başarısız çalışma tekrar denenebilir (--retries 0-3); son denemeden sonra
schedule_failed bildirimi gelir (zamanında başlayamayan çalışma için schedule_missed). Düğümlü paketlerde ücretsiz;
Sunucusuz'da her çalışma görev gibi ücretlenir. Sınırlar: en kısa aralık 1 dakika (Sunucusuz 5), ortam başına
zamanlama pakete göre 5-10 (bkz. "gyro limits").`),
		Example: "  gyro container schedule create 17000 nightly-report --cron \"0 3 * * *\" --service api -- npm run report\n" +
			"  gyro container schedule create 17000 queue-clean --every 15m --service worker --retries 2 -- php artisan queue:prune-failed\n" +
			"  gyro container schedule create 17000 ping-app --every 5m --url https://example.com/cron --method POST --header \"Authorization: Bearer $TOKEN\"\n" +
			"  gyro container schedule create 17000 campaign --at 2026-12-01T09:00 --service api -- node send.js",
		Args: cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !isNumeric(args[0]) {
				return &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("the environment id is a number (gyro service list)", "ortam numarası sayı olmalı (gyro service list)"))}
			}
			if d := cmd.ArgsLenAtDash(); d >= 0 && d < 2 {
				return &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("give the environment id and a name before \"--\"", "\"--\" öncesinde ortam numarası ve ad verin"))}
			}
			b, err := cf.request(cmd, args, false)
			if err != nil {
				return &ExitErr{Code: ExitUsage, Err: err}
			}
			b["name"] = args[1]
			if cf.paused {
				b["paused"] = true
			}
			return do(cmd, "POST", "/cloud/containers/"+args[0]+"/schedules", b)
		},
	}
	cf.bind(create, false)
	grp.AddCommand(create)

	var uf schedFlags
	update := &cobra.Command{
		Use:   "update <ENVIRONMENT_ID> <NAME|ID> [flags] [-- COMMAND...]",
		Short: i18n.S("Change a scheduled job (only the given flags change)", "Zamanlanmış görevi değiştir (yalnız verilen seçenekler değişir)"),
		Example: "  gyro container schedule update 17000 nightly-report --cron \"30 2 * * *\" --timezone UTC\n" +
			"  gyro container schedule update 17000 nightly-report --service api -- npm run report:full\n" +
			"  gyro container schedule update 17000 ping-app --header \"Authorization: Bearer $NEW\"",
		Args: cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if d := cmd.ArgsLenAtDash(); d >= 0 && d < 2 {
				return &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("give the environment id and the schedule before \"--\"", "\"--\" öncesinde ortam numarası ve zamanlamayı verin"))}
			}
			b, err := uf.request(cmd, args, true)
			if err != nil {
				return &ExitErr{Code: ExitUsage, Err: err}
			}
			if d := cmd.ArgsLenAtDash(); d >= 0 && len(args) > d && uf.svc == "" {
				b["command"] = args[d:]
			}
			if len(b) == 0 {
				return &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("nothing to change", "değişecek bir şey yok"))}
			}
			return do(cmd, "PATCH", "/cloud/containers/"+args[0]+"/schedules/"+args[1], b)
		},
	}
	uf.bind(update, true)
	grp.AddCommand(update)

	list := &cobra.Command{
		Use:     "list <ENVIRONMENT_ID>",
		Short:   i18n.S("List scheduled jobs with their next and last run", "Zamanlanmış görevleri sonraki ve son çalışmalarıyla listele"),
		Example: "  gyro container schedule list 17000\n  gyro container schedule list 17000 -o json",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := a.Ctx()
			defer cancel()
			c, err := a.Client()
			if err != nil {
				return err
			}
			resp, err := c.Do(ctx, api.Request{Method: "GET", Path: "/cloud/containers/" + args[0] + "/schedules"})
			if err != nil {
				return err
			}
			p := a.Printer()
			if a.Opt.Query != "" || (p.Format != "table" && p.Format != "text" && !p.Quiet) {
				return a.PrintRaw(resp.Data, nil, "")
			}
			var d struct {
				Schedules []map[string]any `json:"schedules"`
				Limit     int              `json:"limit"`
			}
			if err := json.Unmarshal(resp.Data, &d); err != nil {
				return a.PrintRaw(resp.Data, nil, "")
			}
			rows := make([]any, 0, len(d.Schedules))
			for _, s := range d.Schedules {
				s["when"], s["what"] = schedWhen(s), schedTarget(s)
				if lr, ok := s["last_run"].(map[string]any); ok {
					s["last"] = lr["state"]
				}
				rows = append(rows, s)
			}
			cols := []output.Column{{Header: "name", Path: "name"}, {Header: "state", Path: "state"}, {Header: i18n.S("when", "zaman"), Path: "when"},
				{Header: "tz", Path: "timezone"}, {Header: i18n.S("next", "sonraki"), Path: "next_run_local"}, {Header: i18n.S("last", "son"), Path: "last"},
				{Header: i18n.S("target", "hedef"), Path: "what"}, {Header: "id", Path: "id"}}
			if err := p.Print(rows, cols, "id"); err != nil {
				return err
			}
			if len(rows) == 0 {
				a.Info("%s", i18n.S("No schedules yet.", "Henüz zamanlama yok."))
			}
			if d.Limit > 0 && !p.Quiet {
				a.Info(i18n.S("%d of %d schedules used in this environment.", "Bu ortamda %d / %d zamanlama kullanılıyor."), len(rows), d.Limit)
			}
			return nil
		},
	}
	grp.AddCommand(list)
}
