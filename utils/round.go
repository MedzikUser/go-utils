package utils

// Math round down
func Round(val float64) int {
	if val < 0 {
		return int(val - 1.0)
	}

	return int(val)
}
