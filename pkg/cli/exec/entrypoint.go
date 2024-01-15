package exec

import (
	"errors"
	"os"

	"github.com/restechnica/semverbot/pkg/cli"
	"github.com/restechnica/semverbot/pkg/cli/commands"
	"github.com/rs/zerolog/log"
)

// Run will execute the CLI root command.
func Run() (err error) {
	var command = commands.NewRootCommand()

	if err = command.Execute(); err != nil {
		if errors.As(err, &cli.CommandError{}) {
			log.Error().Err(err).Msg("")
		}

		os.Exit(1)
	}

	return err
}
