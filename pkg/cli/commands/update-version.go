package commands

import (
	"github.com/restechnica/semverbot/pkg/core"

	"github.com/spf13/cobra"
)

// NewUpdateVersionCommand creates a new update version command.
// Returns the new spf13/cobra command.
func NewUpdateVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "version",
		RunE: UpdateVersionCommandRunE,
	}

	return command
}

// UpdateVersionCommandRunE runs the commands.
// Returns an error if it fails.
func UpdateVersionCommandRunE(cmd *cobra.Command, args []string) (err error) {
	return core.UpdateVersion()
}
