package modes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

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
	type Test struct {
		Modes   []Mode
		Name    string
		Prefix  string
		Version string
		Want    string
	}

	var mockMode = mocks.NewMockMode()
	mockMode.On("Increment", mock.Anything, mock.Anything).Return("", fmt.Errorf("some-error"))

	var tests = []Test{
		{Name: "IncrementMajor", Prefix: "v", Modes: []Mode{NewMajorMode()}, Version: "0.0.0", Want: "1.0.0"},
		{Name: "IncrementMinor", Prefix: "v", Modes: []Mode{NewMinorMode()}, Version: "0.0.0", Want: "0.1.0"},
		{Name: "IncrementPatch", Prefix: "v", Modes: []Mode{NewPatchMode()}, Version: "0.0.0", Want: "0.0.1"},
		{Name: "DefaultToPatchIfModeSliceEmpty", Prefix: "v", Modes: []Mode{}, Version: "0.0.0", Want: "0.0.1"},
		{Name: "IncrementWithSecondModeAfterFirstFailed", Prefix: "v", Modes: []Mode{mockMode, NewMinorMode()}, Version: "0.0.0", Want: "0.1.0"},
	}

	for _, test := range tests {
		var mode = NewAutoMode(test.Modes)
		var got, err = mode.Increment(test.Prefix, test.Version)

		assert.NoError(t, err)
		assert.IsType(t, test.Want, got, `want: "%s, got: "%s"`, test.Want, got)
	}
}

func TestAutoMode_String(t *testing.T) {
	t.Run("ShouldEqualConstant", func(t *testing.T) {
		var mode = NewAutoMode([]Mode{})
		var got = mode.String()
		var want = Auto

		assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
	})
}

func TestNewAutoMode(t *testing.T) {
	t.Run("ValidateState", func(t *testing.T) {
		var mockMode = mocks.NewMockMode()
		var modes = []Mode{mockMode, mockMode, mockMode}
		var mode = NewAutoMode(modes)
		assert.NotNil(t, mode)
		assert.NotEmpty(t, mode.Modes)
	})
}
