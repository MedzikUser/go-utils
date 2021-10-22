package utils

import "testing"

func TestBytes(t *testing.T) {
	// 1
	r := Bytes(53)
	if r != "53 B" {
		t.Fatal(r)
	}

	// 2
	r = Bytes(1522)
	if r != "1.5 kB" {
		t.Fatal(r)
	}
}
