package cli

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// 1.44.0 (API 2.46.0): account context. --account / GYRO_ACCOUNT / profile default send X-Account-Id; GET /accounts is never sent with it.
func TestAccountContext(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	m.routes["GET /accounts"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": []any{
			map[string]any{"id": 35554, "name": "G*** P***", "email": "g***@example.com", "is_self": true, "role": "owner", "permissions": []any{},
				"access": map[string]any{"services": "manage", "tickets": "manage", "passwords": true, "account_admin": true}},
			map[string]any{"id": 35553, "name": "G*** P*** M***", "email": "g***@example.com", "is_self": false, "role": "member",
				"permissions": []any{"view_services", "view_tickets"}, "access": map[string]any{"services": "read", "tickets": "manage", "invoices": "none", "orders": false}}},
			"meta": map[string]any{"total": 2, "page": 1, "limit": 100, "next_page": 0}})
	}
	m.routes["GET /cloud/ssh-keys"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": []any{}, "meta": map[string]any{"total": 0, "page": 1, "limit": 50, "next_page": 0}})
	}

	// no flag = no header (unchanged behaviour)
	r := run(t, srv, nil, "", false, "ssh-key", "list")
	if c := m.find("GET", "/cloud/ssh-keys"); r.code != 0 || c == nil || c.Acct != "" {
		t.Fatalf("no account: %d %s %+v", r.code, r.err, c)
	}
	// --account sends the header
	r = run(t, srv, nil, "", false, "ssh-key", "list", "--account", "35553")
	if c := m.find("GET", "/cloud/ssh-keys"); r.code != 0 || c.Acct != "35553" {
		t.Fatalf("--account: %d %s %+v", r.code, r.err, c)
	}
	// env, and --account self wins over env
	r = run(t, srv, map[string]string{"GYRO_ACCOUNT": "35553"}, "", false, "ssh-key", "list")
	if c := m.find("GET", "/cloud/ssh-keys"); c.Acct != "35553" {
		t.Fatalf("GYRO_ACCOUNT: %+v", c)
	}
	r = run(t, srv, map[string]string{"GYRO_ACCOUNT": "35553"}, "", false, "ssh-key", "list", "--account", "self")
	if c := m.find("GET", "/cloud/ssh-keys"); c.Acct != "" {
		t.Fatalf("--account self: %+v", c)
	}
	// invalid value: usage error, no request
	n := len(m.calls)
	r = run(t, srv, nil, "", false, "ssh-key", "list", "--account", "abc")
	if r.code != ExitUsage || len(m.calls) != n {
		t.Fatalf("invalid --account: %d %s", r.code, r.err)
	}
	// account list: table, current marker, never with the header
	r = run(t, srv, nil, "", false, "account", "list", "--account", "35553")
	if c := m.find("GET", "/accounts"); r.code != 0 || c.Acct != "" || !strings.Contains(r.out, "35553") || !strings.Contains(r.out, "member") {
		t.Fatalf("account list: %d %s %s", r.code, r.out, r.err)
	}
	lines := strings.Split(r.out, "\n")
	if len(lines) < 3 || !strings.HasPrefix(strings.TrimSpace(lines[2]), "*") {
		t.Fatalf("current marker on 35553 expected:\n%s", r.out)
	}
	r = run(t, srv, nil, "", false, "account", "list", "-o", "json")
	if r.code != 0 || !strings.Contains(r.out, `"view_services"`) {
		t.Fatalf("account list json: %s", r.out)
	}
	// account use: profile default, then every call carries it; unknown id refused; self clears
	r = run(t, srv, nil, "", false, "account", "use", "35553")
	raw, _ := os.ReadFile(os.Getenv("GYRO_CONFIG_FILE"))
	if r.code != 0 || !strings.Contains(string(raw), "account = 35553") {
		t.Fatalf("account use: %d %s cfg=%s", r.code, r.err, raw)
	}
	run(t, srv, nil, "", false, "ssh-key", "list")
	if c := m.find("GET", "/cloud/ssh-keys"); c.Acct != "35553" {
		t.Fatalf("profile default not sent: %+v", c)
	}
	r = run(t, srv, nil, "", false, "account", "use", "99999")
	if r.code != ExitAuth {
		t.Fatalf("unknown account: %d %s", r.code, r.err)
	}
	r = run(t, srv, nil, "", false, "account", "use", "self")
	raw, _ = os.ReadFile(os.Getenv("GYRO_CONFIG_FILE"))
	if r.code != 0 || strings.Contains(string(raw), "account =") {
		t.Fatalf("account use self: %d cfg=%s", r.code, raw)
	}
	run(t, srv, nil, "", false, "ssh-key", "list")
	if c := m.find("GET", "/cloud/ssh-keys"); c.Acct != "" {
		t.Fatalf("after self: %+v", c)
	}
	// server error codes pass through (403 account_permission_denied -> exit 3)
	m.routes["POST /cloud/ssh-keys"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 403, map[string]any{"error": map[string]any{"code": "account_permission_denied", "message": "Your user in this account lacks: manage_services.",
			"details": map[string]any{"missing": []any{"manage_services"}}}})
	}
	r = run(t, srv, nil, "", false, "ssh-key", "create", pubFile(t), "--name", "gyro-proje", "--account", "35553")
	if r.code != ExitAuth || !strings.Contains(r.err, "account_permission_denied") {
		t.Fatalf("403 passthrough: %d %s", r.code, r.err)
	}
}

// Device sign-in with a default account picked on the approval page: saved to the profile, shown by auth status.
func TestAuthLoginAccountContext(t *testing.T) {
	var gotAcct string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/auth/device/code":
			_, _ = w.Write([]byte(`{"data":{"device_code":"gdc_x","user_code":"BCDF-GHJK","verification_uri":"https://p/device","verification_uri_complete":"https://p/device?code=BCDF-GHJK","expires_in":300,"interval":5}}`))
		case "/auth/device/token":
			_, _ = w.Write([]byte(`{"data":{"access_token":"wck_new","token_type":"Bearer","key":{"id":7,"name":"k","permissions":3,"scopes":["cloud_read"]},"account":{"id":35554,"name":"G*** P***","email":"g***@example.com"},"account_context":{"id":35553,"name":"G*** P*** M***","email":"g***@example.com","role":"member"}}}`))
		case "/auth/session":
			gotAcct = r.Header.Get("X-Account-Id")
			_, _ = w.Write([]byte(`{"data":{"key":{"id":7,"name":"k","source":"device","permissions":3,"scopes":["cloud_read"]},"account":{"id":35554,"name":"G***","email":"g***@example.com"},"account_context":{"id":35553,"name":"G*** P*** M***","email":"g***@example.com"}}}`))
		default:
			w.WriteHeader(404)
		}
	}))
	defer srv.Close()
	cfg := filepath.Join(t.TempDir(), "config.toml")
	t.Setenv("GYRO_CONFIG_FILE", cfg)
	openBrowser = func(u string) error { return nil }
	r := run(t, srv, map[string]string{"GYRO_TOKEN": "", "GYRO_CONFIG_FILE": cfg}, "", false, "auth", "login", "--no-browser")
	b, _ := os.ReadFile(cfg)
	if r.code != 0 || !strings.Contains(string(b), "account = 35553") || !strings.Contains(r.err, "#35553") {
		t.Fatalf("login: %d %s cfg=%s", r.code, r.err, b)
	}
	r = run(t, srv, map[string]string{"GYRO_TOKEN": "", "GYRO_CONFIG_FILE": cfg}, "", false, "auth", "status")
	if r.code != 0 || gotAcct != "35553" || !strings.Contains(r.out, "#35553") {
		t.Fatalf("status: %d hdr=%q %s %s", r.code, gotAcct, r.out, r.err)
	}
}
