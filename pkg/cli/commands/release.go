package commands

import "github.com/spf13/cobra"

func NewReleaseCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "release",
	}

	command.AddCommand(NewReleaseVersionCommand())

	return command
}
