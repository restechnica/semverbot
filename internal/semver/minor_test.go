package semver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinorMode_MinorConstant(t *testing.T) {
	t.Run("CheckConstant", func(t *testing.T) {
		var want = "minor"
		var got = Minor

		assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
	})
}

func TestMinorMode_Increment(t *testing.T) {
	type Test struct {
		Name          string
		TargetVersion string
		Want          string
	}

	var tests = []Test{
		{Name: "HappyPath", TargetVersion: "0.1.0", Want: "0.2.0"},
		{Name: "ResetPatch", TargetVersion: "0.2.3", Want: "0.3.0"},
		{Name: "NoResetMajor", TargetVersion: "6.7.0", Want: "6.8.0"},
		{Name: "DiscardPreBuild", TargetVersion: "0.6.0-pre+001", Want: "0.7.0"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var want = test.Want
			var mode = NewMinorMode()
			var got, err = mode.Increment(test.TargetVersion)

			assert.NoError(t, err)
			assert.Equal(t, want, got, `want: %s, got: %s`, want, got)
		})
	}

	type ErrorTest struct {
		Name          string
		TargetVersion string
	}

	var errorTests = []ErrorTest{
		{Name: "ReturnErrorOnInvalidTargetVersion", TargetVersion: "invalid"},
		{Name: "ReturnErrorOnInvalidCharacter", TargetVersion: "v1.2.3"},
	}

	for _, test := range errorTests {
		t.Run(test.Name, func(t *testing.T) {
			var mode = NewMinorMode()
			var _, got = mode.Increment(test.TargetVersion)
			assert.Error(t, got)
		})
	}
}
