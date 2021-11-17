package core

import (
	"github.com/restechnica/semverbot/pkg/api"
)

type GetVersionOptions struct {
	DefaultVersion string
}

// GetVersion gets the current version based on the latest annotated git tag.
// Returns the current version.
func GetVersion(options *GetVersionOptions) string {
	var versionAPI = api.NewVersionAPI()
	return versionAPI.GetVersionOrDefault(options.DefaultVersion)
}
