package core

import (
	"github.com/restechnica/semverbot/pkg/versions"
)

type PushVersionOptions struct {
	DefaultVersion string
	GitTagsPrefix  string
}

// PushVersion pushes the latest annotated git tag to the git origin.
// Returns an error if pushing the tag went wrong.
func PushVersion(options *PushVersionOptions) (err error) {
	var versionAPI = versions.NewAPI()
	var version = versionAPI.GetVersionOrDefault(options.DefaultVersion)
	return versionAPI.PushVersion(version, options.GitTagsPrefix)
}
