package core

import (
	"github.com/restechnica/semverbot/pkg/modes"
	"github.com/restechnica/semverbot/pkg/versions"
)

type PredictVersionOptions struct {
	DefaultVersion      string
	GitBranchDelimiters string
	GitCommitDelimiters string
	Mode                string
	SemverMap           modes.SemverMap
}

// PredictVersion predicts a version based on a modes.Mode and a modes.SemverMap.
// The modes.SemverMap values will be matched against git information to detect which semver level to increment.
// Returns the next version or an error if the prediction failed.
func PredictVersion(options *PredictVersionOptions) (prediction string, err error) {
	var modeAPI = modes.NewAPI(options.SemverMap, options.GitBranchDelimiters, options.GitCommitDelimiters)
	var mode = modeAPI.SelectMode(options.Mode)

	var versionAPI = versions.NewAPI()
	var version = versionAPI.GetVersionOrDefault(options.DefaultVersion)

	return versionAPI.PredictVersion(version, mode)
}
