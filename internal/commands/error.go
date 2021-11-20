package commands

import (
	"fmt"
)

var CommandErrorTemplate = "command %s exited with %s, output: \n%s"

type CommandError struct {
	Arguments []string
	Err       error
	Output    string
}

func NewCommandError(arguments []string, output string, err error) CommandError {
	return CommandError{
		Arguments: arguments,
		Err:       err,
		Output:    output,
	}
}

func (e CommandError) Error() string {
	return fmt.Sprintf(CommandErrorTemplate, e.Arguments, e.Err.Error(), e.Output)
}
