package v1

import (
	"github.com/spf13/cobra"
)

// NewPushCommand creates a new push command.
// Returns the new spf13/cobra command.
func NewPushCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "push",
	}

	command.AddCommand(NewPushVersionCommand())

	return command
}
