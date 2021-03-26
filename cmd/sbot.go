package cmd

import (
	"github.com/spf13/cobra"

	"github.com/restechnica/semverbot/pkg/cli/commands"
)

func NewApp() *cobra.Command {
	return commands.NewRootCommand()
}
