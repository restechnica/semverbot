package core

import (
	"github.com/restechnica/semverbot/pkg/versions"
)

type ReleaseVersionOptions struct {
	DefaultVersion string
	GitTagsPrefix  string
	SemverMatchMap map[string][]string
	SemverMode     string
}

// ReleaseVersion releases a new version by incrementing the latest annotated git tag.
// It creates an annotated git tag for the new version.
// Returns an error if anything went wrong with incrementing or tagging.
func ReleaseVersion(options *ReleaseVersionOptions) (err error) {
	var versionAPI = versions.NewAPI()

	var version = versionAPI.GetVersionOrDefault(options.DefaultVersion)

	if version, err = versionAPI.PredictVersion(
		version,
		options.SemverMatchMap,
		options.SemverMode,
	); err != nil {
		return err
	}

	return versionAPI.ReleaseVersion(version)
}
