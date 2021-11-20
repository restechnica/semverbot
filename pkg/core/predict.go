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

// PredictVersion predicts the next version.
// Returns the predicted version or an error if anything went wrong with the prediction.
func PredictVersion(options *PredictVersionOptions) (prediction string, err error) {
	var modeAPI = modes.NewAPI(options.SemverMap, options.GitBranchDelimiters, options.GitCommitDelimiters)
	var versionAPI = versions.NewAPI(modeAPI)
	var version = versionAPI.GetVersionOrDefault(options.DefaultVersion)
	return versionAPI.PredictVersion(version, options.SemverMap, options.Mode)
}
