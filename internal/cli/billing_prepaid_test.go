package cli

import (
	"encoding/base64"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPrepaidBilling(t *testing.T) {
	setup(t)
	m, srv := newMock(t)
	m.routes["GET /cloud/billing/balance"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"balance": map[string]any{"amount": 424.48, "currency": "TRY"}, "hourly_cost": map[string]any{"amount": 1.48, "currency": "TRY"},
			"daily_cost": map[string]any{"amount": 35.54, "currency": "TRY"}, "upcoming_invoices_7d": map[string]any{"amount": 0, "currency": "TRY"}, "hours_left": 286.6,
			"runs_out_at": "2026-11-23T13:42:40+03:00", "level": "ok", "services": map[string]any{"active": 2, "suspended": 0}, "topup_url": "https://x/ac-credits"}})
	}
	m.routes["GET /cloud/billing/usage"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		if r.URL.Query().Get("granularity") != "day" || r.URL.Query().Get("service_id") != "17" {
			j(w, 422, errBody("bad", r.URL.RawQuery))
			return
		}
		j(w, 200, map[string]any{"data": map[string]any{"items": []any{map[string]any{"period": "2026-11-10", "hours": 24, "amount": map[string]any{"amount": 13.44, "currency": "TRY"}}},
			"totals": []any{map[string]any{"amount": 13.44, "currency": "TRY"}}}, "meta": map[string]any{"total": 1, "page": 1, "limit": 100, "next_page": 0}})
	}
	var lastEst map[string]any
	m.routes["POST /cloud/billing/estimate"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		lastEst = b
		items, _ := b["items"].([]any)
		if len(items) != 2 {
			j(w, 422, errBody("bad", "items"))
			return
		}
		j(w, 200, map[string]any{"data": map[string]any{"items": []any{}, "hourly": map[string]any{"amount": 1.68, "currency": "TRY"}, "daily": map[string]any{"amount": 40.32, "currency": "TRY"},
			"monthly_cap": map[string]any{"amount": 1128.91, "currency": "TRY"}, "balance": map[string]any{"amount": 10, "currency": "TRY"}, "hours_covered": 5.9, "can_order": false,
			"order_requires": map[string]any{"amount": 40.32, "currency": "TRY"}}})
	}
	m.routes["GET /cloud/billing/statement/2026-11"] = func(w http.ResponseWriter, r *http.Request, b map[string]any) {
		j(w, 200, map[string]any{"data": map[string]any{"filename": "statement-2026-11.pdf", "content_base64": base64.StdEncoding.EncodeToString([]byte("%PDF-1.7 test"))}})
	}
	r := run(t, srv, map[string]string{"GYRO_LANG": "tr"}, "", false, "billing", "balance")
	if r.code != 0 || !strings.Contains(r.out, "11 gün 22 saat") || !strings.Contains(r.out, "424.48 TRY") {
		t.Fatalf("balance: %d %q %q", r.code, r.out, r.err)
	}
	r = run(t, srv, nil, "", false, "billing", "usage", "--daily", "--service", "17")
	if r.code != 0 || !strings.Contains(r.out, "2026-11-10") || !strings.Contains(r.out, "Total: 13.44 TRY") {
		t.Fatalf("usage: %d %q %q", r.code, r.out, r.err)
	}
	r = run(t, srv, nil, "", false, "billing", "estimate", "--item", "standart-2-4:2", "--plan", "standart-4-8")
	if r.code != 0 || !strings.Contains(r.out, "first 24 hours") {
		t.Fatalf("estimate: %d %q %q", r.code, r.out, r.err)
	}
	r = run(t, srv, nil, "", false, "billing", "estimate", "--item", "250/512:3", "--cpu-m", "500", "--memory-mi", "1024", "--hours", "100")
	if r.code != 0 {
		t.Fatalf("estimate serverless: %d %q %q", r.code, r.out, r.err)
	}
	if it, _ := lastEst["items"].([]any); len(it) != 2 || it[0].(map[string]any)["type"] != "serverless" || it[0].(map[string]any)["cpu_m"].(float64) != 250 ||
		it[0].(map[string]any)["count"].(float64) != 3 || it[1].(map[string]any)["memory_mi"].(float64) != 1024 || it[1].(map[string]any)["hours"].(float64) != 100 {
		t.Fatalf("serverless body: %v", lastEst)
	}
	if r = run(t, srv, nil, "", false, "billing", "estimate", "--item", "x/512"); r.code == 0 {
		t.Fatalf("bad serverless item accepted")
	}
	dst := filepath.Join(t.TempDir(), "s.pdf")
	r = run(t, srv, nil, "", false, "billing", "statement", "2026-11", "--pdf", dst)
	if b, _ := os.ReadFile(dst); r.code != 0 || string(b) != "%PDF-1.7 test" {
		t.Fatalf("statement: %d %q %q", r.code, r.out, r.err)
	}
	if r = run(t, srv, nil, "", false, "billing", "statement", "2026-13"); r.code == 0 {
		t.Fatalf("bad month accepted")
	}
}
