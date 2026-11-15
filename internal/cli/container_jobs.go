package cli

// GYR-17 (API 2.32.0): one-off jobs of a container environment.
//   gyro container run <ENV> --service api [--wait|--follow] [--time-limit 30m] [--cpu 0.5] [--memory 512] -- npm run migrate
//   gyro container jobs <ENV> [--service api]
// get / logs / cancel come from the spec (gyro container job get|logs|cancel). The command after "--" is sent as a list:
// no shell is involved unless you write one (sh -c "a && b"), so quotes, <, > and | reach the program unchanged.

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
	"github.com/Gyro-Hosting/gyro-cli/internal/output"
)

const (
	opRunContainerJob   = "runContainerJob"
	opListContainerJobs = "listContainerJobs"
)

type jobDoc struct {
	ID       string   `json:"id"`
	Service  string   `json:"service"`
	State    string   `json:"state"`
	Exit     *int     `json:"exit_code"`
	Error    string   `json:"error"`
	Detail   string   `json:"detail"`
	Command  []string `json:"command"`
	Image    string   `json:"image"`
	Pinned   bool     `json:"pinned"`
	Duration int      `json:"duration_seconds"`
}

func (j *jobDoc) terminal() bool {
	switch j.State {
	case "succeeded", "failed", "timeout", "cancelled":
		return true
	}
	return false
}

func jobColumns() []output.Column {
	return []output.Column{
		{Header: "id", Path: "id"},
		{Header: i18n.S("service", "servis"), Path: "service"},
		{Header: "state", Path: "state"},
		{Header: i18n.S("exit", "çıkış"), Path: "exit_code"},
		{Header: i18n.S("seconds", "saniye"), Path: "duration_seconds"},
		{Header: i18n.S("command", "komut"), Path: "cmd"},
		{Header: i18n.S("started", "başlangıç"), Path: "started"},
	}
}

func jobCmdText(l []string) string {
	if len(l) == 0 {
		return i18n.S("(image command)", "(imajın komutu)")
	}
	parts := make([]string, len(l))
	for i, a := range l {
		if strings.ContainsAny(a, " \t\"'") {
			a = `"` + strings.ReplaceAll(a, `"`, `\"`) + `"`
		}
		parts[i] = a
	}
	s := strings.Join(parts, " ")
	if len(s) > 60 {
		s = s[:57] + "..."
	}
	return s
}

// jobFrom reads a job from an answer: {"job": {...}} (run) or the job fields at the top level (get).
func jobFrom(raw json.RawMessage) (*jobDoc, error) {
	var w struct {
		Job *jobDoc `json:"job"`
	}
	if err := json.Unmarshal(raw, &w); err == nil && w.Job != nil && w.Job.ID != "" {
		return w.Job, nil
	}
	var j jobDoc
	if err := json.Unmarshal(raw, &j); err != nil || j.ID == "" {
		return nil, errors.New(i18n.S("unexpected answer from the API", "API'den beklenmeyen yanıt"))
	}
	return &j, nil
}

// jobExit maps the final state to the exit code of the CLI: the job's own exit code, 124 on timeout (like `timeout`), 130 when stopped.
func jobExit(j *jobDoc) error {
	switch j.State {
	case "succeeded":
		return nil
	case "timeout":
		return &ExitErr{Code: 124, Err: fmt.Errorf(i18n.S("job %s stopped: time limit reached", "görev %s durduruldu: zaman sınırı doldu"), j.ID)}
	case "cancelled":
		return &ExitErr{Code: 130, Err: fmt.Errorf(i18n.S("job %s was stopped", "görev %s durduruldu"), j.ID)}
	}
	code := 1
	if j.Exit != nil && *j.Exit > 0 && *j.Exit < 256 {
		code = *j.Exit
	}
	why := j.Error
	if j.Detail != "" {
		why += " (" + j.Detail + ")"
	}
	return &ExitErr{Code: code, Err: fmt.Errorf(i18n.S("job %s failed: %s", "görev %s başarısız: %s"), j.ID, why)}
}

func addContainerJobs(a *App, root *cobra.Command) {
	markHandled(opRunContainerJob)
	grp := ensureGroup(root, []string{"container"})

	var svc, cpu string
	var memory int
	var limit time.Duration
	var wait, follow bool
	run := &cobra.Command{
		Use:   "run <ENVIRONMENT_ID> --service <NAME> [flags] [-- COMMAND [ARG...]]",
		Short: i18n.S("Run a one-off job (migration, seed) with a service's image, variables and secrets", "Servisin imajı, değişkenleri ve gizli değerleriyle tek seferlik görev çalıştır (migration, seed)"),
		Long: i18n.S(`Runs a command once with the image (pinned to the digest of the running version), variables and secrets of an
existing service, then stops. Everything after "--" is the command; it replaces the image ENTRYPOINT and is passed as
a list (no shell: write sh -c "a && b" for shell features). Without a command the image's own command runs.
--follow streams the output; --wait waits silently. With either, the CLI exits with the job's exit code
(124 = time limit reached, 130 = stopped). Free on node plans (uses your nodes); on Serverless the job is billed per
vCPU-second / GB-second like a task (billing line "<service>-job"). Parallel jobs per environment are limited
(2-3 by plan); see "gyro limits".`,
			`Mevcut bir servisin imajı (çalışan sürümün digest'ine sabitlenir), değişkenleri ve gizli değerleriyle bir komutu
bir kez çalıştırır, sonra durur. "--" sonrası komuttur; imajın ENTRYPOINT'inin yerine geçer ve liste olarak gider
(kabuk yok: kabuk özellikleri için sh -c "a && b" yazın). Komut yoksa imajın kendi komutu çalışır.
--follow çıktıyı akıtır; --wait sessizce bekler. İkisinde de CLI görevin çıkış koduyla biter (124 = zaman sınırı
doldu, 130 = durduruldu). Düğümlü paketlerde ücretsizdir (düğümlerinizi kullanır); Sunucusuz'da görev vCPU-saniye /
GB-saniye ile görev (task) gibi ücretlenir (fatura satırı "<servis>-job"). Ortam başına eşzamanlı görev sayısı
pakete göre 2-3; bkz. "gyro limits".`),
		Example: "  gyro container run 17000 --service api --follow -- npm run migrate\n" +
			"  gyro container run 17000 --service web --wait --time-limit 10m -- sh -c \"php artisan migrate --force && php artisan db:seed --force\"\n" +
			"  gyro container run 17000 --service api --cpu 1 --memory 1024 -- python manage.py migrate",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			env := args[0]
			if !isNumeric(env) {
				return &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("the environment id is a number (gyro container list)", "ortam numarası sayı olmalı (gyro container list)"))}
			}
			if svc == "" {
				return &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("--service is required", "--service gerekli"))}
			}
			body := map[string]any{}
			if dash := cmd.ArgsLenAtDash(); dash >= 0 && len(args) > dash {
				body["command"] = args[dash:]
			} else if len(args) > 1 {
				return &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("put the command after \"--\": gyro container run 17000 --service api -- npm run migrate", "komutu \"--\" sonrasına yazın: gyro container run 17000 --service api -- npm run migrate"))}
			}
			if limit > 0 {
				body["timeout_seconds"] = int(limit.Seconds())
			}
			if cpu != "" {
				var f float64
				if _, err := fmt.Sscanf(cpu, "%g", &f); err != nil || f <= 0 {
					return &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("--cpu: vCPU as a number, e.g. 0.5", "--cpu: vCPU sayı olarak, ör. 0.5"))}
				}
				body["cpu_m"] = int(f*1000 + 0.5)
			}
			if memory > 0 {
				body["memory_mi"] = memory
			}
			ctx, cancel := a.Ctx()
			defer cancel()
			c, err := a.Client()
			if err != nil {
				return err
			}
			resp, err := c.Do(ctx, api.Request{Method: "POST", Path: "/cloud/containers/" + env + "/services/" + url.PathEscape(svc) + "/run", Body: body, IdempotencyKey: api.NewID()})
			if err != nil {
				return err
			}
			j, err := jobFrom(resp.Data)
			if err != nil {
				return err
			}
			if !wait && !follow {
				return a.PrintRaw(resp.Data, nil, "job.id")
			}
			a.Info(i18n.S("job %s started (%s)", "görev %s başladı (%s)"), j.ID, map[bool]string{true: i18n.S("pinned image", "sabit imaj"), false: i18n.S("image by tag", "etiketli imaj")}[j.Pinned])
			printed := ""
			for {
				if follow {
					q := url.Values{"tail": {"2000"}}
					if lr, err := c.Do(ctx, api.Request{Method: "GET", Path: "/cloud/containers/" + env + "/jobs/" + j.ID + "/logs", Query: q}); err == nil {
						var l struct {
							Log string `json:"log"`
						}
						_ = json.Unmarshal(lr.Data, &l)
						switch {
						case strings.HasPrefix(l.Log, printed):
							fmt.Fprint(a.Out, l.Log[len(printed):])
						case l.Log != printed: // tail window moved: print the lines not seen yet
							seen := map[string]bool{}
							for _, ln := range strings.Split(printed, "\n") {
								seen[ln] = true
							}
							for _, ln := range strings.Split(strings.TrimRight(l.Log, "\n"), "\n") {
								if !seen[ln] {
									fmt.Fprintln(a.Out, ln)
								}
							}
						}
						printed = l.Log
					}
				}
				if j.terminal() {
					break
				}
				select {
				case <-ctx.Done():
					return &ExitErr{Code: ExitTimeout, Err: fmt.Errorf(i18n.S("stopped waiting; job %s keeps running (gyro container job get %s %s)", "bekleme bitti; görev %s sürüyor (gyro container job get %s %s)"), j.ID, env, j.ID)}
				case <-time.After(2 * time.Second):
				}
				gr, err := c.Do(ctx, api.Request{Method: "GET", Path: "/cloud/containers/" + env + "/jobs/" + j.ID})
				if err != nil {
					return err
				}
				if nj, err := jobFrom(gr.Data); err == nil {
					j = nj
				}
				if j.terminal() && follow { // one more log read: the final lines after the end
					continue
				}
			}
			code := "-"
			if j.Exit != nil {
				code = fmt.Sprint(*j.Exit)
			}
			a.Info(i18n.S("job %s %s (exit code %s, %d s)", "görev %s %s (çıkış kodu %s, %d sn)"), j.ID, j.State, code, j.Duration)
			if a.Machine() {
				b, _ := json.Marshal(j)
				_ = a.PrintRaw(b, nil, "id")
			}
			return jobExit(j)
		},
	}
	run.Flags().StringVar(&svc, "service", "", i18n.S("service whose image, variables and secrets are used (required)", "imajı, değişkenleri ve gizli değerleri kullanılan servis (zorunlu)"))
	run.Flags().DurationVar(&limit, "time-limit", 0, i18n.S("stop the job after this time (10s-6h, default 30m)", "görevi bu süre sonunda durdur (10s-6h, varsayılan 30m)"))
	run.Flags().StringVar(&cpu, "cpu", "", i18n.S("vCPU for the job (default: same as the service)", "görevin vCPU'su (varsayılan: servisle aynı)"))
	run.Flags().IntVar(&memory, "memory", 0, i18n.S("memory in MB for the job (default: same as the service)", "görevin belleği MB (varsayılan: servisle aynı)"))
	run.Flags().BoolVar(&wait, "wait", false, i18n.S("wait until the job ends; exit with its exit code", "görev bitene kadar bekle; onun çıkış koduyla bit"))
	run.Flags().BoolVar(&follow, "follow", false, i18n.S("stream the output until the job ends (implies --wait)", "görev bitene kadar çıktıyı akıt (--wait içerir)"))
	grp.AddCommand(run)

	var lsvc string
	jobs := &cobra.Command{
		Use:     "jobs <ENVIRONMENT_ID>",
		Short:   i18n.S("List one-off jobs (last 50)", "Tek seferlik görevleri listele (son 50)"),
		Example: "  gyro container jobs 17000\n  gyro container jobs 17000 --service api -o json",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := a.Ctx()
			defer cancel()
			c, err := a.Client()
			if err != nil {
				return err
			}
			q := url.Values{}
			if lsvc != "" {
				q.Set("service", lsvc)
			}
			resp, err := c.Do(ctx, api.Request{Method: "GET", Path: "/cloud/containers/" + args[0] + "/jobs", Query: q})
			if err != nil {
				return err
			}
			p := a.Printer()
			if a.Opt.Query != "" || (p.Format != "table" && p.Format != "text" && !p.Quiet) {
				return a.PrintRaw(resp.Data, nil, "")
			}
			var d struct {
				Jobs  []map[string]any `json:"jobs"`
				Limit int              `json:"limit"`
			}
			if err := json.Unmarshal(resp.Data, &d); err != nil {
				return a.PrintRaw(resp.Data, nil, "")
			}
			rows := make([]any, 0, len(d.Jobs))
			for _, j := range d.Jobs {
				var l []string
				for _, x := range anyList(j["command"]) {
					l = append(l, fmt.Sprint(x))
				}
				j["cmd"] = jobCmdText(l)
				rows = append(rows, j)
			}
			if err := p.Print(rows, jobColumns(), "id"); err != nil {
				return err
			}
			if len(rows) == 0 {
				a.Info("%s", i18n.S("No jobs yet.", "Henüz görev yok."))
			}
			if d.Limit > 0 && !p.Quiet {
				a.Info(i18n.S("At most %d jobs run at the same time in this environment.", "Bu ortamda aynı anda en çok %d görev çalışır."), d.Limit)
			}
			return nil
		},
	}
	jobs.Flags().StringVar(&lsvc, "service", "", i18n.S("only the jobs of this service", "yalnız bu servisin görevleri"))
	grp.AddCommand(jobs)
}

func anyList(v any) []any {
	l, _ := v.([]any)
	return l
}
