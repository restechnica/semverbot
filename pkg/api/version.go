package api

import (
	"fmt"
	"github.com/restechnica/semverbot/internal/semver"
	"github.com/restechnica/semverbot/pkg/cli"
)

type VersionAPI struct {
	GitAPI GitAPI
}

func NewVersionAPI() VersionAPI {
	return VersionAPI{NewGitAPI()}
}

func (api VersionAPI) GetVersion() (version string, err error) {
	version, err = api.GitAPI.GetLatestAnnotatedTag()
	return semver.Trim(version)
}

func (api VersionAPI) GetVersionOrDefault(defaultVersion string) (version string) {
	var err error

	if version, err = api.GetVersion(); err != nil {
		version = defaultVersion
	}

	return version
}

func (api VersionAPI) PredictVersion(mode semver.Mode) (version string, err error) {
	version = api.GetVersionOrDefault(cli.DefaultVersion)
	return mode.Increment(version)
}

func (api VersionAPI) PushVersion(prefix string) (err error) {
	var version = api.GetVersionOrDefault(cli.DefaultVersion)
	var prefixedVersion = fmt.Sprintf("%s%s", prefix, version)
	return api.GitAPI.PushTag(prefixedVersion)
}
