package commands

import (
	"github.com/spf13/cobra"
)

// NewPushCommand creates a new push command.
// returns the new spf13/cobra command.
func NewPushCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "push",
	}

	command.AddCommand(NewPushVersionCommand())

	return command
}
