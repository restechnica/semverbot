package api

import (
	"fmt"

	"github.com/restechnica/semverbot/internal/semver"
	"github.com/restechnica/semverbot/pkg/cli"
)

// VersionAPI an API to work with versions.
type VersionAPI struct {
	GitAPI GitAPI
}

// NewVersionAPI creates a new VersionAPI.
// returns the new VersionAPI.
func NewVersionAPI() VersionAPI {
	return VersionAPI{NewGitAPI()}
}

// GetVersion gets the current version.
// returns the current version or an error if the GitAPI failed.
func (api VersionAPI) GetVersion() (version string, err error) {
	version, err = api.GitAPI.GetLatestAnnotatedTag()
	return semver.Trim(version)
}

// GetVersionOrDefault gets the current version.
// Defaults to a provided default version if the GitAPI failed.
// returns the current version.
func (api VersionAPI) GetVersionOrDefault(defaultVersion string) (version string) {
	var err error

	if version, err = api.GetVersion(); err != nil {
		version = defaultVersion
	}

	return version
}

// PredictVersion predicts the next version with a provided semver mode.
// returns the next version or an error if increment the current version failed.
func (api VersionAPI) PredictVersion(mode semver.Mode) (version string, err error) {
	version = api.GetVersionOrDefault(cli.DefaultVersion)
	return mode.Increment(version)
}

// PushVersion pushes a version with a provided version prefix.
// returns an error if the the GitAPI failed.
func (api VersionAPI) PushVersion(prefix string) (err error) {
	var version = api.GetVersionOrDefault(cli.DefaultVersion)
	var prefixedVersion = fmt.Sprintf("%s%s", prefix, version)
	return api.GitAPI.PushTag(prefixedVersion)
}
