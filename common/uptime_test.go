package common

import (
"testing"
"time"
)

func TestCheckUptime(T *testing.T) {
	Uptime(time.Date(2016, 6, 2, 1, 1, 1, 1, time.UTC))
}
