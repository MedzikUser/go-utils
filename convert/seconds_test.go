package convert

import (
	"testing"
	"time"
)

func TestSeconds(t *testing.T) {
	d, err := time.ParseDuration("48h")
	t.Error("1", err)
	out := Seconds(d)
	t.Error("2", out)
}
