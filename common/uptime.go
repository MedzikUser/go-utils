package common

import (
	"time"

	"github.com/MedzikUser/go-utils/utils"
)

func Uptime(start time.Time) string {
	return utils.FormatSeconds(time.Since(start))
}
