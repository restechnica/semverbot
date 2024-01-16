package v1

import (
	"github.com/spf13/cobra"
)

// NewV1Command creates a new V1 root command.
// Returns the new V1 root command.
func NewV1Command() *cobra.Command {
	var command = &cobra.Command{
		Use:   "v1",
		Short: "v1 sbot API",
	}

	command.AddCommand(NewGetCommand())
	command.AddCommand(NewInitCommand())
	command.AddCommand(NewPredictCommand())
	command.AddCommand(NewPushCommand())
	command.AddCommand(NewReleaseCommand())
	command.AddCommand(NewUpdateCommand())
	command.AddCommand(NewVersionCommand())

	return command
}
