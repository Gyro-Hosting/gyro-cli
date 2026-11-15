package cli

// GYR-19 (API 2.33.0): container logs over the whole environment with filters, paging and live tail.
//   gyro container logs <ENV> [SERVICE] [--since 1h|--start/--end] [--filter t] [--exclude t] [--level error] [--where 'status>=500']... [-f]
// The platform builds every query (no raw query language). Package, rules and statistics come from the spec:
//   gyro container log-usage | log-plans | log-retention set | log-stats | log-rules list|set|test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
)

const (
	opSearchContainerLogs = "searchContainerLogs"
	opServiceLogsLegacy   = "typed_get_containers__id__services__name__logs"
)

type logPage struct {
	Lines []struct {
		Time     string `json:"time"`
		Instance string `json:"instance"`
		Text     string `json:"text"`
	} `json:"lines"`
	Next   string `json:"next_cursor"`
	Tail   string `json:"tail_cursor"`
	Cursor string `json:"cursor"`
}

func addContainerLogs(a *App, root *cobra.Command) {
	markHandled(opSearchContainerLogs, opServiceLogsLegacy)
	grp := ensureGroup(root, []string{"container"})
	var svc, since, start, end, filter, exclude, level, parse string
	var where []string
	var limit int
	var follow, previous, all bool
	cmd := &cobra.Command{
		Use:   "logs <ENVIRONMENT_ID> [SERVICE] [flags]",
		Short: i18n.S("Search and follow the logs of an environment or one service", "Ortamın ya da bir servisin loglarında ara, canlı izle"),
		Long: i18n.S(`Shows what your containers wrote to stdout / stderr, oldest first. Without a service every service of the
environment is searched. Filters: --filter (line contains, case-insensitive), --exclude, --level (error|warn|info|debug)
and --where field conditions on JSON / logfmt lines (status>=500, method=POST, path~/api; repeat the flag for more).
--follow keeps printing new lines (polls every 2 s, Ctrl-C stops). --all pages back through the whole window
(up to 100,000 lines). How far back you can go depends on the log package (free: 7 days; gyro container log-usage).`,
			`Kapsayıcılarınızın stdout / stderr'e yazdıklarını eskiden yeniye gösterir. Servis verilmezse ortamın tüm servisleri
aranır. Filtreler: --filter (satır içerir, büyük/küçük harf duyarsız), --exclude, --level (error|warn|info|debug) ve
JSON / logfmt satırlarda --where alan koşulları (status>=500, method=POST, path~/api; birden fazlası için bayrağı tekrarlayın).
--follow yeni satırları basmayı sürdürür (2 sn'de bir, Ctrl-C durdurur). --all tüm aralığı geriye doğru sayfalar
(en çok 100.000 satır). Ne kadar geriye gidebileceğiniz log paketinize bağlıdır (ücretsiz: 7 gün; gyro container log-usage).`),
		Example: "  gyro container logs 17000 api --since 30m\n" +
			"  gyro container logs 17000 --level error --since 24h\n" +
			"  gyro container logs 17000 api --where 'status>=500' --where 'path~/api' --since 2h\n" +
			"  gyro container logs 17000 api -f",
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			env := args[0]
			if !isNumeric(env) {
				return &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("the environment id is a number (gyro service list)", "ortam numarası sayı olmalı (gyro service list)"))}
			}
			if len(args) == 2 {
				svc = args[1]
			}
			q := url.Values{}
			set := func(k, v string) {
				if v = strings.TrimSpace(v); v != "" {
					q.Set(k, v)
				}
			}
			set("service", svc)
			set("filter", filter)
			set("exclude", exclude)
			set("level", level)
			set("parse", parse)
			if len(where) > 0 {
				q.Set("where", strings.Join(where, ";"))
			}
			if start != "" || end != "" {
				for k, v := range map[string]string{"start": start, "end": end} {
					if v == "" {
						continue
					}
					t, err := parseWhen(v)
					if err != nil {
						return &ExitErr{Code: ExitUsage, Err: fmt.Errorf(i18n.S("--%s: use RFC3339 (2026-11-15T10:00:00+03:00) or a duration ago (2h)", "--%s: RFC3339 (2026-11-15T10:00:00+03:00) ya da geçen süre (2h) yazın"), k)}
					}
					q.Set(k, t.UTC().Format(time.RFC3339))
				}
			} else {
				d := time.Hour
				if follow && !cmd.Flags().Changed("since") {
					d = 15 * time.Minute
				}
				if since != "" {
					v, err := parseDur(since)
					if err != nil {
						return &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("--since: e.g. 30m, 6h, 7d", "--since: ör. 30m, 6h, 7d"))}
					}
					d = v
				}
				q.Set("since_minutes", strconv.Itoa(int(d.Minutes())))
			}
			if limit > 0 {
				q.Set("limit", strconv.Itoa(limit))
			}
			ctx, cancel := a.Ctx()
			defer cancel()
			c, err := a.Client()
			if err != nil {
				return err
			}
			path := "/cloud/containers/" + env + "/logs"
			if previous {
				if svc == "" {
					return &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("--previous needs a service", "--previous için servis gerekir"))}
				}
				path = "/cloud/containers/" + env + "/services/" + url.PathEscape(svc) + "/logs"
				q.Del("service")
				q.Set("previous", "1")
			}
			get := func(qq url.Values) (*logPage, json.RawMessage, error) {
				r, err := c.Do(ctx, api.Request{Method: "GET", Path: path, Query: qq})
				if err != nil {
					return nil, nil, err
				}
				var p logPage
				if err := json.Unmarshal(r.Data, &p); err != nil {
					return nil, nil, err
				}
				return &p, r.Data, nil
			}
			p, raw, err := get(q)
			if err != nil {
				return err
			}
			if all && !previous {
				for n := 0; p.Next != "" && len(p.Lines) < 100000 && n < 50; n++ {
					q2 := cloneQ(q)
					q2.Set("before", p.Next)
					older, _, err := get(q2)
					if err != nil {
						return err
					}
					p.Lines = append(older.Lines, p.Lines...)
					p.Next = older.Next
					if len(older.Lines) == 0 {
						break
					}
				}
				b, _ := json.Marshal(p)
				raw = b
			}
			if a.Machine() && !follow {
				return a.PrintRaw(raw, nil, "")
			}
			show := func(pg *logPage) {
				for _, l := range pg.Lines {
					if a.Machine() {
						b, _ := json.Marshal(l)
						fmt.Fprintln(a.Out, string(b))
						continue
					}
					fmt.Fprintf(a.Out, "%s %s %s\n", shortTime(l.Time), l.Instance, l.Text)
				}
			}
			show(p)
			if !follow {
				if p.Next != "" && !all {
					a.Info(i18n.S("(older lines exist in the window: add --all or narrow --since)", "(aralıkta daha eski satır var: --all ekleyin ya da --since'i daraltın)"))
				}
				return nil
			}
			if previous {
				return &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("--follow and --previous cannot be used together", "--follow ile --previous birlikte kullanılamaz"))}
			}
			cur := p.Tail
			for {
				select {
				case <-ctx.Done():
					return nil
				case <-time.After(2 * time.Second):
				}
				q2 := cloneQ(q)
				q2.Del("since_minutes")
				q2.Del("start")
				q2.Del("end")
				q2.Set("after", cur)
				np, _, err := get(q2)
				if err != nil {
					if ctx.Err() != nil {
						return nil
					}
					a.Info("%v", err)
					continue
				}
				show(np)
				if np.Cursor != "" {
					cur = np.Cursor
				}
			}
		},
	}
	f := cmd.Flags()
	f.StringVar(&svc, "service", "", i18n.S("only this service (same as the second argument)", "yalnız bu servis (ikinci argümanla aynı)"))
	f.StringVar(&since, "since", "", i18n.S("window ending now: 30m, 6h, 7d (default 1h; 15m with --follow)", "şimdi biten aralık: 30m, 6h, 7d (varsayılan 1h; --follow ile 15m)"))
	f.StringVar(&start, "start", "", i18n.S("start time (RFC3339 or a duration ago, e.g. 3d); one query spans at most 31 days", "başlangıç (RFC3339 ya da geçen süre, ör. 3d); bir sorgu en çok 31 gün"))
	f.StringVar(&end, "end", "", i18n.S("end time (RFC3339 or a duration ago)", "bitiş (RFC3339 ya da geçen süre)"))
	f.StringVar(&filter, "filter", "", i18n.S("line contains this text (case-insensitive)", "satır bu metni içerir (büyük/küçük harf duyarsız)"))
	f.StringVar(&exclude, "exclude", "", i18n.S("line does not contain this text", "satır bu metni içermez"))
	f.StringVar(&level, "level", "", "error | warn | info | debug")
	f.StringArrayVar(&where, "where", nil, i18n.S("field condition on JSON / logfmt lines: status>=500, method=POST, path~/api (repeatable)", "JSON / logfmt satırlarda alan koşulu: status>=500, method=POST, path~/api (tekrarlanabilir)"))
	f.StringVar(&parse, "parse", "", "json | logfmt")
	f.IntVar(&limit, "limit", 0, i18n.S("lines per page (default 300, max 5000)", "sayfa başına satır (varsayılan 300, en çok 5000)"))
	f.BoolVarP(&follow, "follow", "f", false, i18n.S("keep printing new lines (Ctrl-C stops)", "yeni satırları basmayı sürdür (Ctrl-C durdurur)"))
	f.BoolVar(&all, "all", false, i18n.S("page back through the whole window (max 100,000 lines)", "tüm aralığı geriye doğru sayfala (en çok 100.000 satır)"))
	f.BoolVar(&previous, "previous", false, i18n.S("output of the last crashed replica of the service", "servisin son çöken kopyasının çıktısı"))
	grp.AddCommand(cmd)
}

func cloneQ(q url.Values) url.Values {
	o := url.Values{}
	for k, v := range q {
		o[k] = append([]string(nil), v...)
	}
	return o
}

// parseDur: Go durations plus "d" days (7d, 1d12h).
func parseDur(s string) (time.Duration, error) {
	s = strings.TrimSpace(s)
	var days float64
	if i := strings.Index(s, "d"); i > 0 {
		v, err := strconv.ParseFloat(s[:i], 64)
		if err != nil {
			return 0, err
		}
		days, s = v, s[i+1:]
	}
	var d time.Duration
	if s != "" {
		v, err := time.ParseDuration(s)
		if err != nil {
			return 0, err
		}
		d = v
	}
	d += time.Duration(days * 24 * float64(time.Hour))
	if d <= 0 {
		return 0, errors.New("empty")
	}
	return d, nil
}

func parseWhen(s string) (time.Time, error) {
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t, nil
	}
	d, err := parseDur(s)
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().Add(-d), nil
}

func shortTime(s string) string {
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return s
	}
	return t.Local().Format("2006-01-02 15:04:05.000")
}
