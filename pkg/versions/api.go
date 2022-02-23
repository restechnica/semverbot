package versions

import (
	"github.com/restechnica/semverbot/pkg/git"
	"github.com/restechnica/semverbot/pkg/modes"
	"github.com/restechnica/semverbot/pkg/semver"
	"strings"
)

// API an API to work with versions.
type API struct {
	GitAPI git.API
}

// NewAPI creates a new version API.
// Returns the new API.
func NewAPI() API {
	return API{GitAPI: git.NewCLI()}
}

// GetVersion gets the latest valid semver version from the git tags.
// The tag is trimmed because git adds newlines to the underlying command.
// Returns the current version or an error if the GitAPI failed.
func (api API) GetVersion() (currentVersion string, err error) {
	var tags string

	if tags, err = api.GitAPI.GetTags(); err != nil {
		return currentVersion, err
	}

	// strip all newlines
	var versions = strings.Fields(tags)

	if currentVersion, err = semver.Find(versions); err != nil {
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

// PredictVersion increments a version based on a modes.Mode.
// Returns the next version or an error if the increment failed.
func (api API) PredictVersion(version string, mode modes.Mode) (string, error) {
	return mode.Increment(version)
}

// ReleaseVersion releases a version by creating an annotated git tag with a prefix.
// Returns an error if the tag creation failed.
func (api API) ReleaseVersion(version string, prefix string) (err error) {
	var prefixedVersion = AddPrefix(version, prefix)
	return api.GitAPI.CreateAnnotatedTag(prefixedVersion)
}

// PushVersion pushes a version by pushing a git tag with a prefix.
// Returns an error if pushing the tag failed.
func (api API) PushVersion(version string, prefix string) (err error) {
	var prefixedVersion = AddPrefix(version, prefix)
	return api.GitAPI.PushTag(prefixedVersion)
}

// UpdateVersion updates the version by making the git repo unshallow and by fetching all git tags.
// Returns and error if anything went wrong. Errors from making the git repo unshallow are ignored.
func (api API) UpdateVersion() (err error) {
	err = api.GitAPI.FetchUnshallow()
	err = api.GitAPI.FetchTags()
	return err
}
