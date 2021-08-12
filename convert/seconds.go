package convert

import (
	"strconv"
	"time"
)

// FIXME:
func Seconds(s time.Duration) string {
	hours := Round(s.Hours())
	minutes := Round(s.Minutes())
	seconds := Round(s.Seconds())

	var (
		days   int
		format string
	)

	for hours/24 > 0 {
		days++
		hours -= 24
	}

	if days > 0 {
		format += strconv.Itoa(days) + "d"
	}

	if hours > 0 {
		if len(format) >= 1 {
			format += " "
		}

		format += strconv.Itoa(hours) + "h"
	} else if days > 0 {
		hours = days * 24
	}

	minutes -= Round(s.Hours()) * 60

	if minutes > 0 {
		if len(format) >= 1 {
			format += " "
		}

		format += strconv.Itoa(minutes) + "m"
	}

	seconds -= Round(s.Minutes()) * 60

	if seconds > 0 {
		if len(format) >= 1 {
			format += " "
		}

		format += strconv.Itoa(seconds) + "s"
	}

	return format
}
