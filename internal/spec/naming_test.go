package spec

import (
	"strings"
	"testing"
)

func TestSingular(t *testing.T) {
	cases := map[string]string{
		"servers": "server", "ssh-keys": "ssh-key", "invoices": "invoice", "addresses": "address", "aliases": "alias",
		"policies": "policy", "metrics": "metrics", "rdns": "rdns", "dns-records": "dns-record", "mailboxes": "mailbox",
		"buckets": "bucket", "settings": "setting", "status": "status", "ips": "ip", "boxes": "box", "keys": "key",
	}
	for in, want := range cases {
		if got := Singular(in); got != want {
			t.Errorf("Singular(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestKebab(t *testing.T) {
	for in, want := range map[string]string{"listSshKeys": "list-ssh-keys", "reset_password": "reset-password", "enable_rescue": "enable-rescue", "X.Y": "x-y"} {
		if got := Kebab(in); got != want {
			t.Errorf("Kebab(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestName(t *testing.T) {
	type tc struct {
		method, path, id                string
		list, hasGet, others, childParm bool
		want                            string
	}
	cases := []tc{
		{"GET", "/cloud/servers", "listServers", true, true, false, true, "server list"},
		{"POST", "/cloud/servers", "createServer", false, true, false, true, "server create"},
		{"GET", "/cloud/servers/{id}", "getServer", false, true, true, false, "server get"},
		{"PATCH", "/cloud/servers/{id}", "updateServer", false, true, true, false, "server update"},
		{"DELETE", "/cloud/servers/{id}", "deleteServer", false, true, true, false, "server delete"},
		{"GET", "/cloud/servers/{id}/snapshots", "listSnapshots", true, true, false, true, "server snapshot list"},
		{"POST", "/cloud/servers/{id}/snapshots/{sid}/restore", "restoreSnapshot", false, false, false, false, "server snapshot restore"},
		{"PUT", "/cloud/servers/{id}/rdns", "setRdns", false, true, false, false, "server rdns set"},
		{"GET", "/balance", "getBalance", false, true, false, false, "balance get"},
		{"GET", "/balance/transactions", "listBalanceTransactions", true, true, false, false, "balance transaction list"},
		{"GET", "/whoami", "whoami", false, true, false, false, "whoami get"},
		{"POST", "/domains/check", "coreCheckDomain", false, false, false, false, "domain check"},
		{"GET", "/domains", "coreGetDomains", false, false, false, true, "domain list"},
		{"GET", "/tickets", "coreGetTickets", false, true, false, true, "ticket list"},
		{"GET", "/domains/{domain}/dns-records", "coreGetDnsRecords", false, true, true, false, "domain dns-record list"},
		{"POST", "/domains/{domain}/dns-records", "coreCreateDnsRecord", false, true, true, false, "domain dns-record create"},
		{"GET", "/domains/{domain}/whois", "coreGetWhoisContacts", false, true, true, false, "domain whois get"},
		{"POST", "/domains/{domain}/auth-code", "coreSendAuthCode", false, false, true, false, "domain auth-code send"},
		{"POST", "/domains/{domain}/addons/{key}", "coreBuyDomainAddon", false, false, false, false, "domain addon buy"},
		{"POST", "/cloud/servers/{id}/firewall/rules", "addFirewallRule", false, false, false, true, "server firewall rule add"},
		{"POST", "/services/{id}/cancellation", "coreCreateCancellation", false, false, true, false, "service cancellation create"},
		{"POST", "/cloud/invoices/{id}/pay", "payInvoice", false, false, false, false, "invoice pay"},
		{"POST", "/tickets/{id}/reply", "replyTicket", false, false, false, false, "ticket reply"},
		{"POST", "/cloud/servers/{id}/console", "openConsole", false, false, false, false, "server console"},
		{"GET", "/cloud/reseller/customers", "listResellerCustomers", true, true, false, true, "reseller customer list"},
		{"POST", "/cloud/reseller/customers/{id}/keys", "createKey", false, true, false, true, "reseller customer key create"},
		{"GET", "/cloud/services/{id}/tools", "listServiceTools", false, false, false, true, "service tool list"},
		{"POST", "/cloud/services/{id}/tools/{tool}/{action}", "runServiceToolAction", false, false, false, false, "service tool run"},
		{"GET", "/sms/messages", "coreGetSmsMessages", false, false, false, true, "sms message list"},
		{"POST", "/cloud/bulk", "bulkServerAction", false, false, false, false, "bulk"},
		{"POST", "/cloud/webhooks/{id}/test", "testWebhook", false, false, false, false, "webhook test"},
	}
	for _, c := range cases {
		res, act := Name(NameInput{Method: c.method, Path: c.path, OpID: c.id, ListResp: c.list, HasGet: c.hasGet, HasOtherMethods: c.others, HasChildParam: c.childParm})
		got := strings.TrimSpace(strings.Join(append(res, act), " "))
		if got != c.want {
			t.Errorf("%s %s: got %q want %q", c.method, c.path, got, c.want)
		}
	}
	res, act := Name(NameInput{Method: "POST", Path: "/cloud/servers/{id}/actions/{action}", OpID: "serverAction", TrailingEnum: true})
	if strings.Join(res, " ") != "server action" || act != "" {
		t.Errorf("enum action: %v %q", res, act)
	}
}

func TestDangerous(t *testing.T) {
	if !Dangerous("DELETE", "deleteSshKey", "delete", "cloud_write", false) {
		t.Error("DELETE must confirm")
	}
	if !Dangerous("POST", "createServer", "create", "cloud_order", false) {
		t.Error("orders must confirm")
	}
	if Dangerous("POST", "createSnapshot", "create", "cloud_write", false) {
		t.Error("snapshot create should not confirm")
	}
	if !Dangerous("POST", "payInvoice", "pay", "", false) {
		t.Error("pay must confirm")
	}
	if Dangerous("GET", "listOrders", "list", "cloud_order", false) {
		t.Error("GET never confirms")
	}
}

func TestGeneratedTableSane(t *testing.T) {
	if len(Ops) < 25 {
		t.Fatalf("only %d ops", len(Ops))
	}
	seen := map[string]bool{}
	for _, o := range Ops {
		k := strings.Join(o.Command, " ")
		if seen[k] {
			t.Errorf("duplicate command %q", k)
		}
		seen[k] = true
		if o.Method == "DELETE" && !o.Confirm {
			t.Errorf("%s: DELETE without confirm", o.ID)
		}
		if o.Scope == "cloud_order" && o.Method != "GET" && !o.Money {
			t.Errorf("%s: order without money flag", o.ID)
		}
	}
	if ByID("createServer") == nil || !ByID("createServer").Money {
		t.Error("createServer missing or not money")
	}
}
