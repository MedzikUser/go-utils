package updater

import (
	"os"

	"github.com/MedzikUser/go-github-selfupdate/selfupdate"
	"github.com/blang/semver/v4"
)

/*
	Checks for available updates

	If there is an update, this function will automatically update your program
*/
func (c *Client) Update() error {
	updater, err := selfupdate.NewUpdater(selfupdate.Config{
		APIToken: c.GitHubToken,
	})

	release, found, err := updater.DetectLatest(c.GitHub)
	if err != nil {
		return err
	}

	version, err := semver.Parse(c.Version)
	if err != nil || !found || release.Version.LTE(version) {
		return err
	}

	exe, err := os.Executable()
	if err != nil {
		return err
	}

	err = updater.UpdateTo(release, exe, c.Binary)
	if err != nil {
		return err
	}

	c.AfterUpdate()

	return nil
}
