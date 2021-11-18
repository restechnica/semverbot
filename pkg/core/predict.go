package core

import (
	"github.com/restechnica/semverbot/pkg/semver"
	"github.com/restechnica/semverbot/pkg/version"
)

type PredictVersionOptions struct {
	DefaultVersion string
	SemverMatchMap map[string][]string
	SemverMode     string
}

// PredictVersion predicts a version based on the latest annotated git tag and a map of semver levels
// matched to specific strings.
// Returns the predicted version or an error if anything went wrong with the increment.
func PredictVersion(options *PredictVersionOptions) (prediction string, err error) {
	var versionAPI = version.NewAPI()
	var version = versionAPI.GetVersionOrDefault(options.DefaultVersion)

	var modeDetector = semver.NewModeDetector(options.SemverMatchMap)

	var semverModeAPI = semver.NewModeAPI(modeDetector)
	var semverMode = semverModeAPI.SelectMode(options.SemverMode)

	return semverMode.Increment(version)
}
