package v1

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/restechnica/semverbot/pkg/cli"
	"github.com/restechnica/semverbot/pkg/core"
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
	log.Debug().Str("command", "init").Msg("starting run...")

	var options = &core.InitOptions{
		Config:         cli.GetDefaultConfig(),
		ConfigFilePath: cli.ConfigFlag,
	}

	log.Debug().Str("config", options.ConfigFilePath).Msg("options")

	if err = core.Init(options); err != nil {
		err = cli.NewCommandError(err)
	}

	return err
}
