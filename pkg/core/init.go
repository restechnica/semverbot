package core

import (
	"io"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

type InitOptions struct {
	ConfigFilePath string
	Config         string
}

// Init initializes a config file with defaults.
// It will prompt for confirmation before overwriting existing files.
// Returns an error if something went wrong with IO operations or the prompt.
func Init(options *InitOptions) (err error) {
	var file *os.File

	if _, err = os.Stat(options.ConfigFilePath); !os.IsNotExist(err) {
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

	if file, err = os.Create(options.ConfigFilePath); err != nil {
		return err
	}

	if _, err = io.WriteString(file, options.Config); err != nil {
		_ = file.Close()
		return err
	}

	return file.Close()
}
