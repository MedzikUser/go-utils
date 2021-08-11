package common

import (
	"errors"
	"testing"
)

func TestCheckErrTrue(T *testing.T) {
	err := errors.New("Test")
	e := CheckErr(err, "test err")
	if e != true {
		T.Error(e)
	}
}

func TestCheckErrFalse(T *testing.T) {
	e := CheckErr(nil, "test err")
	if e != false {
		T.Error(e)
	}
}
