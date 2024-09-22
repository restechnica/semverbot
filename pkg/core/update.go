package core

import (
	"github.com/restechnica/semverbot/pkg/versions"
)

type UpdateVersionOptions struct {
	GitTagsPrefix string
	GitTagsSuffix string
}

// UpdateVersion updates to the latest version.
// Returns an error if updating the version went wrong.
func UpdateVersion(updateOptions *UpdateVersionOptions) error {
	var versionAPI = versions.NewAPI(updateOptions.GitTagsPrefix, updateOptions.GitTagsSuffix)
	return versionAPI.UpdateVersion()
}
