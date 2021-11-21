package modes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPI_SelectMode(t *testing.T) {
	var semverMap = SemverMap{
		Patch: {"fix", "bug"},
		Minor: {"feature"},
		Major: {"release"},
	}

	var gitBranchDelimiters = "/"
	var gitCommitDelimiters = "[]():"

	type Test struct {
		Mode string
		Name string
		Want Mode
	}

	var tests = []Test{
		{Name: "SelectPatchMode", Mode: Patch, Want: NewPatchMode()},
		{Name: "SelectPatchModeIfInvalidMode", Mode: "invalid", Want: NewPatchMode()},
		{Name: "SelectMinorMode", Mode: Minor, Want: NewMinorMode()},
		{Name: "SelectMajorMode", Mode: Major, Want: NewMajorMode()},
		{Name: "SelectAutoMode", Mode: Auto, Want: AutoMode{}},
		{Name: "SelectGitBranchMode", Mode: GitBranch, Want: GitBranchMode{}},
		{Name: "SelectGitCommitMode", Mode: GitCommit, Want: GitCommitMode{}},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var gitBranchMode = NewGitBranchMode(gitBranchDelimiters, semverMap)
			var gitCommitMode = NewGitCommitMode(gitCommitDelimiters, semverMap)

			var modeAPI = NewAPI(gitBranchMode, gitCommitMode)
			var got = modeAPI.SelectMode(test.Mode)

			assert.IsType(t, test.Want, got, `want: "%s, got: "%s"`, test.Want, got)
		})
	}
}

func TestNewAPI(t *testing.T) {
	t.Run("ValidateState", func(t *testing.T) {
		var gitBranchDelimiters = "/"
		var gitCommitDelimiters = "[]"

		var semverMap = SemverMap{}
		var gitBranchMode = NewGitBranchMode(gitBranchDelimiters, semverMap)
		var gitCommitMode = NewGitCommitMode(gitCommitDelimiters, semverMap)
		var modeAPI = NewAPI(gitBranchMode, gitCommitMode)

		assert.NotNil(t, modeAPI)
		assert.NotNil(t, modeAPI.GitBranchMode)
		assert.NotNil(t, modeAPI.GitCommitMode)
	})
}
