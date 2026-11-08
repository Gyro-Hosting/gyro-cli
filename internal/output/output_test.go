package output

import (
	"bytes"
	"strings"
	"testing"
)

func TestScalar(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{nil, "-"}, {"", "-"}, {true, "yes"}, {float64(47466), "47466"}, {553.49, "553.49"},
		{map[string]any{"amount": 553.49, "currency": "TRY"}, "553.49 TRY"},
		{map[string]any{"amount": float64(0), "currency": nil}, "0.00"},
		{[]any{"a", "b"}, "a,b"},
	}
	for _, c := range cases {
		if got := Scalar(c.in); got != c.want {
			t.Errorf("Scalar(%v) = %q want %q", c.in, got, c.want)
		}
	}
}

func TestDeriveOrder(t *testing.T) {
	rows := []any{map[string]any{"zeta": 1.0, "status": "active", "id": 1.0, "name": "x", "nested": map[string]any{"a": 1.0}, "price": map[string]any{"amount": 1.0, "currency": "TRY"}}}
	var hs []string
	for _, c := range Derive(rows, false) {
		hs = append(hs, c.Header)
	}
	if strings.Join(hs, ",") != "id,name,status,price,zeta" {
		t.Errorf("derive order: %v", hs)
	}
}

func TestFormats(t *testing.T) {
	data := []any{map[string]any{"id": 1.0, "name": "a"}, map[string]any{"id": 2.0, "name": "b"}}
	var b bytes.Buffer
	p := &Printer{W: &b, Format: "table"}
	_ = p.Print(data, nil, "")
	if !strings.Contains(b.String(), "ID   NAME") {
		t.Errorf("table: %q", b.String())
	}
	b.Reset()
	p.NoHeader = true
	_ = p.Print(data, nil, "")
	if strings.Contains(b.String(), "NAME") {
		t.Errorf("no-header: %q", b.String())
	}
	b.Reset()
	p = &Printer{W: &b, Quiet: true}
	_ = p.Print(data, nil, "")
	if b.String() != "1\n2\n" {
		t.Errorf("quiet: %q", b.String())
	}
	b.Reset()
	p = &Printer{W: &b, Format: "json", Query: "[?name=='b'].id | [0]"}
	_ = p.Print(data, nil, "")
	if strings.TrimSpace(b.String()) != "2" {
		t.Errorf("query: %q", b.String())
	}
	b.Reset()
	p = &Printer{W: &b, Format: "yaml"}
	_ = p.Print(map[string]any{"amount": 553.49}, nil, "")
	if strings.TrimSpace(b.String()) != "amount: 553.49" {
		t.Errorf("yaml: %q", b.String())
	}
	b.Reset()
	p = &Printer{W: &b, Format: "json", Query: "[["}
	if err := p.Print(data, nil, ""); err == nil {
		t.Error("bad query should fail")
	}
}

func TestColorize(t *testing.T) {
	if Colorize("active") == "active" || Colorize("whatever") != "whatever" {
		t.Error("colorize")
	}
}
