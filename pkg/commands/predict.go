package commands

import (
	"github.com/spf13/cobra"
)

// NewPredictCommand creates a new predict command.
// returns the new spf13/cobra command.
func NewPredictCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "predict",
	}

	command.AddCommand(NewPredictVersionCommand())

	return command
}
