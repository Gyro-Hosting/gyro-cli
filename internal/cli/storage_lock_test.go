package cli

import (
	"net/http"
	"strings"
	"testing"
)

func TestStorageLockEndHint(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	m.routes["PUT /cloud/storage/12/buckets/arsiv-ab12/object-lock"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 422, map[string]any{"error": map[string]any{"code": "retention_exceeds_service_end",
			"message": "The lock period cannot go beyond the service end date (2026-12-24T21:22:00+03:00); at most 29 days now. Extend the service period (longer billing cycle) for a longer lock.",
			"details": map[string]any{"field": "days", "service_end": "2026-12-24T21:22:00+03:00", "max_days": 29}}})
	}
	r := run(t, srv, map[string]string{"GYRO_LANG": "tr"}, "", false, "storage", "object-lock", "set", "12", "arsiv-ab12", "--mode", "GOVERNANCE", "--days", "400", "--yes")
	if r.code == 0 || !strings.Contains(r.err, "retention_exceeds_service_end") || !strings.Contains(r.err, "en çok 29 gün") || !strings.Contains(r.err, "triennially") {
		t.Errorf("tr hint missing: %d %s", r.code, r.err)
	}
	r = run(t, srv, nil, "", false, "storage", "object-lock", "set", "12", "arsiv-ab12", "--mode", "GOVERNANCE", "--days", "400", "--yes")
	if r.code == 0 || !strings.Contains(r.err, "at most 29 days now") || !strings.Contains(r.err, "2026-12-24T21:22:00+03:00") {
		t.Errorf("en hint missing: %d %s", r.code, r.err)
	}
}
