package modes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatchMode_PatchConstant(t *testing.T) {
	t.Run("CheckConstant", func(t *testing.T) {
		var want = "patch"
		var got = Patch

		assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
	})
}

func TestPatchMode_Increment(t *testing.T) {
	type Test struct {
		Name          string
		TargetVersion string
		Want          string
	}

	var tests = []Test{
		{Name: "HappyPath", TargetVersion: "0.0.1", Want: "0.0.2"},
		{Name: "DiscardPreBuild", TargetVersion: "0.0.8-pre+001", Want: "0.0.9"},
		{Name: "NoResetMajorMinor", TargetVersion: "5.4.3", Want: "5.4.4"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var want = test.Want
			var mode = NewPatchMode()
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
		//{Name: "ReturnErrorOnInvalidCharacter", TargetVersion: "v1.2.3"},
	}

	for _, test := range errorTests {
		t.Run(test.Name, func(t *testing.T) {
			var mode = NewPatchMode()
			var _, got = mode.Increment(test.TargetVersion)
			assert.Error(t, got)
		})
	}
}
