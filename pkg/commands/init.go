package commands

import (
	"io"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

const InitCommandSemverbotConfig = `[git]

[git.config]
email = "semverbot@github.com"
name = "semverbot"

[git.tags]
prefix = "v"

[semver]
mode = "auto"

[semver.detection]
patch = ["fix/"]
minor = ["feature/"]
major = ["release/"]
`

func NewInitCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "init",
		RunE:  InitCommandRunE,
		Short: "Creates a default .semverbot.toml config",
	}

	return command
}

func InitCommandRunE(cmd *cobra.Command, args []string) (err error) {
	var file *os.File
	var path = ".semverbot.toml"

	if _, err = os.Stat(path); !os.IsNotExist(err) {
		var prompt = &survey.Confirm{
			Message: "Do you wish to overwrite your current semverbot config?",
		}

		var isOk = false

		if err = survey.AskOne(prompt, &isOk); err != nil {
			return err
		}

		if !isOk {
			return
		}
	}

	if file, err = os.Create(path); err != nil {
		return err
	}

	if _, err = io.WriteString(file, InitCommandSemverbotConfig); err != nil {
		_ = file.Close()
		return err
	}

	return file.Close()
}
