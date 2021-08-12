package convert

import (
	"testing"
	"time"
)

func TestSeconds1(t *testing.T) {
	d, err := time.ParseDuration("48h1m2s")
	if err != nil {
		t.Fatal(err)
	}
	out := Seconds(d)
	if out != "2d 1m 2s" {
		t.Fatal(out)
	}
}

func TestSeconds2(t *testing.T) {
	d, err := time.ParseDuration("48h5h1m2s")
	if err != nil {
		t.Fatal(err)
	}
	out := Seconds(d)
	if out != "2d 5h 1m 2s" {
		t.Fatal(out)
	}
}

func TestSeconds3(t *testing.T) {
	d, err := time.ParseDuration("48h5h60s")
	if err != nil {
		t.Fatal(err)
	}
	out := Seconds(d)
	if out != "2d 5h 1m" {
		t.Fatal(out)
	}
}
