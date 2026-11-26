package cli

import (
	"strings"

	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
)

// cycleHint explains 422 cycle_invalid (API 2.47.0, `gyro template deploy --cycle`): a Turkey VDS plan is billed with
// its normal invoice and only the periods the plan offers are accepted; details.cycles lists them.
func cycleHint(e *api.Error) string {
	var cs []string
	if l, ok := e.Details["cycles"].([]any); ok {
		for _, c := range l {
			if s, ok := c.(string); ok && s != "" {
				cs = append(cs, s)
			}
		}
	}
	if len(cs) > 0 {
		return i18n.F("This plan offers the periods: %s. Use one of them with --cycle (default monthly; Gyro Cloud Istanbul plans are hourly and ignore --cycle).",
			"Bu paketin dönemleri: %s. --cycle ile bunlardan birini verin (varsayılan monthly; Gyro Bulut İstanbul paketleri saatliktir, --cycle yok sayılır).",
			strings.Join(cs, ", "))
	}
	return i18n.S("Use a period the plan offers with --cycle (e.g. monthly, semiannually, annually).",
		"--cycle ile paketin sunduğu bir dönemi verin (ör. monthly, semiannually, annually).")
}
