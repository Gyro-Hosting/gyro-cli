// Package output renders API data as a table (default), JSON or YAML, with optional JMESPath --query,
// --no-header and quiet (ids only) modes. Colours honour NO_COLOR and are off when stdout is not a terminal.
package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"unicode/utf8"

	"github.com/jmespath/go-jmespath"
	"golang.org/x/term"
	"gopkg.in/yaml.v3"
)

// Column is one table column: Header is shown, Path is a dot path into each item (a.b.c).
type Column struct {
	Header string
	Path   string
	Format func(v any) string // optional
}

// Printer holds the output options.
type Printer struct {
	W        io.Writer
	Format   string // table | json | yaml | text
	Query    string
	NoHeader bool
	Quiet    bool
	Color    bool
	Wide     bool
}

// ColorEnabled decides colouring for a writer.
func ColorEnabled(w io.Writer, flagNoColor bool) bool {
	if flagNoColor || os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb" {
		return false
	}
	if f, ok := w.(*os.File); ok {
		return term.IsTerminal(int(f.Fd()))
	}
	return false
}

// Decode turns raw JSON into generic Go values (float64 numbers, as JMESPath expects).
func Decode(raw []byte) (any, error) {
	if len(raw) == 0 {
		return nil, nil
	}
	var v any
	err := json.Unmarshal(raw, &v)
	return v, err
}

// Print renders data. cols may be nil (columns are then derived from the data). idPath is used by
// quiet mode ("id" by default). For lists wrapped in an object, pass the list itself.
func (p *Printer) Print(data any, cols []Column, idPath string) error {
	if p.Query != "" {
		q, err := jmespath.Search(p.Query, data)
		if err != nil {
			return fmt.Errorf("--query: %w", err)
		}
		data = q
		cols = nil // the shape changed; derive columns from the result
	}
	if p.Quiet {
		return p.quiet(data, idPath)
	}
	switch p.Format {
	case "json":
		enc := json.NewEncoder(p.W)
		enc.SetIndent("", "  ")
		enc.SetEscapeHTML(false)
		return enc.Encode(data)
	case "yaml":
		enc := yaml.NewEncoder(p.W)
		enc.SetIndent(2)
		if err := enc.Encode(data); err != nil {
			return err
		}
		return enc.Close()
	case "text":
		return p.text(data)
	default:
		return p.table(data, cols)
	}
}

func (p *Printer) quiet(data any, idPath string) error {
	if idPath == "" {
		idPath = "id"
	}
	emit := func(v any) {
		if m, ok := v.(map[string]any); ok {
			if x := Lookup(m, idPath); x != nil {
				fmt.Fprintln(p.W, Scalar(x))
				return
			}
			for _, k := range []string{"id", "slug", "code", "name"} {
				if x, ok := m[k]; ok {
					fmt.Fprintln(p.W, Scalar(x))
					return
				}
			}
			return
		}
		fmt.Fprintln(p.W, Scalar(v))
	}
	if arr, ok := data.([]any); ok {
		for _, v := range arr {
			emit(v)
		}
		return nil
	}
	if data != nil {
		emit(data)
	}
	return nil
}

// text: tab separated values, no alignment (like aws --output text).
func (p *Printer) text(data any) error {
	switch t := data.(type) {
	case []any:
		cols := Derive(t, true)
		if len(cols) == 0 {
			for _, v := range t {
				fmt.Fprintln(p.W, Scalar(v))
			}
			return nil
		}
		for _, row := range t {
			var cells []string
			for _, c := range cols {
				cells = append(cells, Cell(row, c, false))
			}
			fmt.Fprintln(p.W, strings.Join(cells, "\t"))
		}
	case map[string]any:
		for _, k := range sortedKeys(t) {
			fmt.Fprintf(p.W, "%s\t%s\n", k, Scalar(t[k]))
		}
	default:
		fmt.Fprintln(p.W, Scalar(t))
	}
	return nil
}

func (p *Printer) table(data any, cols []Column) error {
	tw := tabwriter.NewWriter(p.W, 0, 0, 3, ' ', 0)
	defer tw.Flush()
	switch t := data.(type) {
	case []any:
		if len(cols) == 0 {
			cols = Derive(t, p.Wide)
		}
		if len(t) == 0 {
			return nil
		}
		if len(cols) == 0 { // list of scalars
			for _, v := range t {
				fmt.Fprintln(tw, Scalar(v))
			}
			return nil
		}
		if !p.NoHeader {
			var hs []string
			for _, c := range cols {
				hs = append(hs, strings.ToUpper(c.Header))
			}
			fmt.Fprintln(tw, strings.Join(hs, "\t"))
		}
		for _, row := range t {
			var cells []string
			for _, c := range cols {
				cells = append(cells, Cell(row, c, p.Color))
			}
			fmt.Fprintln(tw, strings.Join(cells, "\t"))
		}
	case map[string]any:
		if len(cols) > 0 {
			for _, c := range cols {
				fmt.Fprintf(tw, "%s:\t%s\n", c.Header, Cell(t, c, p.Color))
			}
			return nil
		}
		p.kv(tw, t, "")
	case nil:
	default:
		fmt.Fprintln(tw, Scalar(t))
	}
	return nil
}

// kv prints an object as "key: value" lines, flattening nested objects with dotted keys.
func (p *Printer) kv(w io.Writer, m map[string]any, prefix string) {
	for _, k := range sortedKeys(m) {
		v := m[k]
		key := prefix + k
		if sub, ok := v.(map[string]any); ok && !IsMoney(sub) && len(sub) > 0 {
			p.kv(w, sub, key+".")
			continue
		}
		val := Scalar(v)
		if arr, ok := v.([]any); ok && len(arr) > 0 {
			if _, isObj := arr[0].(map[string]any); isObj {
				val = fmt.Sprintf("[%d]", len(arr))
			} else {
				var parts []string
				for _, x := range arr {
					parts = append(parts, Scalar(x))
				}
				val = strings.Join(parts, ", ")
			}
		}
		if p.Color && (k == "status" || k == "power_state" || k == "state") {
			val = Colorize(val)
		}
		fmt.Fprintf(w, "%s:\t%s\n", key, val)
	}
}

var preferred = []string{"id", "name", "hostname", "slug", "code", "domain", "email", "number", "title", "subject",
	"status", "state", "power_state", "server_type", "type", "location", "ipv4", "city", "country", "price", "total",
	"amount", "balance", "due_date", "created_at"}

// Derive picks columns from list data: preferred keys first, then the rest alphabetically; nested
// objects other than money and arrays of objects are skipped. At most 8 columns unless wide.
func Derive(rows []any, wide bool) []Column {
	seen := map[string]bool{}
	var keys []string
	for _, r := range rows {
		m, ok := r.(map[string]any)
		if !ok {
			return nil
		}
		for k, v := range m {
			if seen[k] {
				continue
			}
			if sub, ok := v.(map[string]any); ok && !IsMoney(sub) {
				continue
			}
			if arr, ok := v.([]any); ok && len(arr) > 0 {
				if _, isObj := arr[0].(map[string]any); isObj {
					continue
				}
			}
			seen[k] = true
			keys = append(keys, k)
		}
	}
	rank := func(k string) int {
		for i, p := range preferred {
			if p == k {
				return i
			}
		}
		return len(preferred)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		ri, rj := rank(keys[i]), rank(keys[j])
		if ri != rj {
			return ri < rj
		}
		return keys[i] < keys[j]
	})
	if !wide && len(keys) > 8 {
		keys = keys[:8]
	}
	var cols []Column
	for _, k := range keys {
		cols = append(cols, Column{Header: k, Path: k})
	}
	return cols
}

// FromNames builds columns from spec property names.
func FromNames(names []string, max int) []Column {
	var cols []Column
	for _, n := range names {
		if max > 0 && len(cols) >= max {
			break
		}
		cols = append(cols, Column{Header: n, Path: n})
	}
	return cols
}

// Cell formats one cell.
func Cell(row any, c Column, color bool) string {
	m, _ := row.(map[string]any)
	v := Lookup(m, c.Path)
	var s string
	if c.Format != nil {
		s = c.Format(v)
	} else {
		s = Scalar(v)
	}
	if color && (c.Path == "status" || c.Path == "state" || c.Path == "power_state" || strings.HasSuffix(c.Path, ".power_state")) {
		s = Colorize(s)
	}
	return s
}

// Lookup follows a dot path.
func Lookup(m map[string]any, path string) any {
	var cur any = m
	for _, part := range strings.Split(path, ".") {
		mm, ok := cur.(map[string]any)
		if !ok {
			return nil
		}
		cur = mm[part]
	}
	return cur
}

// IsMoney reports {"amount": n, "currency": c}.
func IsMoney(m map[string]any) bool {
	_, a := m["amount"]
	_, c := m["currency"]
	return a && c && len(m) <= 3
}

// Scalar renders a value for a table cell.
func Scalar(v any) string {
	switch t := v.(type) {
	case nil:
		return "-"
	case string:
		if t == "" {
			return "-"
		}
		return t
	case bool:
		if t {
			return "yes"
		}
		return "no"
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	case json.Number:
		return t.String()
	case map[string]any:
		if IsMoney(t) {
			cur := ""
			if c, ok := t["currency"].(string); ok {
				cur = " " + c
			}
			if a, ok := t["amount"].(float64); ok {
				return strconv.FormatFloat(a, 'f', 2, 64) + cur
			}
			return Scalar(t["amount"]) + cur
		}
		b, _ := json.Marshal(t)
		return truncate(string(b), 60)
	case []any:
		var parts []string
		for _, x := range t {
			parts = append(parts, Scalar(x))
		}
		return truncate(strings.Join(parts, ","), 60)
	default:
		return fmt.Sprint(t)
	}
}

func truncate(s string, n int) string {
	if utf8.RuneCountInString(s) <= n {
		return s
	}
	r := []rune(s)
	return string(r[:n-1]) + "…"
}

const (
	green  = "\x1b[32m"
	yellow = "\x1b[33m"
	red    = "\x1b[31m"
	reset  = "\x1b[0m"
)

// Colorize paints well-known status words.
func Colorize(s string) string {
	switch strings.ToLower(s) {
	case "active", "running", "paid", "available", "ok", "up", "open", "yes", "enabled", "accepted":
		return green + s + reset
	case "waiting", "inprocess", "starting", "stopping", "creating", "pending", "unpaid", "suspended", "stopped", "replied":
		return yellow + s + reset
	case "cancelled", "overdue", "failed", "deleting", "down", "error", "closed":
		return red + s + reset
	}
	return s
}

func sortedKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		ri, rj := len(preferred), len(preferred)
		for n, p := range preferred {
			if p == keys[i] {
				ri = n
			}
			if p == keys[j] {
				rj = n
			}
		}
		if ri != rj {
			return ri < rj
		}
		return keys[i] < keys[j]
	})
	return keys
}
