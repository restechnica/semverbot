package core

import (
	"github.com/restechnica/semverbot/pkg/versions"
)

// ReleaseVersion releases a new version.
// Returns an error if anything went wrong with the prediction or releasing.
func ReleaseVersion(predictOptions *PredictVersionOptions) error {
	var versionAPI = versions.NewAPI(predictOptions.GitTagsPrefix)
	var predictedVersion, err = PredictVersion(predictOptions)

	if err != nil {
		return err
	}

	return versionAPI.ReleaseVersion(predictedVersion)
}
