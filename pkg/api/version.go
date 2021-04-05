package api

import (
	"strings"

	"github.com/restechnica/semverbot/internal/commands"
	"github.com/restechnica/semverbot/internal/semver"
	"github.com/restechnica/semverbot/pkg/cli"
)

type VersionAPI struct {
	Commander commands.Commander
}

func NewVersionAPI() VersionAPI {
	return VersionAPI{commands.NewExecCommander()}
}

func (api VersionAPI) GetVersion() (version string, err error) {
	version, err = api.Commander.Output("git", "describe", "--tags")
	return strings.TrimSpace(version), err
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
