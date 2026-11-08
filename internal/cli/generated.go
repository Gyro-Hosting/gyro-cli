package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
	"github.com/Gyro-Hosting/gyro-cli/internal/output"
	"github.com/Gyro-Hosting/gyro-cli/internal/spec"
)

// handled lists operationIds covered by hand-written commands; the generator skips them
// (they stay reachable through `gyro api call <operationId>`).
var handled = map[string]bool{}

// handledCmd lists generated command paths ("server action poweron") replaced by hand-written ones;
// used for operations expanded per enum value, where new values must still appear automatically.
var handledCmd = map[string]bool{}

func markHandled(ids ...string) {
	for _, id := range ids {
		handled[id] = true
	}
}

func markHandledCmd(cmds ...string) {
	for _, c := range cmds {
		handledCmd[c] = true
	}
}

var globalFlags = map[string]bool{"profile": true, "api-url": true, "output": true, "query": true, "no-header": true,
	"quiet": true, "no-color": true, "lang": true, "yes": true, "debug": true, "wide": true, "timeout": true, "help": true,
	"data": true, "set": true, "dry-run": true, "wait": true, "all": true, "page": true, "limit": true, "idempotency-key": true}

// resourceAliases are extra names for generated group commands.
var resourceAliases = map[string][]string{
	"server": {"servers", "srv"}, "ssh-key": {"ssh-keys", "sshkey"}, "invoice": {"invoices", "inv"},
	"domain": {"domains"}, "ticket": {"tickets"}, "service": {"services", "svc"}, "webhook": {"webhooks", "hook"},
	"order": {"orders"}, "location": {"locations", "loc"}, "image": {"images"},
	"snapshot": {"snapshots"}, "backup": {"backups"}, "mail": {"email"}, "monitor": {"monitors"},
}

func summary(o *spec.Op) string {
	s := i18n.S(o.Summary, o.SummaryTR)
	if s == "" {
		s = o.ID
	}
	return s
}

func describe(o *spec.Op) string {
	d := i18n.S(o.Desc, o.DescTR)
	long := summary(o)
	if d != "" {
		long += "\n\n" + d
	}
	long += fmt.Sprintf("\n\nAPI: %s %s (operationId %s", o.Method, o.Path, o.ID)
	if o.Scope != "" {
		long += ", scope " + o.Scope
	}
	return long + ")"
}

// ensureGroup returns (creating if needed) the command for a word path under root.
func ensureGroup(root *cobra.Command, words []string) *cobra.Command {
	cur := root
	for _, w := range words {
		var next *cobra.Command
		for _, c := range cur.Commands() {
			if c.Name() == w || c.HasAlias(w) {
				next = c
				break
			}
		}
		if next == nil {
			next = &cobra.Command{
				Use:         w,
				Short:       i18n.F("Manage %s", "%s işlemleri", strings.ReplaceAll(w, "-", " ")),
				Annotations: map[string]string{"autoShort": "1"},
				Aliases:     resourceAliases[w],
				Args:        cobra.NoArgs,
				RunE: func(cmd *cobra.Command, args []string) error {
					return cmd.Help()
				},
			}
			cur.AddCommand(next)
		}
		cur = next
	}
	return cur
}

func findChild(parent *cobra.Command, name string) *cobra.Command {
	for _, c := range parent.Commands() {
		if c.Name() == name || c.HasAlias(name) {
			return c
		}
	}
	return nil
}

// addGenerated creates a command for every operation not covered by a hand-written one.
func addGenerated(a *App, root *cobra.Command) {
	for i := range spec.Ops {
		op := &spec.Ops[i]
		if handled[op.ID] || handledCmd[strings.Join(op.Command, " ")] {
			continue
		}
		words := op.Command
		parent := ensureGroup(root, words[:len(words)-1])
		if ex := findChild(parent, words[len(words)-1]); ex != nil {
			if ex.Annotations["autoShort"] != "" && ex.Annotations["operationId"] == "" {
				// an auto-created group with the same name (ticket reply + ticket reply rating): make it runnable
				gc := genCommand(a, op)
				ex.Use, ex.Short, ex.Long, ex.Args, ex.RunE, ex.ValidArgsFunction = gc.Use, gc.Short, gc.Long, gc.Args, gc.RunE, gc.ValidArgsFunction
				ex.Flags().AddFlagSet(gc.Flags())
				ex.Annotations["operationId"] = op.ID
				delete(ex.Annotations, "autoShort")
			}
			continue // hand-written command with the same name wins
		}
		cmd := genCommand(a, op)
		cmd.Hidden = op.Hidden
		parent.AddCommand(cmd)
		// a parameterless "get" also runs on the bare group: `gyro balance`
		if cmd.Name() == "get" && len(op.PathParams()) == 0 && parent != root {
			if parent.RunE == nil || parent.Annotations["generated"] == "" {
				if parent.Annotations == nil {
					parent.Annotations = map[string]string{}
				}
				parent.Annotations["generated"] = "1"
				delete(parent.Annotations, "autoShort")
				parent.Short = summary(op)
				parent.Args = cobra.NoArgs
				parent.RunE = cmd.RunE
				parent.Flags().AddFlagSet(cmd.Flags())
			}
		}
	}
}

// groupShorts describes well-known generated groups (others take the summary of their list/get child).
var groupShorts = map[string][2]string{
	"domain":    {"Domains: search, register, transfer, nameservers, DNS, WHOIS", "Alan adları: sorgu, kayıt, transfer, ad sunucuları, DNS, WHOIS"},
	"ticket":    {"Support tickets", "Destek talepleri"},
	"order":     {"Orders (preview and create)", "Siparişler (önizleme ve oluşturma)"},
	"reseller":  {"Reseller: end-customers, assignments, keys, report", "Bayi: uç-müşteriler, atamalar, anahtarlar, rapor"},
	"service":   {"All services of every product line (generic bridge)", "Tüm ürün hatlarındaki hizmetler (genel köprü)"},
	"key":       {"Sub-account API keys", "Alt hesap API anahtarları"},
	"webhook":   {"Webhooks (signed event callbacks)", "Webhook'lar (imzalı olay bildirimleri)"},
	"me":        {"Your profile, notifications and security", "Profiliniz, bildirimler ve güvenlik"},
	"address":   {"Billing addresses", "Fatura adresleri"},
	"affiliate": {"Affiliate programme", "Satış ortaklığı programı"},
	"product":   {"Full product catalog (all product lines)", "Tüm ürün kataloğu (tüm ürün hatları)"},
	"reference": {"Countries, states and cities", "Ülkeler, iller ve ilçeler"},
	"sms":       {"SMS messages", "SMS mesajları"},
	"activity":  {"Write log of your API keys", "API anahtarlarınızın yazma kaydı"},
	"image":     {"Operating system images", "İşletim sistemi imajları"},
	"location":  {"Locations", "Lokasyonlar"},
}

// fillGroupShorts gives auto-created groups a description taken from their list/get child.
func fillGroupShorts(c *cobra.Command) {
	for _, ch := range c.Commands() {
		fillGroupShorts(ch)
	}
	if c.Annotations["autoShort"] == "" {
		return
	}
	if gs, ok := groupShorts[c.Name()]; ok && c.Parent() != nil && c.Parent().Parent() == nil {
		c.Short = i18n.S(gs[0], gs[1])
		return
	}
	for _, pref := range []string{"list", "get"} {
		if ch := findChild(c, pref); ch != nil {
			c.Short = ch.Short
			return
		}
	}
	for _, ch := range c.Commands() {
		if !ch.Hidden && ch.Name() != "help" {
			c.Short = ch.Short
			return
		}
	}
}

// genOpts are the per-command flag values of a generated command.
type genOpts struct {
	data     string
	set      []string
	dryRun   bool
	wait     bool
	waitFor  time.Duration
	all      bool
	page     int
	limit    int
	idem     string
	flagVals map[string]any // name -> pointer
	flagName map[string]string
}

func flagNameFor(name string, used map[string]bool) string {
	f := spec.Kebab(name)
	if globalFlags[f] || used[f] {
		f = "param-" + f
	}
	used[f] = true
	return f
}

func addTypedFlag(fs *pflag.FlagSet, name, typ, items, usage string, enum []string) any {
	if len(enum) > 1 {
		usage += " (" + strings.Join(enum, "|") + ")"
	}
	switch typ {
	case "integer":
		return fs.Int64(name, 0, usage)
	case "number":
		return fs.Float64(name, 0, usage)
	case "boolean":
		return fs.Bool(name, false, usage)
	case "array":
		if items == "integer" {
			return fs.Int64Slice(name, nil, usage)
		}
		return fs.StringSlice(name, nil, usage)
	case "object":
		return fs.String(name, "", usage+" (JSON)")
	default:
		return fs.String(name, "", usage)
	}
}

func flagValue(p any, typ string) any {
	switch v := p.(type) {
	case *int64:
		return *v
	case *float64:
		return *v
	case *bool:
		return *v
	case *[]int64:
		return *v
	case *[]string:
		return *v
	case *string:
		if typ == "object" {
			var x any
			if json.Unmarshal([]byte(*v), &x) == nil {
				return x
			}
		}
		return *v
	}
	return nil
}

// listPathFor finds the list endpoint for a path parameter (/cloud/servers/{id} -> /cloud/servers).
func listPathFor(path, param string) (string, *spec.Op) {
	i := strings.Index(path, "/{"+param+"}")
	if i < 0 {
		return "", nil
	}
	base := path[:i]
	if strings.Contains(base, "{") {
		return "", nil // nested ids need the parent id; keep it simple
	}
	for j := range spec.Ops {
		o := &spec.Ops[j]
		if o.Method == "GET" && o.Path == base && o.ListResponse {
			return base, o
		}
	}
	return "", nil
}

var nameFields = []string{"hostname", "host", "name", "slug", "domain", "email", "number", "label", "title", "code", "description"}

func genCommand(a *App, op *spec.Op) *cobra.Command {
	pp := op.PathParams()
	var argNames []string
	var fixed = map[string]string{}
	for _, p := range pp {
		if len(p.Enum) == 1 && p.Default != "" {
			fixed[p.Name] = p.Default
			continue
		}
		argNames = append(argNames, strings.ToUpper(spec.Kebab(p.Name)))
	}
	o := &genOpts{flagVals: map[string]any{}, flagName: map[string]string{}}
	use := op.Command[len(op.Command)-1]
	for _, n := range argNames {
		use += " <" + n + ">"
	}
	cmd := &cobra.Command{
		Use:   use,
		Short: summary(op),
		Long:  describe(op),
		Args:  cobra.ExactArgs(len(argNames)),
		Annotations: map[string]string{
			"operationId": op.ID,
		},
	}
	fs := cmd.Flags()
	used := map[string]bool{}
	for _, p := range op.Params {
		if p.In == "path" || p.Name == "confirm" || (op.Paginated && (p.Name == "page" || p.Name == "limit")) {
			continue
		}
		if p.In == "header" && strings.EqualFold(p.Name, "Idempotency-Key") {
			continue
		}
		fn := flagNameFor(p.Name, used)
		o.flagName["q:"+p.Name] = fn
		o.flagVals["q:"+p.Name] = addTypedFlag(fs, fn, p.Type, p.Items, i18n.S(p.Desc, p.DescTR), p.Enum)
		if len(p.Enum) > 1 {
			_ = cmd.RegisterFlagCompletionFunc(fn, fixedCompletion(p.Enum...))
		}
	}
	for _, f := range op.Body {
		if f.Name == "confirm" || f.Name == "dry_run" {
			continue
		}
		fn := flagNameFor(f.Name, used)
		o.flagName["b:"+f.Name] = fn
		usage := i18n.S(f.Desc, f.DescTR)
		if usage == "" {
			usage = f.Name
		}
		o.flagVals["b:"+f.Name] = addTypedFlag(fs, fn, f.Type, f.Items, usage, f.Enum)
		if len(f.Enum) > 1 {
			_ = cmd.RegisterFlagCompletionFunc(fn, fixedCompletion(f.Enum...))
		}
	}
	if op.HasBody {
		fs.StringVar(&o.data, "data", "", i18n.S("request body as JSON, @file or @- (stdin)", "istek gövdesi JSON, @dosya ya da @- (stdin)"))
		fs.StringArrayVar(&o.set, "set", nil, i18n.S("extra body field key=value (repeatable)", "ek gövde alanı anahtar=değer (tekrarlanabilir)"))
	}
	if op.Method != "GET" {
		fs.BoolVar(&o.dryRun, "dry-run", false, i18n.S("show the price quote / request without changing anything", "hiçbir şeyi değiştirmeden teklifi / isteği göster"))
		fs.StringVar(&o.idem, "idempotency-key", "", i18n.S("idempotency key (automatic for money operations)", "tekrar güvenliği anahtarı (para işlemlerinde otomatik)"))
		if strings.HasPrefix(op.Path, "/cloud/servers/{id}") {
			fs.BoolVar(&o.wait, "wait", false, i18n.S("wait until the server is no longer busy", "sunucu meşgul olmayana kadar bekle"))
			fs.DurationVar(&o.waitFor, "wait-timeout", 15*time.Minute, i18n.S("maximum time for --wait", "--wait için en uzun süre"))
		}
	}
	if op.Paginated {
		fs.BoolVar(&o.all, "all", false, i18n.S("fetch every page", "tüm sayfaları getir"))
		fs.IntVar(&o.page, "page", 1, i18n.S("page number", "sayfa numarası"))
		fs.IntVar(&o.limit, "limit", 0, i18n.S("items per page (max 100)", "sayfa başına kayıt (en çok 100)"))
	}
	if len(argNames) > 0 {
		if lp, _ := listPathFor(op.Path, pp[0].Name); lp != "" {
			cmd.ValidArgsFunction = completeFromList(a, lp)
		}
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runGenerated(a, cmd, op, o, args, fixed)
	}
	return cmd
}

func completeFromList(a *App, listPath string) func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) > 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		c, err := a.Client()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		c.MaxRetries = 0
		items, err := c.All(ctx, listPath, nil, listItemsKey(listPath), 3)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		var out []string
		for _, it := range items {
			v, _ := output.Decode(it)
			m, _ := v.(map[string]any)
			if m == nil {
				continue
			}
			label := ""
			for _, f := range nameFields {
				if s, ok := m[f].(string); ok && s != "" {
					label = s
					break
				}
			}
			out = append(out, output.Scalar(m["id"])+"\t"+label)
		}
		return out, cobra.ShellCompDirectiveNoFileComp
	}
}

func listItemsKey(path string) string {
	for i := range spec.Ops {
		if spec.Ops[i].Method == "GET" && spec.Ops[i].Path == path {
			return spec.Ops[i].ItemsKey
		}
	}
	return ""
}

// resolvePathArg turns a name into an id when the parameter is numeric and a list endpoint exists.
func (a *App) resolvePathArg(ctx context.Context, op *spec.Op, p spec.Param, arg string) (string, error) {
	if p.Type != "integer" || isNumeric(arg) {
		return arg, nil
	}
	lp, lo := listPathFor(op.Path, p.Name)
	if lp == "" {
		return arg, nil
	}
	id, _, err := a.resolveFrom(ctx, lp, lo.ItemsKey, arg, nameFields...)
	return id, err
}

func runGenerated(a *App, cmd *cobra.Command, op *spec.Op, o *genOpts, args []string, fixed map[string]string) error {
	ctx, cancel := a.Ctx()
	defer cancel()
	c, err := a.Client()
	if err != nil {
		return err
	}
	// path
	path := op.Path
	ai := 0
	var firstID string
	for _, p := range op.PathParams() {
		val, ok := fixed[p.Name]
		if !ok {
			val, err = a.resolvePathArg(ctx, op, p, args[ai])
			if err != nil {
				return err
			}
			ai++
		}
		if firstID == "" {
			firstID = val
		}
		path = strings.Replace(path, "{"+p.Name+"}", url.PathEscape(val), 1)
	}
	// query
	q := url.Values{}
	for _, p := range op.Params {
		key := "q:" + p.Name
		ptr, ok := o.flagVals[key]
		if !ok || p.In != "query" || !cmd.Flags().Changed(o.flagName[key]) {
			continue
		}
		switch v := flagValue(ptr, p.Type).(type) {
		case []string:
			for _, x := range v {
				q.Add(p.Name, x)
			}
		case []int64:
			for _, x := range v {
				q.Add(p.Name, strconv.FormatInt(x, 10))
			}
		default:
			q.Set(p.Name, fmt.Sprint(v))
		}
	}
	if op.Paginated {
		if o.limit > 0 {
			q.Set("limit", strconv.Itoa(o.limit))
		}
		if o.page > 1 {
			q.Set("page", strconv.Itoa(o.page))
		}
	}
	// body
	var body map[string]any
	if op.HasBody {
		body, err = a.ReadData(o.data)
		if err != nil {
			return err
		}
		if body == nil {
			body = map[string]any{}
		}
		for _, f := range op.Body {
			key := "b:" + f.Name
			if ptr, ok := o.flagVals[key]; ok && cmd.Flags().Changed(o.flagName[key]) {
				body[f.Name] = flagValue(ptr, f.Type)
			}
		}
		extra, err := ParseKV(o.set)
		if err != nil {
			return err
		}
		for k, v := range extra {
			body[k] = v
		}
		for _, f := range op.Body {
			if f.Required && f.Name != "confirm" && f.Name != "dry_run" {
				if _, ok := body[f.Name]; !ok {
					return &ExitErr{ExitUsage, fmt.Errorf("--%s %s", o.flagName["b:"+f.Name], i18n.S("is required", "zorunlu"))}
				}
			}
		}
	}
	_, bodyConfirm := op.Field("confirm")
	_, queryConfirm := op.Param("confirm")
	_, bodyDry := op.Field("dry_run")
	idem := o.idem
	if idem == "" && op.Method != "GET" {
		if _, required := op.Param("Idempotency-Key"); required || op.Money {
			idem = api.NewID()
		}
	}

	// --dry-run: server-side quote when supported, otherwise show the request
	if o.dryRun {
		if bodyDry {
			body["dry_run"] = true
			resp, err := c.Do(ctx, api.Request{Method: op.Method, Path: path, Query: q, Body: body})
			if err != nil {
				return err
			}
			a.Info("%s", i18n.S("Dry run: nothing was changed.", "Deneme: hiçbir şey değiştirilmedi."))
			return a.PrintRaw(resp.Data, nil, "")
		}
		plan := map[string]any{"method": op.Method, "url": c.URL(path, q), "operation": op.ID, "sent": false}
		if body != nil {
			plan["body"] = body
		}
		if idem != "" {
			plan["idempotency_key"] = idem
		}
		a.Info("%s", i18n.S("Dry run: the request below was NOT sent.", "Deneme: aşağıdaki istek GÖNDERİLMEDİ."))
		return a.Printer().Print(toGeneric(plan), nil, "")
	}

	if op.Confirm {
		question := i18n.F("Run %s on %s?", "%s çalıştırılsın mı (%s)?", strings.Join(op.Command, " "), path)
		if op.Money {
			question = i18n.F("This spends money from your balance (%s). Continue?", "Bu işlem bakiyenizden para harcar (%s). Devam edilsin mi?", strings.Join(op.Command, " "))
		}
		if err := a.Confirm(question); err != nil {
			return err
		}
		if bodyConfirm {
			body["confirm"] = true
		}
		if queryConfirm {
			q.Set("confirm", "true")
		}
	}

	var reqBody any
	if op.HasBody {
		reqBody = body
	}
	if op.Paginated && o.all {
		items, err := c.All(ctx, path, q, op.ItemsKey, 0)
		if err != nil {
			return err
		}
		return a.PrintList(items, cols(op, a.Opt.Wide), "")
	}
	resp, err := c.Do(ctx, api.Request{Method: op.Method, Path: path, Query: q, Body: reqBody, IdempotencyKey: idem})
	if err != nil {
		return err
	}
	if resp.Replayed {
		a.Info("%s", i18n.S("(idempotent replay: the server returned the stored answer)", "(tekrar: sunucu kayıtlı ilk yanıtı döndürdü)"))
	}
	if err := a.printOpResult(op, resp); err != nil {
		return err
	}
	warnSecrets(a, resp.Data)
	if o.wait && firstID != "" {
		return a.waitServerIdle(ctx, firstID, o.waitFor)
	}
	return nil
}

func toGeneric(v any) any {
	b, _ := json.Marshal(v)
	x, _ := output.Decode(b)
	return x
}

// columnOverrides hand-tunes default tables of generated list commands (by operationId).
var columnOverrides = map[string][]string{
	"listSshKeys":  {"id", "name", "fingerprint", "created_at"},
	"listActivity": {"id", "created_at", "method", "path", "status_code", "scope", "ip", "key_id"},
}

func cols(op *spec.Op, wide bool) []output.Column {
	if wide {
		return nil
	}
	if c, ok := columnOverrides[op.ID]; ok {
		return output.FromNames(c, 0)
	}
	if len(op.Columns) == 0 {
		return nil
	}
	return output.FromNames(op.Columns, 8)
}

func (a *App) printOpResult(op *spec.Op, resp *api.Response) error {
	p := a.Printer()
	if op.ListResponse && op.ItemsKey != "" && (p.Format == "table" || p.Format == "text" || p.Quiet) && a.Opt.Query == "" {
		items, err := api.Items(resp.Data, op.ItemsKey)
		if err != nil {
			return err
		}
		return a.PrintList(items, cols(op, a.Opt.Wide), "")
	}
	if op.ListResponse {
		return a.PrintRaw(resp.Data, cols(op, a.Opt.Wide), "")
	}
	return a.PrintRaw(resp.Data, nil, "")
}

// warnSecrets reminds the user that one-time secrets are not shown again.
func warnSecrets(a *App, data json.RawMessage) {
	v, _ := output.Decode(data)
	m, _ := v.(map[string]any)
	if m == nil {
		return
	}
	var found []string
	for k, val := range m {
		lk := strings.ToLower(k)
		if (strings.Contains(lk, "password") || strings.Contains(lk, "secret") || lk == "token" || strings.HasSuffix(lk, "_token") || lk == "secret_key") && !strings.HasSuffix(lk, "withheld") {
			if s, ok := val.(string); ok && s != "" {
				found = append(found, k)
			}
		}
		if strings.HasSuffix(lk, "_withheld") && val == true {
			a.Info("%s", i18n.F("Note: %s was withheld (idempotent request); repeat without --idempotency-key to receive it.",
				"Not: %s gizlendi (tekrar güvenli istek); almak için --idempotency-key olmadan tekrarlayın.", strings.TrimSuffix(k, "_withheld")))
		}
	}
	sort.Strings(found)
	if len(found) > 0 {
		a.Info("%s", i18n.F("Store %s now: it is shown only once.", "%s bilgisini şimdi saklayın: yalnız bir kez gösterilir.", strings.Join(found, ", ")))
	}
}
