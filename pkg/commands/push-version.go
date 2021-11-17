package commands

import (
	"github.com/restechnica/semverbot/pkg/core"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/restechnica/semverbot/pkg/cli"
)

// NewPushVersionCommand creates a new push version command.
// Returns the new spf13/cobra command.
func NewPushVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "version",
		RunE: PushVersionCommandRunE,
	}

	return command
}

// PushVersionCommandRunE runs the command.
// Returns an error if the command fails.
func PushVersionCommandRunE(cmd *cobra.Command, args []string) (err error) {
	var options = &core.PushVersionOptions{
		DefaultVersion: cli.DefaultVersion,
		GitTagsPrefix:  viper.GetString(cli.GitTagsPrefixConfigKey),
	}

	return core.PushVersion(options)
}
