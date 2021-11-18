package semver

import (
	"fmt"

	"github.com/restechnica/semverbot/pkg/git"
)

// GitBranch mode name for GitBranchMode.
const GitBranch = "git-branch"

// GitBranchMode implementation of the Mode interface.
// It increments the semver level based on the naming of the source branch of a git merge.
type GitBranchMode struct {
	GitAPI       git.API
	ModeDetector ModeDetector
}

// NewGitBranchMode creates a new GitBranchMode.
// Returns the new GitBranchMode.
func NewGitBranchMode(detector ModeDetector) GitBranchMode {
	return GitBranchMode{GitAPI: git.NewAPI(), ModeDetector: detector}
}

// Increment increments the semver level based on the naming of the source branch of a git merge.
// Returns the incremented version or an error if the last git commit is not a merge or if no mode was detected
// based on the branch name.
func (mode GitBranchMode) Increment(targetVersion string) (nextVersion string, err error) {
	var branchName string
	var matchedMode Mode

	if branchName, err = mode.GitAPI.GetMergedBranchName(); err != nil {
		return
	}

	var isMergeCommit = branchName != ""

	if !isMergeCommit {
		return nextVersion, fmt.Errorf("failed to increment version because the latest git commit is not a merge commit")
	}

	if matchedMode, err = mode.ModeDetector.DetectMode(branchName); err != nil {
		return nextVersion, err
	}

	return matchedMode.Increment(targetVersion)
}
