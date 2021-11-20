package modes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/restechnica/semverbot/internal/mocks"
)

func TestAutoMode_AutoConstant(t *testing.T) {
	t.Run("CheckConstant", func(t *testing.T) {
		var want = "auto"
		var got = Auto

		assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
	})
}

func TestAutoMode_Increment(t *testing.T) {
	t.Run("IncrementWithGitCommitMode", func(t *testing.T) {
		const target = "0.0.0"
		const want = "1.0.0"

		var gitCommitMode = mocks.NewMockSemverMode()
		gitCommitMode.On("Increment", target).Return(want, nil)

		var autoMode = NewAutoMode([]Mode{gitCommitMode})
		var got, err = autoMode.Increment(target)

		assert.NoError(t, err)
		assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
	})

	t.Run("IncrementWithPatchMode", func(t *testing.T) {
		const target = "0.0.0"
		const want = "0.0.1"

		var gitCommitMode = mocks.NewMockSemverMode()
		gitCommitMode.On("Increment", target).Return("", fmt.Errorf("some-error"))

		var autoMode = NewAutoMode([]Mode{gitCommitMode})
		var got, err = autoMode.Increment(target)

		assert.NoError(t, err)
		assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
	})

	type ErrorTest struct {
		Name    string
		Version string
	}

	var errorTests = []ErrorTest{
		{Name: "ReturnErrorOnInvalidVersion", Version: "invalid"},
		//{Name: "ReturnErrorOnInvalidCharacter", Version: "v1.0.0"}, // commented due to TolerantParse
	}

	for _, test := range errorTests {
		t.Run(test.Name, func(t *testing.T) {
			var gitCommitMode = mocks.NewMockSemverMode()
			gitCommitMode.On("Increment", test.Version).Return("", fmt.Errorf("some-error"))

			var autoMode = NewAutoMode([]Mode{gitCommitMode})
			var _, err = autoMode.Increment(test.Version)

			assert.Error(t, err)
		})
	}
}
