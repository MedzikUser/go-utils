package utils

import (
	"testing"
	"time"
)

func TestFormatSeconds(t *testing.T) {
	// 1
	d, err := time.ParseDuration("48h1m2s")
	if err != nil {
		t.Fatal(err)
	}

	out := FormatSeconds(d)
	if out != "2d 1m 2s" {
		t.Fatal(out)
	}

	// 2
	d, err = time.ParseDuration("48h5h1m2s")
	if err != nil {
		t.Fatal(err)
	}

	out = FormatSeconds(d)
	if out != "2d 5h 1m 2s" {
		t.Fatal(out)
	}

	// 3
	d, err = time.ParseDuration("48h5h60s")
	if err != nil {
		t.Fatal(err)
	}

	out = FormatSeconds(d)
	if out != "2d 5h 1m" {
		t.Fatal(out)
	}

	// 4
	d, err = time.ParseDuration("0s")
	if err != nil {
		t.Fatal(err)
	}

	out = FormatSeconds(d)
	if out != "0s" {
		t.Fatal(out)
	}
}
