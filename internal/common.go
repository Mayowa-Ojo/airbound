package internal

import (
	"airbound/internal/config"
	"strings"
	"time"
)

func SanitizePhone(s string) string {
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")

	return s
}

func ContainsString(s string, slice []string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}

	return false
}

func CheckInternalEmail(email string) bool {
	return strings.HasSuffix(email, "@airbound.com")
}

func CheckAircraftAge(manufactureDate time.Time, cfg *config.Config) bool {
	return time.Since(manufactureDate) < cfg.MaxAircraftAge
}
