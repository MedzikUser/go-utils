package utils

import "fmt"

/*
	Format bytes e.g.
		Bytes(53) // out: "53 B"
		Bytes(1522) // out: "1.5 kB"
*/
func Bytes(b uint64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}

	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}
