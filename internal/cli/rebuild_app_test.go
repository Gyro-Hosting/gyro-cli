package cli

import (
	"net/http"
	"testing"
)

// rebuild + one-click app (API 2.19.1): --app / --app-domain go into the action body; --app with --user-data is refused locally.
func TestRebuildWithApp(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	m.routes["POST /cloud/servers/47466/actions/rebuild"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 202, map[string]any{"data": map[string]any{"action": "rebuild", "status": "accepted", "server_id": 47466, "app_install_id": 7}})
	}
	r := run(t, srv, nil, "", false, "server", "rebuild", "47466", "--image", "debian-12", "--app", "uptime-kuma", "--app-domain", "up.example.com", "--yes")
	if r.code != 0 {
		t.Fatalf("rebuild --app: %d %s", r.code, r.err)
	}
	b := m.find("POST", "/cloud/servers/47466/actions/rebuild").Body
	if b["app"] != "uptime-kuma" || b["app_domain"] != "up.example.com" || b["confirm"] != true || b["image"] != "debian-12" {
		t.Fatalf("body: %v", b)
	}
	if r = run(t, srv, nil, "", false, "server", "rebuild", "47466", "--image", "debian-12", "--app", "docker", "--user-data", "#!/bin/sh", "--yes"); r.code == 0 {
		t.Fatal("--app with --user-data must fail")
	}
	if r = run(t, srv, nil, "", false, "server", "rebuild", "47466", "--image", "debian-12", "--app-domain", "x.example.com", "--yes"); r.code == 0 {
		t.Fatal("--app-domain without --app must fail")
	}
}
