package v1

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/restechnica/semverbot/pkg/cli"
	"github.com/restechnica/semverbot/pkg/core"
)

// NewUpdateVersionCommand creates a new update version command.
// Returns the new spf13/cobra command.
func NewUpdateVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "version",
		RunE: UpdateVersionCommandRunE,
	}

	return command
}

// UpdateVersionCommandRunE runs the commands.
// Returns an error if it fails.
func UpdateVersionCommandRunE(cmd *cobra.Command, args []string) (err error) {
	log.Debug().Str("command", "update version").Msg("starting run...")

	if err = core.UpdateVersion(); err != nil {
		err = cli.NewCommandError(err)
	}

	return err
}
