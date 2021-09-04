package common

import (
	"fmt"
	"time"
)

/*
	If there is an error in the performed function, it will be restarted, this process will be repeated the given number of times

	e.g.

		err := Retry(3, 5*time.Second, connectToDB)

		if err != nil {
			Log.Error(err)
		}
*/
func Retry(attempts int, sleep time.Duration, f func() error) (err error) {
	for i := 0; ; i++ {
		err = f()
		if err == nil {
			return
		}

		if i >= (attempts - 1) {
			break
		}

		time.Sleep(sleep)
		sleep *= 2

		Log.Error("Retrying after error: ", err)
	}

	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
