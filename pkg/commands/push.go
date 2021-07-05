package commands

import (
	"github.com/spf13/cobra"
)

func NewPushCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "push",
	}

	command.AddCommand(NewPushVersionCommand())

	return command
}
