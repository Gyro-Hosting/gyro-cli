package spec

import (
	"regexp"
	"strings"
	"unicode"
)

// Naming rules (path → command). Used by cmd/gyro-gen; unit tested in naming_test.go.
//
//	GET    /cloud/servers                     server list
//	POST   /cloud/servers                     server create
//	GET    /cloud/servers/{id}                server get <id>
//	PATCH  /cloud/servers/{id}                server update <id>
//	DELETE /cloud/servers/{id}                server delete <id>
//	POST   /cloud/servers/{id}/actions/{a}    server action <enum value> <id>   (one command per enum value)
//	GET    /cloud/servers/{id}/snapshots      server snapshot list <id>
//	POST   /cloud/servers/{id}/snapshots/{sid}/restore   server snapshot restore <id> <sid>  (POST verb, no GET on path)
//	GET    /balance                           balance get   (also runnable as bare "gyro balance")
//	POST   /domains/check                     domain check  (POST on a literal without a GET twin = verb)
//
// The leading "cloud" segment is dropped. An OpenAPI extension `x-cli: {command: "a b c"}` overrides everything.

var irregular = map[string]string{
	"metrics": "metrics", "rdns": "rdns", "dns": "dns", "status": "status", "settings": "setting",
	"aliases": "alias", "addresses": "address", "news": "news", "sso": "sso", "analytics": "analytics",
	"activity": "activity", "balance": "balance", "whoami": "whoami", "series": "series", "ips": "ip",
	"statuses": "status", "boxes": "box", "mailboxes": "mailbox", "catch-all": "catch-all", "iso": "iso",
	"access": "access", "s3": "s3", "dnssec": "dnssec", "tools": "tool", "methods": "method",
	"deleted": "deleted", "archive": "archive", "quarantine": "quarantine", "two-factor": "two-factor",
	"delivery-log": "delivery-log", "plan": "plan", "migrations": "migration", "bulk": "bulk",
	"report": "report", "me": "me", "console": "console", "firewall": "firewall", "resize": "resize",
	"keys": "key", "ssh-keys": "ssh-key", "whois": "whois", "sms": "sms", "dashboard": "dashboard", "affiliate": "affiliate", "forwarding": "forwarding", "cancellation": "cancellation", "security": "security", "billing": "billing", "default-nameservers": "default-nameserver", "dns-records": "dns-record", "nameservers": "nameserver",
}

// Singular converts a plural path noun to its singular command name.
func Singular(s string) string {
	s = strings.ToLower(s)
	if v, ok := irregular[s]; ok {
		return v
	}
	// compound nouns: singularise the last word (dns-records -> dns-record)
	if i := strings.LastIndexAny(s, "-_"); i > 0 {
		return s[:i+1] + Singular(s[i+1:])
	}
	switch {
	case strings.HasSuffix(s, "ies") && len(s) > 4:
		return s[:len(s)-3] + "y"
	case strings.HasSuffix(s, "sses"), strings.HasSuffix(s, "xes"), strings.HasSuffix(s, "ches"), strings.HasSuffix(s, "shes"):
		return s[:len(s)-2]
	case strings.HasSuffix(s, "ss"), strings.HasSuffix(s, "us"):
		return s
	case strings.HasSuffix(s, "s") && len(s) > 2:
		return s[:len(s)-1]
	}
	return s
}

var camelRe = regexp.MustCompile(`([a-z0-9])([A-Z])`)

// Kebab turns operationIds / snake_case / camelCase into kebab-case.
func Kebab(s string) string {
	s = camelRe.ReplaceAllString(s, "${1}-${2}")
	s = strings.NewReplacer("_", "-", " ", "-", ".", "-", "/", "-").Replace(s)
	return strings.Trim(strings.ToLower(s), "-")
}

// FirstWord returns the leading verb of an operationId (runServiceTool -> run).
func FirstWord(id string) string {
	for i, r := range id {
		if i > 0 && unicode.IsUpper(r) {
			return strings.ToLower(id[:i])
		}
	}
	return strings.ToLower(id)
}

type seg struct {
	name  string
	param bool
}

func split(path string) []seg {
	var out []seg
	for _, p := range strings.Split(strings.Trim(path, "/"), "/") {
		if p == "" {
			continue
		}
		if strings.HasPrefix(p, "{") && strings.HasSuffix(p, "}") {
			out = append(out, seg{p[1 : len(p)-1], true})
		} else {
			out = append(out, seg{p, false})
		}
	}
	if len(out) > 0 && !out[0].param && out[0].name == "cloud" {
		out = out[1:]
	}
	return out
}

// NameInput describes an operation and its neighbourhood in the spec.
type NameInput struct {
	Method, Path, OpID string
	ListResp           bool // data is an array
	HasGet             bool // a GET exists on the same path
	HasOtherMethods    bool // PUT/PATCH/DELETE exist on the same path
	HasChildParam      bool // a path "<this>/{param}" exists (so this path is a collection)
	TrailingEnum       bool // last path parameter has an enum (one command per value)
}

// Verb returns the leading verb of an operationId, ignoring a "core" prefix (coreSendAuthCode -> send).
func Verb(opID string) string {
	id := opID
	if strings.HasPrefix(id, "core") && len(id) > 4 && unicode.IsUpper(rune(id[4])) {
		id = strings.ToLower(id[4:5]) + id[5:]
	}
	return FirstWord(id)
}

// pluralGetter reports operationIds like getDomains / coreGetTickets / listX.
func pluralGetter(opID string) bool {
	id := strings.TrimPrefix(opID, "core")
	if strings.HasPrefix(id, "list") || strings.HasPrefix(id, "List") {
		return true
	}
	return (strings.HasPrefix(id, "Get") || strings.HasPrefix(id, "get")) && strings.HasSuffix(id, "s") && !strings.HasSuffix(id, "ss") && !strings.HasSuffix(id, "Status") && !strings.HasSuffix(id, "Settings")
}

// Name derives the command words for an operation.
func Name(in NameInput) (resource []string, action string) {
	segs := split(in.Path)
	method := strings.ToUpper(in.Method)
	if len(segs) == 0 {
		return []string{Kebab(in.OpID)}, ""
	}
	last := segs[len(segs)-1]
	nouns := func(ss []seg) []string {
		var r []string
		for _, s := range ss {
			if !s.param {
				r = append(r, Singular(s.name))
			}
		}
		return r
	}
	verb := Verb(in.OpID)
	if last.param {
		if in.TrailingEnum {
			// .../actions/{action}: resource = nouns, action = enum value (filled by the generator)
			return nouns(segs), ""
		}
		switch method {
		case "GET":
			action = "get"
		case "DELETE":
			action = "delete"
		case "PUT":
			action = "set"
		case "PATCH":
			action = "update"
		default:
			action = verb
			if action == "" || action == "post" {
				action = "run"
			}
		}
		return nouns(segs), action
	}
	// last segment is a literal
	if method == "POST" && !in.HasGet {
		switch {
		case in.HasChildParam:
			// POST /firewall/rules with DELETE /firewall/rules/{n}: a collection
			action = "create"
			if verb == "add" {
				action = "add"
			}
			return nouns(segs), action
		case in.HasOtherMethods:
			// POST /auth-code next to PUT /auth-code: name it after the operation verb
			action = verb
			if action == "" || action == "post" {
				action = "create"
			}
			return nouns(segs), action
		case len(segs) > 1:
			return nouns(segs[:len(segs)-1]), Kebab(last.name)
		default:
			return []string{Kebab(last.name)}, "" // POST /cloud/bulk -> gyro bulk
		}
	}
	switch method {
	case "GET":
		lastPlural := strings.ToLower(Singular(last.name)) != strings.ToLower(last.name)
		if in.ListResp || in.HasChildParam || (lastPlural && pluralGetter(in.OpID)) {
			action = "list"
		} else {
			action = "get"
		}
	case "POST":
		action = "create"
	case "PUT":
		action = "set"
	case "PATCH":
		action = "update"
	case "DELETE":
		action = "delete"
	}
	return nouns(segs), action
}

var dangerousWords = regexp.MustCompile(`(?i)(delete|remove|destroy|restore|rebuild|reinstall|reset|terminate|cancel|detach|unmount|wipe|purge|pay|renew|register|transfer|resize|order|upgrade|disable|poweroff|send|bulk|withdraw|assign|cancellation|password|sso|buy|purchase|addon|method|nameserver|privacy|tool)`)

// Dangerous reports whether running the operation should ask for confirmation.
func Dangerous(method, opID, action, scope string, hasConfirm bool) bool {
	if strings.EqualFold(method, "DELETE") || hasConfirm {
		return true
	}
	if scope == "cloud_order" || scope == "cloud_delete" || scope == "billing_pay" {
		return method != "GET"
	}
	if method != "GET" && (dangerousWords.MatchString(opID) || dangerousWords.MatchString(action)) {
		return true
	}
	return false
}
