package core

import (
	"github.com/rs/zerolog/log"

	"github.com/restechnica/semverbot/pkg/versions"
)

type GetVersionOptions struct {
	DefaultVersion string
}

// GetVersion gets the current version.
// Returns the current version.
func GetVersion(options *GetVersionOptions) string {
	log.Debug().Str("default", options.DefaultVersion).Msg("options:")
	var versionAPI = versions.NewAPI()
	return versionAPI.GetVersionOrDefault(options.DefaultVersion)
}
