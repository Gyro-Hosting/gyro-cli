package cli

import (
	"net/http"
	"strings"
	"testing"
)

// 1.29.3 (#3892): archive pipelines need no --repo; git providers do.
func TestPipelineCreateSourceRules(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	var got map[string]any
	m.routes["POST /cloud/containers/17000/pipelines"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		got = b
		j(w, 201, map[string]any{"data": map[string]any{"pipeline": map[string]any{"id": "pabc12345", "name": b["name"]}}})
	}
	// archive without --repo: sent, no repo key
	r := run(t, srv, nil, "", false, "container", "pipeline", "create", "17000", "--name", "site", "--provider", "archive", "--image", "site")
	if r.code != 0 || got == nil || got["provider"] != "archive" {
		t.Fatalf("archive create: code %d err %s body %v", r.code, r.err, got)
	}
	if _, has := got["repo"]; has {
		t.Fatalf("archive create must not send repo: %v", got)
	}
	// archive with --repo "": empty value dropped
	got = nil
	r = run(t, srv, nil, "", false, "container", "pipeline", "create", "17000", "--name", "site", "--provider", "archive", "--repo", "", "--image", "site")
	if r.code != 0 || got == nil {
		t.Fatalf("archive create with empty repo: %d %s", r.code, r.err)
	}
	if _, has := got["repo"]; has {
		t.Fatalf("empty repo must not be sent: %v", got)
	}
	// archive with a repository: refused locally, explained
	got = nil
	r = run(t, srv, nil, "", false, "container", "pipeline", "create", "17000", "--name", "site", "--provider", "archive", "--repo", "https://github.com/gyro-ornek/uygulama", "--image", "site")
	if r.code != ExitUsage || got != nil || !strings.Contains(r.err, "--provider archive") {
		t.Fatalf("archive + repo must be a usage error: %d %q", r.code, r.err)
	}
	// git (default github) without --repo: usage error naming the archive alternative
	r = run(t, srv, nil, "", false, "container", "pipeline", "create", "17000", "--name", "api", "--image", "api")
	if r.code != ExitUsage || got != nil || !strings.Contains(r.err, "--repo") || !strings.Contains(r.err, "archive") {
		t.Fatalf("github without repo must be a usage error: %d %q", r.code, r.err)
	}
	// github with --repo: sent
	r = run(t, srv, nil, "", false, "container", "pipeline", "create", "17000", "--name", "api", "--repo", "https://github.com/gyro-ornek/uygulama", "--image", "api")
	if r.code != 0 || got == nil || got["repo"] != "https://github.com/gyro-ornek/uygulama" {
		t.Fatalf("github create: %d %s %v", r.code, r.err, got)
	}
	// Turkish message
	got = nil
	r = run(t, srv, map[string]string{"GYRO_LANG": "tr"}, "", false, "container", "pipeline", "create", "17000", "--name", "api", "--provider", "gitlab", "--image", "api")
	if r.code != ExitUsage || !strings.Contains(r.err, "--provider gitlab için --repo zorunlu") {
		t.Fatalf("tr message: %q", r.err)
	}
}
