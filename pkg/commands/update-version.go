package commands

import (
	"github.com/spf13/cobra"

	"github.com/restechnica/semverbot/pkg/api"
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
	return gitAPI.FetchTags()
}
