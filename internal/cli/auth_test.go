package cli

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
)

func authServer(t *testing.T, deny bool, sessionSource string) (*httptest.Server, *int32, *int32) {
	var polls, revoked int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == "POST" && r.URL.Path == "/auth/device/code":
			if r.Header.Get("Authorization") != "" {
				t.Errorf("device code request must not carry a key")
			}
			var b map[string]any
			_ = json.NewDecoder(r.Body).Decode(&b)
			if b["client_name"] != "gyro-cli" || b["template"] != "readonly" {
				t.Errorf("unexpected body %v", b)
			}
			_, _ = w.Write([]byte(`{"data":{"device_code":"gdc_x","user_code":"BCDF-GHJK","verification_uri":"https://p/device","verification_uri_complete":"https://p/device?code=BCDF-GHJK","expires_in":300,"interval":5}}`))
		case r.Method == "POST" && r.URL.Path == "/auth/device/token":
			n := atomic.AddInt32(&polls, 1)
			if n == 1 {
				w.WriteHeader(400)
				_, _ = w.Write([]byte(`{"error":{"code":"authorization_pending","message":"wait","details":{"interval":5}}}`))
				return
			}
			if n == 2 {
				w.WriteHeader(400)
				_, _ = w.Write([]byte(`{"error":{"code":"slow_down","message":"slow","details":{"interval":10}}}`))
				return
			}
			if deny {
				w.WriteHeader(400)
				_, _ = w.Write([]byte(`{"error":{"code":"access_denied","message":"The sign-in request was denied in the browser."}}`))
				return
			}
			_, _ = w.Write([]byte(`{"data":{"access_token":"wck_new","token_type":"Bearer","key":{"id":7,"name":"Gyro CLI · h · 2026-11-15","permissions":53,"scopes":["cloud_read"],"expires_at":"2027-11-15T00:00:00+03:00"},"account":{"id":12345,"name":"G*** P***","email":"g***@example.com"}}}`))
		case r.Method == "GET" && r.URL.Path == "/auth/session":
			_, _ = w.Write([]byte(`{"data":{"key":{"id":7,"name":"k","source":"` + sessionSource + `","expires_at":null,"permissions":3,"scopes":["cloud_read"]},"account":{"id":12345,"name":"G***","email":"g***@example.com"}}}`))
		case r.Method == "DELETE" && r.URL.Path == "/auth/session":
			atomic.AddInt32(&revoked, 1)
			_, _ = w.Write([]byte(`{"data":{"revoked":true,"id":7}}`))
		default:
			w.WriteHeader(404)
			_, _ = w.Write([]byte(`{"error":{"code":"endpoint_not_found","message":"x"}}`))
		}
	}))
	return srv, &polls, &revoked
}

func TestAuthLoginStoresKey(t *testing.T) {
	srv, polls, _ := authServer(t, false, "device")
	defer srv.Close()
	cfg := filepath.Join(t.TempDir(), "config.toml")
	t.Setenv("GYRO_CONFIG_FILE", cfg)
	opened := ""
	openBrowser = func(u string) error { opened = u; return nil }
	r := run(t, srv, map[string]string{"GYRO_TOKEN": "", "GYRO_CONFIG_FILE": cfg, "DISPLAY": ":0"}, "", false, "auth", "login", "--template", "readonly", "--profile", "work")
	if r.code != 0 {
		t.Fatalf("code %d: %s", r.code, r.err)
	}
	if *polls != 3 || opened != "https://p/device?code=BCDF-GHJK" || !strings.Contains(r.err, "BCDF-GHJK") || strings.Contains(r.err, "wck_new") {
		t.Fatalf("polls=%d opened=%q err=%s", *polls, opened, r.err)
	}
	b, _ := os.ReadFile(cfg)
	if !strings.Contains(string(b), `token = "wck_new"`) || !strings.Contains(string(b), "[profiles.work]") {
		t.Fatalf("config: %s", b)
	}
	if st, _ := os.Stat(cfg); st.Mode().Perm() != 0o600 {
		t.Fatalf("mode %v", st.Mode().Perm())
	}
}

func TestAuthLoginDeniedAndHeadless(t *testing.T) {
	srv, _, _ := authServer(t, true, "device")
	defer srv.Close()
	opened := false
	t.Setenv("GYRO_CONFIG_FILE", filepath.Join(t.TempDir(), "c.toml"))
	openBrowser = func(string) error { opened = true; return nil }
	r := run(t, srv, map[string]string{"GYRO_TOKEN": "", "GYRO_CONFIG_FILE": filepath.Join(t.TempDir(), "c.toml"), "SSH_CONNECTION": "1 2 3 4"}, "", false, "auth", "login", "--template", "readonly")
	if r.code != ExitAuth || !strings.Contains(r.err, "denied") || opened {
		t.Fatalf("code %d opened %v: %s", r.code, opened, r.err)
	}
}

func TestAuthLogoutRevokesDeviceKeyOnly(t *testing.T) {
	for _, tc := range []struct {
		src  string
		want int32
	}{{"device", 1}, {"panel", 0}} {
		srv, _, revoked := authServer(t, false, tc.src)
		cfg := filepath.Join(t.TempDir(), "c.toml")
		t.Setenv("GYRO_CONFIG_FILE", cfg)
		_ = os.WriteFile(cfg, []byte("default_profile = \"default\"\n[profiles.default]\ntoken = \"wck_old\"\n"), 0o600)
		r := run(t, srv, map[string]string{"GYRO_TOKEN": "", "GYRO_CONFIG_FILE": cfg}, "", false, "auth", "logout")
		srv.Close()
		b, _ := os.ReadFile(cfg)
		if r.code != 0 || *revoked != tc.want || strings.Contains(string(b), "wck_old") {
			t.Fatalf("%s: code %d revoked %d cfg %s err %s", tc.src, r.code, *revoked, b, r.err)
		}
	}
}

func TestAuthStatusJSON(t *testing.T) {
	t.Setenv("GYRO_CONFIG_FILE", filepath.Join(t.TempDir(), "c.toml"))
	srv, _, _ := authServer(t, false, "device")
	defer srv.Close()
	r := run(t, srv, nil, "", false, "auth", "status", "-o", "json")
	if r.code != 0 || !strings.Contains(r.out, `"source": "device"`) {
		t.Fatalf("code %d out %s err %s", r.code, r.out, r.err)
	}
}
