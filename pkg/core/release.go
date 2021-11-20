package core

import (
	"github.com/restechnica/semverbot/pkg/versions"
)

type ReleaseVersionOptions struct {
	DefaultVersion string
	GitTagsPrefix  string
	Mode           string
	SemverMap      map[string][]string
}

// ReleaseVersion releases a new version.
// Returns an error if anything went wrong with the prediction or releasing.
func ReleaseVersion(options *ReleaseVersionOptions) error {
	var versionAPI = versions.NewAPI()

	var currentVersion = versionAPI.GetVersionOrDefault(options.DefaultVersion)
	var predictedVersion, err = versionAPI.PredictVersion(currentVersion, options.SemverMap, options.Mode)

	if err != nil {
		return err
	}

	return versionAPI.ReleaseVersion(predictedVersion)
}
