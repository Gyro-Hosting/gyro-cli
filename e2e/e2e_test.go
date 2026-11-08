//go:build e2e

// End-to-end tests against a real Gyro API installation. They build the real binary and run it.
//
//	GYRO_E2E_URL=https://.../api/v1/client GYRO_E2E_KEYS=keys.json go test -tags e2e ./e2e -v
//	GYRO_E2E_REAL=1 ...   # also orders, uses and deletes a real (smallest) cloud server
//
// Keys are read from the 0600 key file and passed through the environment only; they are never logged.
package e2e

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"
)

var devURL = os.Getenv("GYRO_E2E_URL")

var bin string

func TestMain(m *testing.M) {
	dir, _ := os.MkdirTemp("", "gyro-e2e")
	bin = filepath.Join(dir, "gyro")
	out, err := exec.Command("go", "build", "-o", bin, "../cmd/gyro").CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		os.Exit(1)
	}
	code := m.Run()
	os.RemoveAll(dir)
	os.Exit(code)
}

func keys(t *testing.T) map[string]any {
	p := os.Getenv("GYRO_E2E_KEYS")
	if p == "" || devURL == "" {
		t.Skip("set GYRO_E2E_URL and GYRO_E2E_KEYS")
	}
	b, err := os.ReadFile(p)
	if err != nil {
		t.Skip("no dev key file:", err)
	}
	var k map[string]any
	if err := json.Unmarshal(b, &k); err != nil {
		t.Fatal(err)
	}
	return k
}

type res struct {
	code     int
	out, err string
}

var secretRe = regexp.MustCompile(`wck_[A-Za-z0-9]+`)

func gyro(t *testing.T, key string, stdin string, args ...string) res {
	t.Helper()
	k := keys(t)
	cmd := exec.Command(bin, args...)
	cmd.Env = append(os.Environ(), "GYRO_API_URL="+devURL, "GYRO_LANG=en", "NO_COLOR=1",
		"GYRO_CONFIG_FILE="+filepath.Join(t.TempDir(), "c.toml"))
	if key != "" {
		cmd.Env = append(cmd.Env, "GYRO_TOKEN="+fmt.Sprint(k[key]))
	}
	cmd.Stdin = strings.NewReader(stdin)
	var o, e strings.Builder
	cmd.Stdout, cmd.Stderr = &o, &e
	err := cmd.Run()
	code := 0
	if ee, ok := err.(*exec.ExitError); ok {
		code = ee.ExitCode()
	} else if err != nil {
		t.Fatal(err)
	}
	r := res{code, o.String(), e.String()}
	if secretRe.MatchString(r.out + r.err) {
		t.Fatal("an API key appeared in the output")
	}
	t.Logf("gyro %s -> %d", strings.Join(args, " "), code)
	return r
}

func mustJSON(t *testing.T, s string, v any) {
	t.Helper()
	if err := json.Unmarshal([]byte(s), v); err != nil {
		t.Fatalf("not JSON: %v\n%s", err, s)
	}
}

func TestReadCommands(t *testing.T) {
	for _, args := range [][]string{
		{"whoami"}, {"context"}, {"server", "list"}, {"plan", "list"}, {"plan", "list", "--location", "falkenstein"},
		{"location", "list"}, {"image", "list"}, {"ssh-key", "list"}, {"balance"}, {"balance", "transaction", "list"},
		{"invoice", "list", "--all"}, {"invoice", "list", "--unpaid"}, {"activity", "list", "--limit", "5"}, {"api", "operations"},
	} {
		r := gyro(t, "a_full", "", args...)
		if r.code != 0 {
			t.Errorf("%v: %d %s", args, r.code, r.err)
		}
	}
	var who struct {
		OwnerID int `json:"owner_id"`
	}
	r := gyro(t, "a_full", "", "whoami", "-o", "json")
	mustJSON(t, r.out, &who)
	if who.OwnerID == 0 {
		t.Error("whoami owner_id")
	}
	r = gyro(t, "a_full", "", "plan", "list", "-o", "json", "--query", "length(@)")
	if strings.TrimSpace(r.out) == "0" {
		t.Error("no plans")
	}
}

func TestErrorsAndIsolation(t *testing.T) {
	r := gyro(t, "a_ro", "", "server", "create", "--plan", "standart-1-1", "--location", "fsn1", "--image", "debian-12", "--yes")
	if r.code != 3 || !strings.Contains(r.err, "insufficient_scope") {
		t.Errorf("read-only key ordering: %d %s", r.code, r.err)
	}
	// B must not see A's servers
	var ids []int
	mustJSON(t, gyro(t, "a_full", "", "server", "list", "-o", "json", "--query", "[].id").out, &ids)
	if len(ids) > 0 {
		r = gyro(t, "b_full", "", "server", "get", fmt.Sprint(ids[0]))
		if r.code != 4 || !strings.Contains(r.err, "request id") {
			t.Errorf("cross-tenant get: %d %s", r.code, r.err)
		}
	}
	r = gyro(t, "a_ipblock", "", "server", "list")
	if r.code != 3 || !strings.Contains(r.err, "ip_not_allowed") {
		t.Errorf("ip block: %d %s", r.code, r.err)
	}
	r = gyro(t, "", "", "server", "list")
	if r.code != 3 {
		t.Errorf("no key: %d", r.code)
	}
	r = gyro(t, "a_full", "", "server", "create", "--plan", "standart-1-1", "--location", "fsn1", "--image", "debian-12")
	if r.code != 7 {
		t.Errorf("order without --yes and no terminal must abort: %d", r.code)
	}
}

func TestDryRunQuote(t *testing.T) {
	r := gyro(t, "a_full", "", "server", "create", "e2e-quote", "--plan", "standart-1-1", "--location", "falkenstein", "--image", "ubuntu-24.04", "--dry-run", "-o", "json")
	if r.code != 0 {
		t.Fatalf("quote: %d %s", r.code, r.err)
	}
	var q struct {
		DryRun bool `json:"dry_run"`
		Quote  struct {
			Total struct{ Amount float64 } `json:"total"`
		} `json:"quote"`
	}
	mustJSON(t, r.out, &q)
	if !q.DryRun || q.Quote.Total.Amount <= 0 {
		t.Errorf("quote: %s", r.out)
	}
}

func TestSSHKeyLifecycle(t *testing.T) {
	dir := t.TempDir()
	kf := filepath.Join(dir, "k")
	if out, err := exec.Command("ssh-keygen", "-q", "-t", "ed25519", "-N", "", "-C", "gyro-cli-e2e", "-f", kf).CombinedOutput(); err != nil {
		t.Skip("ssh-keygen:", string(out))
	}
	name := fmt.Sprintf("gyro-cli-e2e-%d", time.Now().Unix())
	r := gyro(t, "a_full", "", "ssh-key", "create", kf+".pub", "--name", name, "-o", "json")
	if r.code != 0 {
		t.Fatalf("create: %d %s", r.code, r.err)
	}
	var k struct{ ID int }
	mustJSON(t, r.out, &k)
	r = gyro(t, "a_full", "", "ssh-key", "create", kf+".pub", "--name", name+"-dup")
	if r.code != 5 {
		t.Errorf("duplicate key should be 409: %d %s", r.code, r.err)
	}
	r = gyro(t, "a_full", "", "ssh-key", "get", name, "-o", "json")
	if r.code != 0 || !strings.Contains(r.out, "gyro-cli-e2e") {
		t.Errorf("get by name: %d %s", r.code, r.err)
	}
	r = gyro(t, "a_full", "", "ssh-key", "delete", name)
	if r.code != 7 {
		t.Errorf("delete without --yes: %d", r.code)
	}
	r = gyro(t, "a_full", "", "ssh-key", "delete", fmt.Sprint(k.ID), "--yes")
	if r.code != 0 {
		t.Errorf("delete: %d %s", r.code, r.err)
	}
	r = gyro(t, "a_full", "", "ssh-key", "get", fmt.Sprint(k.ID))
	if r.code != 4 {
		t.Errorf("deleted key still there: %d", r.code)
	}
}

// TestRealServer orders the smallest plan, waits, logs in over SSH, power cycles and deletes it.
func TestRealServer(t *testing.T) {
	if os.Getenv("GYRO_E2E_REAL") != "1" {
		t.Skip("set GYRO_E2E_REAL=1 to order a real server on dev")
	}
	dir := t.TempDir()
	kf := filepath.Join(dir, "k")
	if out, err := exec.Command("ssh-keygen", "-q", "-t", "ed25519", "-N", "", "-C", "gyro-cli-e2e-real", "-f", kf).CombinedOutput(); err != nil {
		t.Fatal(string(out))
	}
	host := fmt.Sprintf("gyro-cli-e2e-%d.gyro.example", time.Now().Unix())
	plan := os.Getenv("GYRO_E2E_PLAN")
	if plan == "" {
		plan = "standart-1-1"
	}
	start := time.Now()
	r := gyro(t, "a_full", "", "server", "create", host, "--plan", plan, "--location", "falkenstein", "--image", "debian-12",
		"--ssh-key", kf+".pub", "--wait", "--yes", "-o", "json")
	if r.code != 0 {
		t.Fatalf("create: %d %s", r.code, r.err)
	}
	var s struct {
		ID   int    `json:"id"`
		IPv4 string `json:"ipv4"`
		Live struct {
			PowerState string `json:"power_state"`
		} `json:"live"`
	}
	mustJSON(t, r.out, &s)
	sid := fmt.Sprint(s.ID)
	t.Logf("server %s created and running in %s", sid, time.Since(start).Round(time.Second))
	defer func() {
		// always clean up the server and the uploaded key
		r := gyro(t, "a_full", "", "server", "delete", sid, "--yes")
		t.Logf("cleanup delete -> %d %s", r.code, strings.TrimSpace(r.err))
		// The API only requests cancellation (grace period). On DEV, GYRO_E2E_PURGE destroys the machine now,
		// a command given in GYRO_E2E_PURGE (test installations only).
		if p := os.Getenv("GYRO_E2E_PURGE"); p != "" {
			f := strings.Fields(p)
			out, err := exec.Command(f[0], append(f[1:], sid)...).CombinedOutput()
			t.Logf("purge %s -> %v %s", sid, err, strings.TrimSpace(string(out)))
		}
		gyro(t, "a_full", "", "ssh-key", "delete", "gyro-cli-e2e-real", "--yes")
	}()
	if s.Live.PowerState != "running" || s.IPv4 == "" {
		t.Fatalf("not running: %s", r.out)
	}
	// SSH (the image may need a few seconds for sshd)
	var ok bool
	for i := 0; i < 12 && !ok; i++ {
		r = gyro(t, "a_full", "", "server", "ssh", host, "-i", kf, "--", "-o", "StrictHostKeyChecking=no", "-o", "UserKnownHostsFile=/dev/null",
			"-o", "ConnectTimeout=10", "-o", "BatchMode=yes", "hostname")
		ok = r.code == 0
		if !ok {
			time.Sleep(10 * time.Second)
		}
	}
	if !ok {
		t.Errorf("ssh login failed: %s", r.err)
	}
	if r := gyro(t, "a_full", "", "server", "ip", sid); strings.TrimSpace(r.out) != s.IPv4 {
		t.Errorf("server ip: %q", r.out)
	}
	if r := gyro(t, "a_full", "", "server", "power", "shutdown", sid, "--wait", "--wait-timeout", "5m"); r.code != 0 {
		t.Errorf("shutdown: %d %s", r.code, r.err)
	}
	if r := gyro(t, "a_full", "", "server", "get", sid, "-o", "text", "--query", "live.power_state"); strings.TrimSpace(r.out) != "stopped" {
		t.Errorf("after shutdown: %q", r.out)
	}
	if r := gyro(t, "a_full", "", "server", "power", "on", sid, "--wait"); r.code != 0 {
		t.Errorf("power on: %d %s", r.code, r.err)
	}
	if r := gyro(t, "a_full", "", "server", "rename", sid, "renamed-"+host); r.code != 0 {
		t.Errorf("rename: %d %s", r.code, r.err)
	}
	if r := gyro(t, "a_full", "", "server", "metrics", sid, "--period", "24h"); r.code != 0 {
		t.Errorf("metrics: %d %s", r.code, r.err)
	}
	if r := gyro(t, "a_full", "", "server", "rdns", "list", sid); r.code != 0 {
		t.Errorf("rdns: %d %s", r.code, r.err)
	}
	if r := gyro(t, "a_full", "", "server", "activity", "list", sid); r.code != 0 {
		t.Errorf("activity: %d %s", r.code, r.err)
	}
	if r := gyro(t, "a_full", "", "server", "delete", sid); r.code != 7 {
		t.Errorf("delete without --yes must abort: %d", r.code)
	}
	t.Logf("lifecycle done in %s", time.Since(start).Round(time.Second))
}

func TestV2WebhookAndKeys(t *testing.T) {
	r := gyro(t, "a_full", "", "webhook", "create", "--url", "https://example.com/gyro-cli-e2e", "--events", "api.test", "--description", "gyro-cli-e2e", "-o", "json")
	if r.code != 0 {
		t.Skipf("webhooks not available: %d %s", r.code, r.err)
	}
	var wh struct{ ID int }
	mustJSON(t, r.out, &wh)
	id := fmt.Sprint(wh.ID)
	defer gyro(t, "a_full", "", "webhook", "delete", id, "--yes")
	if !strings.Contains(r.err, "only once") {
		t.Errorf("secret warning missing: %s", r.err)
	}
	if r := gyro(t, "a_full", "", "webhook", "test", id); r.code != 0 {
		t.Errorf("webhook test: %d %s", r.code, r.err)
	}
	if r := gyro(t, "a_full", "", "webhook", "delivery", "list", id); r.code != 0 || !strings.Contains(r.out, "api.test") {
		t.Errorf("deliveries: %d %s", r.code, r.out)
	}
	if r := gyro(t, "a_full", "", "webhook", "update", id, "--description", "renamed"); r.code != 0 {
		t.Errorf("update: %d %s", r.code, r.err)
	}
	// sub-account key: the token appears once; our harness fails if any wck_ token is printed, so use --query
	r = gyro(t, "a_full", "", "key", "create", "--name", "gyro-cli-e2e", "--scopes", "cloud_read", "-o", "json", "--query", "id")
	if r.code != 0 {
		t.Fatalf("key create: %d %s", r.code, r.err)
	}
	kid := strings.TrimSpace(r.out)
	if r := gyro(t, "a_full", "", "key", "delete", kid, "--yes"); r.code != 0 {
		t.Errorf("key delete: %d %s", r.code, r.err)
	}
	for _, args := range [][]string{{"service", "list"}, {"order", "list"}, {"ping"}, {"mail", "list"}, {"api", "operations"}} {
		if r := gyro(t, "a_full", "", args...); r.code != 0 {
			t.Errorf("%v: %d %s", args, r.code, r.err)
		}
	}
	if r := gyro(t, "a_full", "", "reseller", "customer", "list"); r.code != 3 || !strings.Contains(r.err, "not_reseller") {
		t.Errorf("non-reseller: %d %s", r.code, r.err)
	}
}
