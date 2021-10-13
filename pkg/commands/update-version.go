package commands

import (
	"fmt"

	"github.com/restechnica/semverbot/pkg/api"
	"github.com/spf13/cobra"
)

// NewUpdateVersionCommand creates a new update version command.
// returns the new spf13/cobra command.
func NewUpdateVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "version",
		RunE: UpdateVersionCommandRunE,
	}

	return command
}

// UpdateVersionCommandRunE runs before the commands runs.
// returns an error if it fails.
func UpdateVersionCommandRunE(cmd *cobra.Command, args []string) (err error) {
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
