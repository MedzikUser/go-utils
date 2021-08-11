package common

import (
	"github.com/sirupsen/logrus"
)

/*
	This function checks for an error.
	If the error isn't nil it return true and print error otherwise false.

		err := errors.New("Test Error")
		if common.CheckErr(err, "example error") {
			return
		}

		fmt.Println("No error was found!")
*/
func CheckErr(err error, trace string) bool {
	if err != nil {
		Log.WithFields(logrus.Fields{
			"trace": trace,
		}).Error(err)
	}

	return err != nil
}
