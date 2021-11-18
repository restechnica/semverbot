package core

import (
	"fmt"

	"github.com/restechnica/semverbot/pkg/git"
	"github.com/restechnica/semverbot/pkg/semver"
	"github.com/restechnica/semverbot/pkg/version"
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
	var versionAPI = version.NewAPI()
	var version = versionAPI.GetVersionOrDefault(options.DefaultVersion)

	var modeDetector = semver.NewModeDetector(options.SemverMatchMap)

	var modeAPI = semver.NewModeAPI(modeDetector)
	var semverMode = modeAPI.SelectMode(options.SemverMode)

	var incrementedVersion string

	if incrementedVersion, err = semverMode.Increment(version); err != nil {
		return err
	}

	incrementedVersion = fmt.Sprintf("%s%s", options.GitTagsPrefix, incrementedVersion)

	var gitAPI = git.NewAPI()
	return gitAPI.CreateAnnotatedTag(incrementedVersion)
}
