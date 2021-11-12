package commands

import (
	"fmt"
)

type CommandError struct {
	Arguments []string
	Output    string
	Err       error
}

func (e CommandError) Error() string {
	return fmt.Sprintf("command %s exited with %s, output: \n%s", e.Arguments, e.Err.Error(), e.Output)
}
