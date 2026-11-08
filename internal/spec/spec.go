// Package spec holds the operation table generated from the Gyro Public API OpenAPI document
// (internal/spec/openapi.yaml, synced from moduller/gyro-api/openapi.yaml by scripts/sync-spec.sh).
//
// The table (zz_generated_ops.go) is produced by cmd/gyro-gen. Every operation becomes a CLI command
// automatically (see internal/cli/generated.go); hand-written commands only add nicer UX on top.
package spec

// Param is a path, query or header parameter.
type Param struct {
	Name     string
	In       string // path | query | header
	Type     string // string | integer | number | boolean | array | object
	Items    string // element type for arrays
	Enum     []string
	Default  string
	Required bool
	Desc     string
	DescTR   string
}

// Field is one top-level property of a JSON request body.
type Field struct {
	Name     string
	Type     string
	Items    string
	Enum     []string
	Default  string
	Required bool
	Desc     string
	DescTR   string
}

// Op is one API operation plus the CLI command derived from it.
type Op struct {
	ID        string
	Method    string
	Path      string
	Tag       string
	Summary   string
	SummaryTR string
	Desc      string
	DescTR    string
	Scope     string // x-scope (cloud_read, cloud_order, ...)

	Params       []Param
	Body         []Field
	HasBody      bool
	BodyRequired bool

	ListResponse bool     // data is an array (or data.<ItemsKey> is)
	ItemsKey     string   // "items" when the list is wrapped: {"data": {"items": [...], "quota_gb": 5}}
	Paginated    bool     // has a page query parameter
	Columns      []string // default table columns (response item properties, spec order)

	Command []string // e.g. ["server", "snapshot", "list"]; derived by Name() unless x-cli.command is set
	Confirm bool     // destructive or money: ask before running
	Money   bool     // spends money (dry-run gives a quote when the body supports it)
	Hidden  bool     // x-cli.hidden
}

// PathParams returns the path parameters in URL order.
func (o *Op) PathParams() []Param {
	var out []Param
	for _, p := range o.Params {
		if p.In == "path" {
			out = append(out, p)
		}
	}
	return out
}

// Field returns the body field with the given name.
func (o *Op) Field(name string) (Field, bool) {
	for _, f := range o.Body {
		if f.Name == name {
			return f, true
		}
	}
	return Field{}, false
}

// Param returns the parameter with the given name.
func (o *Op) Param(name string) (Param, bool) {
	for _, p := range o.Params {
		if p.Name == name {
			return p, true
		}
	}
	return Param{}, false
}

// ByID returns the operation with the given operationId.
func ByID(id string) *Op {
	for i := range Ops {
		if Ops[i].ID == id {
			return &Ops[i]
		}
	}
	return nil
}
