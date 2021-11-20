package core

import (
	"github.com/restechnica/semverbot/pkg/modes"
	"github.com/restechnica/semverbot/pkg/versions"
)

type GetVersionOptions struct {
	DefaultVersion string
}

// GetVersion gets the current version.
// Returns the current version.
func GetVersion(options *GetVersionOptions) string {
	var versionAPI = versions.NewAPI(modes.API{})
	return versionAPI.GetVersionOrDefault(options.DefaultVersion)
}
