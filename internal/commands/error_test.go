package commands

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandError_Error(t *testing.T) {
	t.Run("ValidateMessage", func(t *testing.T) {
		var args = []string{"some", "command", "arguments"}
		var err = fmt.Errorf("some-error")
		var output = "some-output"

		var want = fmt.Sprintf(CommandErrorTemplate, args, err, output)

		var cmdError = CommandError{Arguments: args, Err: err, Output: output}.Error()
		assert.Equal(t, want, cmdError, `want: "%s", got: "%s"`, want, cmdError)
	})
}

func TestNewCommandError(t *testing.T) {
	t.Run("ValidateState", func(t *testing.T) {
		var args = []string{"some", "command", "arguments"}
		var err = fmt.Errorf("some-error")
		var output = "some-output"

		var cmdError = NewCommandError(args, output, err)

		assert.Equal(t, args, cmdError.Arguments, `want: "%s", got: "%s"`, args, cmdError.Arguments)
		assert.Equal(t, err, cmdError.Err, `want: "%s", got: "%s"`, err, cmdError.Err)
		assert.Equal(t, output, cmdError.Output, `want: "%s", got: "%s"`, output, cmdError.Output)
	})
}
