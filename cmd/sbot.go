package cmd

import (
	"github.com/restechnica/semverbot/pkg/cli/commands"
	"github.com/spf13/cobra"
)

// NewApp creates a new semverbot CLI app
// Returns a spf13/cobra command.
func NewApp() *cobra.Command {
	return commands.NewRootCommand()
}
