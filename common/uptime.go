package common

import (
	"time"

	"github.com/MedzikUser/go-utils/utils"
)

// Return uptime of program e.g. "10m 4s"
func Uptime(start time.Time) string {
	return utils.FormatSeconds(time.Since(start))
}
