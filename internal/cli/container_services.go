package cli

import (
	"encoding/json"
	"fmt"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
	"github.com/Gyro-Hosting/gyro-cli/internal/output"
)

// 1.26.5 (gece testi): `gyro container service list` printed the wrapped answer ({"services": [...], "limits": ...}) as key: value
// lines, so the services themselves never appeared in the table. The table now lists the services; JSON/YAML/--query keep the
// full answer unchanged (scripts that read .services keep working).
const opListContainerServices = "typed_get_containers__id__services"

func containerServiceColumns() []output.Column {
	return []output.Column{
		{Header: "name", Path: "name"},
		{Header: "image", Path: "image"},
		{Header: i18n.S("copies", "kopya"), Path: "copies"},
		{Header: "state", Path: "state"},
		{Header: i18n.S("problem", "sorun"), Path: "problem"},
		{Header: "url", Path: "url"},
	}
}

// containerServiceRows flattens the services for the table: copies = ready/desired (+ autoscaling range).
func containerServiceRows(data json.RawMessage) (rows []any, limits map[string]any, ok bool) {
	v, err := output.Decode(data)
	if err != nil {
		return nil, nil, false
	}
	m, isMap := v.(map[string]any)
	if !isMap {
		return nil, nil, false
	}
	list, isList := m["services"].([]any)
	if !isList {
		return nil, nil, false
	}
	limits, _ = m["limits"].(map[string]any)
	rows = make([]any, 0, len(list))
	for _, it := range list {
		s, _ := it.(map[string]any)
		if s == nil {
			continue
		}
		row := map[string]any{}
		for k, val := range s {
			row[k] = val
		}
		want := s["current_replicas"]
		if want == nil {
			want = s["replicas"]
		}
		copies := fmt.Sprintf("%s/%s", output.Scalar(s["ready"]), output.Scalar(want))
		if as, _ := s["autoscaling"].(map[string]any); as != nil && as["enabled"] != false {
			if mn, mx := as["min"], as["max"]; mn != nil && mx != nil {
				copies += fmt.Sprintf(" (%s-%s)", output.Scalar(mn), output.Scalar(mx))
			}
		}
		row["copies"] = copies
		rows = append(rows, row)
	}
	return rows, limits, true
}

// printContainerServices returns done=false when the answer does not have the expected shape (generic printing then).
func (a *App) printContainerServices(resp *api.Response) (bool, error) {
	p := a.Printer()
	if a.Opt.Query != "" || (p.Format != "table" && p.Format != "text" && !p.Quiet) {
		return false, nil
	}
	rows, limits, ok := containerServiceRows(resp.Data)
	if !ok {
		return false, nil
	}
	if p.Quiet {
		return true, p.Print(rows, nil, "name")
	}
	cols := containerServiceColumns()
	if a.Opt.Wide {
		cols = append(cols, output.Column{Header: "cpu_m", Path: "cpu_m"}, output.Column{Header: "memory_mi", Path: "memory_mi"},
			output.Column{Header: "port", Path: "port"}, output.Column{Header: "public", Path: "public"}, output.Column{Header: "revision", Path: "revision"})
	}
	if err := p.Print(rows, cols, "name"); err != nil {
		return true, err
	}
	if p.Format == "table" && !a.Machine() {
		if len(rows) == 0 {
			a.Info("%s", i18n.S("No services in this environment yet. Create one: gyro container service create <ENV> --name web --image nginx --port 80",
				"Bu ortamda henüz servis yok. Oluşturmak için: gyro container service create <ORTAM> --name web --image nginx --port 80"))
		}
		if limits != nil && limits["max_services"] != nil {
			a.Info("%s", i18n.F("Plan: at most %s services, %s copies per service.", "Plan: en fazla %s servis, servis başına en fazla %s kopya.",
				output.Scalar(limits["max_services"]), output.Scalar(limits["max_replicas"])))
		}
	}
	return true, nil
}

// replicaHint explains 422 invalid_replicas with the plan limit (details.max, API 2.24.4+).
func replicaHint(e *api.Error) string {
	if mx, ok := e.Details["max"]; ok && mx != nil {
		return i18n.F("Your plan allows at most %s copies per service: use a number from 0 to %s (or switch on autoscaling within that range).",
			"Planınız servis başına en fazla %s kopyaya izin veriyor: 0 ile %s arasında bir sayı verin (ya da bu aralıkta otomatik ölçeklemeyi açın).",
			output.Scalar(mx), output.Scalar(mx))
	}
	return i18n.S("The copy count is outside your plan's per-service limit; see `gyro container service list <ENV>` for the limit.",
		"Kopya sayısı planınızın servis başına sınırının dışında; sınır için `gyro container service list <ORTAM>`.")
}
