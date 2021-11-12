package commands

import (
	"fmt"
	"os/exec"
)

type CommandError struct {
	Command *exec.Cmd
	Output  string
	Err     error
}

func (e CommandError) Error() string {
	return fmt.Sprintf("command %s exited with %s, output: \n%s", e.Command.Args, e.Err.Error(), e.Output)
}
