package commands

import (
	"github.com/spf13/cobra"
)

// NewUpdateCommand creates a new update command.
// Returns the new spf13/cobra command.
func NewUpdateCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "update",
	}

	command.AddCommand(NewUpdateVersionCommand())

	return command
}
