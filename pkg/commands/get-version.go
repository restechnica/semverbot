package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/restechnica/semverbot/pkg/api"
	"github.com/restechnica/semverbot/pkg/cli"
)

// NewGetVersionCommand creates a new get version command.
// returns the new spf13/cobra command.
func NewGetVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "version",
		Run: GetVersionCommandRun,
	}

	return command
}

// GetVersionCommandRun runs the command.
func GetVersionCommandRun(cmd *cobra.Command, args []string) {
	var versionAPI = api.NewVersionAPI()
	var version = versionAPI.GetVersionOrDefault(cli.DefaultVersion)
	fmt.Println(version)
}
