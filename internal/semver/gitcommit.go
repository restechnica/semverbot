package semver

import (
	"github.com/restechnica/semverbot/internal/commands"
)

// GitCommit mode name for GitCommitMode.
const GitCommit = "git-commit"

// GitCommitMode implementation of the Mode interface.
// It increments the semver level based on the latest git commit messages.
type GitCommitMode struct {
	Commander    commands.Commander
	ModeDetector ModeDetector
}

// NewGitCommitMode creates a new GitCommitMode.
// Returns the new GitCommitMode.
func NewGitCommitMode(detector ModeDetector) GitCommitMode {
	return GitCommitMode{Commander: commands.ExecCommander{}, ModeDetector: detector}
}

// Increment increments a given version based on the latest git commit message.
// Returns the incremented version or an error if it failed to detect the mode based on the git commit.
func (mode GitCommitMode) Increment(targetVersion string) (nextVersion string, err error) {
	var message string
	var matchedMode Mode

	if message, err = mode.Commander.Output("git", "show", "-s", "--format=%s"); err != nil {
		return
	}

	if matchedMode, err = mode.ModeDetector.DetectMode(message); err != nil {
		return
	}

	return matchedMode.Increment(targetVersion)
}
