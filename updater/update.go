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
	release, found, err := selfupdate.DetectLatest(c.GitHub)
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

	err = selfupdate.UpdateTo(release.AssetURL, exe, c.Binary)
	if err != nil {
		return err
	}

	return nil
}
