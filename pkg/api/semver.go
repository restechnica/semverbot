package api

import (
	"github.com/restechnica/semverbot/internal/semver"
)

type SemverModeAPI struct {
	GitBranchMatchers map[string]string
	GitCommitMatchers map[string]string
}

func NewSemverModeAPI(gitBranchMatchers map[string]string, gitCommitMatchers map[string]string) SemverModeAPI {
	return SemverModeAPI{GitBranchMatchers: gitBranchMatchers, GitCommitMatchers: gitCommitMatchers}
}

func (api SemverModeAPI) SelectMode(mode string) semver.Mode {
	switch mode {
	case semver.Auto:
		return semver.NewAutoMode([]semver.Mode{
			semver.NewGitCommitMode(api.GitCommitMatchers),
		})
	case semver.GitCommit:
		return semver.NewGitCommitMode(api.GitCommitMatchers)
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
