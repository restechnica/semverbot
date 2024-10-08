package modes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatchMode_Increment(t *testing.T) {
	type Test struct {
		Name    string
		Prefix  string
		Suffix  string
		Version string
		Want    string
	}

	var tests = []Test{
		{Name: "IncrementPatch", Prefix: "v", Version: "0.0.0", Want: "0.0.1"},
		{Name: "DiscardPrefix", Prefix: "v", Version: "v0.0.1", Want: "0.0.2"},
		{Name: "DiscardSuffix", Prefix: "v", Suffix: "a", Version: "v0.0.1a", Want: "0.0.2"},
		{Name: "DiscardSuffixAlt", Prefix: "v", Suffix: "-alt", Version: "v0.0.1-alt", Want: "0.0.2"},
		{Name: "DiscardPrebuild", Prefix: "v", Version: "0.0.2-pre+001", Want: "0.0.3"},
		{Name: "NoResets", Prefix: "v", Version: "3.2.0", Want: "3.2.1"},
	}

	for _, test := range tests {
		var mode = NewPatchMode()
		var got, err = mode.Increment(test.Prefix, test.Suffix, test.Version)

		assert.NoError(t, err)
		assert.IsType(t, test.Want, got, `want: "%s, got: "%s"`, test.Want, got)
	}

	t.Run("ReturnErrorOnInvalidVersion", func(t *testing.T) {
		var mode = NewPatchMode()
		var _, got = mode.Increment("v", "", "invalid")
		assert.Error(t, got)
	})
}

func TestPatchMode_PatchConstant(t *testing.T) {
	t.Run("CheckConstant", func(t *testing.T) {
		var want = "patch"
		var got = Patch
		assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
	})
}

func TestPatchMode_String(t *testing.T) {
	t.Run("ShouldEqualConstant", func(t *testing.T) {
		var mode = NewPatchMode()
		var got = mode.String()
		var want = Patch

		assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
	})
}

func TestNewPatchMode(t *testing.T) {
	t.Run("ValidateState", func(t *testing.T) {
		var mode = NewPatchMode()
		assert.NotNil(t, mode)
		assert.IsType(t, PatchMode{}, mode)
	})
}
