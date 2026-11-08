// Package api is the HTTP client used by the gyro CLI: bearer auth, JSON envelope, retries with
// backoff on 429/5xx, automatic Idempotency-Key and client request ids, pagination.
package api

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// DefaultURL is the production client API base address.
const DefaultURL = "https://www.gyrohosting.com/api/v1/client"

// Client talks to the Gyro Public API.
type Client struct {
	BaseURL    string
	Token      string
	UserAgent  string
	Lang       string // Accept-Language
	HTTP       *http.Client
	MaxRetries int
	Debug      io.Writer // when set, requests/responses are logged here (Authorization masked)
	Sleep      func(time.Duration)
}

// New returns a client with sane defaults.
func New(baseURL, token, userAgent string) *Client {
	if baseURL == "" {
		baseURL = DefaultURL
	}
	return &Client{
		BaseURL:    strings.TrimRight(baseURL, "/"),
		Token:      token,
		UserAgent:  userAgent,
		HTTP:       &http.Client{Timeout: 120 * time.Second},
		MaxRetries: 4,
		Sleep:      time.Sleep,
	}
}

// Error is an API error ({"error": {...}}) or a transport failure after retries.
type Error struct {
	Status    int            `json:"status"`
	Code      string         `json:"code"`
	Message   string         `json:"message"`
	Details   map[string]any `json:"details,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
	Method    string         `json:"-"`
	Path      string         `json:"-"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d %s: %s", e.Status, e.Code, e.Message)
}

// AsError unwraps an *Error.
func AsError(err error) (*Error, bool) {
	var e *Error
	ok := errors.As(err, &e)
	return e, ok
}

// PageMeta is the pagination block of list answers.
type PageMeta struct {
	Total    int `json:"total"`
	Page     int `json:"page"`
	Limit    int `json:"limit"`
	NextPage int `json:"next_page"`
}

// Request describes one call.
type Request struct {
	Method string
	Path   string // relative to BaseURL, e.g. /cloud/servers
	Query  url.Values
	Body   any
	// IdempotencyKey makes a write safe to retry: the server stores the first answer for 24 h and replays it.
	// The CLI sets one for money operations (and whenever the spec requires the header). It is NOT sent for
	// calls that return one-time secrets (root password, console ticket), because the server withholds
	// secrets on idempotent requests.
	IdempotencyKey string
}

// Response is a decoded success answer.
type Response struct {
	Status    int
	Data      json.RawMessage
	Meta      *PageMeta
	RequestID string
	Replayed  bool
	Header    http.Header
}

// NewID returns a random UUIDv4.
func NewID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	h := hex.EncodeToString(b)
	return h[0:8] + "-" + h[8:12] + "-" + h[12:16] + "-" + h[16:20] + "-" + h[20:32]
}

// URL builds the absolute URL of a request (used by --dry-run printing too).
func (c *Client) URL(path string, q url.Values) string {
	u := c.BaseURL + "/" + strings.TrimLeft(path, "/")
	if len(q) > 0 {
		u += "?" + q.Encode()
	}
	return u
}

// Do performs the call. GETs (and writes carrying an Idempotency-Key, which the server replays instead
// of repeating) are retried on network errors, 429 and 5xx with exponential backoff (Retry-After honoured).
func (c *Client) Do(ctx context.Context, r Request) (*Response, error) {
	var raw []byte
	if r.Body != nil {
		var err error
		if raw, err = json.Marshal(r.Body); err != nil {
			return nil, err
		}
	}
	method := strings.ToUpper(r.Method)
	if method == "" {
		method = http.MethodGet
	}
	u := c.URL(r.Path, r.Query)
	reqID := NewID()
	safe := method == http.MethodGet || r.IdempotencyKey != ""
	var lastErr error
	for attempt := 0; attempt <= c.MaxRetries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, method, u, bytes.NewReader(raw))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("X-Request-Id", reqID)
		if c.Token != "" {
			req.Header.Set("Authorization", "Bearer "+c.Token)
		}
		if c.UserAgent != "" {
			req.Header.Set("User-Agent", c.UserAgent)
		}
		if c.Lang != "" {
			req.Header.Set("Accept-Language", c.Lang)
		}
		if r.Body != nil {
			req.Header.Set("Content-Type", "application/json")
		}
		if r.IdempotencyKey != "" {
			req.Header.Set("Idempotency-Key", r.IdempotencyKey)
		}
		c.logRequest(req, raw, attempt)
		resp, err := c.HTTP.Do(req)
		if err != nil {
			lastErr = err
			if !safe || ctx.Err() != nil {
				return nil, &Error{Code: "network_error", Message: err.Error(), RequestID: reqID, Method: method, Path: r.Path}
			}
			c.sleep(ctx, backoff(attempt, 0))
			continue
		}
		data, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
		resp.Body.Close()
		c.logResponse(resp, data)
		if rid := firstHeader(resp.Header, "X-Request-Id", "X-Gyro-Request-Id", "Request-Id"); rid != "" {
			reqID = rid
		}
		var env struct {
			Data  json.RawMessage `json:"data"`
			Meta  *PageMeta       `json:"meta"`
			Error *Error          `json:"error"`
		}
		decodeErr := error(nil)
		if len(bytes.TrimSpace(data)) > 0 {
			decodeErr = json.Unmarshal(data, &env)
		}
		if resp.StatusCode >= 200 && resp.StatusCode < 300 && env.Error == nil {
			if decodeErr != nil {
				return nil, &Error{Status: resp.StatusCode, Code: "bad_response", Message: snippet(data), RequestID: reqID, Method: method, Path: r.Path}
			}
			return &Response{Status: resp.StatusCode, Data: env.Data, Meta: env.Meta, RequestID: reqID,
				Replayed: resp.Header.Get("Idempotency-Replayed") == "true", Header: resp.Header}, nil
		}
		apiErr := env.Error
		if apiErr == nil {
			apiErr = &Error{Code: strings.ToLower(strings.ReplaceAll(http.StatusText(resp.StatusCode), " ", "_")), Message: snippet(data)}
		}
		apiErr.Status, apiErr.RequestID, apiErr.Method, apiErr.Path = resp.StatusCode, reqID, method, r.Path
		lastErr = apiErr
		retryAfter, _ := strconv.Atoi(resp.Header.Get("Retry-After"))
		switch {
		case resp.StatusCode == http.StatusTooManyRequests:
		case resp.StatusCode == http.StatusConflict && apiErr.Code == "idempotency_in_flight":
		case resp.StatusCode >= 500 && safe:
		default:
			return nil, apiErr
		}
		if attempt == c.MaxRetries || ctx.Err() != nil {
			break
		}
		c.sleep(ctx, backoff(attempt, retryAfter))
	}
	if e, ok := AsError(lastErr); ok {
		return nil, e
	}
	return nil, &Error{Code: "network_error", Message: fmt.Sprint(lastErr), RequestID: reqID, Method: method, Path: r.Path}
}

// Get decodes data of a GET into out.
func (c *Client) Get(ctx context.Context, path string, q url.Values, out any) (*Response, error) {
	resp, err := c.Do(ctx, Request{Method: "GET", Path: path, Query: q})
	if err != nil {
		return nil, err
	}
	if out != nil && len(resp.Data) > 0 {
		if err := json.Unmarshal(resp.Data, out); err != nil {
			return resp, fmt.Errorf("decoding %s: %w", path, err)
		}
	}
	return resp, nil
}

// All walks every page of a paginated GET (meta.next_page) and returns the concatenated items.
// itemsKey is "" for {"data": [...]} or e.g. "items" for {"data": {"items": [...]}}.
func (c *Client) All(ctx context.Context, path string, q url.Values, itemsKey string, maxPages int) ([]json.RawMessage, error) {
	q2 := url.Values{}
	for k, v := range q {
		q2[k] = v
	}
	if q2.Get("limit") == "" {
		q2.Set("limit", "100")
	}
	page := 1
	if p, _ := strconv.Atoi(q2.Get("page")); p > 0 {
		page = p
	}
	var out []json.RawMessage
	for n := 0; maxPages <= 0 || n < maxPages; n++ {
		q2.Set("page", strconv.Itoa(page))
		resp, err := c.Do(ctx, Request{Method: "GET", Path: path, Query: q2})
		if err != nil {
			return nil, err
		}
		items, err := Items(resp.Data, itemsKey)
		if err != nil {
			return nil, err
		}
		out = append(out, items...)
		if resp.Meta == nil || resp.Meta.NextPage <= page || len(items) == 0 {
			break
		}
		page = resp.Meta.NextPage
	}
	return out, nil
}

// Items extracts the list from a data block.
func Items(data json.RawMessage, itemsKey string) ([]json.RawMessage, error) {
	if itemsKey != "" {
		var wrap map[string]json.RawMessage
		if err := json.Unmarshal(data, &wrap); err != nil {
			return nil, err
		}
		data = wrap[itemsKey]
	}
	var items []json.RawMessage
	if len(data) == 0 || string(data) == "null" {
		return nil, nil
	}
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (c *Client) sleep(ctx context.Context, d time.Duration) {
	if c.Sleep != nil {
		c.Sleep(d)
		return
	}
	select {
	case <-ctx.Done():
	case <-time.After(d):
	}
}

func backoff(attempt, retryAfter int) time.Duration {
	if retryAfter > 0 {
		if retryAfter > 60 {
			retryAfter = 60
		}
		return time.Duration(retryAfter) * time.Second
	}
	d := time.Duration(1<<attempt) * 500 * time.Millisecond
	if d > 10*time.Second {
		d = 10 * time.Second
	}
	return d
}

func firstHeader(h http.Header, names ...string) string {
	for _, n := range names {
		if v := h.Get(n); v != "" {
			return v
		}
	}
	return ""
}

func snippet(b []byte) string {
	s := strings.TrimSpace(string(b))
	if len(s) > 200 {
		s = s[:200] + "..."
	}
	if s == "" {
		s = "empty response"
	}
	return s
}

var secretKeys = []string{"password", "secret", "token", "private_key", "auth_code", "epp"}

// Redact masks secret-looking JSON fields (used for --debug logging).
func Redact(b []byte) []byte {
	var v any
	if json.Unmarshal(b, &v) != nil {
		return b
	}
	var walk func(any) any
	walk = func(x any) any {
		switch t := x.(type) {
		case map[string]any:
			for k, val := range t {
				lk := strings.ToLower(k)
				masked := false
				for _, s := range secretKeys {
					if strings.Contains(lk, s) && val != nil && val != "" {
						if _, isBool := val.(bool); !isBool {
							t[k] = "***"
							masked = true
						}
						break
					}
				}
				if !masked {
					t[k] = walk(val)
				}
			}
		case []any:
			for i := range t {
				t[i] = walk(t[i])
			}
		}
		return x
	}
	out, _ := json.Marshal(walk(v))
	return out
}

func (c *Client) logRequest(req *http.Request, body []byte, attempt int) {
	if c.Debug == nil {
		return
	}
	fmt.Fprintf(c.Debug, "> %s %s (attempt %d)\n", req.Method, req.URL, attempt+1)
	for k, v := range req.Header {
		val := strings.Join(v, ",")
		if k == "Authorization" {
			val = "Bearer " + MaskToken(c.Token)
		}
		fmt.Fprintf(c.Debug, "> %s: %s\n", k, val)
	}
	if len(body) > 0 {
		fmt.Fprintf(c.Debug, "> %s\n", Redact(body))
	}
}

func (c *Client) logResponse(resp *http.Response, body []byte) {
	if c.Debug == nil {
		return
	}
	fmt.Fprintf(c.Debug, "< %s\n", resp.Status)
	for _, k := range []string{"X-Request-Id", "X-Ratelimit-Remaining", "Retry-After", "Idempotency-Replayed", "X-Gyro-Api-Version"} {
		if v := resp.Header.Get(k); v != "" {
			fmt.Fprintf(c.Debug, "< %s: %s\n", k, v)
		}
	}
	b := Redact(body)
	if len(b) > 2000 {
		b = append(b[:2000], "..."...)
	}
	fmt.Fprintf(c.Debug, "< %s\n", b)
}

// MaskToken shows only the prefix and the last 4 characters of a key.
func MaskToken(t string) string {
	if t == "" {
		return ""
	}
	if len(t) <= 10 {
		return "****"
	}
	return t[:4] + "…" + t[len(t)-4:]
}
