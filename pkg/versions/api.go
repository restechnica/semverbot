package versions

import (
	"fmt"

	"github.com/restechnica/semverbot/pkg/git"
	"github.com/restechnica/semverbot/pkg/modes"
	"github.com/restechnica/semverbot/pkg/semver"
)

// API an API to work with versions.
type API struct {
	GitAPI git.API
}

// NewAPI creates a new version API.
// Returns the new API.
func NewAPI() API {
	return API{git.NewAPI()}
}

// GetVersion gets the current version by getting the latest annotated git tag.
// The tag is trimmed because git adds newlines to the underlying command.
// Returns the current version or an error if the GitAPI failed.
func (api API) GetVersion() (currentVersion string, err error) {
	if currentVersion, err = api.GitAPI.GetLatestAnnotatedTag(); err != nil {
		return currentVersion, err
	}
	return semver.Trim(currentVersion)
}

// GetVersionOrDefault gets the current version or a default version if it failed.
// Returns the current version or a default version.
func (api API) GetVersionOrDefault(defaultVersion string) (version string) {
	var err error

	if version, err = api.GetVersion(); err != nil {
		version = defaultVersion
	}

	return version
}

// PredictVersion increments a version based on a semver mode and a map of semver levels with matching strings.
// The matching strings will be matched against git information to detect which semver level to increment.
// Returns the next version or an error if the increment failed.
func (api API) PredictVersion(version string, semverMap modes.SemverMap, mode string) (string, error) {
	var modeDetector = modes.NewModeDetector(semverMap)
	var semverModeAPI = modes.NewAPI(modeDetector)
	var semverMode = semverModeAPI.SelectMode(mode)
	return semverMode.Increment(version)
}

// ReleaseVersion releases a version by creating an annotated git tag.
// Returns an error if the tag creation failed.
func (api API) ReleaseVersion(version string) (err error) {
	return api.GitAPI.CreateAnnotatedTag(version)
}

// PushVersion pushes a version by pushing a git tag with a prefix.
// Returns an error if pushing the tag failed.
func (api API) PushVersion(version string, prefix string) (err error) {
	var prefixedVersion = AddPrefix(version, prefix)
	return api.GitAPI.PushTag(prefixedVersion)
}

// UpdateVersion updates the version by making the git repo unshallow and by fetching all git tags.
// Returns and error if anything went wrong.
func (api API) UpdateVersion() (err error) {
	var gitAPI = git.NewAPI()

	if err = gitAPI.FetchUnshallow(); err != nil {
		fmt.Println("something went wrong while fetching from git, attempting to fetch tags anyway")
	}

	if err = gitAPI.FetchTags(); err != nil {
		fmt.Println("something went wrong while updating the version")
	} else {
		fmt.Println("successfully fetched the latest git tags")
	}

	return err
}
