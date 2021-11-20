package core

import (
	"github.com/restechnica/semverbot/pkg/modes"
	"github.com/restechnica/semverbot/pkg/versions"
)

type PredictVersionOptions struct {
	DefaultVersion string
	Mode           string
	SemverMap      modes.SemverMap
}

// PredictVersion predicts the next version.
// Returns the predicted version or an error if anything went wrong with the prediction.
func PredictVersion(options *PredictVersionOptions) (prediction string, err error) {
	var versionAPI = versions.NewAPI()
	var version = versionAPI.GetVersionOrDefault(options.DefaultVersion)
	return versionAPI.PredictVersion(version, options.SemverMap, options.Mode)
}
