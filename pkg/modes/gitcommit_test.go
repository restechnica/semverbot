package modes

import (
	"fmt"
	"github.com/restechnica/semverbot/internal/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitCommitMode_GitCommitConstant(t *testing.T) {
	t.Run("CheckConstant", func(t *testing.T) {
		var want = "git-commit"
		var got = GitCommit

		assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
	})
}

func TestGitCommitMode_DetectMode(t *testing.T) {
	var semverMap = SemverMap{
		Patch: {"fix", "bug"},
		Minor: {"feature", "feat"},
		Major: {"release"},
	}

	type Test struct {
		CommitMessage string
		Delimiters    string
		Name          string
		SemverMap     SemverMap
		Want          Mode
	}

	var tests = []Test{
		{Name: "DetectPatchMode", CommitMessage: "[bug] some fix", Delimiters: "[]", SemverMap: semverMap, Want: NewPatchMode()},
		{Name: "DetectPatchMode", CommitMessage: "[fix] some bug", Delimiters: "[]", SemverMap: semverMap, Want: NewPatchMode()},
		{Name: "DetectMinorMode", CommitMessage: "feat(subject): some changes", Delimiters: "():", SemverMap: semverMap, Want: NewMinorMode()},
		{Name: "DetectMinorMode", CommitMessage: "[feature] some changes", Delimiters: "[]", SemverMap: semverMap, Want: NewMinorMode()},
		{Name: "DetectMajorMode", CommitMessage: "release/some-bug", Delimiters: "/", SemverMap: semverMap, Want: NewMajorMode()},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var mode = NewGitBranchMode(test.Delimiters, test.SemverMap)
			var got, err = mode.DetectMode(test.CommitMessage)

			assert.NoError(t, err)
			assert.IsType(t, test.Want, got, `want: "%s, got: "%s"`, test.Want, got)
		})
	}

	type ErrorTest struct {
		CommitMessage string
		Delimiters    string
		Error         error
		Name          string
		SemverMap     SemverMap
	}

	var errorTests = []ErrorTest{
		{
			Name:          "DetectNothingWithEmptySemverMap",
			CommitMessage: "[feature] some changes",
			Delimiters:    "[]",
			Error:         fmt.Errorf(`failed to detect mode from git branch name "[feature] some changes" with delimiters "[]"`),
			SemverMap:     SemverMap{},
		},
		{
			Name:          "DetectNothingWithEmptyDelimiters",
			CommitMessage: "[feature] some changes",
			Delimiters:    "",
			Error:         fmt.Errorf(`failed to detect mode from git branch name "[feature] some changes" with delimiters ""`),
			SemverMap:     semverMap,
		},
		{
			Name:          "DetectNothingWithEmptyCommitMessage",
			CommitMessage: "",
			Delimiters:    "/",
			Error:         fmt.Errorf(`failed to detect mode from git branch name "" with delimiters "/"`),
			SemverMap:     semverMap,
		},
		{
			Name:          "DetectNothingWithFaultySemverMap",
			CommitMessage: "[feature] some changes",
			Delimiters:    "[]",
			Error:         fmt.Errorf(`failed to detect mode from git branch name "[feature] some changes" with delimiters "[]"`),
			SemverMap: SemverMap{
				"mnr": {"feature"},
			},
		},
	}

	for _, test := range errorTests {
		t.Run(test.Name, func(t *testing.T) {
			var mode = NewGitBranchMode(test.Delimiters, test.SemverMap)
			var _, got = mode.DetectMode(test.CommitMessage)

			assert.Error(t, got)
			assert.Equal(t, test.Error, got, `want: "%s, got: "%s"`, test.Error, got)
		})
	}
}

func TestGitCommitMode_Increment(t *testing.T) {
	var semverMap = SemverMap{
		Patch: {"fix", "bug"},
		Minor: {"feature"},
		Major: {"release"},
	}

	type Test struct {
		CommitMessage string
		Delimiters    string
		Name          string
		SemverMap     SemverMap
		Version       string
		Want          string
	}

	var tests = []Test{
		{Name: "IncrementPatch", CommitMessage: "[fix] some-bug", Delimiters: "[]", SemverMap: semverMap, Version: "0.0.0", Want: "0.0.1"},
		{Name: "IncrementPatch", CommitMessage: "[fi] some/bug", Delimiters: "/", SemverMap: semverMap, Version: "0.0.0", Want: "0.0.1"},
		{Name: "IncrementMinor", CommitMessage: "[feature] some-feat", Delimiters: "[]", SemverMap: semverMap, Version: "0.0.1", Want: "0.1.0"},
		{Name: "IncrementMajor", CommitMessage: "[release] some-release", Delimiters: "[]", SemverMap: semverMap, Version: "0.1.0", Want: "1.0.0"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var gitAPI = mocks.NewMockGitAPI()
			gitAPI.On("GetLatestCommitMessage").Return(test.CommitMessage, nil)

			var mode = NewGitCommitMode(test.Delimiters, test.SemverMap)
			mode.GitAPI = gitAPI

			var got, err = mode.Increment(test.Version)

			assert.NoError(t, err)
			assert.IsType(t, test.Want, got, `want: "%s, got: "%s"`, test.Want, got)
		})
	}

	t.Run("ReturnErrorOnGitAPIError", func(t *testing.T) {
		var want = fmt.Errorf("some-error")

		var gitAPI = mocks.NewMockGitAPI()
		gitAPI.On("GetLatestCommitMessage").Return("", want)

		var mode = NewGitCommitMode("[]", semverMap)
		mode.GitAPI = gitAPI

		var _, got = mode.Increment("0.0.0")

		assert.Error(t, got)
		assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
	})

	t.Run("ReturnErrorIfNoMatchingMode", func(t *testing.T) {
		var gitAPI = mocks.NewMockGitAPI()
		gitAPI.On("GetLatestCommitMessage").Return("nomatch/some-feature", nil)

		var mode = NewGitCommitMode("/", semverMap)
		mode.GitAPI = gitAPI

		var _, got = mode.Increment("0.0.0")

		assert.Error(t, got)
	})

	t.Run("ReturnErrorIfInvalidVersion", func(t *testing.T) {
		var gitAPI = mocks.NewMockGitAPI()
		gitAPI.On("GetLatestCommitMessage").Return("[feature]some-feature", nil)

		var mode = NewGitCommitMode("[]", semverMap)
		mode.GitAPI = gitAPI

		var _, got = mode.Increment("invalid")

		assert.Error(t, got)
	})
}

func TestNewGitCommitMode(t *testing.T) {
	t.Run("ValidateState", func(t *testing.T) {
		var delimiters = "[]"
		var semverMap = SemverMap{}
		var mode = NewGitCommitMode(delimiters, semverMap)

		assert.NotNil(t, mode)
		assert.NotEmpty(t, mode.Delimiters)
		assert.NotNil(t, mode.SemverMap)
	})
}
