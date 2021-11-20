package core

import (
	"github.com/restechnica/semverbot/pkg/modes"
	"github.com/restechnica/semverbot/pkg/versions"
)

type PushVersionOptions struct {
	DefaultVersion string
	GitTagsPrefix  string
}

// PushVersion pushes the current version.
// Returns an error if the push went wrong.
func PushVersion(options *PushVersionOptions) (err error) {
	var versionAPI = versions.NewAPI(modes.API{})
	var version = versionAPI.GetVersionOrDefault(options.DefaultVersion)
	return versionAPI.PushVersion(version, options.GitTagsPrefix)
}
