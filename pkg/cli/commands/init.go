package commands

import (
	"fmt"

	"github.com/restechnica/semverbot/pkg/cli"
	"github.com/restechnica/semverbot/pkg/core"
	"github.com/spf13/cobra"
)

// NewInitCommand creates a new init command.
// Returns a new init spf13/cobra command.
func NewInitCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "init",
		RunE:  InitCommandRunE,
		Short: fmt.Sprintf(`Creates a default "%s" config`, cli.DefaultConfigFilePath),
	}

	return command
}

// InitCommandRunE runs the init command.
// Returns an error if the command failed.
func InitCommandRunE(cmd *cobra.Command, args []string) (err error) {
	var options = &core.InitOptions{
		Config:         cli.DefaultConfig,
		ConfigFilePath: cli.DefaultConfigFilePath,
	}

	return core.Init(options)
}
