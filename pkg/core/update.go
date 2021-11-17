package core

import (
	"fmt"

	"github.com/restechnica/semverbot/pkg/api"
)

// UpdateVersion changes the current git repo into an unshallow repo and fetches all git tags.
// Returns and error if anything went wrong.
func UpdateVersion() (err error) {
	var gitAPI = api.NewGitAPI()

	if err = gitAPI.FetchUnshallow(); err != nil {
		fmt.Println("something went wrong while fetching from git, attempting to fetch tags anyway")
	}

	if err = gitAPI.FetchTags(); err != nil {
		fmt.Println("something went wrong while updating the version")
	} else {
		fmt.Println("successfully fetched the latest git tags")
	}

	return err
}
