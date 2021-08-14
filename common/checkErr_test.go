package common

import (
	"errors"
	"testing"
)

func TestCheckErrTrue(T *testing.T) {
	// true
	err := errors.New("Test")
	e := CheckErr(err, "test err")
	if e != true {
		T.Error(e)
	}

	// false
	e = CheckErr(nil, "test err")
	if e != false {
		T.Error(e)
	}
}
