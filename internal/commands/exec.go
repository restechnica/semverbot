package commands

import "os/exec"

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
	var output, err = exec.Command(name, arg...).Output()
	return string(output), err
}

// Run runs a command.
// Returns an error if it failed.
func (c ExecCommander) Run(name string, arg ...string) error {
	return exec.Command(name, arg...).Run()
}
