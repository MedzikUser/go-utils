package stats

import "testing"

func TestCPU(t *testing.T) {
	_, err := CPU()
	if err != nil {
		t.Fatal(err)
	}
}
