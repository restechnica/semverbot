package core

import (
	"github.com/restechnica/semverbot/pkg/api"
	"github.com/restechnica/semverbot/pkg/cli"
)

// GetVersion gets the current version based on the latest annotated git tag.
// Returns the current version.
func GetVersion() string {
	var versionAPI = api.NewVersionAPI()
	return versionAPI.GetVersionOrDefault(cli.DefaultVersion)
}
