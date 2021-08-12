package common

import (
	"time"

	"github.com/MedzikUser/go-utils/convert"
)

func Uptime(start time.Time) string {
	return convert.Seconds(time.Since(start))
}
