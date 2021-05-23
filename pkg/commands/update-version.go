package commands

import (
	"fmt"

	"github.com/restechnica/semverbot/pkg/api"
	"github.com/spf13/cobra"
)

func NewUpdateVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "version",
		RunE: UpdateVersionCommandRunE,
	}

	return command
}

func UpdateVersionCommandRunE(cmd *cobra.Command, args []string) (err error) {
	var gitAPI = api.NewGitAPI()

	if err = gitAPI.FetchTags(); err != nil {
		fmt.Println("something went wrong while updating the version, you probably already have the latest version")
	}

	return err
}
