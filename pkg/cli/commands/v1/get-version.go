package v1

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/restechnica/semverbot/pkg/cli"
	"github.com/restechnica/semverbot/pkg/core"
)

// NewGetVersionCommand creates a new get version command.
// Returns the new spf13/cobra command.
func NewGetVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "version",
		Run: GetVersionCommandRun,
	}

	return command
}

// GetVersionCommandRun runs the command.
func GetVersionCommandRun(cmd *cobra.Command, args []string) {
	log.Debug().Str("command", "get version").Msg("starting run...")
	var options = &core.GetVersionOptions{DefaultVersion: cli.DefaultVersion}
	var version = core.GetVersion(options)
	fmt.Println(version)
}
