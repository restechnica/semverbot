package versions

import (
	"github.com/restechnica/semverbot/pkg/git"
	"github.com/restechnica/semverbot/pkg/semver"
)

// API an API to work with versions.
type API struct {
	GitAPI git.API
}

// NewAPI creates a new API.
// Returns the new API.
func NewAPI() API {
	return API{git.NewAPI()}
}

// GetVersion gets the current version.
// Git adds newlines to certain command output, which is why the version is trimmed.
// Returns the current version or an error if the GitAPI failed.
func (api API) GetVersion() (version string, err error) {
	if version, err = api.GitAPI.GetLatestAnnotatedTag(); err != nil {
		return version, err
	}
	return semver.Trim(version)
}

// GetVersionOrDefault gets the current version.
// Defaults to a provided default version if the GitAPI failed.
// Returns the current version.
func (api API) GetVersionOrDefault(defaultVersion string) (version string) {
	var err error

	if version, err = api.GetVersion(); err != nil {
		version = defaultVersion
	}

	return version
}

// PredictVersion predicts the next version with a provided semver mode.
// Returns the next version or an error if incrementing the current version failed.
func (api API) PredictVersion(version string, matchMap map[string][]string, mode string) (string, error) {
	var modeDetector = semver.NewModeDetector(matchMap)
	var semverModeAPI = semver.NewModeAPI(modeDetector)
	var semverMode = semverModeAPI.SelectMode(mode)
	return semverMode.Increment(version)
}

func (api API) ReleaseVersion(version string) (err error) {
	return api.GitAPI.CreateAnnotatedTag(version)
}

// PushVersion pushes a version with a prefix.
// Returns an error if the the GitAPI failed.
func (api API) PushVersion(version string, prefix string) (err error) {
	var prefixedVersion = AddPrefix(version, prefix)
	return api.GitAPI.PushTag(prefixedVersion)
}
