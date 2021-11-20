package git

import "github.com/restechnica/semverbot/internal/commands"

// CLI an CLI to interact with the git CLI.
type CLI struct {
	Commander commands.Commander
}

// NewCLI creates a new CLI with a commander to run git commands.
// Returns the new CLI.
func NewCLI() CLI {
	return CLI{Commander: commands.NewExecCommander()}
}

// CreateAnnotatedTag creates an annotated git tag.
// Returns an error if the command fails.
func (api CLI) CreateAnnotatedTag(tag string) (err error) {
	return api.Commander.Run("git", "tag", "-a", tag, "-m", tag)
}

// FetchTags fetches all tags from the remote origin.
// Returns an error if the command fails.
func (api CLI) FetchTags() (err error) {
	return api.Commander.Run("git", "fetch", "--tags")
}

// FetchUnshallow convert a shallow repository to a complete one.
// Returns an error if the command fails.
func (api CLI) FetchUnshallow() (err error) {
	return api.Commander.Run("git", "fetch", "--unshallow")
}

// GetConfig gets the git config for a specific key.
// Returns the value of the git config as a string and an error if the command failed.
func (api CLI) GetConfig(key string) (value string, err error) {
	return api.Commander.Output("git", "config", "--get", key)
}

// GetLatestAnnotatedTag gets the latest annotated git tag.
// Returns the git tag and an error if the command failed.
func (api CLI) GetLatestAnnotatedTag() (tag string, err error) {
	return api.Commander.Output("git", "describe", "--tags")
}

// GetLatestCommitMessage gets the latest git commit message.
// Returns the git commit message or an error if the command failed.
func (api CLI) GetLatestCommitMessage() (message string, err error) {
	return api.Commander.Output("git", "--no-pager", "show", "-s", "--format=%s")
}

// GetMergedBranchName gets the source branch name if the last commit is a merge.
// Returns the branch name or an error if something went wrong with git.
func (api CLI) GetMergedBranchName() (name string, err error) {
	return api.Commander.Output(
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
func (api CLI) PushTag(tag string) (err error) {
	return api.Commander.Run("git", "push", "origin", tag)
}

// SetConfig sets a git config key and value.
// Returns an error if the command failed.
func (api CLI) SetConfig(key string, value string) (err error) {
	return api.Commander.Run("git", "config", key, value)
}

// SetConfigIfNotSet sets a git config key and value if the config does not exist.
// Returns an error if the command failed.
func (api CLI) SetConfigIfNotSet(key string, value string) (err error) {
	if _, err = api.GetConfig(key); err != nil {
		err = api.SetConfig(key, value)
	}

	return err
}
