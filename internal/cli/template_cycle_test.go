package cli

import (
	"net/http"
	"strings"
	"testing"
)

func TestTemplateDeployCycleRegion(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	var got map[string]any
	m.routes["POST /cloud/templates/9/servers"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		got = b
		if b["cycle"] == "hourly" {
			j(w, 422, map[string]any{"error": map[string]any{"code": "cycle_invalid", "message": "cycle is not offered for this plan.",
				"details": map[string]any{"field": "cycle", "cycles": []any{"monthly", "semiannually", "annually", "biennially"}}}})
			return
		}
		j(w, 202, map[string]any{"data": map[string]any{"dry_run": true, "quote": map[string]any{"product_id": 134, "template_id": 9,
			"cycle": b["cycle"], "billing": "invoice", "total": 386.27, "tax": 0}}})
	}
	r := run(t, srv, nil, "", false, "template", "deploy", "9", "--plan", "gvds-2", "--region", "tr-ist-1", "--cycle", "annually", "--dry-run")
	if r.code != 0 || got["cycle"] != "annually" || got["region"] != "tr-ist-1" || got["plan"] != "gvds-2" || got["dry_run"] != true {
		t.Fatalf("body not sent: %d %v %s", r.code, got, r.err)
	}
	r = run(t, srv, map[string]string{"GYRO_LANG": "tr"}, "", false, "template", "deploy", "9", "--plan", "gvds-2", "--cycle", "hourly", "--dry-run")
	if r.code == 0 || !strings.Contains(r.err, "cycle_invalid") || !strings.Contains(r.err, "monthly, semiannually, annually, biennially") {
		t.Errorf("tr hint missing: %d %s", r.code, r.err)
	}
	r = run(t, srv, nil, "", false, "template", "deploy", "9", "--plan", "gvds-2", "--cycle", "hourly", "--dry-run")
	if r.code == 0 || !strings.Contains(r.err, "This plan offers the periods") {
		t.Errorf("en hint missing: %d %s", r.code, r.err)
	}
}
