package cli

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
)

var update = flag.Bool("update", false, "rewrite golden files")

/* ------------------------------------------------------------------ mock API */

type call struct {
	Method, Path, Query, Idem, Auth string
	Body                            map[string]any
}

type mock struct {
	mu        sync.Mutex
	calls     []call
	routes    map[string]func(w http.ResponseWriter, r *http.Request, body map[string]any)
	serverGet int // number of GET /cloud/servers/47466 calls (for --wait)
}

func j(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func errBody(code, msg string) map[string]any {
	return map[string]any{"error": map[string]any{"code": code, "message": msg}}
}

var fxServers = []any{
	map[string]any{"id": 47466, "name": "Gyro Bulut Standart 2-4", "hostname": "web-1.example.com", "status": "active", "product_id": 696,
		"plan": "standart-2-4", "region": "eu-de-1", "server_type": "cpx22", "location": "fsn1", "ipv4": "203.0.113.10", "ipv6": "2001:db8:1::/64", "provisioned": true, "cycle": "monthly",
		"price": map[string]any{"amount": 1845.99, "currency": "TRY"}, "auto_renew": true, "created_at": "2026-11-09T01:00:00+03:00", "due_date": "2026-12-09", "cancellation": nil},
	map[string]any{"id": 47400, "name": "Gyro Bulut Pro 2-8", "hostname": "db-1.example.com", "status": "cancelled", "product_id": 656,
		"plan": "pro-2-8", "region": "eu-fi-1", "server_type": "ccx13", "location": "hel1", "ipv4": "203.0.113.20", "ipv6": nil, "provisioned": true, "cycle": "monthly",
		"price": map[string]any{"amount": 4015.99, "currency": "TRY"}, "auto_renew": false, "created_at": "2026-11-01T01:00:00+03:00", "due_date": "2026-12-01", "cancellation": nil},
}

func newMock(t *testing.T) (*mock, *httptest.Server) {
	m := &mock{routes: map[string]func(http.ResponseWriter, *http.Request, map[string]any){}}
	m.routes["GET /whoami"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"id": 6, "type": "client", "name": "ci", "owner_id": 35071, "permissions": []string{"Module:Addons/GyroAPI/*", "Invoices/*"}}})
	}
	m.routes["GET /cloud/servers"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": fxServers, "meta": map[string]any{"total": 2, "page": 1, "limit": 100, "next_page": 0}})
	}
	m.routes["GET /cloud/servers/47466"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		m.mu.Lock()
		m.serverGet++
		n := m.serverGet
		m.mu.Unlock()
		s := map[string]any{}
		for k, v := range fxServers[0].(map[string]any) {
			s[k] = v
		}
		power, locked := "running", false
		if r.URL.Query().Get("_") == "" && n <= 1 && strings.Contains(t.Name(), "Wait") {
			power, locked = "starting", true
		}
		s["live"] = map[string]any{"power_state": power, "locked": locked, "ipv4": "203.0.113.10", "ipv6_network": "2001:db8:1::/64", "cores": 2, "memory_gb": 4, "disk_gb": 80}
		j(w, 200, map[string]any{"data": s})
	}
	m.routes["GET /cloud/servers/999"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 404, errBody("server_not_found", "Server not found."))
	}
	m.routes["GET /cloud/locations"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": []any{map[string]any{"region": "eu-de-1", "code": "fsn1", "city": "Falkenstein", "country": "DE", "zone": "eu-central", "products": 2},
			map[string]any{"region": "eu-fi-1", "code": "hel1", "city": "Helsinki", "country": "FI", "zone": "eu-central", "products": 1}}})
	}
	m.routes["GET /cloud/images"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": []any{map[string]any{"slug": "ubuntu-24.04", "name": "Ubuntu 24.04", "os_flavor": "ubuntu"}, map[string]any{"slug": "debian-12", "name": "Debian 12", "os_flavor": "debian"}}})
	}
	m.routes["GET /cloud/products"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		all := []any{
			map[string]any{"id": 696, "name": "Gyro Bulut Standart 2-4", "plan": "standart-2-4", "region": "eu-de-1", "server_type": "cpx22", "location": "fsn1", "cores": 2, "memory_gb": 4, "disk_gb": 80, "cpu_type": "shared", "available": true,
				"prices": []any{map[string]any{"cycle": "monthly", "amount": 1845.99, "setup": 0, "currency": "TRY"}}},
			map[string]any{"id": 656, "name": "Gyro Bulut Pro 2-8", "plan": "pro-2-8", "region": "eu-de-1", "server_type": "ccx13", "location": "fsn1", "cores": 2, "memory_gb": 8, "disk_gb": 80, "cpu_type": "dedicated", "available": true,
				"prices": []any{map[string]any{"cycle": "monthly", "amount": 4015.99, "setup": 0, "currency": "TRY"}}},
			map[string]any{"id": 657, "name": "Gyro Bulut Pro 2-8", "plan": "pro-2-8", "region": "eu-fi-1", "server_type": "ccx13", "location": "hel1", "cores": 2, "memory_gb": 8, "disk_gb": 80, "cpu_type": "dedicated", "available": true,
				"prices": []any{map[string]any{"cycle": "monthly", "amount": 4015.99, "setup": 0, "currency": "TRY"}}},
		}
		loc := r.URL.Query().Get("location")
		var out []any
		for _, p := range all {
			if loc == "" || p.(map[string]any)["location"] == loc || p.(map[string]any)["region"] == loc {
				out = append(out, p)
			}
		}
		j(w, 200, map[string]any{"data": out, "meta": map[string]any{"total": len(out), "page": 1, "limit": 100, "next_page": 0}})
	}
	m.routes["GET /cloud/ssh-keys"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": []any{map[string]any{"id": 4, "name": "laptop", "fingerprint": "SHA256:abc", "public_key": "ssh-ed25519 AAAAkey laptop", "created_at": "2026-11-09T01:57:49+03:00"}},
			"meta": map[string]any{"total": 1, "page": 1, "limit": 100, "next_page": 0}})
	}
	m.routes["DELETE /cloud/ssh-keys/4"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"deleted": true, "id": 4}})
	}
	m.routes["POST /cloud/servers"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		if b["dry_run"] == true {
			j(w, 200, map[string]any{"data": map[string]any{"dry_run": true, "quote": map[string]any{"product_id": b["product_id"], "cycle": "monthly",
				"total": map[string]any{"amount": 2215.19, "currency": "TRY"}, "tax": map[string]any{"amount": 369.2, "currency": "TRY"}, "balance": map[string]any{"amount": 5000, "currency": "TRY"}}}})
			return
		}
		if b["confirm"] != true {
			j(w, 422, errBody("confirm_required", "confirm must be true"))
			return
		}
		if r.Header.Get("Idempotency-Key") == "" {
			j(w, 400, errBody("idempotency_key_required", "Idempotency-Key header is required"))
			return
		}
		j(w, 201, map[string]any{"data": map[string]any{"server": fxServers[0], "order": map[string]any{"id": 9001, "number": "#9001", "status": "active"},
			"payment": map[string]any{"method": "balance", "status": "paid"}}})
	}
	m.routes["DELETE /cloud/servers/47466"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		if r.URL.Query().Get("confirm") != "true" {
			j(w, 422, errBody("confirm_required", "confirm"))
			return
		}
		j(w, 200, map[string]any{"data": map[string]any{"id": 47466, "cancellation": "approved"}})
	}
	m.routes["POST /cloud/servers/47466/actions/poweroff"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 202, map[string]any{"data": map[string]any{"action": "poweroff", "status": "accepted", "server_id": 47466, "message": "Power off started."}})
	}
	m.routes["POST /cloud/servers/47466/actions/reset_password"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		if r.Header.Get("Idempotency-Key") != "" {
			j(w, 202, map[string]any{"data": map[string]any{"action": "reset_password", "status": "accepted", "server_id": 47466, "root_password_withheld": true}})
			return
		}
		j(w, 202, map[string]any{"data": map[string]any{"action": "reset_password", "status": "accepted", "server_id": 47466, "root_password": "S3cret-once"}})
	}
	m.routes["GET /cloud/servers/47466/snapshots"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"quota_gb": 20, "items": []any{map[string]any{"id": 11, "description": "before-upgrade", "status": "available", "size_bytes": 1048576, "created_at": "2026-11-09T02:00:00+03:00", "type": "manual"}}}})
	}
	m.routes["GET /balance/transactions"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		page := r.URL.Query().Get("page")
		if page == "2" {
			j(w, 200, map[string]any{"data": []any{map[string]any{"id": 3, "amount": map[string]any{"amount": -10, "currency": "TRY"}, "description": "c"}}, "meta": map[string]any{"total": 3, "page": 2, "limit": 2, "next_page": 0}})
			return
		}
		j(w, 200, map[string]any{"data": []any{map[string]any{"id": 1, "amount": map[string]any{"amount": 100, "currency": "TRY"}, "description": "a"}, map[string]any{"id": 2, "amount": map[string]any{"amount": -50, "currency": "TRY"}, "description": "b"}},
			"meta": map[string]any{"total": 3, "page": 1, "limit": 2, "next_page": 2}})
	}
	m.routes["GET /balance"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"balance": map[string]any{"amount": 4280.61, "currency": "TRY"}, "low_balance": false, "auto_pay": map[string]any{"from_balance": false}}})
	}
	m.routes["GET /invoices"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": []any{
			map[string]any{"invoice_id": 45575, "number": "#45575", "status": "paid", "state": "paid", "total": map[string]any{"amount": 719.39, "currency": "TRY"}, "amount_due": map[string]any{"amount": 0, "currency": "TRY"}, "created_at": "2026-11-09", "due_date": "2026-11-16", "first_item": "Standart 1-1"},
			map[string]any{"invoice_id": 45580, "number": "#45580", "status": "unpaid", "state": "overdue", "total": map[string]any{"amount": 100, "currency": "TRY"}, "amount_due": map[string]any{"amount": 100, "currency": "TRY"}, "created_at": "2026-11-01", "due_date": "2026-11-08", "first_item": "Domain"},
		}, "meta": map[string]any{"total": 2, "page": 1, "limit": 100, "next_page": 0}})
	}
	flaky := 0
	m.routes["GET /cloud/images-flaky"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		flaky++
		if flaky < 3 {
			w.Header().Set("Retry-After", "1")
			j(w, 429, errBody("rate_limited", "Too many requests"))
			return
		}
		j(w, 200, map[string]any{"data": []any{}})
	}
	m.routes["GET /cloud/services"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": []any{map[string]any{"id": 500, "type": "special", "product_line": "mail", "module": "GyroMail", "name": "Kurumsal Mail 10", "status": "active", "host": "example.com"}},
			"meta": map[string]any{"total": 1, "page": 1, "limit": 100, "next_page": 0}})
	}
	m.routes["GET /cloud/services/500/tools"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"service_id": 500, "product_line": "mail", "tools": []any{
			map[string]any{"tool": "mailboxes", "label": "Mailboxes", "actions": []any{
				map[string]any{"action": "create", "read": false, "confirm": false, "secret": false},
				map[string]any{"action": "delete", "read": false, "confirm": true, "secret": false}}}}}})
	}
	m.routes["GET /cloud/services/500/tools/mailboxes"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"mailboxes": []any{map[string]any{"address": "info@example.com", "quota_mb": 1024}}, "used": 1}})
	}
	m.routes["POST /cloud/services/500/tools/mailboxes/create"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"ok": true, "address": fmt.Sprint(b["local_part"], "@example.com")}})
	}
	m.routes["POST /cloud/services/500/tools/mailboxes/delete"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"ok": b["confirm"] == true}})
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		var body map[string]any
		_ = json.Unmarshal(raw, &body)
		m.mu.Lock()
		m.calls = append(m.calls, call{Method: r.Method, Path: r.URL.Path, Query: r.URL.RawQuery, Idem: r.Header.Get("Idempotency-Key"), Auth: r.Header.Get("Authorization"), Body: body})
		m.mu.Unlock()
		if r.Header.Get("Authorization") != "Bearer wck_test" {
			j(w, 401, errBody("invalid_token", "Invalid API key."))
			return
		}
		h := m.routes[r.Method+" "+r.URL.Path]
		if h == nil {
			j(w, 404, errBody("not_found", "no route "+r.Method+" "+r.URL.Path))
			return
		}
		h(w, r, body)
	}))
	t.Cleanup(srv.Close)
	return m, srv
}

func (m *mock) find(method, path string) *call {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := len(m.calls) - 1; i >= 0; i-- {
		if m.calls[i].Method == method && m.calls[i].Path == path {
			c := m.calls[i]
			return &c
		}
	}
	return nil
}

/* ------------------------------------------------------------------ harness */

type result struct {
	code     int
	out, err string
	execd    []string
}

func run(t *testing.T, srv *httptest.Server, env map[string]string, stdin string, tty bool, args ...string) result {
	t.Helper()
	i18n.Set("en")
	if l := env["GYRO_LANG"]; l != "" {
		i18n.Set(l)
	}
	handled = map[string]bool{}
	handledCmd = map[string]bool{}
	var out, errb bytes.Buffer
	var execd []string
	a := NewApp()
	a.In, a.Out, a.Err = strings.NewReader(stdin), &out, &errb
	a.IsTTY = func() bool { return tty }
	a.Sleep = func(time.Duration) {}
	a.Exec = func(name string, args []string) error { execd = append([]string{name}, args...); return nil }
	base := map[string]string{"GYRO_TOKEN": "wck_test"}
	if srv != nil {
		base["GYRO_API_URL"] = srv.URL
	}
	for k, v := range env {
		base[k] = v
	}
	a.Env = func(k string) string { return base[k] }
	a.NewClient = func(baseURL, token string) *api.Client {
		c := api.New(baseURL, token, "gyro-cli/test")
		c.Sleep = func(time.Duration) {}
		return c
	}
	code := Execute(a, args)
	return result{code, out.String(), errb.String(), execd}
}

func golden(t *testing.T, name, got string) {
	t.Helper()
	p := filepath.Join("testdata", name+".golden")
	if *update {
		_ = os.MkdirAll("testdata", 0o755)
		if err := os.WriteFile(p, []byte(got), 0o644); err != nil {
			t.Fatal(err)
		}
		return
	}
	want, err := os.ReadFile(p)
	if err != nil {
		t.Fatalf("missing golden %s (run go test ./internal/cli -update): %v", p, err)
	}
	if string(want) != got {
		t.Errorf("%s mismatch\n--- want\n%s\n--- got\n%s", name, want, got)
	}
}

func setup(t *testing.T) {
	t.Setenv("GYRO_CONFIG_FILE", filepath.Join(t.TempDir(), "config.toml"))
	t.Setenv("NO_COLOR", "1")
}

/* ------------------------------------------------------------------ tests */

func TestServerListGolden(t *testing.T) {
	setup(t)
	_, srv := newMock(t)
	r := run(t, srv, nil, "", false, "server", "list")
	if r.code != 0 {
		t.Fatalf("code %d: %s", r.code, r.err)
	}
	golden(t, "server_list", r.out)
	r = run(t, srv, nil, "", false, "server", "list", "-o", "json")
	golden(t, "server_list_json", r.out)
	r = run(t, srv, nil, "", false, "server", "list", "-q")
	if r.out != "47466\n47400\n" {
		t.Errorf("quiet: %q", r.out)
	}
	r = run(t, srv, nil, "", false, "servers", "ls", "--query", "[?status=='active'].ipv4", "-o", "text")
	if strings.TrimSpace(r.out) != "203.0.113.10" {
		t.Errorf("query: %q", r.out)
	}
	r = run(t, srv, nil, "", false, "server", "list", "--no-header", "-o", "table")
	if strings.Contains(r.out, "HOSTNAME") {
		t.Errorf("no-header: %q", r.out)
	}
}

func TestServerGetByName(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	r := run(t, srv, nil, "", false, "server", "get", "web-1", "-o", "yaml")
	if r.code != 0 || !strings.Contains(r.out, "power_state: running") {
		t.Fatalf("code %d out %q err %q", r.code, r.out, r.err)
	}
	if m.find("GET", "/cloud/servers/47466") == nil {
		t.Error("did not resolve web-1 to 47466")
	}
	r = run(t, srv, nil, "", false, "server", "get", "nope")
	if r.code != ExitNotFound {
		t.Errorf("unknown name code %d", r.code)
	}
	r = run(t, srv, nil, "", false, "server", "get", "999")
	if r.code != ExitNotFound || !strings.Contains(r.err, "server_not_found") || !strings.Contains(r.err, "request id") {
		t.Errorf("404: %d %q", r.code, r.err)
	}
	golden(t, "server_get_404_err", stripReqID(r.err))
}

func stripReqID(s string) string {
	i := strings.Index(s, "request id: ")
	if i < 0 {
		return s
	}
	j := strings.Index(s[i:], ")")
	return s[:i] + "request id: <id>" + s[i+j:]
}

func TestServerCreateDryRun(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	r := run(t, srv, nil, "", false, "server", "create", "web-2", "--plan", "pro-2-8", "--location", "Falkenstein", "--image", "Ubuntu 24.04", "--ssh-key", "laptop", "--dry-run")
	if r.code != 0 {
		t.Fatalf("code %d: %s", r.code, r.err)
	}
	golden(t, "server_create_dry_run_err", r.err)
	c := m.find("POST", "/cloud/servers")
	if c == nil || c.Body["dry_run"] != true || c.Body["product_id"].(float64) != 656 || c.Body["image"] != "ubuntu-24.04" || c.Body["hostname"] != "web-2" {
		t.Errorf("quote body: %+v", c)
	}
	if ks, _ := c.Body["ssh_keys"].([]any); len(ks) != 1 || ks[0].(float64) != 4 {
		t.Errorf("ssh_keys: %v", c.Body["ssh_keys"])
	}
	if len(m.calls) > 0 {
		for _, cl := range m.calls {
			if cl.Method == "POST" && cl.Body["confirm"] == true {
				t.Error("dry run must not order")
			}
		}
	}
}

func TestServerCreateNeedsConfirmation(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	r := run(t, srv, nil, "", false, "server", "create", "--plan", "standart-2-4", "--location", "eu-de-1", "--image", "debian-12")
	if r.code != ExitAborted || !strings.Contains(r.err, "--yes") {
		t.Fatalf("expected abort, got %d %q", r.code, r.err)
	}
	for _, c := range m.calls {
		if c.Body["confirm"] == true {
			t.Fatal("ordered without confirmation")
		}
	}
	// interactive "n"
	r = run(t, srv, nil, "n\n", true, "server", "create", "--plan", "standart-2-4", "--location", "eu-de-1", "--image", "debian-12")
	if r.code != ExitAborted {
		t.Fatalf("expected abort on n, got %d", r.code)
	}
	// ambiguous plan without location
	r = run(t, srv, nil, "", false, "server", "create", "--plan", "pro-2-8", "--image", "debian-12", "--yes")
	if r.code != ExitUsage || !strings.Contains(r.err, "656") || !strings.Contains(r.err, "657") {
		t.Fatalf("ambiguous: %d %q", r.code, r.err)
	}
}

func TestServerCreateYesWait(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	r := run(t, srv, nil, "e\n", true, "server", "create", "web-3", "--plan", "standart-2-4", "--location", "fsn1", "--image", "ubuntu-24.04", "--wait", "-o", "json", "--lang", "tr")
	if r.code != 0 {
		t.Fatalf("code %d: %s", r.code, r.err)
	}
	c := m.find("POST", "/cloud/servers")
	if c.Body["confirm"] != true || c.Idem == "" {
		t.Errorf("order must carry confirm + Idempotency-Key: %+v", c)
	}
	if m.serverGet < 2 {
		t.Errorf("--wait polled %d times", m.serverGet)
	}
	if !strings.Contains(r.out, `"power_state": "running"`) {
		t.Errorf("final server not printed: %s", r.out)
	}
	if !strings.Contains(r.err, "hazır") {
		t.Errorf("turkish messages expected: %s", r.err)
	}
}

func TestServerDelete(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	r := run(t, srv, nil, "", false, "server", "delete", "web-1")
	if r.code != ExitAborted {
		t.Fatalf("delete without --yes: %d", r.code)
	}
	r = run(t, srv, nil, "web-2\n", true, "server", "delete", "web-1")
	if r.code != ExitAborted {
		t.Fatalf("wrong typed name must abort: %d", r.code)
	}
	if m.find("DELETE", "/cloud/servers/47466") != nil {
		t.Fatal("deleted without confirmation")
	}
	r = run(t, srv, nil, "web-1.example.com\n", true, "server", "delete", "web-1")
	if r.code != 0 {
		t.Fatalf("typed confirm: %d %s", r.code, r.err)
	}
	c := m.find("DELETE", "/cloud/servers/47466")
	if c == nil || c.Query != "confirm=true" {
		t.Errorf("delete call: %+v", c)
	}
}

func TestPowerAndSecrets(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	r := run(t, srv, nil, "", false, "server", "power", "off", "47466")
	if r.code != ExitAborted {
		t.Fatalf("hard power off needs --yes: %d", r.code)
	}
	r = run(t, srv, nil, "", false, "server", "power", "off", "47466", "--yes")
	if r.code != 0 || m.find("POST", "/cloud/servers/47466/actions/poweroff").Body["confirm"] != true {
		t.Fatalf("poweroff: %d %s", r.code, r.err)
	}
	r = run(t, srv, nil, "", false, "server", "reset-password", "web-1", "--yes")
	if r.code != 0 || !strings.Contains(r.out, "S3cret-once") || !strings.Contains(r.err, "only once") {
		t.Fatalf("reset-password: %d out=%q err=%q", r.code, r.out, r.err)
	}
	if m.find("POST", "/cloud/servers/47466/actions/reset_password").Idem != "" {
		t.Error("secret-returning action must not send Idempotency-Key")
	}
}

func TestGeneratedCommands(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	r := run(t, srv, nil, "", false, "server", "snapshot", "list", "web-1")
	if r.code != 0 {
		t.Fatalf("snapshot list: %d %s", r.code, r.err)
	}
	golden(t, "snapshot_list", r.out)
	r = run(t, srv, nil, "", false, "ssh-key", "delete", "laptop", "--yes")
	if r.code != 0 || m.find("DELETE", "/cloud/ssh-keys/4") == nil {
		t.Fatalf("ssh-key delete by name: %d %s", r.code, r.err)
	}
	r = run(t, srv, nil, "", false, "balance", "transaction", "list", "--all")
	if r.code != 0 || strings.Count(r.out, "\n") != 4 {
		t.Fatalf("pagination --all: %d %q", r.code, r.out)
	}
	golden(t, "balance_transactions_all", r.out)
	r = run(t, srv, nil, "", false, "ssh-key", "delete", "4", "--dry-run")
	if r.code != 0 || !strings.Contains(r.out, "DELETE") || !strings.Contains(r.err, "NOT sent") {
		t.Fatalf("generic dry-run: %d %q %q", r.code, r.out, r.err)
	}
}

func TestBillingAndBalance(t *testing.T) {
	setup(t)
	_, srv := newMock(t)
	r := run(t, srv, nil, "", false, "invoice", "list", "--unpaid")
	golden(t, "invoice_unpaid", r.out)
	r = run(t, srv, map[string]string{"GYRO_LANG": "tr"}, "", false, "balance")
	if r.code != 0 || r.out != "Bakiye: 4280.61 TRY\nFaturaları bakiyeden otomatik öde: no\n" {
		t.Fatalf("balance: %q", r.out)
	}
}

func TestRetryOn429(t *testing.T) {
	setup(t)
	_, srv := newMock(t)
	r := run(t, srv, nil, "", false, "api", "request", "GET", "/cloud/images-flaky")
	if r.code != 0 {
		t.Fatalf("429 retry: %d %s", r.code, r.err)
	}
}

func TestAuthErrors(t *testing.T) {
	setup(t)
	_, srv := newMock(t)
	r := run(t, srv, map[string]string{"GYRO_TOKEN": "wck_wrong"}, "", false, "server", "list")
	if r.code != ExitAuth || !strings.Contains(r.err, "gyro context") {
		t.Fatalf("401: %d %q", r.code, r.err)
	}
	r = run(t, srv, map[string]string{"GYRO_TOKEN": ""}, "", false, "server", "list")
	if r.code != ExitAuth || !strings.Contains(r.err, "gyro configure") {
		t.Fatalf("no token: %d %q", r.code, r.err)
	}
}

func TestConfigureAndContext(t *testing.T) {
	setup(t)
	_, srv := newMock(t)
	r := run(t, srv, map[string]string{"GYRO_TOKEN": ""}, "wck_test\n", false, "configure", "--token-stdin", "--profile", "ci", "--set-api-url", srv.URL, "--set-output", "json")
	if r.code != 0 {
		t.Fatalf("configure: %d %s", r.code, r.err)
	}
	st, err := os.Stat(os.Getenv("GYRO_CONFIG_FILE"))
	if err != nil || (runtime.GOOS != "windows" && st.Mode().Perm() != 0o600) { // Windows has no POSIX modes
		t.Fatalf("config perms: %v %v", err, st.Mode())
	}
	b, _ := os.ReadFile(os.Getenv("GYRO_CONFIG_FILE"))
	if !strings.Contains(string(b), `[profiles.ci]`) || !strings.Contains(string(b), `output = "json"`) {
		t.Errorf("config: %s", b)
	}
	// use the profile (no env token / url)
	r = run(t, nil, map[string]string{"GYRO_TOKEN": "", "GYRO_PROFILE": "ci"}, "", false, "context")
	if r.code != 0 || !strings.Contains(r.out, `"key_status": "ok"`) || strings.Contains(r.out, "wck_test") {
		t.Fatalf("context: %d %s %s", r.code, r.out, r.err)
	}
	r = run(t, nil, map[string]string{"GYRO_TOKEN": ""}, "", false, "configure", "list", "-o", "text")
	if !strings.Contains(r.out, "ci") {
		t.Errorf("configure list: %q", r.out)
	}
	r = run(t, nil, map[string]string{"GYRO_TOKEN": ""}, "", false, "--profile", "ci", "configure", "get", "output")
	if strings.TrimSpace(r.out) != "json" {
		t.Errorf("configure get: %q", r.out)
	}
}

func TestSSHHelper(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("argv uses ssh.exe on Windows")
	}
	setup(t)
	_, srv := newMock(t)
	r := run(t, srv, nil, "", false, "server", "ssh", "web-1", "-i", "/k/id", "--", "uptime")
	if r.code != 0 || strings.Join(r.execd, " ") != "ssh -i /k/id root@203.0.113.10 uptime" {
		t.Fatalf("ssh: %d %v %s", r.code, r.execd, r.err)
	}
	r = run(t, srv, nil, "", false, "server", "ssh", "web-1", "--ipv6", "-u", "deploy")
	if strings.Join(r.execd, " ") != "ssh deploy@2001:db8:1::1" {
		t.Fatalf("ssh v6: %v", r.execd)
	}
}

func TestHelpLanguagesAndCompletion(t *testing.T) {
	setup(t)
	r := run(t, nil, map[string]string{"GYRO_LANG": "tr"}, "", false, "server", "--help")
	if !strings.Contains(r.out, "Sunucuya SSH bağlantısı aç") {
		t.Errorf("tr help: %s", r.out)
	}
	for _, sh := range []string{"bash", "zsh", "fish", "powershell"} {
		r = run(t, nil, nil, "", false, "completion", sh)
		if r.code != 0 || len(r.out) < 500 {
			t.Errorf("completion %s: %d len %d", sh, r.code, len(r.out))
		}
	}
	r = run(t, nil, nil, "", false, "nosuch")
	if r.code != ExitUsage {
		t.Errorf("unknown command code %d", r.code)
	}
	r = run(t, nil, nil, "", false, "api", "operations", "-o", "json")
	var ops []map[string]any
	if json.Unmarshal([]byte(r.out), &ops); len(ops) < 25 {
		t.Errorf("operations: %d", len(ops))
	}
	for _, o := range ops {
		if !strings.HasPrefix(fmt.Sprint(o["command"]), "gyro ") {
			t.Errorf("operation without command: %v", o)
		}
	}
}

func TestVersionNewer(t *testing.T) {
	if !newer("v1.2.0", "1.1.9") || newer("v1.0.0", "1.0.0") || newer("v2.0.0", "dev") || !newer("v1.10.0", "v1.9.3") {
		t.Error("newer()")
	}
}

func TestSlugs(t *testing.T) {
	if PlanSlug("Gyro Bulut Pro 2-8") != "pro-2-8" || Slug("İstanbul Şişli") != "istanbul-sisli" {
		t.Error("slug")
	}
}

func TestProductLineBridge(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	r := run(t, srv, nil, "", false, "mail", "mailbox", "list", "example.com")
	if r.code != 0 || !strings.Contains(r.out, "info@example.com") {
		t.Fatalf("mailbox list: %d %q %q", r.code, r.out, r.err)
	}
	r = run(t, srv, nil, "", false, "mail", "mailbox", "create", "example.com", "--set", "local_part=sales", "-o", "json")
	if r.code != 0 || !strings.Contains(r.out, "sales@example.com") {
		t.Fatalf("mailbox create: %d %q %q", r.code, r.out, r.err)
	}
	if c := m.find("POST", "/cloud/services/500/tools/mailboxes/create"); c == nil || c.Body["confirm"] != nil {
		t.Errorf("create must not need confirm: %+v", c)
	}
	r = run(t, srv, nil, "", false, "mail", "mailbox", "delete", "500", "--set", "address=sales@example.com")
	if r.code != ExitAborted {
		t.Fatalf("delete needs --yes (policy confirm): %d", r.code)
	}
	r = run(t, srv, nil, "", false, "mail", "mailbox", "delete", "500", "--set", "address=sales@example.com", "--yes", "-o", "json")
	if r.code != 0 || !strings.Contains(r.out, "true") {
		t.Fatalf("delete --yes: %d %q", r.code, r.out)
	}
	r = run(t, srv, nil, "", false, "mail", "tools", "example.com")
	if r.code != 0 || !strings.Contains(r.out, "mailboxes") {
		t.Fatalf("tools: %d %q", r.code, r.out)
	}
	r = run(t, srv, nil, "", false, "mail", "alias", "list", "500")
	if r.code != ExitNotFound {
		t.Fatalf("missing tool must be 404-ish: %d %q", r.code, r.err)
	}
}
