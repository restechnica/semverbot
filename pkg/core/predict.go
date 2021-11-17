package core

import (
	"github.com/restechnica/semverbot/pkg/api"
	"github.com/restechnica/semverbot/pkg/semver"
)

type PredictVersionOptions struct {
	DefaultVersion string
	SemverMatchMap map[string][]string
	SemverMode     string
}

func PredictVersion(options *PredictVersionOptions) (prediction string, err error) {
	var versionAPI = api.NewVersionAPI()
	var version = versionAPI.GetVersionOrDefault(options.DefaultVersion)

	var modeDetector = semver.NewModeDetector(options.SemverMatchMap)

	var semverModeAPI = api.NewSemverModeAPI(modeDetector)
	var semverMode = semverModeAPI.SelectMode(options.SemverMode)

	return semverMode.Increment(version)
}
