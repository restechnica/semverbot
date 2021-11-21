package modes

import (
	"fmt"

	"github.com/restechnica/semverbot/internal/util"
	"github.com/restechnica/semverbot/pkg/git"
)

// GitBranch mode name for GitBranchMode.
const GitBranch = "git-branch"

// GitBranchMode implementation of the Mode interface.
// It increments the semver level based on the naming of the source branch of a git merge.
type GitBranchMode struct {
	Delimiters string
	GitAPI     git.API
	SemverMap  SemverMap
}

// NewGitBranchMode creates a new GitBranchMode.
// Returns the new GitBranchMode.
func NewGitBranchMode(delimiters string, semverMap SemverMap) GitBranchMode {
	return GitBranchMode{Delimiters: delimiters, GitAPI: git.NewCLI(), SemverMap: semverMap}
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

	if matchedMode, err = mode.DetectMode(branchName); err != nil {
		return nextVersion, err
	}

	return matchedMode.Increment(targetVersion)
}

// DetectMode detects the mode (patch, minor, major) based on a git branch name.
// Returns the detected mode.
func (mode GitBranchMode) DetectMode(branchName string) (detected Mode, err error) {
	for key, values := range mode.SemverMap {
		for _, value := range values {
			if mode.isMatch(branchName, value) {
				switch key {
				case Patch:
					return NewPatchMode(), err
				case Minor:
					return NewMinorMode(), err
				case Major:
					return NewMajorMode(), err
				}
			}
		}
	}

	return detected, fmt.Errorf(`failed to detect mode from git branch name "%s" with delimiters "%s"`,
		branchName, mode.Delimiters)
}

// isMatch returns true if a string is part of branch name, after splitting the branch name with delimiters
func (mode GitBranchMode) isMatch(branchName string, value string) bool {
	return util.Contains(branchName, value, mode.Delimiters)
}
