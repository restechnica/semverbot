package modes

import (
	"fmt"

	"github.com/restechnica/semverbot/pkg/git"
	"github.com/restechnica/semverbot/pkg/semver"
)

// GitBranch mode name for GitBranchMode.
const GitBranch = "git-branch"

// GitBranchMode implementation of the Mode interface.
// It increments the semver level based on the naming of the source branch of a git merge.
type GitBranchMode struct {
	Delimiters string
	GitAPI     git.API
	SemverMap  semver.Map
}

// NewGitBranchMode creates a new GitBranchMode.
// Returns the new GitBranchMode.
func NewGitBranchMode(delimiters string, semverMap semver.Map) GitBranchMode {
	return GitBranchMode{Delimiters: delimiters, GitAPI: git.NewCLI(), SemverMap: semverMap}
}

// Increment increments the semver level based on the naming of the source branch of a git merge.
// Returns the incremented version or an error if the last git commit is not a merge or if no mode was detected
// based on the branch name.
func (mode GitBranchMode) Increment(prefix string, suffix string, targetVersion string) (nextVersion string, err error) {
	var branchName string
	var matchedMode Mode

	if branchName, err = mode.GitAPI.GetMergedBranchName(); err != nil {
		return
	}

	var isMergeCommit = branchName != ""

	if !isMergeCommit {
		return nextVersion, fmt.Errorf("failed to increment version because the latest git commit is not a merge commit")
	}

	if matchedMode, err = DetectModeFromString(branchName, mode.SemverMap, mode.Delimiters); err != nil {
		return nextVersion, err
	}

	return matchedMode.Increment(prefix, suffix, targetVersion)
}

// String returns a string representation of an instance.
func (mode GitBranchMode) String() string {
	return GitBranch
}
