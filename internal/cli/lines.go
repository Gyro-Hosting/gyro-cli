package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/spf13/cobra"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
	"github.com/Gyro-Hosting/gyro-cli/internal/output"
	"github.com/Gyro-Hosting/gyro-cli/internal/spec"
)

// Product-line commands (mail, s3, monitor, license, app, vds, dedicated) sit on the generic service
// bridge of API v2: /cloud/services?product_line=..., /cloud/services/{id}/tools[/{tool}[/{action}]].
// Tool and action names are discovered at run time from GET /cloud/services/{id}/tools, so new module
// tools work without a CLI release; the nouns below only add friendly names.

type lineNoun struct {
	name    string
	aliases []string
	tools   []string // candidate tool names, first one present on the service wins
	short   [2]string
}

type productLine struct {
	name, line string
	aliases    []string
	short      [2]string
	nouns      []lineNoun
}

var productLines = []productLine{
	{name: "mail", line: "mail", aliases: []string{"email"}, short: [2]string{"Corporate mail: mailboxes, aliases, forwarding ...", "Kurumsal mail: kutular, takma adlar, yönlendirme ..."},
		nouns: []lineNoun{
			{"mailbox", []string{"mailboxes", "box"}, []string{"mailboxes", "mailbox"}, [2]string{"Mailboxes", "Posta kutuları"}},
			{"alias", []string{"aliases"}, []string{"aliases", "alias"}, [2]string{"Aliases (forwarders)", "Takma adlar (yönlendirmeler)"}},
			{"catch-all", nil, []string{"catch-all", "catchall"}, [2]string{"Catch-all address", "Tümünü yakala adresi"}},
			{"autoresponder", []string{"vacation"}, []string{"autoresponder", "autoresponders"}, [2]string{"Auto-responders", "Otomatik yanıtlar"}},
			{"migration", []string{"sync", "imapsync"}, []string{"sync-jobs", "migrations", "imap-sync"}, [2]string{"IMAP migration jobs", "IMAP taşıma işleri"}},
			{"dns", nil, []string{"dns-setup", "dns"}, [2]string{"Required DNS records and their status", "Gerekli DNS kayıtları ve durumları"}},
			{"device", []string{"devices"}, []string{"devices"}, [2]string{"Device setup / app passwords", "Cihaz kurulumu / uygulama şifreleri"}},
			{"spam", nil, []string{"spam-filter", "spam"}, [2]string{"Spam filter", "Spam filtresi"}},
		}},
	{name: "s3", line: "storage", aliases: []string{"storage"}, short: [2]string{"S3 object storage: buckets and access keys", "S3 nesne depolama: kovalar ve erişim anahtarları"},
		nouns: []lineNoun{
			{"bucket", []string{"buckets"}, []string{"buckets", "bucket"}, [2]string{"Buckets", "Kovalar"}},
			{"key", []string{"keys", "access-key"}, []string{"keys", "access-keys", "credentials"}, [2]string{"Access keys (secret shown once)", "Erişim anahtarları (gizli anahtar bir kez gösterilir)"}},
		}},
	{name: "monitor", line: "monitor", aliases: []string{"monitors", "uptime"}, short: [2]string{"Uptime monitors", "Erişilebilirlik izleme"},
		nouns: []lineNoun{
			{"check", []string{"checks"}, []string{"monitors", "checks"}, [2]string{"Monitor checks", "İzleme kontrolleri"}},
			{"incident", []string{"incidents", "event"}, []string{"incidents", "events"}, [2]string{"Incidents", "Olaylar"}},
			{"contact", []string{"contacts"}, []string{"contacts", "alert-contacts"}, [2]string{"Alert contacts", "Uyarı kişileri"}},
		}},
	{name: "license", line: "license", aliases: []string{"licenses", "lic"}, short: [2]string{"Software licenses (cPanel, Plesk, LiteSpeed ...)", "Yazılım lisansları (cPanel, Plesk, LiteSpeed ...)"}},
	{name: "app", line: "app", aliases: []string{"apps"}, short: [2]string{"One-click applications", "Tek tık uygulamalar"}},
	{name: "vds", line: "vds", short: [2]string{"VDS servers", "VDS sunucular"}},
	{name: "dedicated", line: "dedicated_server", aliases: []string{"bare-metal"}, short: [2]string{"Dedicated servers", "Fiziksel sunucular"}},
}

func addProductLines(a *App, root *cobra.Command) {
	if spec.ByID("listServiceTools") == nil || spec.ByID("listServices") == nil {
		return // API without the service bridge
	}
	for _, pl := range productLines {
		pl := pl
		g := &cobra.Command{
			Use: pl.name, Aliases: pl.aliases, Short: i18n.S(pl.short[0], pl.short[1]),
			Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error { return cmd.Help() },
		}
		g.AddCommand(lineList(a, pl), lineTools(a), lineTool(a), lineCall(a))
		for _, n := range pl.nouns {
			g.AddCommand(lineNounCmd(a, n))
		}
		if pl.name == "license" {
			for _, m := range []struct{ use, method string }{{"status", "license_status"}, {"change-ip", "change_ip"}} {
				m := m
				c := &cobra.Command{
					Use: m.use + " <SERVICE>", Short: i18n.F("Run %s on a license", "Lisansta %s çalıştır", m.method),
					Args: cobra.ExactArgs(1), ValidArgsFunction: completeServices(a, pl.line),
				}
				set := c.Flags().StringArray("set", nil, i18n.S("method field key=value", "metot alanı anahtar=değer"))
				c.RunE = func(cmd *cobra.Command, args []string) error {
					return runMethod(a, args[0], m.method, *set, m.method == "change_ip")
				}
				g.AddCommand(c)
			}
		}
		root.AddCommand(g)
	}
}

var serviceNameFields = []string{"host", "hostname", "domain", "name", "label"}

func (a *App) resolveService(ctx context.Context, arg string) (string, error) {
	if isNumeric(arg) {
		return arg, nil
	}
	id, _, err := a.resolveFrom(ctx, "/cloud/services", "", arg, serviceNameFields...)
	return id, err
}

func completeServices(a *App, line string) func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	path := "/cloud/services"
	_ = line
	return completeFromList(a, path)
}

func lineList(a *App, pl productLine) *cobra.Command {
	var all bool
	var status string
	c := &cobra.Command{
		Use: "list", Aliases: []string{"ls"}, Short: i18n.S("List your services of this product line", "Bu ürün hattındaki hizmetlerinizi listele"),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := a.Ctx()
			defer cancel()
			c, err := a.Client()
			if err != nil {
				return err
			}
			q := url.Values{"product_line": {pl.line}}
			if status != "" {
				q.Set("status", status)
			}
			items, err := c.All(ctx, "/cloud/services", q, "", pagesFor(all))
			if err != nil {
				return err
			}
			cols := []output.Column{{Header: "id", Path: "id"}, {Header: "name", Path: "name"}, {Header: "host", Path: "host"},
				{Header: "status", Path: "status"}, {Header: "module", Path: "module"}, {Header: "due", Path: "due_date"}}
			return a.PrintList(items, cols, "id")
		},
	}
	c.Flags().BoolVar(&all, "all", false, i18n.S("fetch every page", "tüm sayfaları getir"))
	c.Flags().StringVar(&status, "status", "", "active|suspended|cancelled ...")
	return c
}

type discovery struct {
	ProductLine string `json:"product_line"`
	Tools       []struct {
		Tool    string `json:"tool"`
		Group   string `json:"group"`
		Label   string `json:"label"`
		Actions []struct {
			Action  string `json:"action"`
			Read    bool   `json:"read"`
			Confirm bool   `json:"confirm"`
			Secret  bool   `json:"secret"`
		} `json:"actions"`
	} `json:"tools"`
	Methods []struct {
		Method  string `json:"method"`
		Confirm bool   `json:"confirm"`
	} `json:"methods"`
}

func (a *App) discover(ctx context.Context, sid string) (*discovery, json.RawMessage, error) {
	c, err := a.Client()
	if err != nil {
		return nil, nil, err
	}
	resp, err := c.Do(ctx, api.Request{Method: "GET", Path: "/cloud/services/" + sid + "/tools"})
	if err != nil {
		return nil, nil, err
	}
	var d discovery
	if err := json.Unmarshal(resp.Data, &d); err != nil {
		return nil, resp.Data, err
	}
	return &d, resp.Data, nil
}

func lineTools(a *App) *cobra.Command {
	return &cobra.Command{
		Use: "tools <SERVICE>", Short: i18n.S("Show the tools, actions and safety policy of a service", "Hizmetin araçlarını, eylemlerini ve güvenlik politikasını göster"),
		Args: cobra.ExactArgs(1), ValidArgsFunction: completeServices(a, ""),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := a.Ctx()
			defer cancel()
			sid, err := a.resolveService(ctx, args[0])
			if err != nil {
				return err
			}
			d, raw, err := a.discover(ctx, sid)
			if err != nil {
				return err
			}
			if a.Machine() || a.Opt.Query != "" {
				return a.PrintRaw(raw, nil, "")
			}
			var rows []any
			for _, t := range d.Tools {
				if len(t.Actions) == 0 {
					rows = append(rows, map[string]any{"tool": t.Tool, "label": t.Label, "action": "-"})
				}
				for _, ac := range t.Actions {
					rows = append(rows, map[string]any{"tool": t.Tool, "label": t.Label, "action": ac.Action, "read": ac.Read, "confirm": ac.Confirm, "secret": ac.Secret})
				}
			}
			for _, m := range d.Methods {
				rows = append(rows, map[string]any{"tool": "(method)", "label": "", "action": m.Method, "read": false, "confirm": m.Confirm, "secret": false})
			}
			cols := []output.Column{{Header: "tool", Path: "tool"}, {Header: "action", Path: "action"}, {Header: "read", Path: "read"},
				{Header: "confirm", Path: "confirm"}, {Header: "secret", Path: "secret"}, {Header: "label", Path: "label"}}
			return a.Printer().Print(toGeneric(rows), cols, "tool")
		},
	}
}

// runTool reads tool data (no action / read action) or runs an action, honouring the discovered policy.
func runTool(a *App, svcArg, tool, action string, set []string, data string) error {
	ctx, cancel := a.Ctx()
	defer cancel()
	sid, err := a.resolveService(ctx, svcArg)
	if err != nil {
		return err
	}
	c, _ := a.Client()
	body, err := a.ReadData(data)
	if err != nil {
		return err
	}
	extra, err := ParseKV(set)
	if err != nil {
		return err
	}
	if body == nil {
		body = map[string]any{}
	}
	for k, v := range extra {
		body[k] = v
	}
	if action == "" {
		q := url.Values{}
		for k, v := range body {
			q.Set(k, fmt.Sprint(v))
		}
		resp, err := c.Do(ctx, api.Request{Method: "GET", Path: "/cloud/services/" + sid + "/tools/" + url.PathEscape(tool), Query: q})
		if err != nil {
			return err
		}
		return a.printToolData(resp.Data)
	}
	confirm, secret := true, false // unknown actions are treated as needing confirmation (API policy)
	if d, _, err := a.discover(ctx, sid); err == nil {
		for _, t := range d.Tools {
			if t.Tool != tool {
				continue
			}
			for _, ac := range t.Actions {
				if ac.Action == action {
					confirm, secret = ac.Confirm, ac.Secret
				}
			}
		}
	}
	if confirm {
		if err := a.Confirm(i18n.F("Run %s/%s on service %s?", "%s/%s, %s hizmetinde çalıştırılsın mı?", tool, action, sid)); err != nil {
			return err
		}
		body["confirm"] = true
	}
	req := api.Request{Method: "POST", Path: "/cloud/services/" + sid + "/tools/" + url.PathEscape(tool) + "/" + url.PathEscape(action), Body: body}
	_ = secret // secrets: no Idempotency-Key is sent, so the one-time value is returned
	resp, err := c.Do(ctx, req)
	if err != nil {
		return err
	}
	if err := a.printToolData(resp.Data); err != nil {
		return err
	}
	warnSecrets(a, resp.Data)
	return nil
}

func (a *App) printToolData(raw json.RawMessage) error {
	v, err := output.Decode(raw)
	if err != nil {
		return err
	}
	// tool data is often {"items": [...]} or {"<tool>": [...]}; show the list in table mode
	if m, ok := v.(map[string]any); ok && !a.Machine() && a.Opt.Query == "" && a.Printer().Format != "text" {
		var lists []string
		for k, x := range m {
			if arr, ok := x.([]any); ok && len(arr) > 0 {
				if _, obj := arr[0].(map[string]any); obj {
					lists = append(lists, k)
				}
			}
		}
		if len(lists) == 1 {
			return a.Printer().Print(m[lists[0]], nil, "")
		}
	}
	return a.Printer().Print(v, nil, "")
}

func lineTool(a *App) *cobra.Command {
	var set []string
	var data string
	c := &cobra.Command{
		Use: "tool <SERVICE> <TOOL> [ACTION]", Short: i18n.S("Read a tool's data or run one of its actions (see 'tools')", "Bir aracın verisini oku ya da eylemini çalıştır ('tools' ile bakın)"),
		Example: "  gyro mail tool example.com mailboxes\n  gyro mail tool example.com mailboxes create --set local_part=info --set password=@-",
		Args:    cobra.RangeArgs(2, 3), ValidArgsFunction: completeServices(a, ""),
		RunE: func(cmd *cobra.Command, args []string) error {
			action := ""
			if len(args) == 3 {
				action = args[2]
			}
			return runTool(a, args[0], args[1], action, set, data)
		},
	}
	c.Flags().StringArrayVar(&set, "set", nil, i18n.S("field key=value (repeatable)", "alan anahtar=değer (tekrarlanabilir)"))
	c.Flags().StringVar(&data, "data", "", i18n.S("JSON body, @file or @-", "JSON gövde, @dosya ya da @-"))
	return c
}

func runMethod(a *App, svcArg, method string, set []string, confirm bool) error {
	ctx, cancel := a.Ctx()
	defer cancel()
	sid, err := a.resolveService(ctx, svcArg)
	if err != nil {
		return err
	}
	body, err := ParseKV(set)
	if err != nil {
		return err
	}
	if confirm {
		if err := a.Confirm(i18n.F("Call %s on service %s?", "%s, %s hizmetinde çağrılsın mı?", method, sid)); err != nil {
			return err
		}
		body["confirm"] = true
	}
	c, _ := a.Client()
	resp, err := c.Do(ctx, api.Request{Method: "POST", Path: "/cloud/services/" + sid + "/methods/" + url.PathEscape(method), Body: body})
	if err != nil {
		return err
	}
	return a.PrintRaw(resp.Data, nil, "")
}

func lineCall(a *App) *cobra.Command {
	var set []string
	c := &cobra.Command{
		Use: "call <SERVICE> <METHOD>", Short: i18n.S("Call a module method of a service", "Hizmetin modül metodunu çağır"),
		Args: cobra.ExactArgs(2), ValidArgsFunction: completeServices(a, ""),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runMethod(a, args[0], args[1], set, true)
		},
	}
	c.Flags().StringArrayVar(&set, "set", nil, i18n.S("field key=value", "alan anahtar=değer"))
	return c
}

var verbSynonyms = map[string]*regexp.Regexp{
	"create": regexp.MustCompile(`^(create|add|new|insert)`),
	"update": regexp.MustCompile(`^(update|edit|save|set|change)`),
	"delete": regexp.MustCompile(`^(delete|remove|destroy|del)`),
}

// lineNounCmd builds e.g. `gyro mail mailbox list|create|update|delete <SERVICE>`.
func lineNounCmd(a *App, n lineNoun) *cobra.Command {
	g := &cobra.Command{Use: n.name, Aliases: n.aliases, Short: i18n.S(n.short[0], n.short[1]),
		Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error { return cmd.Help() }}
	pickTool := func(ctx context.Context, sid string) (string, *discovery, error) {
		d, _, err := a.discover(ctx, sid)
		if err != nil {
			return "", nil, err
		}
		for _, cand := range n.tools {
			for _, t := range d.Tools {
				if t.Tool == cand {
					return cand, d, nil
				}
			}
		}
		var names []string
		for _, t := range d.Tools {
			names = append(names, t.Tool)
		}
		return "", d, &ExitErr{ExitNotFound, fmt.Errorf("%s (%s): %s", i18n.S("this service has no such tool", "bu hizmette böyle bir araç yok"), strings.Join(n.tools, "/"), strings.Join(names, ", "))}
	}
	list := &cobra.Command{Use: "list <SERVICE>", Aliases: []string{"ls", "get"}, Short: i18n.S("Show", "Göster"), Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := a.Ctx()
			sid, err := a.resolveService(ctx, args[0])
			if err != nil {
				cancel()
				return err
			}
			tool, _, err := pickTool(ctx, sid)
			cancel()
			if err != nil {
				return err
			}
			return runTool(a, sid, tool, "", nil, "")
		}}
	g.AddCommand(list)
	for _, verb := range []string{"create", "update", "delete"} {
		verb := verb
		var set []string
		var data string
		c := &cobra.Command{Use: verb + " <SERVICE>", Short: i18n.F("%s (fields with --set key=value)", "%s (alanlar --set anahtar=değer)", verb), Args: cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				ctx, cancel := a.Ctx()
				sid, err := a.resolveService(ctx, args[0])
				if err != nil {
					cancel()
					return err
				}
				tool, d, err := pickTool(ctx, sid)
				cancel()
				if err != nil {
					return err
				}
				action := ""
				var acts []string
				for _, t := range d.Tools {
					if t.Tool != tool {
						continue
					}
					for _, ac := range t.Actions {
						acts = append(acts, ac.Action)
						if action == "" && verbSynonyms[verb].MatchString(ac.Action) {
							action = ac.Action
						}
					}
				}
				if action == "" {
					return &ExitErr{ExitUsage, fmt.Errorf("%s %q: %s (gyro ... tool %s %s <ACTION>)", i18n.S("no matching action for", "eşleşen eylem yok:"), verb, strings.Join(acts, ", "), args[0], tool)}
				}
				return runTool(a, sid, tool, action, set, data)
			}}
		c.Flags().StringArrayVar(&set, "set", nil, i18n.S("field key=value (repeatable)", "alan anahtar=değer (tekrarlanabilir)"))
		c.Flags().StringVar(&data, "data", "", i18n.S("JSON body, @file or @-", "JSON gövde, @dosya ya da @-"))
		g.AddCommand(c)
	}
	return g
}

/* ------------------------------------------------------------------ dns shortcut */

// addDNS adds `gyro dns record list|add|update|delete <DOMAIN>` on top of the generated domain DNS commands.
func addDNS(a *App, root *cobra.Command) {
	ops := map[string]string{"list": "coreGetDnsRecords", "add": "coreCreateDnsRecord", "update": "coreUpdateDnsRecord", "delete": "coreDeleteDnsRecord"}
	if spec.ByID(ops["list"]) == nil {
		return
	}
	dns := &cobra.Command{Use: "dns", Short: i18n.S("DNS records of your domains", "Alan adlarınızın DNS kayıtları"), Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error { return cmd.Help() }}
	rec := &cobra.Command{Use: "record", Aliases: []string{"records", "rr"}, Short: i18n.S("List, add, change and delete DNS records", "DNS kayıtlarını listele, ekle, değiştir, sil"),
		Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error { return cmd.Help() }}
	for _, name := range []string{"list", "add", "update", "delete"} {
		op := spec.ByID(ops[name])
		if op == nil {
			continue
		}
		c := genCommand(a, op)
		c.Use = strings.Replace(c.Use, op.Command[len(op.Command)-1], name, 1)
		if name == "add" {
			c.Aliases = []string{"create"}
			c.Example = "  gyro dns record add example.com --type A --name www --value 203.0.113.10 --ttl 3600"
		}
		if name == "list" {
			c.Aliases = []string{"ls"}
		}
		rec.AddCommand(c)
	}
	dns.AddCommand(rec)
	root.AddCommand(dns)
}

// addCoreShortcuts exposes "gyro core X ..." operations at "gyro X ..." when that path is free.
func addCoreShortcuts(a *App, root *cobra.Command) {
	for i := range spec.Ops {
		op := &spec.Ops[i]
		if len(op.Command) < 2 || op.Command[0] != "core" || handled[op.ID] {
			continue
		}
		words := op.Command[1:]
		if c, rest, err := root.Find(words); err == nil && c != root && len(rest) == 0 && c.Name() == words[len(words)-1] {
			continue // taken
		}
		parent := ensureGroup(root, words[:len(words)-1])
		if findChild(parent, words[len(words)-1]) != nil {
			continue
		}
		parent.AddCommand(genCommand(a, op))
	}
	if core := findChild(root, "core"); core != nil {
		core.Short = i18n.S("Raw WISECP core endpoints shadowed by Gyro cloud endpoints", "Gyro bulut uçlarının gölgelediği ham WISECP çekirdek uçları")
		core.Hidden = true
		delete(core.Annotations, "autoShort")
	}
}
