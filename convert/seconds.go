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

	hours, days := HoursToDays(hours)
	days, months := DaysToMonths(days)
	months, years := MonthsToYears(months)

	var format string

	if years > 0 {
		format += strconv.Itoa(years) + " years "
		months -= years * 12
	}

	if months > 0 {
		format += strconv.Itoa(months) + " months "
		days -= months * 30
	}

	if days > 0 {
		format += strconv.Itoa(days) + " days "
		hours -= days * 24
	}

	if hours > 0 {
		format += strconv.Itoa(hours) + " hours "
		minutes -= hours * 24
	}

	if minutes > 0 {
		format += strconv.Itoa(minutes) + " minutes "
		seconds -= minutes * 60
	}

	if seconds > 0 {
		format += strconv.Itoa(seconds) + " seconds"
	}

	return format
}

func HoursToDays(hours int) (int, int) {
	var days int

	for hours/24 > 0 {
		days++
		hours -= 24
	}

	return hours, days
}

func DaysToMonths(days int) (int, int) {
	var months int

	for days/30 > 0 {
		months++
		days -= 30
	}

	return days, months
}

func MonthsToYears(months int) (int, int) {
	var years int

	for months/12 > 0 {
		years++
		months -= 12
	}

	return months, years
}
