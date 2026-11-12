package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
	"github.com/Gyro-Hosting/gyro-cli/internal/output"
	"github.com/Gyro-Hosting/gyro-cli/internal/spec"
)

var serverColumns = []output.Column{
	{Header: "id", Path: "id"},
	{Header: "hostname", Path: "hostname"},
	{Header: "status", Path: "status"},
	{Header: "plan", Path: "plan"},
	{Header: "region", Path: "region"},
	{Header: "ipv4", Path: "ipv4"},
	{Header: "price", Path: "price"},
	{Header: "due", Path: "due_date"},
}

// resolveServer accepts an id or a hostname (full or first label).
func (a *App) resolveServer(ctx context.Context, arg string) (string, error) {
	if isNumeric(arg) {
		return arg, nil
	}
	id, _, err := a.resolveFrom(ctx, "/cloud/servers", "", arg, "hostname", "name")
	return id, err
}

func completeServers(a *App) func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	return completeFromList(a, "/cloud/servers")
}

func addServer(a *App, root *cobra.Command) {
	srv := ensureGroup(root, []string{"server"})
	srv.Short = i18n.S("Cloud servers: create, power, snapshots, SSH ...", "Bulut sunucular: oluşturma, güç, snapshot, SSH ...")
	delete(srv.Annotations, "autoShort")
	markHandled("listServers", "getServer", "createServer", "deleteServer", "getMetrics")

	// list
	var status string
	var all bool
	list := &cobra.Command{
		Use: "list", Aliases: []string{"ls"},
		Short:   i18n.S("List your servers", "Sunucularınızı listeleyin"),
		Example: "  gyro server list\n  gyro server list --status active -o json\n  gyro server list --query \"[?region=='eu-de-1'].ipv4\" -o text",
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
			items, err := c.All(ctx, "/cloud/servers", q, "", pagesFor(all))
			if err != nil {
				return err
			}
			if !all && !a.Opt.Quiet && a.Opt.Output == "" && len(items) == 100 {
				a.Info("%s", i18n.S("(first 100 servers; use --all for more)", "(ilk 100 sunucu; tümü için --all)"))
			}
			return a.PrintList(items, serverColumns, "id")
		},
	}
	list.Flags().StringVar(&status, "status", "", "waiting|inprocess|active|suspended|cancelled")
	list.Flags().BoolVar(&all, "all", false, i18n.S("fetch every page", "tüm sayfaları getir"))
	_ = list.RegisterFlagCompletionFunc("status", fixedCompletion("waiting", "inprocess", "active", "suspended", "cancelled"))
	srv.AddCommand(list)

	// get
	get := &cobra.Command{
		Use: "get <SERVER>", Aliases: []string{"describe", "show"},
		Short:             i18n.S("Show a server with its live state", "Sunucuyu canlı durumuyla göster"),
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: completeServers(a),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := a.Ctx()
			defer cancel()
			id, err := a.resolveServer(ctx, args[0])
			if err != nil {
				return err
			}
			c, _ := a.Client()
			resp, err := c.Do(ctx, api.Request{Method: "GET", Path: "/cloud/servers/" + id})
			if err != nil {
				return err
			}
			return a.PrintRaw(resp.Data, nil, "id")
		},
	}
	srv.AddCommand(get)

	// ip (handy in scripts)
	ipCmd := &cobra.Command{
		Use:               "ip <SERVER>",
		Short:             i18n.S("Print the public IPv4 (or --ipv6) of a server", "Sunucunun genel IPv4 adresini (ya da --ipv6) yaz"),
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: completeServers(a),
	}
	var ip6 bool
	ipCmd.Flags().BoolVarP(&ip6, "ipv6", "6", false, "IPv6")
	ipCmd.RunE = func(cmd *cobra.Command, args []string) error {
		ctx, cancel := a.Ctx()
		defer cancel()
		s, err := a.fetchServer(ctx, args[0])
		if err != nil {
			return err
		}
		addr := serverAddr(s, ip6)
		if addr == "" {
			return fmt.Errorf("%s", i18n.S("the server has no such address yet", "sunucunun henüz böyle bir adresi yok"))
		}
		fmt.Fprintln(a.Out, addr)
		return nil
	}
	srv.AddCommand(ipCmd)

	addServerCreate(a, srv)
	addServerDelete(a, srv)
	addServerUpdate(a, srv)
	addServerPower(a, srv)
	addServerMaintenance(a, srv)
	addServerSSH(a, srv)
	addServerWait(a, srv)
	addServerMetrics(a, srv)
}

func pagesFor(all bool) int {
	if all {
		return 0
	}
	return 1
}

type serverDoc struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Hostname    *string `json:"hostname"`
	Status      string  `json:"status"`
	IPv4        *string `json:"ipv4"`
	IPv6        *string `json:"ipv6"`
	Provisioned bool    `json:"provisioned"`
	Region      string  `json:"region"`
	ProductID   int64   `json:"product_id"`
	Live        *struct {
		PowerState string  `json:"power_state"`
		Locked     bool    `json:"locked"`
		IPv4       *string `json:"ipv4"`
		IPv6Net    *string `json:"ipv6_network"`
	} `json:"live"`
	raw json.RawMessage
}

func (a *App) fetchServerByID(ctx context.Context, id string) (*serverDoc, error) {
	c, err := a.Client()
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(ctx, api.Request{Method: "GET", Path: "/cloud/servers/" + id})
	if err != nil {
		return nil, err
	}
	var s serverDoc
	if err := json.Unmarshal(resp.Data, &s); err != nil {
		return nil, err
	}
	s.raw = resp.Data
	return &s, nil
}

func (a *App) fetchServer(ctx context.Context, arg string) (*serverDoc, error) {
	id, err := a.resolveServer(ctx, arg)
	if err != nil {
		return nil, err
	}
	return a.fetchServerByID(ctx, id)
}

func serverAddr(s *serverDoc, v6 bool) string {
	str := func(p *string) string {
		if p == nil {
			return ""
		}
		return *p
	}
	if v6 {
		n := str(s.IPv6)
		if s.Live != nil && str(s.Live.IPv6Net) != "" {
			n = str(s.Live.IPv6Net)
		}
		if n == "" {
			return ""
		}
		n = strings.SplitN(n, "/", 2)[0]
		if strings.HasSuffix(n, "::") {
			n += "1"
		}
		return n
	}
	if s.Live != nil && str(s.Live.IPv4) != "" {
		return str(s.Live.IPv4)
	}
	return str(s.IPv4)
}

func (s *serverDoc) state() string {
	if s.Live == nil {
		if !s.Provisioned {
			return i18n.S("provisioning", "kuruluyor")
		}
		return s.Status
	}
	st := s.Live.PowerState
	if s.Live.Locked {
		st += ", " + i18n.S("busy", "meşgul")
	}
	return st
}

// waitServerIdle waits until no action runs on the server (live.locked=false).
func (a *App) waitServerIdle(ctx context.Context, id string, timeout time.Duration) error {
	return a.waitServer(ctx, id, "", timeout)
}

// waitServer waits for provisioning, an idle server and (optionally) a power state.
func (a *App) waitServer(ctx context.Context, id, power string, timeout time.Duration) error {
	label := i18n.F("server %s", "sunucu %s", id)
	err := a.WaitFor(ctx, label, timeout, func() (bool, string, error) {
		s, err := a.fetchServerByID(ctx, id)
		if err != nil {
			if e, ok := api.AsError(err); ok && (e.Status >= 500 || e.Status == 429) {
				return false, e.Code, nil
			}
			return false, "", err
		}
		if s.Status == "cancelled" && power != "" {
			return false, "", fmt.Errorf("%s", i18n.S("server was cancelled", "sunucu iptal edildi"))
		}
		if !s.Provisioned || s.Live == nil {
			return false, s.state(), nil
		}
		if s.Live.Locked {
			return false, s.state(), nil
		}
		if power != "" && s.Live.PowerState != power {
			return false, s.state(), nil
		}
		return true, s.state(), nil
	})
	if err == nil {
		a.Info("%s", i18n.F("Server %s is ready.", "Sunucu %s hazır.", id))
	}
	return err
}

/* ------------------------------------------------------------------ create */

type product struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Plan       string  `json:"plan"`
	Region     string  `json:"region"`
	ServerType string  `json:"server_type"`
	Location   string  `json:"location"`
	Available  bool    `json:"available"`
	Cores      float64 `json:"cores"`
	MemoryGB   float64 `json:"memory_gb"`
	DiskGB     float64 `json:"disk_gb"`
	Prices     []struct {
		Cycle    string  `json:"cycle"`
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"prices"`
}

// where is the region shown to users (older APIs only had the location code).
func (p product) where() string {
	if p.Region != "" {
		return p.Region
	}
	return p.Location
}

var slugRe = regexp.MustCompile(`[^a-z0-9]+`)

// Slug folds Turkish letters and punctuation: "Gyro Bulut Pro 2-8" -> "gyro-bulut-pro-2-8".
func Slug(s string) string {
	s = strings.NewReplacer("İ", "i", "I", "i", "ı", "i", "Ş", "s", "ş", "s", "Ğ", "g", "ğ", "g", "Ü", "u", "ü", "u",
		"Ö", "o", "ö", "o", "Ç", "c", "ç", "c").Replace(s)
	return strings.Trim(slugRe.ReplaceAllString(strings.ToLower(s), "-"), "-")
}

// PlanSlug is the short plan name shown to users: "Gyro Bulut Pro 2-8" -> "pro-2-8".
func PlanSlug(name string) string {
	s := Slug(name)
	for _, p := range []string{"gyro-bulut-", "gyro-cloud-", "gyro-"} {
		s = strings.TrimPrefix(s, p)
	}
	return s
}

func (a *App) resolveLocation(ctx context.Context, arg string) (string, error) {
	if arg == "" {
		return "", nil
	}
	c, err := a.Client()
	if err != nil {
		return "", err
	}
	var locs []struct {
		Region, Code, City, Country string
	}
	if _, err := c.Get(ctx, "/cloud/locations", nil, &locs); err != nil {
		return "", err
	}
	var names []string
	for _, l := range locs {
		if strings.EqualFold(l.Region, arg) || strings.EqualFold(l.Code, arg) || Slug(l.City) == Slug(arg) {
			if l.Region != "" {
				return l.Region, nil
			}
			return l.Code, nil
		}
		if l.Region != "" {
			names = append(names, l.Region+" ("+l.City+")")
		} else {
			names = append(names, l.Code+" ("+l.City+")")
		}
	}
	return "", &ExitErr{ExitUsage, fmt.Errorf("%s %q; %s: %s", i18n.S("unknown location", "bilinmeyen lokasyon"), arg, i18n.S("choose", "seçenekler"), strings.Join(names, ", "))}
}

// resolvePlan finds the product for --plan (id, server type or plan slug) and --location.
func (a *App) resolvePlan(ctx context.Context, plan, loc string) (*product, error) {
	c, err := a.Client()
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	if loc != "" {
		q.Set("location", loc)
	}
	items, err := c.All(ctx, "/cloud/products", q, "", 0)
	if err != nil {
		return nil, err
	}
	var all, hits []product
	for _, it := range items {
		var p product
		if json.Unmarshal(it, &p) != nil {
			continue
		}
		all = append(all, p)
		key := strings.ToLower(plan)
		if fmt.Sprint(p.ID) == plan || strings.EqualFold(p.ServerType, plan) || strings.EqualFold(p.Plan, plan) || PlanSlug(p.Name) == key || Slug(p.Name) == key {
			hits = append(hits, p)
		}
	}
	if len(hits) > 1 {
		var avail []product
		for _, h := range hits {
			if h.Available {
				avail = append(avail, h)
			}
		}
		if len(avail) > 0 {
			hits = avail
		}
	}
	switch len(hits) {
	case 1:
		if !hits[0].Available {
			return nil, &ExitErr{ExitConflict, fmt.Errorf("%s: %s %s", i18n.S("temporarily out of stock", "geçici olarak stokta yok"), PlanSlug(hits[0].Name), hits[0].where())}
		}
		return &hits[0], nil
	case 0:
		seen := map[string]bool{}
		var names []string
		for _, p := range all {
			if n := PlanSlug(p.Name); !seen[n] {
				seen[n] = true
				names = append(names, n)
			}
		}
		sort.Strings(names)
		return nil, &ExitErr{ExitUsage, fmt.Errorf("%s %q (%s). %s: %s", i18n.S("no plan", "plan yok"), plan, locText(loc), i18n.S("available", "mevcut"), strings.Join(names, ", "))}
	}
	var opts []string
	for _, h := range hits {
		opts = append(opts, fmt.Sprintf("%d=%s/%s", h.ID, PlanSlug(h.Name), h.where()))
	}
	return nil, &ExitErr{ExitUsage, fmt.Errorf("%s %q: %s (%s)", i18n.S("several products match", "birden çok ürün uyuyor"), plan, strings.Join(opts, ", "),
		i18n.S("add --location or use the product id", "--location ekleyin ya da ürün kimliğini kullanın"))}
}

func locText(l string) string {
	if l == "" {
		return i18n.S("any location", "tüm lokasyonlar")
	}
	return l
}

// resolveImage checks an image slug against GET /cloud/images; with region / product (API 2.10.0+) only the images
// offered there (older APIs ignore the filters and return the global list).
func (a *App) resolveImage(ctx context.Context, img, region string, productID int64) (string, error) {
	c, err := a.Client()
	if err != nil {
		return "", err
	}
	q := url.Values{}
	if region != "" {
		q.Set("region", region)
	}
	if productID > 0 {
		q.Set("product_id", fmt.Sprint(productID))
	}
	var imgs []struct{ Slug, Name string }
	if _, err := c.Get(ctx, "/cloud/images", q, &imgs); err != nil {
		return "", err
	}
	var slugs []string
	for _, i := range imgs {
		if strings.EqualFold(i.Slug, img) || Slug(i.Name) == Slug(img) {
			return i.Slug, nil
		}
		slugs = append(slugs, i.Slug)
	}
	return "", &ExitErr{ExitUsage, fmt.Errorf("%s %q: %s", i18n.S("unknown image", "bilinmeyen imaj"), img, strings.Join(slugs, ", "))}
}

// resolveSSHKey accepts an id, a name, a fingerprint or a path to a .pub file (uploaded when new).
func (a *App) resolveSSHKey(ctx context.Context, arg string, dryRun bool) (int64, error) {
	c, err := a.Client()
	if err != nil {
		return 0, err
	}
	items, err := c.All(ctx, "/cloud/ssh-keys", nil, "", 0)
	if err != nil {
		return 0, err
	}
	type key struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Fingerprint string `json:"fingerprint"`
		PublicKey   string `json:"public_key"`
	}
	var keys []key
	for _, it := range items {
		var k key
		if json.Unmarshal(it, &k) == nil {
			keys = append(keys, k)
		}
	}
	pubFromFile := ""
	if b, err := os.ReadFile(expandHome(arg)); err == nil {
		if t := strings.TrimSpace(string(b)); strings.HasPrefix(t, "ssh-") || strings.HasPrefix(t, "ecdsa-") || strings.HasPrefix(t, "sk-") {
			pubFromFile = t
		}
	}
	for _, k := range keys {
		if fmt.Sprint(k.ID) == arg || strings.EqualFold(k.Name, arg) || k.Fingerprint == arg {
			return k.ID, nil
		}
		if pubFromFile != "" && sameKey(k.PublicKey, pubFromFile) {
			return k.ID, nil
		}
	}
	if pubFromFile == "" {
		return 0, &ExitErr{ExitNotFound, fmt.Errorf("%s %q (%s)", i18n.S("SSH key not found", "SSH anahtarı bulunamadı"), arg,
			i18n.S("give an id, a name or a path to a .pub file", "kimlik, ad ya da .pub dosya yolu verin"))}
	}
	if dryRun {
		a.Info("%s", i18n.F("(dry run) %s would be uploaded as a new SSH key", "(deneme) %s yeni SSH anahtarı olarak yüklenecekti", arg))
		return 0, nil
	}
	name := strings.TrimSuffix(filepath.Base(arg), ".pub")
	if f := strings.Fields(pubFromFile); len(f) >= 3 {
		name = f[2]
	}
	resp, err := c.Do(ctx, api.Request{Method: "POST", Path: "/cloud/ssh-keys", Body: map[string]any{"name": name, "public_key": pubFromFile}})
	if err != nil {
		return 0, err
	}
	var k key
	_ = json.Unmarshal(resp.Data, &k)
	a.Info("%s", i18n.F("Uploaded SSH key %q (id %d).", "SSH anahtarı %q yüklendi (kimlik %d).", name, k.ID))
	return k.ID, nil
}

func sameKey(a, b string) bool {
	fa, fb := strings.Fields(a), strings.Fields(b)
	return len(fa) >= 2 && len(fb) >= 2 && fa[0] == fb[0] && fa[1] == fb[1]
}

// appField finds the request body field used for one-click apps (API v2) if the spec has one.
func appField(op *spec.Op) string {
	for _, n := range []string{"app", "app_slug", "application", "apps"} {
		if _, ok := op.Field(n); ok {
			return n
		}
	}
	return ""
}

func addServerCreate(a *App, srv *cobra.Command) {
	var plan, loc, region, image, hostname, cycle, userData, app, data, idem string
	var keys, set []string
	var wait, dry bool
	var waitFor time.Duration
	cmd := &cobra.Command{
		Use:   "create [HOSTNAME]",
		Short: i18n.S("Order a new cloud server (paid from your balance)", "Yeni bulut sunucu siparişi (bakiyeden ödenir)"),
		Long: i18n.S(`Order a cloud server. The price is shown first (a quote from the API) and you are asked to confirm;
the order is paid from your account balance, never from a card. An idempotency key is attached
automatically, so a retried command never orders twice.

--plan accepts a plan name (standart-2-4, pro-2-8 ...) or a product id (see 'gyro plan list').
--region (or --location) accepts a region (eu-de-1, tr-ist-1, see 'gyro location list') or a city (falkenstein, istanbul).
In tr-ist-1 (Istanbul) up to 10 --ssh-key values and --user-data go with cloud-init at creation (SSH in about 25 s).`,
			`Bulut sunucu siparişi verir. Önce API'den fiyat teklifi gösterilir ve onay istenir; ödeme kart değil
hesap bakiyesinden yapılır. Komuta otomatik tekrar güvenliği anahtarı eklenir, tekrar denenen komut
iki kez sipariş vermez.

--plan: plan adı (standart-2-4, pro-2-8 ...) ya da ürün kimliği ('gyro plan list').
--region (ya da --location): bölge (eu-de-1, tr-ist-1, 'gyro location list') ya da şehir (falkenstein, istanbul).
tr-ist-1'de (İstanbul) en çok 10 --ssh-key ve --user-data oluşturmada cloud-init ile kurulur (SSH ~25 sn).`),
		Example: `  gyro server create web-1 --plan pro-2-8 --location falkenstein --image ubuntu-24.04 --ssh-key mykey --wait
  gyro server create --plan standart-2-4 --location eu-de-1 --image debian-12 --ssh-key ~/.ssh/id_ed25519.pub --dry-run
  gyro server create db-1 --plan standart-2-4 --location eu-de-2 --image ubuntu-24.04 --user-data @cloud-init.yaml --yes -o json
  gyro server create ist-1 --plan standart-2-4 --region tr-ist-1 --image ubuntu-24.04 --ssh-key mykey --wait`,
		Args: cobra.MaximumNArgs(1),
	}
	f := cmd.Flags()
	f.StringVar(&plan, "plan", "", i18n.S("plan name or product id (required)", "plan adı ya da ürün kimliği (zorunlu)"))
	f.StringVar(&loc, "location", "", i18n.S("region (eu-de-1) or city", "bölge (eu-de-1) ya da şehir"))
	f.StringVar(&region, "region", "", i18n.S("region (eu-de-1, tr-ist-1); same as --location", "bölge (eu-de-1, tr-ist-1); --location ile aynı"))
	f.StringVar(&image, "image", "", i18n.S("OS image slug, e.g. ubuntu-24.04 (required)", "işletim sistemi imajı, ör. ubuntu-24.04 (zorunlu)"))
	f.StringVar(&hostname, "hostname", "", i18n.S("host name (or give it as the argument)", "sunucu adı (ya da argüman olarak verin)"))
	f.StringSliceVar(&keys, "ssh-key", nil, i18n.S("SSH key id, name or .pub file (uploaded if new)", "SSH anahtarı kimliği, adı ya da .pub dosyası (yeniyse yüklenir)"))
	f.StringVar(&cycle, "cycle", "monthly", "monthly|quarterly|semiannually|annually")
	f.StringVar(&userData, "user-data", "", i18n.S("cloud-init user data, inline or @file", "cloud-init verisi, metin ya da @dosya"))
	f.StringVar(&app, "app", "", i18n.S("one-click app to install (needs API v2)", "kurulacak tek tık uygulama (API v2 gerekir)"))
	f.BoolVar(&dry, "dry-run", false, i18n.S("only show the price quote", "yalnız fiyat teklifini göster"))
	f.BoolVar(&wait, "wait", false, i18n.S("wait until the server is running", "sunucu çalışana kadar bekle"))
	f.DurationVar(&waitFor, "wait-timeout", 15*time.Minute, i18n.S("maximum time for --wait", "--wait için en uzun süre"))
	f.StringVar(&idem, "idempotency-key", "", i18n.S("idempotency key (default: random)", "tekrar güvenliği anahtarı (varsayılan: rastgele)"))
	f.StringVar(&data, "data", "", i18n.S("extra request body JSON / @file", "ek istek gövdesi JSON / @dosya"))
	f.StringArrayVar(&set, "set", nil, i18n.S("extra body field key=value (repeatable)", "ek gövde alanı anahtar=değer (tekrarlanabilir)"))
	_ = cmd.MarkFlagRequired("plan")
	_ = cmd.MarkFlagRequired("image")
	_ = cmd.RegisterFlagCompletionFunc("cycle", fixedCompletion("monthly", "quarterly", "semiannually", "annually"))
	_ = cmd.RegisterFlagCompletionFunc("location", completeField(a, "/cloud/locations", "region"))
	_ = cmd.RegisterFlagCompletionFunc("region", completeField(a, "/cloud/locations", "region"))
	_ = cmd.RegisterFlagCompletionFunc("image", completeField(a, "/cloud/images", "slug"))
	_ = cmd.RegisterFlagCompletionFunc("ssh-key", completeField(a, "/cloud/ssh-keys", "name"))
	_ = cmd.RegisterFlagCompletionFunc("plan", func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
		return []string{"standart-1-1", "standart-2-4", "standart-4-8", "pro-2-8", "pro-4-16", "pro-8-32"}, cobra.ShellCompDirectiveNoFileComp
	})
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		ctx, cancel := a.Ctx()
		defer cancel()
		c, err := a.Client()
		if err != nil {
			return err
		}
		if len(args) == 1 {
			if hostname != "" && hostname != args[0] {
				return &ExitErr{ExitUsage, fmt.Errorf("%s", i18n.S("give the host name once (argument or --hostname)", "sunucu adını bir kez verin (argüman ya da --hostname)"))}
			}
			hostname = args[0]
		}
		op := spec.ByID("createServer")
		if region != "" {
			if loc != "" && !strings.EqualFold(loc, region) {
				return &ExitErr{ExitUsage, fmt.Errorf("%s", i18n.S("give the region once (--region or --location)", "bölgeyi bir kez verin (--region ya da --location)"))}
			}
			loc = region
		}
		code, err := a.resolveLocation(ctx, loc)
		if err != nil {
			return err
		}
		prod, err := a.resolvePlan(ctx, plan, code)
		if err != nil {
			return err
		}
		img, err := a.resolveImage(ctx, image, code, prod.ID)
		if err != nil {
			return err
		}
		body, err := a.ReadData(data)
		if err != nil {
			return err
		}
		if body == nil {
			body = map[string]any{}
		}
		body["product_id"] = prod.ID
		body["image"] = img
		body["cycle"] = cycle
		if hostname != "" {
			body["hostname"] = hostname
		}
		if userData != "" {
			ud, err := a.ReadValue(userData)
			if err != nil {
				return err
			}
			body["user_data"] = ud
		}
		var keyIDs []int64
		for _, k := range keys {
			id, err := a.resolveSSHKey(ctx, k, dry)
			if err != nil {
				return err
			}
			if id > 0 {
				keyIDs = append(keyIDs, id)
			}
		}
		if len(keyIDs) > 0 {
			body["ssh_keys"] = keyIDs
		}
		if app != "" {
			fld := ""
			if op != nil {
				fld = appField(op)
			}
			if fld == "" {
				return &ExitErr{ExitUsage, fmt.Errorf("%s", i18n.S("--app: one-click apps are not available in this API version yet (run 'gyro apps list' when it is)",
					"--app: tek tık uygulamalar bu API sürümünde henüz yok (geldiğinde 'gyro apps list')"))}
			}
			body[fld] = app
		}
		extra, err := ParseKV(set)
		if err != nil {
			return err
		}
		for k, v := range extra {
			body[k] = v
		}

		// quote
		qb := map[string]any{}
		for k, v := range body {
			qb[k] = v
		}
		qb["dry_run"] = true
		qresp, err := c.Do(ctx, api.Request{Method: "POST", Path: "/cloud/servers", Body: qb})
		if err != nil {
			return err
		}
		var quote struct {
			Total   map[string]any `json:"total"`
			Tax     map[string]any `json:"tax"`
			Balance map[string]any `json:"balance"`
			Cycle   string         `json:"cycle"`
			Prepaid *struct {
				Hourly float64 `json:"hourly_eur"`
				Cap    float64 `json:"monthly_cap_eur"`
			} `json:"prepaid"`
		}
		var wrapped struct {
			Quote *json.RawMessage `json:"quote"`
		}
		if json.Unmarshal(qresp.Data, &wrapped) == nil && wrapped.Quote != nil {
			_ = json.Unmarshal(*wrapped.Quote, &quote)
		} else {
			_ = json.Unmarshal(qresp.Data, &quote)
		}
		if quote.Prepaid != nil {
			cycle = "prepaid" // no billing cycle: printQuote skips the order price line
		}
		if dry {
			if a.Machine() || a.Opt.Query != "" {
				return a.PrintRaw(qresp.Data, nil, "")
			}
			printQuote(a, prod, img, hostname, quote.Total, quote.Tax, quote.Balance, cycle)
			if quote.Prepaid != nil {
				fmt.Fprintf(a.Err, "  %-12s %.4f EUR / %s (%s %.2f EUR)\n", i18n.S("Prepaid", "Ön ödemeli")+":", quote.Prepaid.Hourly, i18n.S("hour", "saat"), i18n.S("monthly at most", "aylık en çok"), quote.Prepaid.Cap)
			}
			a.Info("%s", i18n.S("Dry run: nothing was ordered.", "Deneme: sipariş verilmedi."))
			return nil
		}
		if !a.Machine() {
			printQuote(a, prod, img, hostname, quote.Total, quote.Tax, quote.Balance, cycle)
		}
		if quote.Prepaid != nil {
			// prepaid (e.g. tr-ist-1): no order total, usage is charged hourly from the balance
			if !a.Machine() {
				fmt.Fprintf(a.Err, "  %-12s %.4f EUR / %s (%s %.2f EUR)\n", i18n.S("Prepaid", "Ön ödemeli")+":", quote.Prepaid.Hourly, i18n.S("hour", "saat"), i18n.S("monthly at most", "aylık en çok"), quote.Prepaid.Cap)
			}
			if err := a.Confirm(i18n.F("Order this server? It is charged hourly from your balance (%.4f EUR/hour, at most %.2f EUR a month).", "Bu sunucu sipariş edilsin mi? Bakiyenizden saatlik düşülür (%.4f EUR/saat, ayda en çok %.2f EUR).", quote.Prepaid.Hourly, quote.Prepaid.Cap)); err != nil {
				return err
			}
		} else if err := a.Confirm(i18n.F("Order this server and pay %s from your balance?", "Bu sunucu sipariş edilip bakiyenizden %s ödensin mi?", output.Scalar(quote.Total))); err != nil {
			return err
		}
		body["confirm"] = true
		if idem == "" {
			idem = api.NewID()
		}
		resp, err := c.Do(ctx, api.Request{Method: "POST", Path: "/cloud/servers", Body: body, IdempotencyKey: idem})
		if err != nil {
			return err
		}
		var res struct {
			Server  struct{ ID int64 } `json:"server"`
			Payment struct {
				Status string `json:"status"`
				Error  string `json:"error"`
			} `json:"payment"`
			Order struct {
				ID     int64  `json:"id"`
				Number string `json:"number"`
			} `json:"order"`
		}
		_ = json.Unmarshal(resp.Data, &res)
		if resp.Replayed {
			a.Info("%s", i18n.S("(idempotent replay: this order already existed, nothing new was charged)", "(tekrar: bu sipariş zaten vardı, yeni ücret alınmadı)"))
		}
		if res.Payment.Status != "" && res.Payment.Status != "paid" {
			_ = a.PrintRaw(resp.Data, nil, "server.id")
			return &ExitErr{ExitPayment, fmt.Errorf("%s: %s %s", i18n.S("payment not completed", "ödeme tamamlanmadı"), res.Payment.Status, res.Payment.Error)}
		}
		id := fmt.Sprint(res.Server.ID)
		a.Info("%s", i18n.F("Order %s placed, server id %s (idempotency key %s).", "Sipariş %s verildi, sunucu kimliği %s (tekrar anahtarı %s).", res.Order.Number, id, idem))
		if wait && res.Server.ID > 0 {
			if err := a.waitServer(ctx, id, "running", waitFor); err != nil {
				return err
			}
			s, err := a.fetchServerByID(ctx, id)
			if err != nil {
				return err
			}
			if !a.Machine() && a.Opt.Query == "" {
				if ip := serverAddr(s, false); ip != "" {
					a.Info("%s", i18n.F("Connect: gyro server ssh %s   (ssh root@%s)", "Bağlanın: gyro server ssh %s   (ssh root@%s)", id, ip))
				}
			}
			return a.PrintRaw(s.raw, nil, "id")
		}
		if a.Machine() || a.Opt.Query != "" {
			return a.PrintRaw(resp.Data, nil, "server.id")
		}
		v, _ := output.Decode(resp.Data)
		if m, ok := v.(map[string]any); ok {
			if s, ok := m["server"].(map[string]any); ok {
				return a.Printer().Print(s, nil, "id")
			}
		}
		return a.PrintRaw(resp.Data, nil, "server.id")
	}
	srv.AddCommand(cmd)
}

func printQuote(a *App, p *product, img, host string, total, tax, balance map[string]any, cycle string) {
	w := a.Err
	fmt.Fprintln(w, i18n.S("Order summary", "Sipariş özeti"))
	fmt.Fprintf(w, "  %-12s %s (%s, %g vCPU / %g GB / %g GB)\n", i18n.S("Plan", "Plan")+":", PlanSlug(p.Name), p.where(), p.Cores, p.MemoryGB, p.DiskGB)
	fmt.Fprintf(w, "  %-12s %s\n", i18n.S("Image", "İmaj")+":", img)
	if host != "" {
		fmt.Fprintf(w, "  %-12s %s\n", i18n.S("Hostname", "Ad")+":", host)
	}
	if cycle != "prepaid" {
		fmt.Fprintf(w, "  %-12s %s / %s\n", i18n.S("Price", "Fiyat")+":", output.Scalar(total), cycle)
	}
	if tax != nil && cycle != "prepaid" {
		fmt.Fprintf(w, "  %-12s %s\n", i18n.S("Tax", "Vergi")+":", output.Scalar(tax))
	}
	if balance != nil {
		fmt.Fprintf(w, "  %-12s %s\n", i18n.S("Balance", "Bakiye")+":", output.Scalar(balance))
	}
}

func completeField(a *App, path, field string) func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	return func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
		c, err := a.Client()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		c.MaxRetries = 0
		items, err := c.All(ctx, path, nil, "", 2)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		var out []string
		for _, it := range items {
			v, _ := output.Decode(it)
			if m, ok := v.(map[string]any); ok {
				out = append(out, output.Scalar(m[field]))
			}
		}
		return out, cobra.ShellCompDirectiveNoFileComp
	}
}

/* ------------------------------------------------------------------ delete / update */

func addServerDelete(a *App, srv *cobra.Command) {
	cmd := &cobra.Command{
		Use: "delete <SERVER>", Aliases: []string{"rm", "destroy"},
		Short:             i18n.S("Delete (cancel) a server - its data is destroyed after the grace period", "Sunucuyu sil (iptal) - veri bekleme süresi sonunda silinir"),
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: completeServers(a),
	}
	var dry bool
	cmd.Flags().BoolVar(&dry, "dry-run", false, i18n.S("show what would be deleted", "neyin silineceğini göster"))
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		ctx, cancel := a.Ctx()
		defer cancel()
		s, err := a.fetchServer(ctx, args[0])
		if err != nil {
			return err
		}
		host := fmt.Sprint(s.ID)
		if s.Hostname != nil && *s.Hostname != "" {
			host = *s.Hostname
		}
		if dry {
			a.Info("%s", i18n.F("Dry run: server %d (%s, %s) would be cancelled and its data wiped. Nothing was changed.",
				"Deneme: sunucu %d (%s, %s) iptal edilip verisi silinecekti. Hiçbir şey değişmedi.", s.ID, host, serverAddr(s, false)))
			return nil
		}
		if err := a.ConfirmTyped(i18n.F("Delete server %d (%s, %s)? This cannot be undone.", "Sunucu %d (%s, %s) silinsin mi? Geri alınamaz.", s.ID, host, serverAddr(s, false)), host); err != nil {
			return err
		}
		c, _ := a.Client()
		resp, err := c.Do(ctx, api.Request{Method: "DELETE", Path: fmt.Sprintf("/cloud/servers/%d", s.ID), Query: url.Values{"confirm": {"true"}}})
		if err != nil {
			return err
		}
		var dr struct {
			Status       string `json:"status"`
			Cancellation string `json:"cancellation"`
			Message      string `json:"message"`
		}
		_ = json.Unmarshal(resp.Data, &dr)
		if dr.Message != "" {
			a.Info("%s", dr.Message)
		}
		a.Info("%s", i18n.F("Cancellation of server %d requested: it is powered off when processed and destroyed with its data after the grace period.",
			"Sunucu %d için iptal talebi oluşturuldu: işlendiğinde kapatılır, bekleme süresi sonunda verisiyle birlikte silinir.", s.ID))
		if a.Machine() || a.Opt.Query != "" {
			return a.PrintRaw(resp.Data, nil, "id")
		}
		return nil
	}
	srv.AddCommand(cmd)
}

func addServerUpdate(a *App, srv *cobra.Command) {
	markHandled("updateServer")
	var host string
	var autoRenew bool
	upd := &cobra.Command{
		Use:               "update <SERVER>",
		Short:             i18n.S("Change host name or auto-renew", "Ad ya da otomatik yenilemeyi değiştir"),
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: completeServers(a),
		Example:           "  gyro server update web-1 --hostname web-01.example.com\n  gyro server update 47466 --auto-renew=false",
	}
	upd.Flags().StringVar(&host, "hostname", "", i18n.S("new host name", "yeni sunucu adı"))
	upd.Flags().BoolVar(&autoRenew, "auto-renew", false, i18n.S("renew automatically from balance", "bakiyeden otomatik yenile"))
	run := func(cmd *cobra.Command, id string, body map[string]any) error {
		ctx, cancel := a.Ctx()
		defer cancel()
		sid, err := a.resolveServer(ctx, id)
		if err != nil {
			return err
		}
		if len(body) == 0 {
			return &ExitErr{ExitUsage, fmt.Errorf("%s", i18n.S("nothing to change", "değişecek bir şey yok"))}
		}
		c, _ := a.Client()
		resp, err := c.Do(ctx, api.Request{Method: "PATCH", Path: "/cloud/servers/" + sid, Body: body})
		if err != nil {
			return err
		}
		return a.PrintRaw(resp.Data, nil, "id")
	}
	upd.RunE = func(cmd *cobra.Command, args []string) error {
		body := map[string]any{}
		if cmd.Flags().Changed("hostname") {
			body["hostname"] = host
		}
		if cmd.Flags().Changed("auto-renew") {
			body["auto_renew"] = autoRenew
		}
		return run(cmd, args[0], body)
	}
	srv.AddCommand(upd)
	srv.AddCommand(&cobra.Command{
		Use: "rename <SERVER> <HOSTNAME>", Short: i18n.S("Change the host name", "Sunucu adını değiştir"),
		Args: cobra.ExactArgs(2), ValidArgsFunction: completeServers(a),
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd, args[0], map[string]any{"hostname": args[1]})
		},
	})
}

/* ------------------------------------------------------------------ actions */

// runAction posts /cloud/servers/{id}/actions/{action}.
func (a *App) runAction(cmd *cobra.Command, arg, action string, body map[string]any, confirm string, wait bool, power string, waitFor time.Duration) error {
	ctx, cancel := a.Ctx()
	defer cancel()
	id, err := a.resolveServer(ctx, arg)
	if err != nil {
		return err
	}
	if confirm != "" {
		if err := a.Confirm(fmt.Sprintf(confirm, id)); err != nil {
			return err
		}
		if body == nil {
			body = map[string]any{}
		}
		body["confirm"] = true
	}
	c, _ := a.Client()
	req := api.Request{Method: "POST", Path: "/cloud/servers/" + id + "/actions/" + action}
	if body != nil {
		req.Body = body
	}
	resp, err := c.Do(ctx, req)
	if err != nil {
		if e, ok := api.AsError(err); ok && e.Code == "server_locked" && wait {
			a.Info("%s", i18n.S("Server busy, waiting and retrying ...", "Sunucu meşgul, bekleyip tekrar deneniyor ..."))
			if werr := a.waitServerIdle(ctx, id, waitFor); werr != nil {
				return werr
			}
			resp, err = c.Do(ctx, req)
		}
		if err != nil {
			return err
		}
	}
	v, _ := output.Decode(resp.Data)
	m, _ := v.(map[string]any)
	if a.Machine() || a.Opt.Query != "" || (m != nil && m["root_password"] != nil) {
		if err := a.Printer().Print(v, nil, "server_id"); err != nil {
			return err
		}
		warnSecrets(a, resp.Data)
	} else if m != nil {
		msg := output.Scalar(m["message"])
		if msg == "-" {
			msg = i18n.F("%s accepted", "%s kabul edildi", action)
		}
		a.Info("%s", msg)
	}
	if wait {
		if power != "" {
			// give the provider a moment to start the transition before polling
			a.Sleep(2 * time.Second)
		}
		return a.waitServer(ctx, id, power, waitFor)
	}
	return nil
}

func addServerPower(a *App, srv *cobra.Command) {
	power := &cobra.Command{
		Use:   "power",
		Short: i18n.S("Power on, off, shutdown, reboot or reset a server", "Sunucuyu aç, kapat, ACPI kapat, yeniden başlat ya da resetle"),
		Args:  cobra.NoArgs,
		RunE:  func(cmd *cobra.Command, args []string) error { return cmd.Help() },
	}
	type pa struct {
		use, action, power, confirm string
		short                       [2]string
		aliases                     []string
	}
	for _, p := range []pa{
		{"on", "poweron", "running", "", [2]string{"Power on", "Aç"}, []string{"start"}},
		{"shutdown", "shutdown", "stopped", "", [2]string{"Graceful shutdown (ACPI)", "Düzgün kapat (ACPI)"}, []string{"stop"}},
		{"off", "poweroff", "stopped", i18n.S("Cut the power of server %s (like pulling the plug; unsaved data is lost)?", "Sunucu %s elektriği kesilsin mi (fişi çekmek gibi; kaydedilmemiş veri kaybolur)?"), [2]string{"Hard power off", "Elektriği kes"}, nil},
		{"reboot", "reboot", "running", "", [2]string{"Reboot (graceful)", "Yeniden başlat (düzgün)"}, []string{"restart"}},
		{"reset", "reset", "running", i18n.S("Hard reset server %s?", "Sunucu %s donanımsal olarak resetlensin mi?"), [2]string{"Hard reset", "Donanımsal reset"}, nil},
	} {
		p := p
		var wait bool
		var waitFor time.Duration
		c := &cobra.Command{
			Use: p.use + " <SERVER>", Aliases: p.aliases, Short: i18n.S(p.short[0], p.short[1]),
			Args: cobra.ExactArgs(1), ValidArgsFunction: completeServers(a),
			RunE: func(cmd *cobra.Command, args []string) error {
				return a.runAction(cmd, args[0], p.action, nil, p.confirm, wait, p.power, waitFor)
			},
		}
		c.Flags().BoolVar(&wait, "wait", false, i18n.F("wait until the server is %s", "sunucu %s olana kadar bekle", p.power))
		c.Flags().DurationVar(&waitFor, "wait-timeout", 10*time.Minute, i18n.S("maximum time for --wait", "--wait için en uzun süre"))
		power.AddCommand(c)
		markHandledCmd("server action " + p.action)
	}
	srv.AddCommand(power)
	// shortcuts like hcloud: gyro server reboot <id>
	for _, alias := range []struct{ name, sub string }{{"poweron", "on"}, {"poweroff", "off"}, {"shutdown", "shutdown"}, {"reboot", "reboot"}, {"reset", "reset"}} {
		alias := alias
		orig := findChild(power, alias.sub)
		short := &cobra.Command{
			Use: alias.name + " <SERVER>", Short: orig.Short, Hidden: true,
			Args: cobra.ExactArgs(1), ValidArgsFunction: completeServers(a),
			RunE: func(cmd *cobra.Command, args []string) error {
				orig.SetArgs(args)
				return orig.RunE(orig, args)
			},
		}
		short.Flags().AddFlagSet(orig.Flags())
		srv.AddCommand(short)
	}
}

func addServerMaintenance(a *App, srv *cobra.Command) {
	// rebuild
	var image string
	var wait bool
	var waitFor time.Duration
	rebuild := &cobra.Command{
		Use: "rebuild <SERVER>", Aliases: []string{"reinstall"},
		Short: i18n.S("Reinstall the OS - WIPES THE DISK", "İşletim sistemini yeniden kur - DİSK SİLİNİR"),
		Args:  cobra.ExactArgs(1), ValidArgsFunction: completeServers(a),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := a.Ctx()
			reg, pid := "", int64(0)
			if s, err := a.fetchServer(ctx, args[0]); err == nil && s != nil {
				reg, pid = s.Region, s.ProductID
			}
			img, err := a.resolveImage(ctx, image, reg, pid)
			cancel()
			if err != nil {
				return err
			}
			return a.runAction(cmd, args[0], "rebuild", map[string]any{"image": img},
				i18n.F("Reinstall server %%s with %s? ALL DATA ON THE DISK IS LOST.", "Sunucu %%s %s ile yeniden kurulsun mu? DİSKTEKİ TÜM VERİ SİLİNİR.", img), wait, "running", waitFor)
		},
	}
	rebuild.Flags().StringVar(&image, "image", "", i18n.S("OS image slug (required)", "imaj (zorunlu)"))
	_ = rebuild.MarkFlagRequired("image")
	_ = rebuild.RegisterFlagCompletionFunc("image", completeField(a, "/cloud/images", "slug"))
	rebuild.Flags().BoolVar(&wait, "wait", false, i18n.S("wait until the server runs again", "sunucu tekrar çalışana kadar bekle"))
	rebuild.Flags().DurationVar(&waitFor, "wait-timeout", 20*time.Minute, i18n.S("maximum time for --wait", "--wait için en uzun süre"))
	srv.AddCommand(rebuild)

	var wait2 bool
	resetPw := &cobra.Command{
		Use:   "reset-password <SERVER>",
		Short: i18n.S("Set a new random root password (shown once)", "Yeni rastgele root parolası (bir kez gösterilir)"),
		Args:  cobra.ExactArgs(1), ValidArgsFunction: completeServers(a),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.runAction(cmd, args[0], "reset_password", nil, i18n.S("Reset the root password of server %s?", "Sunucu %s root parolası sıfırlansın mı?"), wait2, "", 10*time.Minute)
		},
	}
	resetPw.Flags().BoolVar(&wait2, "wait", false, i18n.S("wait until done", "bitene kadar bekle"))
	srv.AddCommand(resetPw)

	rescue := &cobra.Command{Use: "rescue", Short: i18n.S("Rescue system (boot into a recovery Linux)", "Kurtarma sistemi (kurtarma Linux'u ile açılış)"),
		Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error { return cmd.Help() }}
	var w3 bool
	en := &cobra.Command{Use: "enable <SERVER>", Short: i18n.S("Enable rescue mode (prints a one-time root password)", "Kurtarma modunu aç (tek seferlik root parolası yazar)"),
		Args: cobra.ExactArgs(1), ValidArgsFunction: completeServers(a),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.runAction(cmd, args[0], "enable_rescue", nil, "", w3, "", 10*time.Minute)
		}}
	en.Flags().BoolVar(&w3, "wait", false, i18n.S("wait until done", "bitene kadar bekle"))
	dis := &cobra.Command{Use: "disable <SERVER>", Short: i18n.S("Disable rescue mode", "Kurtarma modunu kapat"),
		Args: cobra.ExactArgs(1), ValidArgsFunction: completeServers(a),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.runAction(cmd, args[0], "disable_rescue", nil, "", w3, "", 10*time.Minute)
		}}
	dis.Flags().BoolVar(&w3, "wait", false, i18n.S("wait until done", "bitene kadar bekle"))
	rescue.AddCommand(en, dis)
	srv.AddCommand(rescue)
	markHandledCmd("server action rebuild", "server action reset-password", "server action enable-rescue", "server action disable-rescue")
}

/* ------------------------------------------------------------------ ssh / wait / metrics */

func addServerSSH(a *App, srv *cobra.Command) {
	var user, identity string
	var port int
	var v6 bool
	cmd := &cobra.Command{
		Use:   "ssh <SERVER> [-- SSH-ARGS/COMMAND...]",
		Short: i18n.S("Open an SSH session to a server", "Sunucuya SSH bağlantısı aç"),
		Long: i18n.S(`Looks up the server's IP and runs your local ssh client. Everything after "--" is passed to ssh,
so you can run a command remotely. Defaults come from the profile (ssh_user, ssh_identity).`,
			`Sunucunun IP adresini bulur ve yerel ssh istemcinizi çalıştırır. "--" sonrasındaki her şey ssh'a geçer,
uzaktan komut çalıştırabilirsiniz. Varsayılanlar profilden gelir (ssh_user, ssh_identity).`),
		Example:           "  gyro server ssh web-1\n  gyro server ssh 47466 -i ~/.ssh/id_ed25519 -- uptime\n  gyro server ssh web-1 --ipv6 -- -L 8080:localhost:80",
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: completeServers(a),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := a.Ctx()
			s, err := a.fetchServer(ctx, args[0])
			cancel()
			if err != nil {
				return err
			}
			addr := serverAddr(s, v6)
			if addr == "" {
				return fmt.Errorf("%s", i18n.S("the server has no IP address yet (still provisioning?)", "sunucunun henüz IP adresi yok (kurulum sürüyor mu?)"))
			}
			if user == "" {
				user = a.Profile().SSHUser
			}
			if user == "" {
				user = "root"
			}
			if identity == "" {
				identity = a.Profile().SSHIdentity
			}
			sshArgs := []string{}
			if identity != "" {
				sshArgs = append(sshArgs, "-i", expandHome(identity))
			}
			if port != 22 && port != 0 {
				sshArgs = append(sshArgs, "-p", fmt.Sprint(port))
			}
			sshArgs = append(sshArgs, user+"@"+addr)
			sshArgs = append(sshArgs, args[1:]...)
			bin := "ssh"
			if runtime.GOOS == "windows" {
				bin = "ssh.exe"
			}
			if a.Opt.Debug {
				fmt.Fprintf(a.Err, "+ %s %s\n", bin, strings.Join(sshArgs, " "))
			}
			return a.Exec(bin, sshArgs)
		},
	}
	cmd.Flags().StringVarP(&user, "user", "u", "", i18n.S("remote user (default root)", "uzak kullanıcı (varsayılan root)"))
	cmd.Flags().StringVarP(&identity, "identity", "i", "", i18n.S("private key file", "özel anahtar dosyası"))
	cmd.Flags().IntVar(&port, "port", 22, i18n.S("SSH port", "SSH portu"))
	cmd.Flags().BoolVarP(&v6, "ipv6", "6", false, i18n.S("connect over IPv6", "IPv6 ile bağlan"))
	srv.AddCommand(cmd)
}

func addServerWait(a *App, srv *cobra.Command) {
	var power string
	var timeout time.Duration
	cmd := &cobra.Command{
		Use:               "wait <SERVER>",
		Short:             i18n.S("Wait until a server is provisioned and idle", "Sunucu kurulup boşa çıkana kadar bekle"),
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: completeServers(a),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := a.Ctx()
			defer cancel()
			id, err := a.resolveServer(ctx, args[0])
			if err != nil {
				return err
			}
			return a.waitServer(ctx, id, power, timeout)
		},
	}
	cmd.Flags().StringVar(&power, "power-state", "", "running|stopped")
	cmd.Flags().DurationVar(&timeout, "timeout-after", 15*time.Minute, i18n.S("give up after", "şu süreden sonra vazgeç"))
	srv.AddCommand(cmd)
}

func addServerMetrics(a *App, srv *cobra.Command) {
	var period string
	cmd := &cobra.Command{
		Use:               "metrics <SERVER>",
		Short:             i18n.S("CPU, disk and network metrics (summary)", "CPU, disk ve ağ ölçümleri (özet)"),
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: completeServers(a),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := a.Ctx()
			defer cancel()
			id, err := a.resolveServer(ctx, args[0])
			if err != nil {
				return err
			}
			c, _ := a.Client()
			resp, err := c.Do(ctx, api.Request{Method: "GET", Path: "/cloud/servers/" + id + "/metrics", Query: url.Values{"period": {period}}})
			if err != nil {
				return err
			}
			if a.Machine() || a.Opt.Query != "" || a.Printer().Format == "text" {
				return a.PrintRaw(resp.Data, nil, "")
			}
			var m struct {
				Period string `json:"period"`
				Series map[string]struct {
					Unit  string                 `json:"unit"`
					Lines map[string][][]float64 `json:"lines"`
				} `json:"series"`
			}
			if err := json.Unmarshal(resp.Data, &m); err != nil {
				return a.PrintRaw(resp.Data, nil, "")
			}
			var rows []any
			names := make([]string, 0, len(m.Series))
			for k := range m.Series {
				names = append(names, k)
			}
			sort.Strings(names)
			for _, name := range names {
				s := m.Series[name]
				lines := make([]string, 0, len(s.Lines))
				for k := range s.Lines {
					lines = append(lines, k)
				}
				sort.Strings(lines)
				for _, ln := range lines {
					pts := s.Lines[ln]
					if len(pts) == 0 {
						continue
					}
					mn, mx, sum := pts[0][1], pts[0][1], 0.0
					for _, p := range pts {
						if len(p) < 2 {
							continue
						}
						mn, mx, sum = minF(mn, p[1]), maxF(mx, p[1]), sum+p[1]
					}
					last := pts[len(pts)-1]
					rows = append(rows, map[string]any{"series": name, "line": ln, "unit": s.Unit,
						"last": round2(last[len(last)-1]), "min": round2(mn), "max": round2(mx), "avg": round2(sum / float64(len(pts))), "points": float64(len(pts))})
				}
			}
			cols := []output.Column{{Header: "series", Path: "series"}, {Header: "line", Path: "line"}, {Header: "unit", Path: "unit"},
				{Header: "last", Path: "last"}, {Header: "avg", Path: "avg"}, {Header: "min", Path: "min"}, {Header: "max", Path: "max"}, {Header: "points", Path: "points"}}
			return a.Printer().Print(rows, cols, "series")
		},
	}
	cmd.Flags().StringVar(&period, "period", "24h", "24h|7d|30d")
	_ = cmd.RegisterFlagCompletionFunc("period", fixedCompletion("24h", "7d", "30d"))
	srv.AddCommand(cmd)
}

func minF(a, b float64) float64 {
	if b < a {
		return b
	}
	return a
}
func maxF(a, b float64) float64 {
	if b > a {
		return b
	}
	return a
}
func round2(f float64) float64 { return float64(int64(f*100+0.5*sign(f))) / 100 }
func sign(f float64) float64 {
	if f < 0 {
		return -1
	}
	return 1
}
