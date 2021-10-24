package common

import (
"fmt"
"testing"
"time"
)

func TestCheckUptime(T *testing.T) {
	uptime := Uptime(time.Date(2016, 6, 2, 1, 1, 1, 1, time.UTC))
fmt.Println(uptime)
}
