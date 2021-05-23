package api

import "github.com/restechnica/semverbot/internal/commands"

type GitAPI struct {
	commander commands.Commander
}

func NewGitAPI() GitAPI {
	return GitAPI{commander: commands.NewExecCommander()}
}

func (api GitAPI) CreateAnnotatedTag(tag string) (err error) {
	return api.commander.Run("git", "tag", "-a", tag, "-m", tag)
}

func (api GitAPI) FetchTags() (err error) {
	return api.commander.Run("git", "fetch", "--tags")
}
