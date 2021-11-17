package core

import (
	"fmt"

	"github.com/restechnica/semverbot/pkg/api"
)

type PushVersionOptions struct {
	DefaultVersion string
	GitTagsPrefix  string
}

func PushVersion(options *PushVersionOptions) (err error) {
	var versionAPI = api.NewVersionAPI()
	var version = versionAPI.GetVersionOrDefault(options.DefaultVersion)

	var prefixedVersion = fmt.Sprintf("%s%s", options.GitTagsPrefix, version)

	var gitAPI = api.NewGitAPI()
	return gitAPI.PushTag(prefixedVersion)
}
