package cli

import (
	"net/http"
	"strings"
	"testing"
)

// GYR-12: gyro container schedule create / update / list
func TestContainerScheduleCreate(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	var got []map[string]any
	m.routes["POST /cloud/containers/77/schedules"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		got = append(got, b)
		j(w, 200, map[string]any{"data": map[string]any{"schedule": map[string]any{"id": "stoj02crbr9i", "name": b["name"], "kind": "cron", "cron": b["cron"], "next_run_local": "2026-11-16T03:00:00+03:00"}}})
	}
	r := run(t, srv, nil, "", false, "container", "schedule", "create", "77", "nightly", "--cron", "0 3 * * *", "--service", "api", "--retries", "2", "--", "sh", "-c", "a < b && c > d")
	if r.code != 0 || !strings.Contains(r.err+r.out, "stoj02crbr9i") {
		t.Fatalf("create: %d %q %q", r.code, r.out, r.err)
	}
	b := got[0]
	if b["name"] != "nightly" || b["cron"] != "0 3 * * *" || b["service"] != "api" || b["retries"] != float64(2) || b["target"] != "service" {
		t.Fatalf("body: %#v", b)
	}
	if c, _ := b["command"].([]any); len(c) != 3 || c[2] != "a < b && c > d" {
		t.Fatalf("command: %#v", b["command"])
	}
	r = run(t, srv, nil, "", false, "container", "schedule", "create", "77", "ping", "--every", "15m", "--url", "https://example.com/cron", "--method", "post", "--header", "Authorization: Bearer x:y")
	b = got[1]
	if r.code != 0 || b["rate_minutes"] != float64(15) || b["target"] != "http" || b["method"] != "POST" || b["headers"].(map[string]any)["Authorization"] != "Bearer x:y" {
		t.Fatalf("http body: %d %#v %s", r.code, b, r.err)
	}
	for _, bad := range [][]string{{"77", "x", "--service", "api"}, {"77", "x", "--cron", "* * * * *", "--every", "5m", "--service", "api"}, {"77", "x", "--cron", "0 3 * * *"},
		{"77", "x", "--every", "90s", "--service", "api"}, {"77", "x", "--cron", "0 3 * * *", "--url", "https://e.com", "--", "true"}} {
		if r := run(t, srv, nil, "", false, append([]string{"container", "schedule", "create"}, bad...)...); r.code != ExitUsage {
			t.Fatalf("%v must be a usage error, got %d %s", bad, r.code, r.err)
		}
	}
}

func TestContainerScheduleUpdateAndList(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	var got map[string]any
	m.routes["PATCH /cloud/containers/77/schedules/nightly"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		got = b
		j(w, 200, map[string]any{"data": map[string]any{"schedule": map[string]any{"id": "stoj02crbr9i", "name": "nightly", "kind": "rate", "rate_minutes": 30}}})
	}
	m.routes["GET /cloud/containers/77/schedules"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"limit": 10, "schedules": []any{
			map[string]any{"id": "stoj02crbr9i", "name": "nightly", "state": "active", "kind": "cron", "cron": "0 3 * * *", "timezone": "UTC", "target": "service", "service": "api",
				"command": []any{"npm", "run", "report"}, "next_run_local": "2026-11-16T03:00:00Z", "last_run": map[string]any{"state": "succeeded"}}}}})
	}
	r := run(t, srv, nil, "", false, "container", "schedule", "update", "77", "nightly", "--every", "30m")
	if r.code != 0 || len(got) != 1 || got["rate_minutes"] != float64(30) {
		t.Fatalf("update sends only what changed: %d %#v %s", r.code, got, r.err)
	}
	r = run(t, srv, nil, "", false, "container", "schedule", "list", "77")
	if r.code != 0 || !strings.Contains(r.out, "nightly") || !strings.Contains(r.out, "npm run report") || !strings.Contains(r.out, "succeeded") || !strings.Contains(r.err, "1 of 10") {
		t.Fatalf("list: %d %q %q", r.code, r.out, r.err)
	}
}
