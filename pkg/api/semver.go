package api

import (
	"github.com/restechnica/semverbot/internal/semver"
)

type SemverModeAPI struct {
	GitBranchMode semver.GitBranchMode
	GitCommitMode semver.GitCommitMode
}

func NewSemverModeAPI(detector semver.ModeDetector) SemverModeAPI {
	return SemverModeAPI{
		GitBranchMode: semver.NewGitBranchMode(detector),
		GitCommitMode: semver.NewGitCommitMode(detector),
	}
}

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
