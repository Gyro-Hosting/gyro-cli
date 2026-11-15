package cli

import (
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
)

// GYR-17: gyro container run / jobs
func TestContainerRunWaitExitCode(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	var got map[string]any
	var polls int32
	m.routes["POST /cloud/containers/77/services/api/run"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		got = b
		j(w, 202, map[string]any{"data": map[string]any{"job": map[string]any{"id": "jabc123456789", "service": "api", "state": "starting", "pinned": true}}})
	}
	m.routes["GET /cloud/containers/77/jobs/jabc123456789"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		st := "running"
		if atomic.AddInt32(&polls, 1) > 1 {
			st = "failed"
		}
		j(w, 200, map[string]any{"data": map[string]any{"id": "jabc123456789", "service": "api", "state": st, "exit_code": 3, "error": "exit_nonzero"}})
	}
	m.routes["GET /cloud/containers/77/jobs/jabc123456789/logs"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"id": "jabc123456789", "state": "running", "log": "migrating\na<b>c\n"}})
	}
	r := run(t, srv, nil, "", false, "container", "run", "77", "--service", "api", "--follow", "--time-limit", "5m", "--cpu", "0.5", "--", "sh", "-c", "test 1 -lt 2 && cat < x > y")
	if r.code != 3 {
		t.Fatalf("exit code must be the job's (3), got %d: %s", r.code, r.err)
	}
	cmd, _ := got["command"].([]any)
	if len(cmd) != 3 || cmd[2] != "test 1 -lt 2 && cat < x > y" || got["timeout_seconds"] != float64(300) || got["cpu_m"] != float64(500) {
		t.Fatalf("body: %#v", got)
	}
	if strings.Count(r.out, "a<b>c") != 1 || !strings.Contains(r.out, "migrating") {
		t.Fatalf("follow output (no duplicates): %q", r.out)
	}
	r = run(t, srv, nil, "", false, "container", "run", "77", "--", "true")
	if r.code != ExitUsage {
		t.Fatalf("--service required: %d", r.code)
	}
}

func TestContainerJobsTable(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	m.routes["GET /cloud/containers/77/jobs"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"limit": 3, "active": 0, "jobs": []any{
			map[string]any{"id": "jabc123456789", "service": "api", "state": "succeeded", "exit_code": 0, "command": []any{"npm", "run", "migrate"}}}}})
	}
	r := run(t, srv, nil, "", false, "container", "jobs", "77")
	if r.code != 0 || !strings.Contains(r.out, "jabc123456789") || !strings.Contains(r.out, "npm run migrate") || !strings.Contains(r.err, "3") {
		t.Fatalf("jobs table: %d %q %q", r.code, r.out, r.err)
	}
}
