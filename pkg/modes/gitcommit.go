package modes

import (
	"fmt"

	"github.com/restechnica/semverbot/internal/util"
	"github.com/restechnica/semverbot/pkg/git"
)

// GitCommit mode name for GitCommitMode.
const GitCommit = "git-commit"

// GitCommitMode implementation of the Mode interface.
// It increments the semver level based on the latest git commit messages.
type GitCommitMode struct {
	Delimiters string
	GitAPI     git.API
	SemverMap  SemverMap
}

// NewGitCommitMode creates a new GitCommitMode.
// Returns the new GitCommitMode.
func NewGitCommitMode(delimiters string, semverMap SemverMap) GitCommitMode {
	return GitCommitMode{Delimiters: delimiters, GitAPI: git.NewCLI(), SemverMap: semverMap}
}

// Increment increments a given version based on the latest git commit message.
// Returns the incremented version or an error if it failed to detect the mode based on the git commit.
func (mode GitCommitMode) Increment(targetVersion string) (nextVersion string, err error) {
	var message string
	var detectedMode Mode

	if message, err = mode.GitAPI.GetLatestCommitMessage(); err != nil {
		return
	}

	if detectedMode, err = mode.DetectMode(message); err != nil {
		return
	}

	return detectedMode.Increment(targetVersion)
}

// DetectMode detects the mode (patch, minor, major) based on a git commit message.
// Returns the detected mode.
func (mode GitCommitMode) DetectMode(commitMessage string) (detected Mode, err error) {
	for level, values := range mode.SemverMap {
		for _, value := range values {
			if mode.isMatch(commitMessage, value) {
				switch level {
				case Patch:
					detected = NewPatchMode()
				case Minor:
					detected = NewMinorMode()
				case Major:
					detected = NewMajorMode()
				}
				return detected, err
			}
		}
	}

	return detected, fmt.Errorf(`failed to detect mode from git commit message "%s" with delimiters "%s"`,
		commitMessage, mode.Delimiters)
}

// isMatch returns true if a string is part of commit message, after splitting the git commit message with delimiters
func (mode GitCommitMode) isMatch(commitMessage string, value string) bool {
	return util.Contains(commitMessage, value, mode.Delimiters)
}
