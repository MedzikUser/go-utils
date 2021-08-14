package updater

import "time"

/*
	Auto checks for available updates

	This function checking latest version from github and automatically update your program
*/
func (c *Client) AutoUpdater() {
	// Check on start
	c.Update()

	ticker := time.NewTicker(c.CheckEvery)

	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			c.Update()

		case <-quit:
			ticker.Stop()

			return
		}
	}
}
