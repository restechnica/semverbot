package core

import (
	"github.com/restechnica/semverbot/pkg/versions"
)

type UpdateVersionOptions struct {
	GitTagsPrefix string
}

// UpdateVersion updates to the latest version.
// Returns an error if updating the version went wrong.
func UpdateVersion(updateOptions *UpdateVersionOptions) error {
	var versionAPI = versions.NewAPI(updateOptions.GitTagsPrefix)
	return versionAPI.UpdateVersion()
}
