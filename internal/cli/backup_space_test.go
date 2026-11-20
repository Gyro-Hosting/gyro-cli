package cli

import (
	"net/http"
	"strings"
	"testing"
)

// 1.42.0: gyro backup-space (Gyro Yedekleme Alanı, API 2.44.0 cloud/backup-spaces).
func TestBackupSpaceCommands(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	conn := map[string]any{"host": "203.0.113.12", "port": 2201, "user": "y-m48201", "path": "/veri", "protocol": "sftp", "size_gb": 100, "size_plan": "100gb",
		"used_bytes": 0, "used_percent": 0, "server_id": 47466, "allowed_ips": []any{"203.0.113.10"}, "state": "active"}
	m.routes["GET /cloud/backup-spaces/plans"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"plans": []any{
			map[string]any{"id": 1925, "plan": "50gb", "name": "Gyro Yedekleme Alanı 50 GB", "region": "tr-ist-1", "protocol": "sftp", "size_gb": 50, "available": true,
				"prices": []any{map[string]any{"cycle": "monthly", "amount": 0.99, "list_amount": 1.29, "currency": "EUR"}}, "tax_included": false},
			map[string]any{"id": 1926, "plan": "100gb", "name": "Gyro Yedekleme Alanı 100 GB", "region": "tr-ist-1", "protocol": "sftp", "size_gb": 100, "available": true,
				"prices": []any{map[string]any{"cycle": "monthly", "amount": 1.49, "list_amount": 1.99, "currency": "EUR"}}, "tax_included": false}},
			"servers": []any{map[string]any{"server_id": 47466, "name": "web-1", "ip": "203.0.113.10"}}}})
	}
	m.routes["GET /cloud/backup-spaces"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": []any{map[string]any{"id": 48201, "name": "gyro-proje-yedek", "status": "active", "product_line": "backup_space"}},
			"meta": map[string]any{"total": 1, "page": 1, "limit": 100, "next_page": 0}})
	}
	m.routes["POST /cloud/backup-spaces"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		if b["dry_run"] == true {
			j(w, 200, map[string]any{"data": map[string]any{"dry_run": true, "quote": map[string]any{"plan": b["plan"], "server_id": b["server_id"], "total": 58.9}}})
			return
		}
		j(w, 201, map[string]any{"data": map[string]any{"backup_space": map[string]any{"id": 48201, "status": "inprocess"}, "plan": b["plan"], "size_gb": 100, "cycle": "monthly", "server_id": 47466}})
	}
	m.routes["GET /cloud/backup-spaces/48201"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"id": 48201, "name": "gyro-proje-yedek", "status": "active", "backup_space": conn}})
	}
	m.routes["GET /cloud/backup-spaces/48201/credentials"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"backup_space_id": 48201, "host": "203.0.113.12", "port": 2201, "user": "y-m48201", "password": "Gyro-Test-Parola-1", "path": "/veri", "protocol": "sftp"}})
	}
	m.routes["POST /cloud/backup-spaces/48201/password"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		if b["confirm"] != true {
			j(w, 422, errBody("confirm_required", "Repeat with confirm."))
			return
		}
		j(w, 200, map[string]any{"data": map[string]any{"backup_space_id": 48201, "host": "203.0.113.12", "port": 2201, "user": "y-m48201", "password": "Gyro-Test-Parola-2"}})
	}
	m.routes["PUT /cloud/backup-spaces/48201/server"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"backup_space_id": 48201, "backup_space": conn}})
	}
	m.routes["DELETE /cloud/backup-spaces/48201"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		if r.URL.Query().Get("confirm") != "true" {
			j(w, 422, errBody("confirm_required", "Repeat with confirm=true."))
			return
		}
		st := "cancellation_requested"
		if r.URL.Query().Get("immediate") == "true" {
			st = "cancelling"
		}
		j(w, 202, map[string]any{"data": map[string]any{"status": st, "backup_space_id": 48201}})
	}

	r := run(t, srv, nil, "", false, "backup-space", "plans", "-o", "json")
	if r.code != 0 || !strings.Contains(r.out, "100gb") || !strings.Contains(r.out, "203.0.113.10") {
		t.Fatalf("plans: %d %q %q", r.code, r.out, r.err)
	}
	r = run(t, srv, nil, "", false, "backup-space", "plans")
	if r.code != 0 || !strings.Contains(r.out, "1.49 EUR") || !strings.Contains(r.out, "100gb") || !strings.Contains(r.err, "203.0.113.10") {
		t.Fatalf("plans table: %d %q %q", r.code, r.out, r.err)
	}
	r = run(t, srv, nil, "", false, "backup-spaces", "list")
	if r.code != 0 || !strings.Contains(r.out, "48201") {
		t.Fatalf("list (alias): %d %q %q", r.code, r.out, r.err)
	}
	// money: dry run gives the server quote, no Idempotency-Key needed, nothing ordered
	r = run(t, srv, nil, "", false, "backup-space", "create", "--plan", "100gb", "--server-id", "47466", "--dry-run", "-o", "json")
	if c := m.find("POST", "/cloud/backup-spaces"); r.code != 0 || c == nil || c.Body["dry_run"] != true || c.Body["confirm"] != nil || !strings.Contains(r.out, "quote") {
		t.Fatalf("create dry-run: %d %q %q %+v", r.code, r.out, r.err, c)
	}
	// without --yes (no TTY) it must not order
	m.calls = nil
	r = run(t, srv, nil, "", false, "backup-space", "create", "--plan", "100gb", "--server-id", "47466")
	if r.code != ExitAborted || m.find("POST", "/cloud/backup-spaces") != nil {
		t.Fatalf("create needs --yes: %d %q", r.code, r.err)
	}
	r = run(t, srv, nil, "", false, "backup-space", "create", "--plan", "100gb", "--server-id", "47466", "--cycle", "monthly", "--name", "gyro-proje-yedek", "--yes", "-o", "json")
	c := m.find("POST", "/cloud/backup-spaces")
	if r.code != 0 || c == nil || c.Body["confirm"] != true || c.Idem == "" || c.Body["server_id"] != float64(47466) || c.Body["plan"] != "100gb" || c.Body["name"] != "gyro-proje-yedek" {
		t.Fatalf("create --yes: %d %q %+v", r.code, r.err, c)
	}
	r = run(t, srv, nil, "", false, "backup-space", "get", "48201", "-o", "json")
	if r.code != 0 || !strings.Contains(r.out, "y-m48201") || strings.Contains(r.out, "password") {
		t.Fatalf("get: %d %q %q", r.code, r.out, r.err)
	}
	r = run(t, srv, nil, "", false, "backup-space", "credentials", "48201", "-o", "json")
	if r.code != 0 || !strings.Contains(r.out, "Gyro-Test-Parola-1") || strings.Contains(r.err, "only once") || !strings.Contains(r.err, "SFTP password") {
		t.Fatalf("credentials: %d %q %q", r.code, r.out, r.err)
	}
	m.calls = nil
	r = run(t, srv, nil, "", false, "backup-space", "reset-password", "48201")
	if r.code != ExitAborted || m.find("POST", "/cloud/backup-spaces/48201/password") != nil {
		t.Fatalf("reset-password needs --yes: %d %q", r.code, r.err)
	}
	r = run(t, srv, nil, "", false, "backup-space", "reset-password", "48201", "--yes", "-o", "json")
	if c := m.find("POST", "/cloud/backup-spaces/48201/password"); r.code != 0 || c == nil || c.Body["confirm"] != true || !strings.Contains(r.out, "Gyro-Test-Parola-2") || !strings.Contains(r.err, "only once") {
		t.Fatalf("reset-password --yes: %d %q %q %+v", r.code, r.out, r.err, c)
	}
	r = run(t, srv, nil, "", false, "backup-space", "attach-server", "48201", "--server-id", "47466", "-o", "json")
	if c := m.find("PUT", "/cloud/backup-spaces/48201/server"); r.code != 0 || c == nil || c.Body["server_id"] != float64(47466) {
		t.Fatalf("attach-server: %d %q %+v", r.code, r.err, c)
	}
	m.calls = nil
	r = run(t, srv, nil, "", false, "backup-space", "delete", "48201", "--immediate")
	if r.code != ExitAborted || m.find("DELETE", "/cloud/backup-spaces/48201") != nil {
		t.Fatalf("delete needs --yes: %d %q", r.code, r.err)
	}
	r = run(t, srv, nil, "", false, "backup-space", "delete", "48201", "--immediate", "--yes", "-o", "json")
	c = m.find("DELETE", "/cloud/backup-spaces/48201")
	if r.code != 0 || c == nil || !strings.Contains(c.Query, "confirm=true") || !strings.Contains(c.Query, "immediate=true") || !strings.Contains(r.out, "cancelling") {
		t.Fatalf("delete --immediate --yes: %d %q %q %+v", r.code, r.out, r.err, c)
	}
	r = run(t, srv, map[string]string{"GYRO_LANG": "tr"}, "", false, "backup-space", "--help")
	if r.code != 0 || !strings.Contains(r.out, "reset-password") || !strings.Contains(r.out, "attach-server") || !strings.Contains(r.out, "credentials") {
		t.Fatalf("help: %d %q", r.code, r.out)
	}
	r = run(t, srv, map[string]string{"GYRO_LANG": "tr"}, "", false, "--help")
	if !strings.Contains(r.out, "Yedekleme Alanı: sunucunuza bağlı SFTP") {
		t.Fatalf("root help group short: %q", r.out)
	}
}
