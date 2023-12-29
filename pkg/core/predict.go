package core

import (
	"github.com/restechnica/semverbot/pkg/modes"
	"github.com/restechnica/semverbot/pkg/semver"
	"github.com/restechnica/semverbot/pkg/versions"
)

type PredictVersionOptions struct {
	DefaultVersion      string
	GitBranchDelimiters string
	GitCommitDelimiters string
	GitTagsPrefix       string
	Mode                string
	SemverMap           semver.Map
}

// PredictVersion predicts a version based on a modes.Mode and a modes.Map.
// The modes.Map values will be matched against git information to detect which semver level to increment.
// Returns the next version or an error if the prediction failed.
func PredictVersion(options *PredictVersionOptions) (prediction string, err error) {
	var gitBranchMode = modes.NewGitBranchMode(options.GitBranchDelimiters, options.SemverMap)
	var gitCommitMode = modes.NewGitCommitMode(options.GitCommitDelimiters, options.SemverMap)

	var versionAPI = versions.NewAPI(options.GitTagsPrefix)
	var version = versionAPI.GetVersionOrDefault(options.DefaultVersion)

	var modeAPI = modes.NewAPI(gitBranchMode, gitCommitMode)
	var mode = modeAPI.SelectMode(options.Mode)

	return versionAPI.PredictVersion(version, mode)
}
