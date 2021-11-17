package semver

import (
	"fmt"

	"github.com/restechnica/semverbot/internal/commands"
)

// GitBranch mode name for GitBranchMode.
const GitBranch = "git-branch"

// GitBranchMode implementation of the Mode interface.
// It increments the semver level based on the naming of the source branch of a git merge.
type GitBranchMode struct {
	Commander    commands.Commander
	ModeDetector ModeDetector
}

// NewGitBranchMode creates a new GitBranchMode.
// Returns the new GitBranchMode.
func NewGitBranchMode(detector ModeDetector) GitBranchMode {
	return GitBranchMode{Commander: commands.ExecCommander{}, ModeDetector: detector}
}

// Increment increments the semver level based on the naming of the source branch of a git merge.
// Returns the incremented version or an error if the last git commit is not a merge or if no mode was detected
// based on the branch name.
func (mode GitBranchMode) Increment(targetVersion string) (nextVersion string, err error) {
	var branchName string
	var matchedMode Mode

	if branchName, err = mode.Commander.Output("git", "name-rev", "--name-only", "--refs=refs/heads/*",
		"--refs=refs/remotes/*", "HEAD^2"); err != nil {
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
