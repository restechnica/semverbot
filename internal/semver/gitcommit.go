package semver

import (
	"fmt"
	"strings"

	"github.com/restechnica/semverbot/internal/commands"
)

// GitCommit mode name for GitCommitMode.
const GitCommit = "git-commit"

// GitCommitMode implementation of the Mode interface.
// It makes use of several matching strategies based on git commit messages.
type GitCommitMode struct {
	Commander commands.Commander
	Matches   map[string]string
}

// NewGitCommitMode creates a new GitCommitMode.
// Returns the new GitCommitMode.
func NewGitCommitMode(matchers map[string]string) GitCommitMode {
	return GitCommitMode{Commander: commands.ExecCommander{}, Matches: matchers}
}

// Increment increments a given version using the GitCommitMode.
// Returns the incremented version.
func (mode GitCommitMode) Increment(targetVersion string) (nextVersion string, err error) {
	var message string
	var matchedMode Mode

	if message, err = mode.Commander.Output("git", "show", "-s", "--format=%s"); err != nil {
		return
	}

	if matchedMode, err = mode.GetMatchedMode(message); err != nil {
		return
	}

	return matchedMode.Increment(targetVersion)
}

// GetMatchedMode gets the mode that testMatches specific tokens within the git commit message.
// It returns the matched mode.
func (mode GitCommitMode) GetMatchedMode(message string) (matched Mode, err error) {
	for match, mode := range mode.Matches {
		if strings.Contains(message, match) {
			switch mode {
			case Patch:
				matched = NewPatchMode()
			case Minor:
				matched = NewMinorMode()
			case Major:
				matched = NewMajorMode()
			}
			return
		}
	}

	return matched, fmt.Errorf(`could not match a mode to the commit message "%s"`, message)
}
