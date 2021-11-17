package core

import (
	"github.com/restechnica/semverbot/internal/semver"
	"github.com/restechnica/semverbot/pkg/api"
)

type PredictVersionOptions struct {
	DefaultVersion  string
	SemverDetection map[string][]string
	SemverMode      string
}

func PredictVersion(options *PredictVersionOptions) (prediction string, err error) {
	var versionAPI = api.NewVersionAPI()
	var version = versionAPI.GetVersionOrDefault(options.DefaultVersion)

	var modeDetector = semver.NewModeDetector(options.SemverDetection)

	var semverModeAPI = api.NewSemverModeAPI(modeDetector)
	var semverMode = semverModeAPI.SelectMode(options.SemverMode)

	return semverMode.Increment(version)
}
