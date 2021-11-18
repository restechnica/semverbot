package semver

import (
	"github.com/restechnica/semverbot/pkg/git"
)

// GitCommit mode name for GitCommitMode.
const GitCommit = "git-commit"

// GitCommitMode implementation of the Mode interface.
// It increments the semver level based on the latest git commit messages.
type GitCommitMode struct {
	GitAPI       git.API
	ModeDetector ModeDetector
}

// NewGitCommitMode creates a new GitCommitMode.
// Returns the new GitCommitMode.
func NewGitCommitMode(detector ModeDetector) GitCommitMode {
	return GitCommitMode{GitAPI: git.NewAPI(), ModeDetector: detector}
}

// Increment increments a given version based on the latest git commit message.
// Returns the incremented version or an error if it failed to detect the mode based on the git commit.
func (mode GitCommitMode) Increment(targetVersion string) (nextVersion string, err error) {
	var message string
	var detectedMode Mode

	if message, err = mode.GitAPI.GetLatestCommitMessage(); err != nil {
		return
	}

	if detectedMode, err = mode.ModeDetector.DetectMode(message); err != nil {
		return
	}

	return detectedMode.Increment(targetVersion)
}
