package core

import (
	"fmt"

	"github.com/restechnica/semverbot/pkg/api"
	"github.com/restechnica/semverbot/pkg/semver"
)

type ReleaseVersionOptions struct {
	DefaultVersion string
	GitTagsPrefix  string
	SemverMatchMap map[string][]string
	SemverMode     string
}

func ReleaseVersion(options *ReleaseVersionOptions) (err error) {
	var versionAPI = api.NewVersionAPI()
	var version = versionAPI.GetVersionOrDefault(options.DefaultVersion)

	var modeDetector = semver.NewModeDetector(options.SemverMatchMap)

	var semverModeAPI = api.NewSemverModeAPI(modeDetector)
	var semverMode = semverModeAPI.SelectMode(options.SemverMode)

	var incrementedVersion string

	if incrementedVersion, err = semverMode.Increment(version); err != nil {
		return err
	}

	incrementedVersion = fmt.Sprintf("%s%s", options.GitTagsPrefix, incrementedVersion)

	var gitAPI = api.NewGitAPI()
	return gitAPI.CreateAnnotatedTag(incrementedVersion)
}