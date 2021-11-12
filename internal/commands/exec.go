package commands

import (
	"bytes"
	"os/exec"
)

// ExecCommander implementation of the Commander interface.
// It makes use of exec.Command to run commands.
type ExecCommander struct{}

// NewExecCommander creates a new ExecCommander.
// Returns the new ExecCommander.
func NewExecCommander() *ExecCommander {
	return &ExecCommander{}
}

// Output runs a command.
// Returns the output of the command or an error if it failed.
func (c ExecCommander) Output(name string, arg ...string) (string, error) {
	var command = exec.Command(name, arg...)
	var buffer bytes.Buffer

	command.Stdout = &buffer
	command.Stderr = &buffer

	var err = command.Run()
	var output = buffer.String()

	if err != nil {
		return "", CommandError{Arguments: command.Args, Err: err, Output: output}
	}

	return output, err
}

// Run runs a command.
// Returns an error if it failed.
func (c ExecCommander) Run(name string, arg ...string) error {
	var _, err = c.Output(name, arg...)
	return err
}
