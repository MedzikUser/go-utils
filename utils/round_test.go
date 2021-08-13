package utils

import "testing"

func TestRound(t *testing.T) {
	// 1
	r := Round(0.1)
	if r != 0 {
		t.Fatal(r)
	}

	// 2
	r = Round(0.9)
	if r != 0 {
		t.Fatal(r)
	}

	// 3
	r = Round(25.2)
	if r != 25 {
		t.Fatal(r)
	}

	// 4
	r = Round(46.9998655)
	if r != 46 {
		t.Fatal(r)
	}
}
