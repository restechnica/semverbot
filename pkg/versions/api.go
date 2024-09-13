package versions

import (
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/restechnica/semverbot/pkg/git"
	"github.com/restechnica/semverbot/pkg/modes"
	"github.com/restechnica/semverbot/pkg/semver"
)

// API an API to work with versions.
type API struct {
	Prefix string
	Suffix string
	GitAPI git.API
}

// NewAPI creates a new version API.
// Returns the new API.
func NewAPI(prefix string, suffix string) API {
	return API{Prefix: prefix, Suffix: suffix, GitAPI: git.NewCLI()}
}

// GetVersion gets the latest valid semver version from the git tags.
// The tag is trimmed because git adds newlines to the underlying command.
// Returns the current version or an error if the GitAPI failed.
func (api API) GetVersion() (currentVersion string, err error) {
	var tags string

	if tags, err = api.GitAPI.GetTags(); err != nil {
		return currentVersion, err
	}

	// strip all newlines
	var versions = strings.Fields(tags)

	if currentVersion, err = semver.Find(api.Prefix, api.Suffix, versions); err != nil {
		return currentVersion, err
	}

	return semver.Trim(api.Prefix, api.Suffix, currentVersion)
}

// GetVersionOrDefault gets the current version or a default version if it failed.
// Returns the current version or a default version.
func (api API) GetVersionOrDefault(defaultVersion string) (version string) {
	var err error

	log.Info().Msg("getting version...")

	if version, err = api.GetVersion(); err != nil {
		log.Debug().Err(err).Msg("")
		log.Warn().Msg("falling back to default version")
		version = defaultVersion
	}

	log.Info().Msg(version)

	return version
}

// PredictVersion increments a version based on a modes.Mode.
// Returns the next version or an error if the increment failed.
func (api API) PredictVersion(version string, mode modes.Mode) (string, error) {
	var err error

	log.Info().Msg("predicting version...")

	version, err = mode.Increment(api.Prefix, api.Suffix, version)

	log.Info().Msg(version)

	return version, err
}

// ReleaseVersion releases a version by creating an annotated git tag with a prefix.
// Returns an error if the tag creation failed.
func (api API) ReleaseVersion(version string) (err error) {
	log.Info().Msg("releasing version...")
	var prefixedVersion = AddPrefix(version, api.Prefix)
	var prefixedAndSuffixedVersion = AddSuffix(prefixedVersion, api.Suffix)
	return api.GitAPI.CreateAnnotatedTag(prefixedAndSuffixedVersion)
}

// PushVersion pushes a version by pushing a git tag with a prefix.
// Returns an error if pushing the tag failed.
func (api API) PushVersion(version string) (err error) {
	log.Info().Msg("pushing version...")
	var prefixedVersion = AddPrefix(version, api.Prefix)
	var prefixedAndSuffixedVersion = AddSuffix(prefixedVersion, api.Suffix)
	return api.GitAPI.PushTag(prefixedAndSuffixedVersion)
}

// UpdateVersion updates the version by making the git repo unshallow and by fetching all git tags.
// Returns and error if anything went wrong. Errors from making the git repo unshallow are ignored.
func (api API) UpdateVersion() (err error) {
	log.Info().Msg("updating version...")

	var output string

	log.Info().Msg("fetching unshallow repository...")

	if output, err = api.GitAPI.FetchUnshallow(); err != nil {
		log.Debug().Err(err).Msg("")
		log.Warn().Msg("ignoring failed unshallow fetch for now, repository might already be complete")
	} else {
		log.Debug().Msg(strings.Trim(output, "\n"))
	}

	log.Info().Msg("fetching tags...")

	if output, err = api.GitAPI.FetchTags(); err == nil {
		log.Debug().Msg(strings.Trim(output, "\n"))
	}

	return err
}
