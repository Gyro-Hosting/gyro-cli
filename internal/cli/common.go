package cli

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/config"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
	"github.com/Gyro-Hosting/gyro-cli/internal/output"
)

// preScan sets the message language before the command tree (and its help texts) is built.
func preScan(a *App, args []string) {
	lang, profile := "", ""
	for i := 0; i < len(args); i++ {
		s := args[i]
		switch {
		case s == "--":
			i = len(args)
		case strings.HasPrefix(s, "--lang="):
			lang = strings.TrimPrefix(s, "--lang=")
		case s == "--lang" && i+1 < len(args):
			lang = args[i+1]
		case strings.HasPrefix(s, "--profile="):
			profile = strings.TrimPrefix(s, "--profile=")
		case (s == "--profile" || s == "-p") && i+1 < len(args):
			profile = args[i+1]
		}
	}
	if lang != "" {
		i18n.Set(lang)
		return
	}
	if a.Env("GYRO_LANG") != "" {
		i18n.Set(a.Env("GYRO_LANG"))
		return
	}
	if c, err := config.Load(); err == nil {
		if p, ok := c.Profiles[c.ActiveName(profile)]; ok && p.Lang != "" {
			i18n.Set(p.Lang)
			return
		}
	}
	i18n.Set(i18n.Detect())
}

// Confirm asks a yes/no question. --yes skips it; without a terminal it fails (exit 7) so scripts
// never spend money silently.
func (a *App) Confirm(question string) error {
	if a.Opt.Yes {
		return nil
	}
	if !a.IsTTY() {
		return &ExitErr{ExitAborted, errors.New(i18n.S(
			"confirmation required: re-run with --yes (no terminal to ask)",
			"onay gerekli: --yes ile tekrar çalıştırın (soracak terminal yok)"))}
	}
	fmt.Fprintf(a.Err, "%s [%s] ", question, i18n.S("y/N", "e/H"))
	line, _ := bufio.NewReader(a.In).ReadString('\n')
	ans := strings.ToLower(strings.TrimSpace(line))
	if ans == "y" || ans == "yes" || ans == "e" || ans == "evet" {
		return nil
	}
	return &ExitErr{ExitAborted, errors.New(i18n.S("aborted", "iptal edildi"))}
}

// ConfirmTyped makes the user type a word (e.g. the server name) before deleting.
func (a *App) ConfirmTyped(question, word string) error {
	if a.Opt.Yes {
		return nil
	}
	if !a.IsTTY() {
		return a.Confirm(question)
	}
	fmt.Fprintf(a.Err, "%s\n%s ", question, i18n.F("Type %q to continue:", "Devam etmek için %q yazın:", word))
	line, _ := bufio.NewReader(a.In).ReadString('\n')
	if strings.TrimSpace(line) == word {
		return nil
	}
	return &ExitErr{ExitAborted, errors.New(i18n.S("aborted", "iptal edildi"))}
}

// Prompt reads one line with a default.
func (a *App) Prompt(r *bufio.Reader, label, def string) string {
	if def != "" {
		fmt.Fprintf(a.Err, "%s [%s]: ", label, def)
	} else {
		fmt.Fprintf(a.Err, "%s: ", label)
	}
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	if line == "" {
		return def
	}
	return line
}

// PrintRaw decodes and prints a data block.
func (a *App) PrintRaw(raw json.RawMessage, cols []output.Column, idPath string) error {
	v, err := output.Decode(raw)
	if err != nil {
		return err
	}
	if err := a.Printer().Print(v, cols, idPath); err != nil {
		return err
	}
	a.noneHint(v)
	return nil
}

// noneHint tells table users that an empty list is empty (stderr, so pipes stay clean).
func (a *App) noneHint(v any) {
	if arr, ok := v.([]any); ok && len(arr) == 0 && !a.Machine() && a.Opt.Query == "" && a.Printer().Format == "table" {
		a.Info("%s", i18n.S("No items found.", "Kayıt bulunamadı."))
	}
}

// PrintList prints raw list items.
func (a *App) PrintList(items []json.RawMessage, cols []output.Column, idPath string) error {
	arr := make([]any, 0, len(items))
	for _, it := range items {
		v, err := output.Decode(it)
		if err != nil {
			return err
		}
		arr = append(arr, v)
	}
	if err := a.Printer().Print(arr, cols, idPath); err != nil {
		return err
	}
	a.noneHint(arr)
	return nil
}

// Spinner shows progress on stderr when it is a terminal.
type Spinner struct {
	a      *App
	label  string
	start  time.Time
	frame  int
	active bool
}

func (a *App) NewSpinner(label string) *Spinner {
	s := &Spinner{a: a, label: label, start: time.Now()}
	if f, ok := a.Err.(*os.File); ok && !a.Opt.Quiet && output.ColorEnabled(f, false) {
		s.active = true
	}
	return s
}

func (s *Spinner) Tick(status string) {
	if !s.active {
		return
	}
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	s.frame++
	fmt.Fprintf(s.a.Err, "\r\x1b[K%s %s %s (%ds)", frames[s.frame%len(frames)], s.label, status, int(time.Since(s.start).Seconds()))
}

func (s *Spinner) Done(msg string) {
	if s.active {
		fmt.Fprint(s.a.Err, "\r\x1b[K")
	}
	if msg != "" {
		s.a.Info("%s", msg)
	}
}

// WaitFor polls check() until it reports done, with a timeout. check returns (done, status text, error).
func (a *App) WaitFor(ctx context.Context, label string, timeout time.Duration, check func() (bool, string, error)) error {
	if timeout <= 0 {
		timeout = 15 * time.Minute
	}
	deadline := time.Now().Add(timeout)
	sp := a.NewSpinner(label)
	interval := 3 * time.Second
	for {
		done, status, err := check()
		if err != nil {
			sp.Done("")
			return err
		}
		if done {
			sp.Done("")
			return nil
		}
		sp.Tick(status)
		if time.Now().After(deadline) {
			sp.Done("")
			return &ExitErr{ExitTimeout, fmt.Errorf("%s: %s (%s)", label, i18n.S("timed out", "zaman aşımı"), status)}
		}
		if ctx.Err() != nil {
			sp.Done("")
			return ctx.Err()
		}
		// poll every 3 s, polling slower after a minute
		if time.Since(deadline.Add(-timeout)) > time.Minute {
			interval = 5 * time.Second
		}
		for i := 0; i < int(interval/(250*time.Millisecond)); i++ {
			a.Sleep(250 * time.Millisecond)
			sp.Tick(status)
			if ctx.Err() != nil {
				break
			}
		}
	}
}

// ParseKV parses k=v pairs (values that look like JSON are decoded: numbers, true/false, [..], {..}).
func ParseKV(pairs []string) (map[string]any, error) {
	out := map[string]any{}
	for _, p := range pairs {
		k, v, ok := strings.Cut(p, "=")
		if !ok || k == "" {
			return nil, fmt.Errorf("--set %q: %s", p, i18n.S("expected key=value", "anahtar=değer bekleniyor"))
		}
		out[k] = guessJSON(v)
	}
	return out, nil
}

func guessJSON(v string) any {
	if v == "true" || v == "false" || v == "null" || strings.HasPrefix(v, "[") || strings.HasPrefix(v, "{") {
		var x any
		if json.Unmarshal([]byte(v), &x) == nil {
			return x
		}
	}
	if n, err := strconv.ParseInt(v, 10, 64); err == nil && (v == "0" || !strings.HasPrefix(strings.TrimPrefix(v, "-"), "0")) {
		return n
	}
	return v
}

// ReadData loads --data: inline JSON, @file or @- (stdin).
func (a *App) ReadData(s string) (map[string]any, error) {
	if s == "" {
		return nil, nil
	}
	var raw []byte
	var err error
	switch {
	case s == "@-":
		raw, err = readAll(a.In)
	case strings.HasPrefix(s, "@"):
		raw, err = os.ReadFile(s[1:])
	default:
		raw = []byte(s)
	}
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, fmt.Errorf("--data: %s: %w", i18n.S("expected a JSON object", "JSON nesnesi bekleniyor"), err)
	}
	return m, nil
}

// ReadValue loads @file / @- values for string flags (user data, public keys).
func (a *App) ReadValue(s string) (string, error) {
	switch {
	case s == "@-":
		b, err := readAll(a.In)
		return string(b), err
	case strings.HasPrefix(s, "@"):
		b, err := os.ReadFile(expandHome(s[1:]))
		return string(b), err
	}
	return s, nil
}

func expandHome(p string) string {
	if strings.HasPrefix(p, "~/") {
		if h, err := os.UserHomeDir(); err == nil {
			return h + p[1:]
		}
	}
	return p
}

func readAll(r interface{ Read([]byte) (int, error) }) ([]byte, error) {
	var out []byte
	buf := make([]byte, 32*1024)
	for {
		n, err := r.Read(buf)
		out = append(out, buf[:n]...)
		if err != nil {
			if err.Error() == "EOF" {
				return out, nil
			}
			return out, err
		}
	}
}

func execProgram(name string, args []string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	err := cmd.Run()
	var ee *exec.ExitError
	if errors.As(err, &ee) {
		return &ExitErr{ee.ExitCode(), fmt.Errorf("%s exited with %d", name, ee.ExitCode())}
	}
	return err
}

// resolveFrom finds an object by numeric id or by name-like fields in a list endpoint.
func (a *App) resolveFrom(ctx context.Context, listPath, itemsKey, arg string, fields ...string) (string, map[string]any, error) {
	c, err := a.Client()
	if err != nil {
		return "", nil, err
	}
	items, err := c.All(ctx, listPath, url.Values{}, itemsKey, 20)
	if err != nil {
		return "", nil, err
	}
	var matches []map[string]any
	for _, it := range items {
		v, _ := output.Decode(it)
		m, ok := v.(map[string]any)
		if !ok {
			continue
		}
		if output.Scalar(m["id"]) == arg {
			return arg, m, nil
		}
		for _, f := range fields {
			if s, ok := m[f].(string); ok && s != "" && (strings.EqualFold(s, arg) || strings.EqualFold(strings.SplitN(s, ".", 2)[0], arg)) {
				matches = append(matches, m)
				break
			}
		}
	}
	if len(matches) > 1 {
		// prefer live objects over cancelled/deleted history with the same name
		var live []map[string]any
		for _, m := range matches {
			switch output.Scalar(m["status"]) {
			case "cancelled", "terminated", "deleted":
			default:
				live = append(live, m)
			}
		}
		if len(live) > 0 {
			matches = live
		}
	}
	switch len(matches) {
	case 0:
		return "", nil, &ExitErr{ExitNotFound, fmt.Errorf("%s: %q", i18n.S("not found", "bulunamadı"), arg)}
	case 1:
		return output.Scalar(matches[0]["id"]), matches[0], nil
	}
	var ids []string
	for _, m := range matches {
		ids = append(ids, output.Scalar(m["id"]))
	}
	return "", nil, &ExitErr{ExitUsage, fmt.Errorf("%q %s: %s", arg, i18n.S("is ambiguous, use an id", "birden çok kayda uyuyor, kimlik kullanın"), strings.Join(ids, ", "))}
}

func isNumeric(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func setFlagIfAny(cmd *cobra.Command, name string) bool {
	f := cmd.Flags().Lookup(name)
	return f != nil && f.Changed
}

// unwrapAPI returns the API error inside err if any.
func unwrapAPI(err error) (*api.Error, bool) { return api.AsError(err) }
