package api

import (
	"github.com/restechnica/semverbot/internal/commands"
)

// GitAPI an API to interact with the git CLI.
type GitAPI struct {
	commander commands.Commander
}

// NewGitAPI creates a new GitAPI with a commander to run git commands.
// returns the new GitAPI.
func NewGitAPI() GitAPI {
	return GitAPI{commander: commands.NewExecCommander()}
}

// CreateAnnotatedTag creates an annotated git tag.
// returns an error if the command fails.
func (api GitAPI) CreateAnnotatedTag(tag string) (err error) {
	return api.commander.Run("git", "tag", "-a", tag, "-m", tag)
}

// FetchTags fetches all tags from the remote origin.
// returns an error if the command fails.
func (api GitAPI) FetchTags() (err error) {
	return api.commander.Run("git", "fetch", "--tags")
}

// FetchUnshallow convert a shallow repository to a complete one.
// returns an error if the command fails.
func (api GitAPI) FetchUnshallow() (err error) {
	return api.commander.Run("git", "fetch", "--unshallow")
}

// GetConfig gets the git config for a specific key.
// returns the value of the git config as a string and an error if the command failed.
func (api GitAPI) GetConfig(key string) (value string, err error) {
	return api.commander.Output("git", "config", "--get", key)
}

// GetLatestAnnotatedTag gets the latest annotated git tag.
// returns the git tag and an error if the command failed.
func (api GitAPI) GetLatestAnnotatedTag() (value string, err error) {
	return api.commander.Output("git", "describe", "--tags")
}

// PushTag pushes a tag to the remote origin.
// returns an error if the command failed.
func (api GitAPI) PushTag(tag string) (err error) {
	return api.commander.Run("git", "push", "origin", tag)
}

// SetConfig sets a git config key and value.
// returns an error if the command failed.
func (api GitAPI) SetConfig(key string, value string) (err error) {
	return api.commander.Run("git", "config", key, value)
}

// SetConfigIfNotSet sets a git config key and value if the config does not exist.
// returns an error if the command failed.
func (api GitAPI) SetConfigIfNotSet(key string, value string) (err error) {
	if _, err = api.GetConfig(key); err != nil {
		err = api.SetConfig(key, value)
	}

	return err
}
