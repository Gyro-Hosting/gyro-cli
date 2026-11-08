// Package i18n picks Turkish or English messages. Messages are written inline as pairs:
//
//	i18n.S("Server deleted", "Sunucu silindi")
//
// The language comes from --lang, GYRO_LANG, the profile, then LC_ALL / LC_MESSAGES / LANG
// (tr* → Turkish, otherwise English).
package i18n

import (
	"fmt"
	"os"
	"strings"
)

var lang = Detect()

// Detect reads the locale from the environment.
func Detect() string {
	for _, k := range []string{"GYRO_LANG", "LC_ALL", "LC_MESSAGES", "LANG", "LANGUAGE"} {
		if v := os.Getenv(k); v != "" {
			return Normalize(v)
		}
	}
	return "en"
}

// Normalize maps tr_TR.UTF-8, tr, turkish ... to "tr"; anything else to "en".
func Normalize(v string) string {
	v = strings.ToLower(strings.TrimSpace(v))
	if strings.HasPrefix(v, "tr") || strings.HasPrefix(v, "turk") {
		return "tr"
	}
	return "en"
}

// Set forces a language ("tr" or "en").
func Set(l string) {
	if l != "" {
		lang = Normalize(l)
	}
}

// Lang returns the current language code.
func Lang() string { return lang }

// TR reports Turkish output.
func TR() bool { return lang == "tr" }

// S returns the message in the current language.
func S(en, tr string) string {
	if lang == "tr" && tr != "" {
		return tr
	}
	return en
}

// F formats the message in the current language.
func F(en, tr string, a ...any) string {
	return fmt.Sprintf(S(en, tr), a...)
}
