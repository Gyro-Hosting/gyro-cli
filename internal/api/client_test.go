package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestRetries(t *testing.T) {
	n := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n++
		if r.Method == "GET" && n < 3 {
			w.WriteHeader(503)
			w.Write([]byte(`{"error":{"code":"backend_error","message":"x"}}`))
			return
		}
		if r.Method == "POST" {
			w.WriteHeader(502)
			w.Write([]byte(`{"error":{"code":"backend_error","message":"y"}}`))
			return
		}
		w.Write([]byte(`{"data":{"ok":true}}`))
	}))
	defer srv.Close()
	c := New(srv.URL, "wck_x", "t")
	c.Sleep = func(time.Duration) {}
	if _, err := c.Do(context.Background(), Request{Method: "GET", Path: "/a"}); err != nil || n != 3 {
		t.Fatalf("GET retry: %v n=%d", err, n)
	}
	n = 0
	_, err := c.Do(context.Background(), Request{Method: "POST", Path: "/b", Body: map[string]any{}})
	if e, ok := AsError(err); !ok || e.Status != 502 || n != 1 {
		t.Fatalf("POST without key must not retry on 5xx: %v n=%d", err, n)
	}
	n = 0
	_, err = c.Do(context.Background(), Request{Method: "POST", Path: "/b", Body: map[string]any{}, IdempotencyKey: "k-123456789"})
	if n != c.MaxRetries+1 || err == nil {
		t.Fatalf("POST with key retries: n=%d err=%v", n, err)
	}
}

func TestRequestIDAndNonJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Request-Id") == "" {
			t.Error("no X-Request-Id sent")
		}
		w.WriteHeader(500)
		w.Write([]byte("<html>gateway</html>"))
	}))
	defer srv.Close()
	c := New(srv.URL, "", "t")
	c.Sleep = func(time.Duration) {}
	c.MaxRetries = 0
	_, err := c.Do(context.Background(), Request{Method: "GET", Path: "/"})
	e, ok := AsError(err)
	if !ok || e.RequestID == "" || !strings.Contains(e.Message, "gateway") {
		t.Fatalf("err %+v", e)
	}
}

func TestRedactAndMask(t *testing.T) {
	out := string(Redact([]byte(`{"root_password":"abc","x":{"secret":"s","n":1},"root_password_withheld":true}`)))
	if strings.Contains(out, "abc") || strings.Contains(out, `"s"`) || !strings.Contains(out, `"root_password_withheld":true`) {
		t.Errorf("redact: %s", out)
	}
	if MaskToken("wck_abcdefghijklmnop") != "wck_…mnop" {
		t.Error(MaskToken("wck_abcdefghijklmnop"))
	}
}
