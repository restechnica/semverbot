package commands

import (
	"io"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"

	"github.com/restechnica/semverbot/pkg/cli"
)

// NewInitCommand creates a new init command.
// Returns a new init spf13/cobra command.
func NewInitCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "init",
		RunE:  InitCommandRunE,
		Short: "Creates a default .semverbot.toml config",
	}

	return command
}

// InitCommandRunE runs the init command.
// Returns an error if the command failed.
func InitCommandRunE(cmd *cobra.Command, args []string) (err error) {
	var file *os.File

	if _, err = os.Stat(cli.DefaultConfigFilePath); !os.IsNotExist(err) {
		var prompt = &survey.Confirm{
			Message: "Do you wish to overwrite your current config?",
		}

		var isOk = false

		if err = survey.AskOne(prompt, &isOk); err != nil {
			return err
		}

		if !isOk {
			return
		}
	}

	if file, err = os.Create(cli.DefaultConfigFilePath); err != nil {
		return err
	}

	if _, err = io.WriteString(file, cli.DefaultConfig); err != nil {
		_ = file.Close()
		return err
	}

	return file.Close()
}
