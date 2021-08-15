package updater

import (
	"time"

	"github.com/MedzikUser/go-utils/common"
)

/*
	Auto checks for available updates

	This function checking latest version from github and automatically update your program
*/
func (c *Client) AutoUpdater() {
	// Check on start
	err := c.Update()
	common.CheckErr(err, "AutoUpdater")

	ticker := time.NewTicker(c.CheckEvery)

	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			err := c.Update()
			common.CheckErr(err, "AutoUpdater")

		case <-quit:
			ticker.Stop()

			return
		}
	}
}
