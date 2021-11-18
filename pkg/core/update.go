package core

import (
	"github.com/restechnica/semverbot/pkg/versions"
)

// UpdateVersion updates to the latest version.
// Returns an error if updating the version went wrong.
func UpdateVersion() error {
	var versionAPI = versions.NewAPI()
	return versionAPI.UpdateVersion()
}
