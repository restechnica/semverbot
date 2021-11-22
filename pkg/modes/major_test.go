package modes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorMode_Increment(t *testing.T) {
	type Test struct {
		Name    string
		Version string
		Want    string
	}

	var tests = []Test{
		{Name: "IncrementMajor", Version: "0.0.0", Want: "1.0.0"},
		{Name: "DiscardPrefix", Version: "v1.0.0", Want: "2.0.0"},
		{Name: "DiscardPrebuild", Version: "2.0.0-pre+001", Want: "3.0.0"},
		{Name: "ResetMinor", Version: "4.5.0", Want: "5.0.0"},
		{Name: "ResetPatch", Version: "3.0.4", Want: "4.0.0"},
		{Name: "ResetPatchAndMinor", Version: "2.1.5", Want: "3.0.0"},
	}

	for _, test := range tests {
		var mode = NewMajorMode()
		var got, err = mode.Increment(test.Version)

		assert.NoError(t, err)
		assert.IsType(t, test.Want, got, `want: "%s, got: "%s"`, test.Want, got)
	}

	t.Run("ReturnErrorOnInvalidVersion", func(t *testing.T) {
		var mode = NewMajorMode()
		var _, got = mode.Increment("invalid")
		assert.Error(t, got)
	})
}

func TestMajorMode_MajorConstant(t *testing.T) {
	t.Run("CheckConstant", func(t *testing.T) {
		var want = "major"
		var got = Major
		assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
	})
}

func TestNewMajorMode(t *testing.T) {
	t.Run("ValidateState", func(t *testing.T) {
		var mode = NewMajorMode()
		assert.NotNil(t, mode)
		assert.IsType(t, MajorMode{}, mode)
	})
}
