package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecCommander_Output(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		var want = "hello world"

		var commander = NewExecCommander()
		var got, err = commander.Output("echo", "-n", want)

		assert.NoError(t, err)
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})

	t.Run("ReturnErrorOnError", func(t *testing.T) {
		var fakeCommand = "lskdf"
		var commander = NewExecCommander()
		var _, err = commander.Output(fakeCommand)
		assert.Error(t, err)
	})
}

func TestExecCommander_Run(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		var commander = NewExecCommander()
		var err = commander.Run("echo")
		assert.NoError(t, err)
	})

	t.Run("ReturnErrorOnError", func(t *testing.T) {
		var fakeCommand = "lskdf"
		var commander = NewExecCommander()
		var err = commander.Run(fakeCommand)
		assert.Error(t, err)
	})
}
