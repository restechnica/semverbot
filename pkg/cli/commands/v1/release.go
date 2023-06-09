package v1

import "github.com/spf13/cobra"

// NewReleaseCommand creates a new release command.
// Returns the new spf13/cobra command.
func NewReleaseCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "release",
	}

	command.AddCommand(NewReleaseVersionCommand())

	return command
}
