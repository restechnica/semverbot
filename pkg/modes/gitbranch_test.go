package modes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/restechnica/semverbot/internal/mocks"
	"github.com/restechnica/semverbot/pkg/semver"
)

func TestGitBranchMode_GitBranchConstant(t *testing.T) {
	t.Run("CheckConstant", func(t *testing.T) {
		var want = "git-branch"
		var got = GitBranch

		assert.Equal(t, want, got, `want: '%s', got: '%s'`, want, got)
	})
}

func TestGitBranchMode_DetectMode(t *testing.T) {
	var semverMap = semver.Map{
		Patch: {"fix", "bug"},
		Minor: {"feature"},
		Major: {"release"},
	}

	type Test struct {
		BranchName string
		Delimiters string
		Name       string
		SemverMap  semver.Map
		Want       Mode
	}

	var tests = []Test{
		{Name: "DetectPatchMode", BranchName: "fix/some-bug", Delimiters: "/", SemverMap: semverMap, Want: NewPatchMode()},
		{Name: "DetectMinorMode", BranchName: "feature/some-bug", Delimiters: "/", SemverMap: semverMap, Want: NewMinorMode()},
		{Name: "DetectMajorMode", BranchName: "release/some-bug", Delimiters: "/", SemverMap: semverMap, Want: NewMajorMode()},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var mode = NewGitBranchMode(test.Delimiters, test.SemverMap)
			var got, err = mode.DetectMode(test.BranchName)

			assert.NoError(t, err)
			assert.IsType(t, test.Want, got, `want: '%s, got: '%s'`, test.Want, got)
		})
	}

	type ErrorTest struct {
		BranchName string
		Delimiters string
		Error      error
		Name       string
		SemverMap  semver.Map
	}

	var errorTests = []ErrorTest{
		{
			Name:       "DetectNothingWithEmptySemverMap",
			BranchName: "feature/some-feature",
			Delimiters: "/",
			Error:      fmt.Errorf(`failed to detect mode from git branch name 'feature/some-feature' with delimiters '/'`),
			SemverMap:  semver.Map{},
		},
		{
			Name:       "DetectNothingWithEmptyDelimiters",
			BranchName: "feature/some-feature",
			Delimiters: "",
			Error:      fmt.Errorf(`failed to detect mode from git branch name 'feature/some-feature' with delimiters ''`),
			SemverMap:  semverMap,
		},
		{
			Name:       "DetectNothingWithEmptyBranchName",
			BranchName: "",
			Delimiters: "/",
			Error:      fmt.Errorf(`failed to detect mode from git branch name '' with delimiters '/'`),
			SemverMap:  semverMap,
		},
		{
			Name:       "DetectNothingWithFaultySemverMap",
			BranchName: "feature/some-feature",
			Delimiters: "/",
			Error:      fmt.Errorf(`failed to detect mode from git branch name 'feature/some-feature' with delimiters '/'`),
			SemverMap: semver.Map{
				"mnr": {"feature"},
			},
		},
	}

	for _, test := range errorTests {
		t.Run(test.Name, func(t *testing.T) {
			var mode = NewGitBranchMode(test.Delimiters, test.SemverMap)
			var _, got = mode.DetectMode(test.BranchName)

			assert.Error(t, got)
			assert.Equal(t, test.Error, got, `want: '%s, got: '%s'`, test.Error, got)
		})
	}
}

func TestGitBranchMode_Increment(t *testing.T) {
	var semverMap = semver.Map{
		Patch: {"fix", "bug"},
		Minor: {"feature"},
		Major: {"release"},
	}

	type Test struct {
		BranchName string
		Delimiters string
		Name       string
		SemverMap  semver.Map
		Version    string
		Want       string
	}

	var tests = []Test{
		{Name: "IncrementPatch", BranchName: "fix/some-bug", Delimiters: "/", SemverMap: semverMap, Version: "0.0.0", Want: "0.0.1"},
		{Name: "IncrementMinor", BranchName: "feature/some-bug", Delimiters: "/", SemverMap: semverMap, Version: "0.0.1", Want: "0.1.0"},
		{Name: "IncrementMajor", BranchName: "release/some-bug", Delimiters: "/", SemverMap: semverMap, Version: "0.1.0", Want: "1.0.0"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var gitAPI = mocks.NewMockGitAPI()
			gitAPI.On("GetMergedBranchName").Return(test.BranchName, nil)

			var mode = NewGitBranchMode(test.Delimiters, test.SemverMap)
			mode.GitAPI = gitAPI

			var got, err = mode.Increment(test.Version)

			assert.NoError(t, err)
			assert.IsType(t, test.Want, got, `want: '%s, got: '%s'`, test.Want, got)
		})
	}

	t.Run("ReturnErrorOnGitAPIError", func(t *testing.T) {
		var want = fmt.Errorf("some-error")

		var gitAPI = mocks.NewMockGitAPI()
		gitAPI.On("GetMergedBranchName").Return("", want)

		var mode = NewGitBranchMode("/", semverMap)
		mode.GitAPI = gitAPI

		var _, got = mode.Increment("0.0.0")

		assert.Error(t, got)
		assert.Equal(t, want, got, `want: '%s, got: '%s'`, want, got)
	})

	t.Run("ReturnErrorIfNoMergeCommit", func(t *testing.T) {
		var want = fmt.Errorf("failed to increment version because the latest git commit is not a merge commit")

		var gitAPI = mocks.NewMockGitAPI()
		gitAPI.On("GetMergedBranchName").Return("", nil)

		var mode = NewGitBranchMode("/", semverMap)
		mode.GitAPI = gitAPI

		var _, got = mode.Increment("0.0.0")

		assert.Error(t, got)
		assert.Equal(t, want, got, `want: '%s, got: '%s'`, want, got)
	})

	t.Run("ReturnErrorIfNoMatchingMode", func(t *testing.T) {
		var gitAPI = mocks.NewMockGitAPI()
		gitAPI.On("GetMergedBranchName").Return("feat/some-feature", nil)

		var mode = NewGitBranchMode("/", semverMap)
		mode.GitAPI = gitAPI

		var _, got = mode.Increment("0.0.0")

		assert.Error(t, got)
	})

	t.Run("ReturnErrorIfInvalidVersion", func(t *testing.T) {
		var gitAPI = mocks.NewMockGitAPI()
		gitAPI.On("GetMergedBranchName").Return("feature/some-feat", nil)

		var mode = NewGitBranchMode("/", semverMap)
		mode.GitAPI = gitAPI

		var _, got = mode.Increment("invalid")

		assert.Error(t, got)
	})
}

func TestGitBranchMode_String(t *testing.T) {
	t.Run("ShouldEqualConstant", func(t *testing.T) {
		var mode = NewGitBranchMode("", semver.Map{})
		var got = mode.String()
		var want = GitBranch

		assert.Equal(t, want, got, `want: '%s, got: '%s'`, want, got)
	})
}

func TestNewGitBranchMode(t *testing.T) {
	t.Run("ValidateState", func(t *testing.T) {
		var delimiters = "/"
		var semverMap = semver.Map{}
		var mode = NewGitBranchMode(delimiters, semverMap)

		assert.NotNil(t, mode)
		assert.NotEmpty(t, mode.Delimiters)
		assert.NotNil(t, mode.SemverMap)
	})
}
