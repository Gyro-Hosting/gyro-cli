package cli

import (
	"fmt"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
	"github.com/Gyro-Hosting/gyro-cli/internal/output"
)

const opListBackupSpacePlans = "listBackupSpacePlans"

// printBackupSpacePlans renders GET /cloud/backup-spaces/plans ({"plans": [...], "servers": [...]}) as a plan table
// with the monthly / annual price, followed by the servers a space can be bound to (table and text output only).
func (a *App) printBackupSpacePlans(resp *api.Response) (bool, error) {
	p := a.Printer()
	if !(p.Format == "table" || p.Format == "text") || p.Quiet || a.Opt.Query != "" {
		return false, nil
	}
	v, err := output.Decode(resp.Data)
	if err != nil {
		return false, nil
	}
	m, _ := v.(map[string]any)
	if m == nil {
		return false, nil
	}
	plans, _ := m["plans"].([]any)
	rows := make([]any, 0, len(plans))
	for _, x := range plans {
		pl, _ := x.(map[string]any)
		if pl == nil {
			continue
		}
		row := map[string]any{"id": pl["id"], "plan": pl["plan"], "size_gb": pl["size_gb"], "available": pl["available"], "region": pl["region"], "monthly": "", "annually": ""}
		prices, _ := pl["prices"].([]any)
		for _, pr := range prices {
			pm, _ := pr.(map[string]any)
			if pm == nil {
				continue
			}
			if c, _ := pm["cycle"].(string); c == "monthly" || c == "annually" {
				row[c] = fmt.Sprintf("%v %v", pm["amount"], pm["currency"])
			}
		}
		rows = append(rows, row)
	}
	if len(rows) == 0 {
		a.Info("%s", i18n.S("No backup space plan is orderable for this account yet.", "Bu hesap için henüz sipariş edilebilir yedekleme alanı paketi yok."))
	} else if err := p.Print(rows, output.FromNames([]string{"id", "plan", "size_gb", "monthly", "annually", "available", "region"}, 0), "plan"); err != nil {
		return true, err
	}
	servers, _ := m["servers"].([]any)
	if len(servers) == 0 {
		a.Info("%s", i18n.S("No active server to bind a backup space to: create a server first.", "Yedekleme alanının bağlanacağı etkin sunucu yok: önce bir sunucu oluşturun."))
		return true, nil
	}
	a.Info("%s", i18n.S("Prices excl. VAT. Servers you can bind a space to (--server-id):", "Fiyatlar KDV hariç. Alanın bağlanabileceği sunucular (--server-id):"))
	for _, x := range servers {
		if s, _ := x.(map[string]any); s != nil {
			a.Info("  %v  %v  %v", output.Scalar(s["server_id"]), output.Scalar(s["ip"]), output.Scalar(s["name"]))
		}
	}
	return true, nil
}
