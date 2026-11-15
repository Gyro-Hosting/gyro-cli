package cli

import (
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
)

// VPC V2: gyro container vpc attach/expose/detach, vpc port-forward update/add --source
func TestContainerVpcAttachWait(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	var got map[string]any
	var polls int32
	m.routes["PUT /cloud/containers/77/vpc"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		got = b
		j(w, 200, map[string]any{"data": map[string]any{"status": "successful", "vpc": map[string]any{"attached": true, "state": "attaching"}}})
	}
	m.routes["GET /cloud/containers/77/vpc"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		st := "attaching"
		if atomic.AddInt32(&polls, 1) > 1 {
			st = "ready"
		}
		j(w, 200, map[string]any{"data": map[string]any{"attached": true, "state": st, "source_range": "10.10.1.16/28", "endpoint_ip": "10.10.1.16"}})
	}
	r := run(t, srv, nil, "", false, "container", "vpc", "attach", "77", "--vpc", "vpc-1", "--subnet", "sn-1", "--expose", "API:8080", "--expose", "web:80,​", "--wait", "-o", "json")
	if r.code != 0 || !strings.Contains(r.out, "10.10.1.16/28") || !strings.Contains(r.out, "ready") {
		t.Fatalf("attach --wait: %d %q %q", r.code, r.out, r.err)
	}
	ex, _ := got["expose"].([]any)
	if got["subnet_id"] != "sn-1" || got["vpc_id"] != "vpc-1" || len(ex) != 2 || ex[0].(map[string]any)["service"] != "api" || ex[1].(map[string]any)["port"] != float64(80) {
		t.Fatalf("body: %#v", got)
	}
	r = run(t, srv, nil, "", false, "container", "vpc", "attach", "77", "--subnet", "sn-1", "--expose", "api:22x")
	if r.code != ExitUsage {
		t.Fatalf("bad expose must be a usage error: %d", r.code)
	}
	r = run(t, srv, nil, "", false, "container", "vpc", "attach", "77", "--subnet", "sn-1", "--expose", "a:80", "--expose", "b:80")
	if r.code != ExitUsage {
		t.Fatalf("duplicate port: %d", r.code)
	}
	r = run(t, srv, nil, "", false, "container", "vpc", "attach", "abc", "--subnet", "sn-1")
	if r.code != ExitUsage {
		t.Fatalf("env id must be numeric: %d", r.code)
	}
}

func TestContainerVpcExposeDetach(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	var patch, del map[string]any
	m.routes["PATCH /cloud/containers/77/vpc"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		patch = b
		j(w, 200, map[string]any{"data": map[string]any{"status": "successful", "vpc": map[string]any{"attached": true, "state": "ready"}}})
	}
	m.routes["DELETE /cloud/containers/77/vpc"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		del = b
		j(w, 200, map[string]any{"data": map[string]any{"status": "successful", "vpc": map[string]any{"attached": false, "state": "detaching"}}})
	}
	r := run(t, srv, nil, "", false, "container", "vpc", "expose", "77", "--clear")
	if l, ok := patch["expose"].([]any); r.code != 0 || !ok || len(l) != 0 {
		t.Fatalf("--clear sends []: %d %#v %q", r.code, patch, r.err)
	}
	r = run(t, srv, nil, "", false, "container", "vpc", "expose", "77")
	if r.code != ExitUsage {
		t.Fatalf("expose without list: %d", r.code)
	}
	r = run(t, srv, nil, "", false, "container", "vpc", "detach", "77")
	if r.code != ExitAborted || del != nil {
		t.Fatalf("detach asks first (no terminal -> aborted): %d", r.code)
	}
	r = run(t, srv, nil, "", false, "container", "vpc", "detach", "77", "--yes")
	if r.code != 0 || del["confirm"] != true {
		t.Fatalf("detach --yes: %d %#v", r.code, del)
	}
}

func TestVpcPortForwardSources(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	var upd, add map[string]any
	m.routes["PATCH /cloud/vpcs/vpc-1/port-forwards/pf-1"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		upd = b
		j(w, 200, map[string]any{"data": map[string]any{"id": "pf-1", "source_cidrs": b["source_cidrs"]}})
	}
	m.routes["POST /cloud/vpcs/vpc-1/port-forwards"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		add = b
		j(w, 201, map[string]any{"data": map[string]any{"id": "pf-2"}})
	}
	r := run(t, srv, nil, "", false, "vpc", "port-forward", "update", "vpc-1", "pf-1", "--source", "203.0.113.10", "--source", "198.51.100.0/24, 203.0.113.10")
	l, _ := upd["source_cidrs"].([]any)
	if r.code != 0 || len(l) != 2 || l[0] != "203.0.113.10/32" || l[1] != "198.51.100.0/24" {
		t.Fatalf("update --source: %d %#v %q", r.code, upd, r.err)
	}
	r = run(t, srv, nil, "", false, "vpc", "port-forward", "update", "vpc-1", "pf-1", "--any")
	if l, ok := upd["source_cidrs"].([]any); r.code != 0 || !ok || len(l) != 0 {
		t.Fatalf("--any sends []: %#v", upd)
	}
	r = run(t, srv, nil, "", false, "vpc", "port-forward", "update", "vpc-1", "pf-1")
	if r.code != ExitUsage {
		t.Fatalf("update needs --source or --any: %d", r.code)
	}
	r = run(t, srv, nil, "", false, "vpc", "port-forward", "add", "vpc-1", "--protocol", "tcp", "--public-port", "2222", "--target-ip", "10.10.1.5", "--target-port", "22", "--source", "203.0.113.10/32", "--yes")
	if l, _ := add["source_cidrs"].([]any); r.code != 0 || len(l) != 1 || l[0] != "203.0.113.10/32" {
		t.Fatalf("add --source alias: %d %#v %q", r.code, add, r.err)
	}
}
