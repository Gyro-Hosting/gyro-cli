package cli

import (
	"github.com/Gyro-Hosting/gyro-cli/internal/api"
	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
	"github.com/Gyro-Hosting/gyro-cli/internal/output"
)

// lockEndHint explains 422 retention_exceeds_service_end (Protected Storage, API 2.46.1): an object lock
// (bucket default retention or a version retention) may not end after the service end (renewal) date.
func lockEndHint(e *api.Error) string {
	end, mx := e.Details["service_end"], e.Details["max_days"]
	if end != nil && mx != nil {
		return i18n.F("A lock may not end after the service end date (%s): use at most %s days now, or renew the service / move it to a longer billing cycle (annually, biennially = 2 years, triennially = 3 years) for a longer lock. See `gyro storage object-lock status <ID>`.",
			"Kilit hizmet bitiş tarihinden (%s) sonra bitemez: şu an en çok %s gün verin ya da daha uzun kilit için hizmeti yenileyin / daha uzun döneme geçirin (annually, biennially = 2 yıl, triennially = 3 yıl). Bkz. `gyro storage object-lock status <ID>`.",
			output.Scalar(end), output.Scalar(mx))
	}
	return i18n.S("A lock may not end after the service end date; renew the service or move it to a longer billing cycle (biennially, triennially) for a longer lock. See `gyro storage object-lock status <ID>`.",
		"Kilit hizmet bitiş tarihinden sonra bitemez; daha uzun kilit için hizmeti yenileyin ya da daha uzun döneme geçirin (biennially, triennially). Bkz. `gyro storage object-lock status <ID>`.")
}
