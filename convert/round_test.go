package convert

import "testing"

func TestRound1(t *testing.T) {
	r := Round(0.1)
	if r != 0 {
		t.Fatal(r)
	}
}

func TestRound2(t *testing.T) {
	r := Round(0.9)
	if r != 0 {
		t.Fatal(r)
	}
}

func TestRound3(t *testing.T) {
	r := Round(25.2)
	if r != 25 {
		t.Fatal(r)
	}
}

func TestRound4(t *testing.T) {
	r := Round(46.9998655)
	if r != 46 {
		t.Fatal(r)
	}
}
