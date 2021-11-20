package core

import (
	"github.com/restechnica/semverbot/pkg/modes"
	"github.com/restechnica/semverbot/pkg/versions"
)

type ReleaseVersionOptions struct {
	DefaultVersion      string
	GitTagsPrefix       string
	GitBranchDelimiters string
	GitCommitDelimiters string
	Mode                string
	SemverMap           modes.SemverMap
}

// ReleaseVersion releases a new version.
// Returns an error if anything went wrong with the prediction or releasing.
func ReleaseVersion(predictOptions *PredictVersionOptions, releaseOptions *ReleaseVersionOptions) error {
	var versionAPI = versions.NewAPI(modes.API{})
	var predictedVersion, err = PredictVersion(predictOptions)

	if err != nil {
		return err
	}

	return versionAPI.ReleaseVersion(predictedVersion, releaseOptions.GitTagsPrefix)
}
