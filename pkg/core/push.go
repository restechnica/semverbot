package core

import (
	"fmt"
	"github.com/restechnica/semverbot/pkg/version"

	"github.com/restechnica/semverbot/pkg/git"
)

type PushVersionOptions struct {
	DefaultVersion string
	GitTagsPrefix  string
}

// PushVersion pushes the latest annotated git tag to the git origin.
// Returns an error if pushing the tag went wrong.
func PushVersion(options *PushVersionOptions) (err error) {
	var versionAPI = version.NewAPI()
	var version = versionAPI.GetVersionOrDefault(options.DefaultVersion)

	var prefixedVersion = fmt.Sprintf("%s%s", options.GitTagsPrefix, version)

	var gitAPI = git.NewAPI()
	return gitAPI.PushTag(prefixedVersion)
}
