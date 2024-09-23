package modes

import (
	"github.com/restechnica/semverbot/pkg/git"
	"github.com/restechnica/semverbot/pkg/semver"
)

// GitCommit mode name for GitCommitMode.
const GitCommit = "git-commit"

// GitCommitMode implementation of the Mode interface.
// It increments the semver level based on the latest git commit messages.
type GitCommitMode struct {
	Delimiters string
	GitAPI     git.API
	SemverMap  semver.Map
}

// NewGitCommitMode creates a new GitCommitMode.
// Returns the new GitCommitMode.
func NewGitCommitMode(delimiters string, semverMap semver.Map) GitCommitMode {
	return GitCommitMode{Delimiters: delimiters, GitAPI: git.NewCLI(), SemverMap: semverMap}
}

// Increment increments a given version based on the latest git commit message.
// Returns the incremented version or an error if it failed to detect the mode based on the git commit.
func (mode GitCommitMode) Increment(prefix string, suffix string, targetVersion string) (nextVersion string, err error) {
	var message string
	var detectedMode Mode

	if message, err = mode.GitAPI.GetLatestCommitMessage(); err != nil {
		return
	}

	if detectedMode, err = DetectModeFromString(message, mode.SemverMap, mode.Delimiters); err != nil {
		return
	}

	return detectedMode.Increment(prefix, suffix, targetVersion)
}

// String returns a string representation of an instance.
func (mode GitCommitMode) String() string {
	return GitCommit
}
