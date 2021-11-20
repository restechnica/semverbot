package modes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/restechnica/semverbot/internal/mocks"
	"github.com/restechnica/semverbot/pkg/git"
)

var semverMap = SemverMap{
	Patch: {"fix"},
	Minor: {"feature"},
	Major: {"release"},
}

func TestGitCommitMode_GitCommitConstant(t *testing.T) {
	t.Run("CheckConstant", func(t *testing.T) {
		var want = "git-commit"
		var got = GitCommit

		assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
	})
}

// TODO move to ModeDetector
//func TestGitCommitMode_GetMatchedMode(t *testing.T) {
//	type Test struct {
//		Message string
//		Name    string
//		Want    Mode
//	}
//
//	var tests = []Test{
//		{Name: "GetPatchModeWithBrackets", Message: "[fix] some message", Want: NewPatchMode()},
//		{Name: "GetPatchModeWithTrailingSlash", Message: "Merged: repo/fix/some-error", Want: NewPatchMode()},
//		{Name: "GetMinorModeWithBrackets", Message: "some [feature] message", Want: NewMinorMode()},
//		{Name: "GetMinorModeWithTrailingSlash", Message: "Merged: repo/feature/some-error", Want: NewMinorMode()},
//		{Name: "GetMajorModeWithBrackets", Message: "some message [release]", Want: NewMajorMode()},
//		{Name: "GetMajorModeWithTrailingSlash", Message: "Merged: repo/release/some-error", Want: NewMajorMode()},
//	}
//
//	for _, test := range tests {
//		t.Run(test.Name, func(t *testing.T) {
//			var want = test.Want
//
//			var gitCommitMode = NewGitCommitMode(NewModeDetector(semverMap))
//			var got, err = gitCommitMode.ModeDetector.(test.Message)
//
//			assert.NoError(t, err)
//			assert.IsType(t, want, got, `want: "%s", got: "%s"`, want, got)
//		})
//	}
//
//	type ErrorTest struct {
//		Message string
//		Name    string
//	}
//
//	var errorTests = []ErrorTest{
//		{Name: "ReturnErrorOnUnmatchedMode", Message: "[fix some message"},
//	}
//
//	for _, test := range errorTests {
//		t.Run(test.Name, func(t *testing.T) {
//			var want = fmt.Sprintf(`could not match a mode to the commit message "%s"`, test.Message)
//
//			var gitCommitMode = NewGitCommitMode(semverMap)
//			var _, err = gitCommitMode.GetMatchedMode(test.Message)
//
//			assert.Error(t, err)
//			assert.Equal(t, err.Error(), want, `want: "%s", got: "%s"`, want, err.Error())
//		})
//	}
//}

func TestGitCommitMode_Increment(t *testing.T) {
	type Test struct {
		Message string
		Name    string
		Version string
		Want    string
	}

	var tests = []Test{
		{Name: "IncrementPatchWithBrackets", Message: "[fix] some message", Version: "0.0.0", Want: "0.0.1"},
		//{Name: "IncrementPatchWithTrailingSlash", Message: "Merged: repo/fix/some-error", Version: "0.0.1", Want: "0.0.2"},
		{Name: "IncrementMinorWithBrackets", Message: "some [feature] message", Version: "0.0.0", Want: "0.1.0"},
		//{Name: "IncrementMinorWithTrailingSlash", Message: "Merged: repo/feature/some-error", Version: "0.1.0", Want: "0.2.0"},
		{Name: "IncrementMajorWithBrackets", Message: "some message [release]", Version: "0.0.0", Want: "1.0.0"},
		//{Name: "IncrementMajorWithTrailingSlash", Message: "Merged: repo/release/some-error", Version: "1.0.0", Want: "2.0.0"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var want = test.Want

			var cmder = mocks.NewMockCommander()
			cmder.On("Output", mock.Anything, mock.Anything).Return(test.Message, nil)

			var gitCommitMode = NewGitCommitMode("[]", semverMap)
			gitCommitMode.GitAPI = git.API{Commander: cmder}
			var got, err = gitCommitMode.Increment(test.Version)

			assert.NoError(t, err)
			assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
		})
	}

	type ErrorTest struct {
		GitError error
		Message  string
		Name     string
		Version  string
	}

	var errorTests = []ErrorTest{
		{Name: "ReturnErrorOnUnmatchedMode", Message: "[fix some message", Version: "0.0.0", GitError: nil},
		{Name: "ReturnErrorOnInvalidVersion", Message: "[fix] some message", Version: "invalid", GitError: nil},
		//{Name: "ReturnErrorOnInvalidCharacter", Message: "[fix] some message", Version: "v1.0.0", GitError: nil}, // commented out due to TolerantParse
		{Name: "ReturnErrorOnGitError", Message: "[fix] some message", Version: "1.0.0",
			GitError: fmt.Errorf("some-error")},
	}

	for _, test := range errorTests {
		t.Run(test.Name, func(t *testing.T) {
			var cmder = mocks.NewMockCommander()
			cmder.On("Output", mock.Anything, mock.Anything).Return(test.Message, test.GitError)

			var gitCommitMode = NewGitCommitMode("[]", semverMap)
			gitCommitMode.GitAPI = git.API{Commander: cmder}
			var _, err = gitCommitMode.Increment(test.Version)

			assert.Error(t, err)
		})
	}
}
