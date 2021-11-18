package git

import "github.com/restechnica/semverbot/internal/commands"

// API an API to interact with the git CLI.
type API struct {
	commander commands.Commander
}

// NewAPI creates a new API with a commander to run git commands.
// Returns the new API.
func NewAPI() API {
	return API{commander: commands.NewExecCommander()}
}

// CreateAnnotatedTag creates an annotated git tag.
// Returns an error if the command fails.
func (api API) CreateAnnotatedTag(tag string) (err error) {
	return api.commander.Run("git", "tag", "-a", tag, "-m", tag)
}

// FetchTags fetches all tags from the remote origin.
// Returns an error if the command fails.
func (api API) FetchTags() (err error) {
	return api.commander.Run("git", "fetch", "--tags")
}

// FetchUnshallow convert a shallow repository to a complete one.
// Returns an error if the command fails.
func (api API) FetchUnshallow() (err error) {
	return api.commander.Run("git", "fetch", "--unshallow")
}

// GetConfig gets the git config for a specific key.
// Returns the value of the git config as a string and an error if the command failed.
func (api API) GetConfig(key string) (value string, err error) {
	return api.commander.Output("git", "config", "--get", key)
}

// GetLatestAnnotatedTag gets the latest annotated git tag.
// Returns the git tag and an error if the command failed.
func (api API) GetLatestAnnotatedTag() (tag string, err error) {
	return api.commander.Output("git", "describe", "--tags")
}

// GetLatestCommitMessage gets the latest git commit message.
// Returns the git commit message or an error if the command failed.
func (api API) GetLatestCommitMessage() (message string, err error) {
	return api.commander.Output("git", "--no-pager", "show", "-s", "--format=%s")
}

// GetMergedBranchName gets the source branch name if the last commit is a merge.
// Returns the branch name or an error if something went wrong with git.
func (api API) GetMergedBranchName() (name string, err error) {
	return api.commander.Output(
		"git",
		"name-rev",
		"--name-only",
		"--refs=refs/heads/*",
		"--refs=refs/remotes/*",
		"HEAD^2",
	)
}

// PushTag pushes a tag to the remote origin.
// Returns an error if the command failed.
func (api API) PushTag(tag string) (err error) {
	return api.commander.Run("git", "push", "origin", tag)
}

// SetConfig sets a git config key and value.
// Returns an error if the command failed.
func (api API) SetConfig(key string, value string) (err error) {
	return api.commander.Run("git", "config", key, value)
}

// SetConfigIfNotSet sets a git config key and value if the config does not exist.
// Returns an error if the command failed.
func (api API) SetConfigIfNotSet(key string, value string) (err error) {
	if _, err = api.GetConfig(key); err != nil {
		err = api.SetConfig(key, value)
	}

	return err
}
