package commands

import (
	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "update",
	}

	command.AddCommand(NewUpdateVersionCommand())

	return command
}
