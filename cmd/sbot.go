package cmd

import (
	"github.com/spf13/cobra"

	"github.com/restechnica/semverbot/pkg/commands"
)

// NewApp creates a new semverbot CLI app
// Returns a spf13/cobra command.
func NewApp() *cobra.Command {
	return commands.NewRootCommand()
}
