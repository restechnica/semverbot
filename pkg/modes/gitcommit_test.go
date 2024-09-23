package modes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/restechnica/semverbot/internal/mocks"
	"github.com/restechnica/semverbot/pkg/semver"
)

func TestGitCommitMode_GitCommitConstant(t *testing.T) {
	t.Run("CheckConstant", func(t *testing.T) {
		var want = "git-commit"
		var got = GitCommit

		assert.Equal(t, want, got, `want: '%s', got: '%s'`, want, got)
	})
}

func TestGitCommitMode_Increment(t *testing.T) {
	var semverMap = semver.Map{
		Patch: {"fix", "bug"},
		Minor: {"feature"},
		Major: {"release"},
	}

	type Test struct {
		CommitMessage string
		Delimiters    string
		Name          string
		Prefix        string
		Suffix        string
		SemverMap     semver.Map
		Version       string
		Want          string
	}

	var tests = []Test{
		{Name: "IncrementPatch", CommitMessage: "fix] some-bug", Delimiters: "[]", Prefix: "v", Suffix: "", SemverMap: semverMap, Version: "0.0.0", Want: "0.0.1"},
		{Name: "IncrementPatch", CommitMessage: "[fi] some/bug", Delimiters: "/", Prefix: "v", Suffix: "", SemverMap: semverMap, Version: "0.0.0", Want: "0.0.1"},
		{Name: "IncrementMinor", CommitMessage: "[feature] some-feat", Delimiters: "[]", Prefix: "v", Suffix: "", SemverMap: semverMap, Version: "0.0.1", Want: "0.1.0"},
		{Name: "IncrementMajor", CommitMessage: "[release] some-release", Delimiters: "[]", Prefix: "v", Suffix: "", SemverMap: semverMap, Version: "0.1.0", Want: "1.0.0"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var gitAPI = mocks.NewMockGitAPI()
			gitAPI.On("GetLatestCommitMessage").Return(test.CommitMessage, nil)

			var mode = NewGitCommitMode(test.Delimiters, test.SemverMap)
			mode.GitAPI = gitAPI

			var got, err = mode.Increment(test.Prefix, test.Suffix, test.Version)

			assert.NoError(t, err)
			assert.IsType(t, test.Want, got, `want: '%s, got: '%s'`, test.Want, got)
		})
	}

	t.Run("ReturnErrorOnGitAPIError", func(t *testing.T) {
		var want = fmt.Errorf("some-error")

		var gitAPI = mocks.NewMockGitAPI()
		gitAPI.On("GetLatestCommitMessage").Return("", want)

		var mode = NewGitCommitMode("[]", semverMap)
		mode.GitAPI = gitAPI

		var _, got = mode.Increment("v", "", "0.0.0")

		assert.Error(t, got)
		assert.Equal(t, want, got, `want: '%s, got: '%s'`, want, got)
	})

	t.Run("ReturnErrorIfNoMatchingMode", func(t *testing.T) {
		var gitAPI = mocks.NewMockGitAPI()
		gitAPI.On("GetLatestCommitMessage").Return("nomatch/some-feature", nil)

		var mode = NewGitCommitMode("/", semverMap)
		mode.GitAPI = gitAPI

		var _, got = mode.Increment("v", "", "0.0.0")

		assert.Error(t, got)
	})

	t.Run("ReturnErrorIfInvalidVersion", func(t *testing.T) {
		var gitAPI = mocks.NewMockGitAPI()
		gitAPI.On("GetLatestCommitMessage").Return("[feature]some-feature", nil)

		var mode = NewGitCommitMode("[]", semverMap)
		mode.GitAPI = gitAPI

		var _, got = mode.Increment("v", "", "invalid")

		assert.Error(t, got)
	})
}

func TestGitCommitMode_String(t *testing.T) {
	t.Run("ShouldEqualConstant", func(t *testing.T) {
		var mode = NewGitCommitMode("", semver.Map{})
		var got = mode.String()
		var want = GitCommit

		assert.Equal(t, want, got, `want: '%s, got: '%s'`, want, got)
	})
}

func TestNewGitCommitMode(t *testing.T) {
	t.Run("ValidateState", func(t *testing.T) {
		var delimiters = "[]"
		var semverMap = semver.Map{}
		var mode = NewGitCommitMode(delimiters, semverMap)

		assert.NotNil(t, mode)
		assert.NotEmpty(t, mode.Delimiters)
		assert.NotNil(t, mode.SemverMap)
	})
}
