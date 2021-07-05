package cmd

import (
	"github.com/spf13/cobra"

	"github.com/restechnica/semverbot/pkg/commands"
)

// NewApp creates a new semverbot CLI app
// returns a spf13/cobra command.
func NewApp() *cobra.Command {
	return commands.NewRootCommand()
}
