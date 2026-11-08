// gyro-gen reads the Gyro Public API OpenAPI document and writes internal/spec/zz_generated_ops.go.
//
//	go run ./cmd/gyro-gen -spec internal/spec/openapi.yaml -out internal/spec/zz_generated_ops.go
//
// It is wired to `go generate ./...` (see internal/spec/gen.go). The CLI turns every operation into a
// command; adding an endpoint to the spec and re-running the generator is all that is needed.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/Gyro-Hosting/gyro-cli/internal/spec"
)

type M = map[string]any

var doc M

func main() {
	in := flag.String("spec", "internal/spec/openapi.yaml", "OpenAPI document (yaml or json)")
	out := flag.String("out", "internal/spec/zz_generated_ops.go", "output Go file")
	flag.Parse()
	raw, err := os.ReadFile(*in)
	check(err)
	var root yaml.Node
	check(yaml.Unmarshal(raw, &root))
	d, _ := convert(&root).(M)
	if d == nil {
		check(fmt.Errorf("%s: not an OpenAPI document", *in))
	}
	doc = d
	ops := build()
	src := render(ops)
	formatted, err := format.Source(src)
	if err != nil {
		os.WriteFile(*out+".broken", src, 0o644)
		check(err)
	}
	check(os.WriteFile(*out, formatted, 0o644))
	fmt.Fprintf(os.Stderr, "gyro-gen: %d operations, spec %v -> %s\n", len(ops), str(get(doc, "info"), "version"), *out)
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "gyro-gen:", err)
		os.Exit(1)
	}
}

func get(m any, k string) any {
	if mm, ok := m.(M); ok {
		return mm[k]
	}
	return nil
}

func str(m any, k string) string {
	if v, ok := get(m, k).(string); ok {
		return v
	}
	return ""
}

// deref follows local $ref pointers (#/components/...).
func deref(v any) any {
	for i := 0; i < 20; i++ {
		ref := str(v, "$ref")
		if ref == "" {
			return v
		}
		cur := any(doc)
		for _, p := range strings.Split(strings.TrimPrefix(ref, "#/"), "/") {
			p = strings.ReplaceAll(strings.ReplaceAll(p, "~1", "/"), "~0", "~")
			cur = get(cur, p)
		}
		v = cur
	}
	return v
}

// props flattens object/allOf/oneOf/anyOf schemas into an ordered property list.
type prop struct {
	name   string
	schema any
}

func props(s any) (out []prop, required map[string]bool) {
	required = map[string]bool{}
	var walk func(s any)
	seen := map[string]bool{}
	walk = func(s any) {
		s = deref(s)
		if s == nil {
			return
		}
		for _, k := range []string{"allOf", "oneOf", "anyOf"} {
			if arr, ok := get(s, k).([]any); ok {
				for _, x := range arr {
					walk(x)
				}
			}
		}
		if req, ok := get(s, "required").([]any); ok {
			for _, r := range req {
				required[fmt.Sprint(r)] = true
			}
		}
		if node, ok := get(s, "properties").(M); ok {
			for _, k := range orderedKeys(s, "properties") {
				if !seen[k] {
					seen[k] = true
					out = append(out, prop{k, node[k]})
				}
			}
		}
	}
	walk(s)
	return
}

// orderedKeys returns map keys in YAML document order.

func orderedKeys(parent any, field string) []string {
	m, _ := get(parent, field).(M)
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	pos := map[string]int{}
	if raw, ok := rawOrder[fmt.Sprintf("%p", m)]; ok {
		for i, k := range raw {
			pos[k] = i
		}
	}
	sort.SliceStable(keys, func(i, j int) bool {
		pi, iok := pos[keys[i]]
		pj, jok := pos[keys[j]]
		if iok && jok {
			return pi < pj
		}
		return keys[i] < keys[j]
	})
	return keys
}

var rawOrder = map[string][]string{}

// convert turns a yaml.Node tree into plain Go values, remembering mapping key order in rawOrder.
func convert(n *yaml.Node) any {
	switch n.Kind {
	case yaml.DocumentNode:
		if len(n.Content) > 0 {
			return convert(n.Content[0])
		}
		return nil
	case yaml.AliasNode:
		return convert(n.Alias)
	case yaml.MappingNode:
		m := M{}
		var keys []string
		for i := 0; i+1 < len(n.Content); i += 2 {
			k := n.Content[i].Value
			keys = append(keys, k)
			m[k] = convert(n.Content[i+1])
		}
		rawOrder[fmt.Sprintf("%p", m)] = keys
		return m
	case yaml.SequenceNode:
		out := make([]any, 0, len(n.Content))
		for _, c := range n.Content {
			out = append(out, convert(c))
		}
		return out
	default:
		var v any
		_ = n.Decode(&v)
		return v
	}
}

func typeOf(s any) (typ, items string, enum []string, def string) {
	s = deref(s)
	switch t := get(s, "type").(type) {
	case string:
		typ = t
	case []any:
		for _, x := range t {
			if fmt.Sprint(x) != "null" {
				typ = fmt.Sprint(x)
				break
			}
		}
	}
	if typ == "" {
		if get(s, "properties") != nil || get(s, "allOf") != nil {
			typ = "object"
		} else {
			typ = "string"
		}
	}
	if typ == "array" {
		it, _, _, _ := typeOf(get(s, "items"))
		items = it
	}
	if e, ok := get(s, "enum").([]any); ok {
		for _, x := range e {
			if x != nil {
				enum = append(enum, fmt.Sprint(x))
			}
		}
	}
	if c := get(s, "const"); c != nil {
		enum = []string{fmt.Sprint(c)}
	}
	if d := get(s, "default"); d != nil {
		def = fmt.Sprint(d)
	}
	return
}

func dataSchema(op any) any {
	resps, _ := get(op, "responses").(M)
	for _, code := range []string{"200", "201", "202"} {
		r := deref(resps[code])
		if r == nil {
			continue
		}
		sch := deref(get(get(get(r, "content"), "application/json"), "schema"))
		for _, p := range func() []prop { ps, _ := props(sch); return ps }() {
			if p.name == "data" {
				return deref(p.schema)
			}
		}
	}
	return nil
}

func build() []spec.Op {
	paths, _ := doc["paths"].(M)
	pathKeys := make([]string, 0, len(paths))
	for k := range paths {
		pathKeys = append(pathKeys, k)
	}
	// GyroAPI (/cloud/...) operations first: they own the short names; core duplicates go to "gyro core ..."
	sort.Slice(pathKeys, func(i, j int) bool {
		ci, cj := strings.HasPrefix(pathKeys[i], "/cloud/"), strings.HasPrefix(pathKeys[j], "/cloud/")
		if ci != cj {
			return ci
		}
		return pathKeys[i] < pathKeys[j]
	})
	hasChildParam := func(p string) bool {
		for _, k := range pathKeys {
			if strings.HasPrefix(k, p+"/{") && !strings.Contains(k[len(p)+2:], "/") {
				return true
			}
		}
		return false
	}
	var ops []spec.Op
	used := map[string]string{}
	for _, path := range pathKeys {
		item := paths[path]
		hasGet := get(item, "get") != nil
		others := map[string]bool{}
		for _, m := range []string{"put", "patch", "delete"} {
			if get(item, m) != nil {
				others[m] = true
			}
		}
		common, _ := get(item, "parameters").([]any)
		for _, method := range []string{"get", "post", "put", "patch", "delete"} {
			o := get(item, method)
			if o == nil {
				continue
			}
			op := spec.Op{
				ID: str(o, "operationId"), Method: strings.ToUpper(method), Path: path,
				Summary: str(o, "summary"), SummaryTR: str(o, "x-summary-tr"),
				Desc: str(o, "description"), DescTR: str(o, "x-description-tr"), Scope: str(o, "x-scope"),
			}
			if op.ID == "" {
				op.ID = spec.Kebab(method + "-" + path)
			}
			if tags, ok := get(o, "tags").([]any); ok && len(tags) > 0 {
				op.Tag = fmt.Sprint(tags[0])
			}
			params, _ := get(o, "parameters").([]any)
			seenP := map[string]bool{}
			for _, p := range append(append([]any{}, params...), common...) {
				p = deref(p)
				key := str(p, "in") + ":" + str(p, "name")
				if seenP[key] {
					continue
				}
				seenP[key] = true
				t, it, en, def := typeOf(get(p, "schema"))
				req, _ := get(p, "required").(bool)
				op.Params = append(op.Params, spec.Param{
					Name: str(p, "name"), In: str(p, "in"), Type: t, Items: it, Enum: en, Default: def, Required: req,
					Desc: oneLine(str(p, "description")), DescTR: oneLine(str(p, "x-description-tr")),
				})
				if str(p, "name") == "page" && str(p, "in") == "query" {
					op.Paginated = true
				}
			}
			// order path params by their position in the URL
			sort.SliceStable(op.Params, func(i, j int) bool {
				a, b := op.Params[i], op.Params[j]
				if a.In == "path" && b.In == "path" {
					return strings.Index(path, "{"+a.Name+"}") < strings.Index(path, "{"+b.Name+"}")
				}
				return a.In == "path" && b.In != "path"
			})
			if rb := deref(get(o, "requestBody")); rb != nil {
				op.HasBody = true
				op.BodyRequired, _ = get(rb, "required").(bool)
				sch := get(get(get(rb, "content"), "application/json"), "schema")
				ps, req := props(sch)
				for _, pr := range ps {
					t, it, en, def := typeOf(pr.schema)
					d := deref(pr.schema)
					if ro, _ := get(d, "readOnly").(bool); ro {
						continue
					}
					op.Body = append(op.Body, spec.Field{
						Name: pr.name, Type: t, Items: it, Enum: en, Default: def, Required: req[pr.name],
						Desc: oneLine(str(d, "description")), DescTR: oneLine(str(d, "x-description-tr")),
					})
				}
			}
			ds := dataSchema(o)
			if t, _, _, _ := typeOf(ds); t == "array" {
				op.ListResponse = true
				ds = deref(get(ds, "items"))
			}
			if ds != nil && !op.ListResponse {
				// {"data": {"items": [...], ...}} is a list too (snapshots, backups)
				ps, _ := props(ds)
				for _, pr := range ps {
					if t, _, _, _ := typeOf(pr.schema); t == "array" && (pr.name == "items" || pr.name == "data") {
						op.ListResponse = true
						op.ItemsKey = pr.name
						ds = deref(get(deref(pr.schema), "items"))
						break
					}
				}
			}
			if ds != nil {
				ps, _ := props(ds)
				for _, pr := range ps {
					t, _, _, _ := typeOf(pr.schema)
					d := deref(pr.schema)
					money := t == "object" && get(get(d, "properties"), "amount") != nil
					if t == "array" || (t == "object" && !money) {
						continue
					}
					op.Columns = append(op.Columns, pr.name)
				}
			}
			_, hasConfirmBody := op.Field("confirm")
			_, hasConfirmQuery := op.Param("confirm")
			_, hasDry := op.Field("dry_run")
			op.Money = op.Scope == "cloud_order" || op.Scope == "billing_pay" || (hasDry && op.Method != "GET")
			x, _ := get(o, "x-cli").(M)
			if b, ok := x["money"].(bool); ok {
				op.Money = b
			}
			if b, ok := x["hidden"].(bool); ok {
				op.Hidden = b
			}

			// trailing enum path parameter → one command per value
			pp := op.PathParams()
			trailingEnum := len(pp) > 0 && strings.HasSuffix(path, "{"+pp[len(pp)-1].Name+"}") && len(pp[len(pp)-1].Enum) > 0
			otherMethods := false
			for m := range others {
				if m != method {
					otherMethods = true
				}
			}
			res, action := spec.Name(spec.NameInput{Method: op.Method, Path: path, OpID: op.ID, ListResp: op.ListResponse,
				HasGet: hasGet, HasOtherMethods: otherMethods, HasChildParam: hasChildParam(path), TrailingEnum: trailingEnum})
			variants := []spec.Op{op}
			if trailingEnum {
				variants = nil
				last := pp[len(pp)-1]
				for _, v := range last.Enum {
					c := op
					c.Params = nil
					for _, p := range op.Params {
						if p.Name == last.Name && p.In == "path" {
							p.Default = v
							p.Enum = []string{v}
						}
						c.Params = append(c.Params, p)
					}
					c.Command = append(append([]string{}, res...), spec.Kebab(v))
					variants = append(variants, c)
				}
			} else {
				op.Command = append(append([]string{}, res...), action)
				if action == "" {
					op.Command = res
				}
				variants[0] = op
			}
			for _, v := range variants {
				if cmd, ok := x["command"].(string); ok && cmd != "" && !trailingEnum {
					v.Command = strings.Fields(cmd)
				}
				key := strings.Join(v.Command, " ")
				if prev, dup := used[key]; dup && !strings.HasPrefix(v.Path, "/cloud/") {
					// a core endpoint shadowed by a GyroAPI one: keep it reachable as "gyro core ..."
					v.Command = append([]string{"core"}, v.Command...)
					key = strings.Join(v.Command, " ")
					_ = prev
				}
				if prev, dup := used[key]; dup {
					// collision: keep the first, name this one after its operationId
					v.Command = append(append([]string{}, v.Command[:len(v.Command)-1]...), spec.Kebab(v.ID))
					fmt.Fprintf(os.Stderr, "gyro-gen: %q used by %s, %s becomes %q\n", key, prev, v.ID, strings.Join(v.Command, " "))
					key = strings.Join(v.Command, " ")
				}
				used[key] = v.ID
				act := v.Command[len(v.Command)-1]
				v.Confirm = spec.Dangerous(v.Method, v.ID, act, v.Scope, hasConfirmBody || hasConfirmQuery)
				if b, ok := x["confirm"].(bool); ok {
					v.Confirm = b
				}
				ops = append(ops, v)
			}
		}
	}
	return ops
}

func oneLine(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func render(ops []spec.Op) []byte {
	var b bytes.Buffer
	b.WriteString("// Code generated by cmd/gyro-gen from openapi.yaml. DO NOT EDIT.\n\npackage spec\n\n")
	fmt.Fprintf(&b, "// APIVersion is info.version of the spec this table was generated from.\nconst APIVersion = %q\n\n", str(get(doc, "info"), "version"))
	b.WriteString("// Ops lists every API operation.\nvar Ops = []Op{\n")
	for _, o := range ops {
		fmt.Fprintf(&b, "\t{ID: %q, Method: %q, Path: %q, Tag: %q,\n", o.ID, o.Method, o.Path, o.Tag)
		fmt.Fprintf(&b, "\t\tSummary: %q, SummaryTR: %q,\n", o.Summary, o.SummaryTR)
		if o.Desc != "" || o.DescTR != "" {
			fmt.Fprintf(&b, "\t\tDesc: %q, DescTR: %q,\n", oneLine(o.Desc), oneLine(o.DescTR))
		}
		fmt.Fprintf(&b, "\t\tScope: %q, Command: %#v,\n", o.Scope, o.Command)
		fmt.Fprintf(&b, "\t\tHasBody: %v, BodyRequired: %v, ListResponse: %v, Paginated: %v, Confirm: %v, Money: %v, Hidden: %v,\n",
			o.HasBody, o.BodyRequired, o.ListResponse, o.Paginated, o.Confirm, o.Money, o.Hidden)
		if o.ItemsKey != "" {
			fmt.Fprintf(&b, "\t\tItemsKey: %q,\n", o.ItemsKey)
		}
		if len(o.Columns) > 0 {
			fmt.Fprintf(&b, "\t\tColumns: %#v,\n", o.Columns)
		}
		if len(o.Params) > 0 {
			b.WriteString("\t\tParams: []Param{\n")
			for _, p := range o.Params {
				fmt.Fprintf(&b, "\t\t\t{Name: %q, In: %q, Type: %q, Items: %q, Enum: %#v, Default: %q, Required: %v, Desc: %q, DescTR: %q},\n",
					p.Name, p.In, p.Type, p.Items, p.Enum, p.Default, p.Required, p.Desc, p.DescTR)
			}
			b.WriteString("\t\t},\n")
		}
		if len(o.Body) > 0 {
			b.WriteString("\t\tBody: []Field{\n")
			for _, f := range o.Body {
				fmt.Fprintf(&b, "\t\t\t{Name: %q, Type: %q, Items: %q, Enum: %#v, Default: %q, Required: %v, Desc: %q, DescTR: %q},\n",
					f.Name, f.Type, f.Items, f.Enum, f.Default, f.Required, f.Desc, f.DescTR)
			}
			b.WriteString("\t\t},\n")
		}
		b.WriteString("\t},\n")
	}
	b.WriteString("}\n")
	return bytes.ReplaceAll(b.Bytes(), []byte("[]string(nil)"), []byte("nil"))
}
