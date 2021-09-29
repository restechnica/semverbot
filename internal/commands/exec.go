package commands

import (
	"bytes"
	"fmt"
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

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	command.Stdout = &stdout
	command.Stderr = &stderr

	var err = command.Run()

	if err != nil || !isEmpty(&stderr) {
		return stdout.String(), fmt.Errorf("%s: %s", err, stderr.String())
	}

	return stdout.String(), err
}

// Run runs a command.
// Returns an error if it failed.
func (c ExecCommander) Run(name string, arg ...string) error {
	var command = exec.Command(name, arg...)

	var stderr bytes.Buffer
	command.Stderr = &stderr

	var err = command.Run()

	if err != nil || !isEmpty(&stderr) {
		return fmt.Errorf("%s: %s", err, stderr.String())
	}

	return nil
}

func isEmpty(buffer *bytes.Buffer) bool {
	return buffer.Len() == 0
}
