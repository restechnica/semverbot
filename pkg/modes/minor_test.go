package modes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinorMode_Increment(t *testing.T) {
	type Test struct {
		Name    string
		Prefix  string
		Suffix  string
		Version string
		Want    string
	}

	var tests = []Test{
		{Name: "IncrementMinor", Prefix: "v", Version: "0.0.0", Want: "0.1.0"},
		{Name: "DiscardPrefix", Prefix: "v", Version: "v0.1.0", Want: "0.2.0"},
		{Name: "DiscardSuffix", Prefix: "v", Suffix: "a", Version: "0.1.0a", Want: "0.2.0"},
		{Name: "DiscardSuffixAlt", Prefix: "v", Suffix: "-alt", Version: "0.1.0-alt", Want: "0.2.0"},
		{Name: "DiscardPrebuild", Prefix: "v", Version: "0.2.0-pre+001", Want: "0.3.0"},
		{Name: "ResetPatch", Prefix: "v", Version: "3.0.4", Want: "3.1.0"},
	}

	for _, test := range tests {
		var mode = NewMinorMode()
		var got, err = mode.Increment(test.Prefix, test.Suffix, test.Version)

		assert.NoError(t, err)
		assert.IsType(t, test.Want, got, `want: "%s, got: "%s"`, test.Want, got)
	}

	t.Run("ReturnErrorOnInvalidVersion", func(t *testing.T) {
		var mode = NewMinorMode()
		var _, got = mode.Increment("v", "", "invalid")
		assert.Error(t, got)
	})
}

func TestMinorMode_MinorConstant(t *testing.T) {
	t.Run("CheckConstant", func(t *testing.T) {
		var want = "minor"
		var got = Minor
		assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
	})
}

func TestMinorMode_String(t *testing.T) {
	t.Run("ShouldEqualConstant", func(t *testing.T) {
		var mode = NewMinorMode()
		var got = mode.String()
		var want = Minor

		assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
	})
}

func TestNewMinorMode(t *testing.T) {
	t.Run("ValidateState", func(t *testing.T) {
		var mode = NewMinorMode()
		assert.NotNil(t, mode)
		assert.IsType(t, MinorMode{}, mode)
	})
}
