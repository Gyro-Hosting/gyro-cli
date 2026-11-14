package cli

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// 1.26.5 (gece testi): a .pub file given to `server create` is uploaded only after the quote and the confirmation, and is
// removed again when the order is refused.
func pubFile(t *testing.T) string {
	p := filepath.Join(t.TempDir(), "id_ed25519.pub")
	if err := os.WriteFile(p, []byte("ssh-ed25519 AAAAnewkey gyro-proje\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	return p
}

func keyRoutes(m *mock) {
	m.routes["POST /cloud/ssh-keys"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 201, map[string]any{"data": map[string]any{"id": 9, "name": "gyro-proje", "fingerprint": "SHA256:new"}})
	}
	m.routes["DELETE /cloud/ssh-keys/9"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"deleted": true, "id": 9}})
	}
}

func TestServerCreateRefusedOrderRemovesUploadedKey(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	keyRoutes(m)
	prev := m.routes["POST /cloud/servers"]
	m.routes["POST /cloud/servers"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		if b["dry_run"] == true {
			prev(w, r, b)
			return
		}
		j(w, 409, errBody("order_refused", "The order was refused."))
	}
	r := run(t, srv, map[string]string{"GYRO_LANG": "tr"}, "", false, "server", "create", "--plan", "standart-2-4", "--location", "eu-de-1", "--image", "debian-12", "--ssh-key", pubFile(t), "--yes")
	if r.code == 0 {
		t.Fatal("refused order must fail")
	}
	up := m.find("POST", "/cloud/ssh-keys")
	if up == nil {
		t.Fatal("key was not uploaded before the real order")
	}
	if m.find("DELETE", "/cloud/ssh-keys/9") == nil || !strings.Contains(r.err, "geri silindi") {
		t.Fatalf("uploaded key not removed after the refusal: %s", r.err)
	}
	ord := m.find("POST", "/cloud/servers")
	if ks, _ := ord.Body["ssh_keys"].([]any); len(ks) != 1 || ks[0] != float64(9) {
		t.Errorf("order must carry the new key id: %+v", ord.Body["ssh_keys"])
	}
}

func TestServerCreateNoUploadWithoutConfirmation(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	keyRoutes(m)
	r := run(t, srv, nil, "", false, "server", "create", "--plan", "standart-2-4", "--location", "eu-de-1", "--image", "debian-12", "--ssh-key", pubFile(t))
	if r.code != ExitAborted {
		t.Fatalf("expected abort, got %d %s", r.code, r.err)
	}
	if m.find("POST", "/cloud/ssh-keys") != nil {
		t.Fatal("key uploaded although nothing was ordered")
	}
	r = run(t, srv, nil, "", false, "server", "create", "--plan", "standart-2-4", "--location", "eu-de-1", "--image", "debian-12", "--ssh-key", pubFile(t), "--dry-run")
	if r.code != 0 || m.find("POST", "/cloud/ssh-keys") != nil || !strings.Contains(r.err, "would be uploaded") {
		t.Fatalf("dry run must not upload: %d %s", r.code, r.err)
	}
	// existing key by name: no upload, id sent
	r = run(t, srv, nil, "", false, "server", "create", "--plan", "standart-2-4", "--location", "eu-de-1", "--image", "debian-12", "--ssh-key", "laptop", "--yes")
	if r.code != 0 || m.find("POST", "/cloud/ssh-keys") != nil {
		t.Fatalf("existing key: %d %s", r.code, r.err)
	}
	if ks, _ := m.find("POST", "/cloud/servers").Body["ssh_keys"].([]any); len(ks) != 1 || ks[0] != float64(4) {
		t.Errorf("existing key id not sent: %+v", ks)
	}
}

func ecsRoutes(m *mock) {
	m.routes["GET /cloud/containers/77/services"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{
			"services": []any{
				map[string]any{"name": "api-gateway", "image": "registry.example/gyro-proje/api:1.2", "replicas": 3, "current_replicas": 3, "ready": 2, "state": "deploying", "problem": "",
					"cpu_m": 500, "memory_mi": 512, "public": true, "port": 8080, "revision": 4, "url": "https://api-gateway-abc123.gyroapp.net"},
				map[string]any{"name": "worker", "image": "busybox", "replicas": 1, "current_replicas": 2, "ready": 2, "state": "running",
					"autoscaling": map[string]any{"enabled": true, "min": 1, "max": 4}},
			},
			"limits": map[string]any{"max_services": 10, "max_replicas": 5}, "suffix": "-abc123.gyroapp.net", "live": true}})
	}
	m.routes["GET /cloud/containers/78/services"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"services": []any{}, "limits": map[string]any{"max_services": 3, "max_replicas": 3}}})
	}
	m.routes["POST /cloud/containers/77/services/worker/scale"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 422, map[string]any{"error": map[string]any{"code": "invalid_replicas", "message": "replicas must be between 0 and 5: your plan allows at most 5 copies (replicas) per service.",
			"details": map[string]any{"field": "replicas", "min": 0, "max": 5}}})
	}
}

func TestContainerServiceListTable(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	ecsRoutes(m)
	r := run(t, srv, nil, "", false, "container", "service", "list", "77")
	if r.code != 0 {
		t.Fatalf("code %d: %s", r.code, r.err)
	}
	for _, want := range []string{"NAME", "COPIES", "api-gateway", "2/3", "deploying", "https://api-gateway-abc123.gyroapp.net", "worker", "2/2 (1-4)", "running"} {
		if !strings.Contains(r.out, want) {
			t.Errorf("table lacks %q:\n%s", want, r.out)
		}
	}
	if !strings.Contains(r.err, "at most 10 services, 5 copies") {
		t.Errorf("plan limits not shown: %s", r.err)
	}
	// JSON keeps the full answer
	r = run(t, srv, nil, "", false, "container", "service", "list", "77", "-o", "json")
	if !strings.Contains(r.out, `"limits"`) || !strings.Contains(r.out, `"services"`) {
		t.Errorf("json changed: %s", r.out)
	}
	r = run(t, srv, nil, "", false, "container", "service", "list", "77", "-q")
	if strings.TrimSpace(r.out) != "api-gateway\nworker" {
		t.Errorf("quiet: %q", r.out)
	}
	r = run(t, srv, map[string]string{"GYRO_LANG": "tr"}, "", false, "container", "service", "list", "78")
	if r.code != 0 || !strings.Contains(r.err, "henüz servis yok") || !strings.Contains(r.err, "en fazla 3 kopya") {
		t.Errorf("empty env: %d %s", r.code, r.err)
	}
}

func TestContainerScaleLimitHint(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	ecsRoutes(m)
	r := run(t, srv, map[string]string{"GYRO_LANG": "tr"}, "", false, "container", "service", "scale", "77", "worker", "--replicas", "9")
	if r.code == 0 || !strings.Contains(r.err, "en fazla 5 kopya") || !strings.Contains(r.err, "0 ile 5") {
		t.Fatalf("limit not explained: %d %s", r.code, r.err)
	}
}
