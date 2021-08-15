package updater

import "time"

type Client struct {
	// GitHub repo with author e.g. "MedzikUser/go-utils"
	GitHub string
	// GitHub Token
	GitHubToken string
	// Application version e.g. "2.15.43"
	Version string
	// Binary name in archive e.g. "utils.out"
	Binary string
	// Check latest version every
	CheckEvery time.Duration
	/*
		After update exec function e.g.
			func() {
				os.Exit(0)
			}
	*/
	AfterUpdate func()
}
