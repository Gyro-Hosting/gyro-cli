package cli

import (
	"net/http"
	"strings"
	"testing"
)

// 1.28.0: gyro limits + the "how to raise it" line under limit errors (API 2.26.0 details.how_to_increase).
func TestLimitsCommandAndLimitErrorHint(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	m.routes["GET /cloud/account/limits"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"account_id": 12345,
			"new_account": map[string]any{"applies": true, "period_days": 7, "ends_at": "2026-11-21 10:00:00", "servers": map[string]any{"limit": 5, "used": 2}, "container_environments": map[string]any{"limit": 1, "used": 0}},
			"limits": []any{map[string]any{"key": "compute.max_vms", "scope": "account", "limit": 10, "used": 2, "unit": "servers", "custom": true, "how_to_increase": "open_support_ticket", "description": "x"},
				map[string]any{"key": "s3.max_buckets", "scope": "service", "limit": nil, "used": nil, "how_to_increase": "upgrade_plan_or_open_support_ticket",
					"services": []any{map[string]any{"service_id": 12345, "name": "Gyro Proje", "limit": 5, "used": 1}}}}}})
	}
	r := run(t, srv, map[string]string{"GYRO_LANG": "en"}, "", false, "limits")
	if r.code != 0 || !strings.Contains(r.out, "compute.max_vms") || !strings.Contains(r.out, "10*") || !strings.Contains(r.out, "2/5") ||
		!strings.Contains(r.out, "gyro ticket create --department-id 3") || !strings.Contains(r.out, "#12345") {
		t.Fatalf("limits output: %d %s %s", r.code, r.out, r.err)
	}
	m.routes["POST /cloud/ssh-keys"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 422, map[string]any{"error": map[string]any{"code": "limit_exceeded", "message": "New accounts can run at most 5 Istanbul cloud servers.",
			"details": map[string]any{"limit_exceeded": true, "limit_key": "new_account.servers", "limit": 5, "current": 5, "how_to_increase": "open_support_ticket",
				"hint": "This limit can be raised with a support ticket.", "support_ticket": map[string]any{"cli": `gyro ticket create --department-id 3 --subject "Limit increase: new_account.servers" --message "x"`}}}})
	}
	r = run(t, srv, map[string]string{"GYRO_LANG": "tr"}, "", false, "ssh-key", "create", pubFile(t), "--name", "gyro-proje")
	if r.code == 0 || !strings.Contains(r.err, "Sınırı artırmak için: gyro ticket create --department-id 3") || strings.Contains(r.err, "support_ticket:") ||
		!strings.Contains(r.err, "limit_key: new_account.servers") || !strings.Contains(r.err, "gyro limits") {
		t.Fatalf("limit error output: %s", r.err)
	}
}
