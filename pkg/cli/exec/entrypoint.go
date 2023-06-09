package exec

import (
	"os"

	"github.com/restechnica/semverbot/pkg/cli/commands"
)

// Run will execute the CLI root command.
func Run() (err error) {
	var command = commands.NewRootCommand()

	if err = command.Execute(); err != nil {
		os.Exit(1)
	}

	return err
}
