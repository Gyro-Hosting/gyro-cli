package cli

// VPC V2 (API 2.34.0): a container environment in a private network (VPC) of tr-ist-1, and the source allow list of
// VPC gateway port forwards. API 2.38.0: serverless environments too; their connection is charged (GET .../vpc "pricing"),
// attach shows the price and asks (--yes skips), then sends "confirm": true. Node environments stay free.
//   gyro container vpc get <ENV>
//   gyro container vpc attach <ENV> --vpc <VPC> --subnet <SUBNET> [--expose api:8080 ...] [--wait]
//   gyro container vpc expose <ENV> --expose api:8080 ... | --clear
//   gyro container vpc detach <ENV> [--wait]
//   gyro vpc port-forward update <VPC> <PF> --source 203.0.113.10 ... | --any
// (gyro vpc port-forward add takes --source too: alias of the generated --source-cidrs flag.)

import (
	"context"
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
	opGetContainerVpc       = "getContainerVpc"
	opAttachContainerVpc    = "attachContainerVpc"
	opSetContainerVpcExpose = "setContainerVpcExpose"
	opDetachContainerVpc    = "detachContainerVpc"
	opUpdateVpcPortForward  = "updateVpcPortForward"
)

type containerVpcDoc struct {
	Attached    bool   `json:"attached"`
	Supported   *bool  `json:"supported"`
	State       string `json:"state"`
	SourceRange string `json:"source_range"`
	EndpointIP  string `json:"endpoint_ip"`
	Error       string `json:"error"`
}

// vpcFrom reads the attachment from {"vpc": {...}} (writes) or the fields at the top level (GET).
func vpcFrom(raw json.RawMessage) (*containerVpcDoc, json.RawMessage) {
	var w struct {
		Vpc json.RawMessage `json:"vpc"`
	}
	if json.Unmarshal(raw, &w) == nil && len(w.Vpc) > 2 {
		var d containerVpcDoc
		_ = json.Unmarshal(w.Vpc, &d)
		return &d, w.Vpc
	}
	var d containerVpcDoc
	_ = json.Unmarshal(raw, &d)
	return &d, raw
}

// vpcPricingDoc is the "pricing" block of GET .../vpc (API 2.38.0); nil on older APIs.
type vpcPricingDoc struct {
	Charged         bool    `json:"charged"`
	MonthlyEUR      float64 `json:"monthly_eur"`
	MonthlyText     string  `json:"monthly_text"`
	HourlyText      string  `json:"hourly_text"`
	ListMonthlyText string  `json:"list_monthly_text"`
	MinimumHours    int     `json:"minimum_hours"`
}

func vpcPricing(raw json.RawMessage) (*vpcPricingDoc, bool) {
	var g struct {
		Attached bool           `json:"attached"`
		Pricing  *vpcPricingDoc `json:"pricing"`
	}
	_, v := vpcFrom(raw)
	_ = json.Unmarshal(v, &g)
	if g.Pricing != nil && g.Pricing.MonthlyText == "" {
		g.Pricing.MonthlyText = strconv.FormatFloat(g.Pricing.MonthlyEUR, 'f', 2, 64) + " EUR"
	}
	return g.Pricing, g.Attached
}

// parseExpose turns "api:8080" values into [{service, port}]. Spaces, invisible characters and a trailing comma are tolerated.
func parseExpose(vals []string) ([]map[string]any, error) {
	out := []map[string]any{}
	seen := map[int]bool{}
	for _, v := range vals {
		for _, p := range strings.FieldsFunc(v, func(r rune) bool { return r == ',' || r == ';' || r == ' ' }) {
			p = strings.Map(func(r rune) rune {
				switch r {
				case '\u200b', '\u200c', '\u200d', '\ufeff', '\u2060', '\u00a0', '\t', '\r', '\n':
					return -1
				}
				return r
			}, p)
			if p == "" {
				continue
			}
			name, port, ok := strings.Cut(p, ":")
			n, err := strconv.Atoi(strings.TrimSpace(port))
			name = strings.ToLower(strings.TrimSpace(name))
			if !ok || name == "" || err != nil || n < 1 || n > 65535 {
				return nil, &ExitErr{Code: ExitUsage, Err: fmt.Errorf(i18n.S("--expose %q: expected service:port, e.g. api:8080", "--expose %q: servis:port bekleniyor, ör. api:8080"), p)}
			}
			if seen[n] {
				return nil, &ExitErr{Code: ExitUsage, Err: fmt.Errorf(i18n.S("--expose: port %d is used twice (each port once)", "--expose: %d portu iki kez kullanılmış (her port bir kez)"), n)}
			}
			seen[n] = true
			out = append(out, map[string]any{"service": name, "port": n})
		}
	}
	if len(out) > 20 {
		return nil, &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("--expose: at most 20 services", "--expose: en çok 20 servis"))}
	}
	return out, nil
}

// parseSources cleans --source values: comma/space separated, invisible characters removed, a single IP becomes /32.
func parseSources(vals []string) []string {
	out := []string{}
	seen := map[string]bool{}
	for _, v := range vals {
		for _, p := range strings.FieldsFunc(v, func(r rune) bool { return r == ',' || r == ';' || r == ' ' || r == '\t' || r == '\n' || r == '\r' }) {
			p = strings.Map(func(r rune) rune {
				switch r {
				case '\u200b', '\u200c', '\u200d', '\ufeff', '\u2060', '\u00a0':
					return -1
				}
				return r
			}, p)
			if p == "" {
				continue
			}
			if !strings.Contains(p, "/") {
				p += "/32"
			}
			if !seen[p] {
				seen[p] = true
				out = append(out, p)
			}
		}
	}
	return out
}

func envArg(s string) (string, error) {
	if !isNumeric(s) {
		return "", &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("the environment id is a number (gyro service list)", "ortam numarası sayı olmalı (gyro service list)"))}
	}
	return s, nil
}

// waitVpc polls GET .../vpc until the attachment is ready (attach) or off (detach). error state stops the wait.
func (a *App) waitVpc(ctx context.Context, c *api.Client, env string, wantOff bool, timeout time.Duration) (json.RawMessage, error) {
	var last json.RawMessage
	label := i18n.F("environment %s VPC", "ortam %s VPC", env)
	err := a.WaitFor(ctx, label, timeout, func() (bool, string, error) {
		resp, err := c.Do(ctx, api.Request{Method: "GET", Path: "/cloud/containers/" + env + "/vpc"})
		if err != nil {
			var ae *api.Error
			if wantOff && errors.As(err, &ae) && ae.Status == 404 {
				return true, "off", nil
			}
			return false, "", err
		}
		last = resp.Data
		d, _ := vpcFrom(resp.Data)
		if d.State == "error" {
			msg := d.Error
			if msg == "" {
				msg = "error"
			}
			return false, "", fmt.Errorf(i18n.S("VPC attachment failed: %s", "VPC bağı başarısız: %s"), msg)
		}
		if wantOff {
			return d.State == "off" || (d.State == "" && !d.Attached), d.State, nil
		}
		return d.State == "ready", d.State, nil
	})
	return last, err
}

func addContainerVpc(a *App, root *cobra.Command) {
	markHandled(opGetContainerVpc, opAttachContainerVpc, opSetContainerVpcExpose, opDetachContainerVpc, opUpdateVpcPortForward)
	grp := ensureGroup(root, []string{"container", "vpc"})
	grp.Short = i18n.S("Private network (VPC) of a container environment: fixed source range, VPC endpoint", "Kapsayıcı ortamının özel ağı (VPC): sabit kaynak aralığı, VPC uç noktası")
	delete(grp.Annotations, "autoShort")
	long := i18n.S(`Connects a container environment to one subnet of your private network (VPC) in tr-ist-1. The environment
gets a fixed VPC source range (/28) inside the subnet: connections from your containers to the VPC servers come
from it and it does not change when nodes are replaced, scaled or renewed. The first address is the VPC endpoint:
VPC servers reach the services you publish (--expose service:port) at endpoint:port. Never reachable from the
internet. The subnet must be /26 or larger.
Price: node environments (Mini, Small, Starter, Standard, Pro) free. Serverless environments: a connector dedicated
to the environment runs, so the connection is charged (2.99 EUR/month, billed hourly from the prepaid balance while
attached); attach shows the price in your currency and asks for confirmation (--yes accepts it).`,
		`Bir Kapsayıcı ortamını tr-ist-1'deki özel ağınızın (VPC) bir subnetine bağlar. Ortama subnet içinde sabit bir
VPC kaynak aralığı (/28) ayrılır: kapsayıcılarınızdan VPC sunucularına giden bağlantılar buradan gelir; düğüm
değişimi, ölçekleme ya da yenilemede değişmez. İlk adres VPC uç noktasıdır: VPC sunucuları yayınladığınız
servislere (--expose servis:port) uç_nokta:port ile ulaşır. İnternetten erişilemez. Subnet en az /26 olmalı.
Ücret: düğümlü ortamlarda (Mini, Küçük, Başlangıç, Standart, Pro) ücretsiz. Sunucusuz ortamlarda ortama ayrılmış
bir bağlayıcı çalıştığı için bağlantı ücretlidir (2,99 €/ay, bağlıyken saatlik, ön ödemeli bakiyeden); attach
fiyatı TL ve € olarak gösterir ve onay ister (--yes onaylar).`)
	grp.Long = long

	get := &cobra.Command{
		Use:     "get <ENVIRONMENT_ID>",
		Short:   i18n.S("Show the environment's VPC attachment (source range, endpoint, published services)", "Ortamın VPC bağını göster (kaynak aralığı, uç nokta, yayınlanan servisler)"),
		Long:    long,
		Example: "  gyro container vpc get 17000\n  gyro container vpc get 17000 -o json",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			env, err := envArg(args[0])
			if err != nil {
				return err
			}
			ctx, cancel := a.Ctx()
			defer cancel()
			c, err := a.Client()
			if err != nil {
				return err
			}
			resp, err := c.Do(ctx, api.Request{Method: "GET", Path: "/cloud/containers/" + env + "/vpc"})
			if err != nil {
				return err
			}
			return a.PrintRaw(resp.Data, nil, "source_range")
		},
	}

	var dry1, dry2, dry3, dry4 bool
	dryUse := i18n.S("show the request without changing anything", "hiçbir şeyi değiştirmeden isteği göster")
	var vpcID, subnetID string
	var expose []string
	var wait bool
	var waitFor time.Duration
	attach := &cobra.Command{
		Use:   "attach <ENVIRONMENT_ID> --subnet <SUBNET_ID> [--vpc <VPC_ID>] [--expose service:port ...]",
		Short: i18n.S("Attach the environment to a VPC subnet", "Ortamı bir VPC subnetine bağla"),
		Long:  long,
		Example: "  gyro container vpc attach 17000 --vpc vpc-1a2b3c4d --subnet sn-5e6f7a8b --expose api:8080 --wait\n" +
			"  gyro vpc subnet list vpc-1a2b3c4d    # subnet ids",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			env, err := envArg(args[0])
			if err != nil {
				return err
			}
			subnetID, vpcID = strings.TrimSpace(subnetID), strings.TrimSpace(vpcID)
			if subnetID == "" {
				return &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("--subnet is required (gyro vpc subnet list <VPC>)", "--subnet gerekli (gyro vpc subnet list <VPC>)"))}
			}
			body := map[string]any{"subnet_id": subnetID}
			if vpcID != "" {
				body["vpc_id"] = vpcID
			}
			if cmd.Flags().Changed("expose") {
				l, err := parseExpose(expose)
				if err != nil {
					return err
				}
				body["expose"] = l
			}
			if dry1 {
				return a.Printer().Print(map[string]any{"method": "PUT", "path": "/cloud/containers/" + env + "/vpc", "body": body}, nil, "")
			}
			ctx, cancel := a.Ctx()
			defer cancel()
			c, err := a.Client()
			if err != nil {
				return err
			}
			// API 2.38.0: a serverless environment's VPC connection is charged -> show the price, ask, send "confirm": true
			if pre, err := c.Do(ctx, api.Request{Method: "GET", Path: "/cloud/containers/" + env + "/vpc"}); err == nil {
				if p, attached := vpcPricing(pre.Data); p != nil && p.Charged && !attached {
					// the price is printed even without a terminal (--yes / CI logs show what was accepted)
					fmt.Fprintln(a.Err, i18n.F("Environment %s is serverless: its VPC connection costs %s per month (%s per hour while attached, from your prepaid balance; at least %d hour per attachment).",
						"%s ortamı sunucusuz: VPC bağlantısı aylık %s (bağlıyken saatlik %s, ön ödemeli bakiyeden; her bağlamada en az %d saat).",
						env, p.MonthlyText, p.HourlyText, max(1, p.MinimumHours)))
					if err := a.Confirm(i18n.S("Attach and accept the charge?", "Bağlansın ve ücret onaylansın mı?")); err != nil {
						return err
					}
					body["confirm"] = true
				}
			}
			resp, err := c.Do(ctx, api.Request{Method: "PUT", Path: "/cloud/containers/" + env + "/vpc", Body: body})
			if err != nil {
				return err
			}
			out := resp.Data
			if wait {
				if raw, err := a.waitVpc(ctx, c, env, false, waitFor); err != nil {
					return err
				} else if raw != nil {
					out = raw
				}
			}
			_, v := vpcFrom(out)
			return a.PrintRaw(v, nil, "source_range")
		},
	}
	attach.Flags().StringVar(&vpcID, "vpc", "", i18n.S("VPC id (gyro vpc list); checked against the subnet", "VPC kimliği (gyro vpc list); subnetle eşleşmesi denetlenir"))
	attach.Flags().StringVar(&subnetID, "subnet", "", i18n.S("subnet id (/26 or larger)", "subnet kimliği (en az /26)"))
	attach.Flags().StringArrayVar(&expose, "expose", nil, i18n.S("publish a service at the VPC endpoint: service:port (repeatable, at most 20)", "servisi VPC uç noktasında yayınla: servis:port (tekrarlanabilir, en çok 20)"))
	attach.Flags().BoolVar(&wait, "wait", false, i18n.S("wait until the nodes are in the subnet (state ready)", "düğümler subnete katılana kadar bekle (state ready)"))
	attach.Flags().BoolVar(&dry1, "dry-run", false, dryUse)
	attach.Flags().DurationVar(&waitFor, "wait-timeout", 10*time.Minute, i18n.S("maximum time for --wait", "--wait için en uzun süre"))

	var exp2 []string
	var clear bool
	exposeCmd := &cobra.Command{
		Use:     "expose <ENVIRONMENT_ID> --expose service:port ... | --clear",
		Short:   i18n.S("Replace the services published at the VPC endpoint", "VPC uç noktasında yayınlanan servisleri değiştir"),
		Long:    long,
		Example: "  gyro container vpc expose 17000 --expose api:8080 --expose web:80\n  gyro container vpc expose 17000 --clear",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			env, err := envArg(args[0])
			if err != nil {
				return err
			}
			if clear == (len(exp2) > 0) {
				return &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("give --expose service:port (the list replaces the old one) or --clear", "--expose servis:port (liste eskisinin yerine geçer) ya da --clear verin"))}
			}
			l := []map[string]any{}
			if !clear {
				if l, err = parseExpose(exp2); err != nil {
					return err
				}
			}
			body := map[string]any{"expose": l}
			if dry2 {
				return a.Printer().Print(map[string]any{"method": "PATCH", "path": "/cloud/containers/" + env + "/vpc", "body": body}, nil, "")
			}
			ctx, cancel := a.Ctx()
			defer cancel()
			c, err := a.Client()
			if err != nil {
				return err
			}
			resp, err := c.Do(ctx, api.Request{Method: "PATCH", Path: "/cloud/containers/" + env + "/vpc", Body: body})
			if err != nil {
				return err
			}
			_, v := vpcFrom(resp.Data)
			return a.PrintRaw(v, nil, "source_range")
		},
	}
	exposeCmd.Flags().StringArrayVar(&exp2, "expose", nil, i18n.S("service:port (repeatable; the list replaces the old one)", "servis:port (tekrarlanabilir; liste eskisinin yerine geçer)"))
	exposeCmd.Flags().BoolVar(&dry2, "dry-run", false, dryUse)
	exposeCmd.Flags().BoolVar(&clear, "clear", false, i18n.S("remove every published service", "yayınlanan tüm servisleri kaldır"))

	var wait3 bool
	var waitFor3 time.Duration
	detach := &cobra.Command{
		Use:     "detach <ENVIRONMENT_ID>",
		Short:   i18n.S("Detach the environment from its VPC (the /28 source range is released)", "Ortamın VPC bağını kaldır (/28 kaynak aralığı serbest kalır)"),
		Long:    long,
		Example: "  gyro container vpc detach 17000 --wait\n  gyro container vpc detach 17000 --yes",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			env, err := envArg(args[0])
			if err != nil {
				return err
			}
			if dry3 {
				return a.Printer().Print(map[string]any{"method": "DELETE", "path": "/cloud/containers/" + env + "/vpc", "body": map[string]any{"confirm": true}}, nil, "")
			}
			if err := a.Confirm(i18n.F("Detach environment %s from its VPC? VPC servers lose access to its published services and the source range is released.",
				"%s ortamının VPC bağı kaldırılsın mı? VPC sunucuları yayınlanan servislere erişemez ve kaynak aralığı serbest kalır.", env)); err != nil {
				return err
			}
			ctx, cancel := a.Ctx()
			defer cancel()
			c, err := a.Client()
			if err != nil {
				return err
			}
			resp, err := c.Do(ctx, api.Request{Method: "DELETE", Path: "/cloud/containers/" + env + "/vpc", Body: map[string]any{"confirm": true}})
			if err != nil {
				return err
			}
			out := resp.Data
			if wait3 {
				if raw, err := a.waitVpc(ctx, c, env, true, waitFor3); err != nil {
					return err
				} else if raw != nil {
					out = raw
				}
			}
			_, v := vpcFrom(out)
			return a.PrintRaw(v, nil, "state")
		},
	}
	detach.Flags().BoolVar(&wait3, "wait", false, i18n.S("wait until it is detached (state off)", "bağ kalkana kadar bekle (state off)"))
	detach.Flags().BoolVar(&dry3, "dry-run", false, dryUse)
	detach.Flags().DurationVar(&waitFor3, "wait-timeout", 10*time.Minute, i18n.S("maximum time for --wait", "--wait için en uzun süre"))
	grp.AddCommand(get, attach, exposeCmd, detach)

	// gyro vpc port-forward update <VPC> <PF> --source ... | --any
	pfg := ensureGroup(root, []string{"vpc", "port-forward"})
	var src []string
	var anyone bool
	var pfName string
	upd := &cobra.Command{
		Use:   "update <VPC_ID> <PORT_FORWARD_ID> [--source <IP|CIDR> ... | --any] [--name <label>]",
		Short: i18n.S("Change a port forward (source allow list, name)", "Port yönlendirmeyi değiştir (kaynak izin listesi, ad)"),
		Long: i18n.S(`Only the given public IPv4 networks may use the forward (a single IP becomes /32; write the network
address, e.g. 198.51.100.0/24; at most 20). The list replaces the old one; --any opens it to everyone.
--name renames the forward (a label of at most 64 characters, "" removes it; the gateway is not touched).
What you leave out does not change.`,
			`Yönlendirmeyi yalnız verilen public IPv4 ağları kullanabilir (tek IP /32 olur; ağ adresi yazılır, ör.
198.51.100.0/24; en çok 20). Liste eskisinin yerine geçer; --any herkese açar.
--name yönlendirmenin adını değiştirir (en çok 64 karakterlik etiket, "" kaldırır; kapıya dokunulmaz).
Verilmeyen değişmez.`),
		Example: "  gyro vpc port-forward update vpc-1a2b3c4d pf-9c8b7a6d --source 203.0.113.10 --source 198.51.100.0/24\n" +
			"  gyro vpc port-forward update vpc-1a2b3c4d pf-9c8b7a6d --any\n" +
			"  gyro vpc port-forward update vpc-1a2b3c4d pf-9c8b7a6d --name ssh-admin",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			l := parseSources(src)
			named := cmd.Flags().Changed("name")
			if anyone && len(l) > 0 || !named && anyone == (len(l) > 0) {
				return &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("give --source <IP|CIDR> (repeatable), --any and/or --name", "--source <IP|CIDR> (tekrarlanabilir), --any ve/veya --name verin"))}
			}
			if len(l) > 20 {
				return &ExitErr{Code: ExitUsage, Err: errors.New(i18n.S("--source: at most 20 networks", "--source: en çok 20 ağ"))}
			}
			body := map[string]any{}
			if anyone || len(l) > 0 {
				body["source_cidrs"] = l
			}
			if named {
				body["name"] = strings.TrimSpace(pfName)
			}
			p := "/cloud/vpcs/" + url.PathEscape(args[0]) + "/port-forwards/" + url.PathEscape(args[1])
			if dry4 {
				return a.Printer().Print(map[string]any{"method": "PATCH", "path": p, "body": body}, nil, "")
			}
			ctx, cancel := a.Ctx()
			defer cancel()
			c, err := a.Client()
			if err != nil {
				return err
			}
			resp, err := c.Do(ctx, api.Request{Method: "PATCH", Path: p, Body: body})
			if err != nil {
				return err
			}
			return a.PrintRaw(resp.Data, nil, "id")
		},
	}
	upd.Flags().StringArrayVar(&src, "source", nil, i18n.S("allowed source IPv4 or CIDR (repeatable or comma separated)", "izin verilen kaynak IPv4 ya da CIDR (tekrarlanabilir ya da virgülle)"))
	upd.Flags().BoolVar(&dry4, "dry-run", false, dryUse)
	upd.Flags().BoolVar(&anyone, "any", false, i18n.S("allow everyone (empty list)", "herkese izin ver (boş liste)"))
	upd.Flags().StringVar(&pfName, "name", "", i18n.S("new name (label, at most 64 characters; \"\" removes it)", "yeni ad (etiket, en çok 64 karakter; \"\" kaldırır)"))
	pfg.AddCommand(upd)
}
