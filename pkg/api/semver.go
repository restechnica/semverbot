package api

import (
	"github.com/restechnica/semverbot/pkg/semver"
)

// SemverModeAPI an API to work with different modes.
type SemverModeAPI struct {
	GitBranchMode semver.GitBranchMode
	GitCommitMode semver.GitCommitMode
}

// NewSemverModeAPI creates a new semver mode API with a mode detector to pass
// it on to the different modes that require it.
// Returns the new SemverModeAPI.
func NewSemverModeAPI(detector semver.ModeDetector) SemverModeAPI {
	return SemverModeAPI{
		GitBranchMode: semver.NewGitBranchMode(detector),
		GitCommitMode: semver.NewGitCommitMode(detector),
	}
}

// SelectMode selects the mode corresponding to the mode string.
// Returns the corresponding mode.
func (api SemverModeAPI) SelectMode(mode string) semver.Mode {
	switch mode {
	case semver.Auto:
		return semver.NewAutoMode([]semver.Mode{api.GitBranchMode, api.GitCommitMode})
	case semver.GitCommit:
		return api.GitCommitMode
	case semver.GitBranch:
		return api.GitBranchMode
	case semver.Patch:
		return semver.NewPatchMode()
	case semver.Minor:
		return semver.NewMinorMode()
	case semver.Major:
		return semver.NewMajorMode()
	default:
		return semver.NewPatchMode()
	}
}
