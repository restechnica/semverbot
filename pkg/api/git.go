package api

import (
	"github.com/restechnica/semverbot/internal/commands"
)

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
	return api.commander.Run("git", "fetch", "--unshallow")
}

func (api GitAPI) GetConfig(key string) (value string, err error) {
	return api.commander.Output("git", "config", "--get", key)
}

func (api GitAPI) GetLatestAnnotatedTag() (value string, err error) {
	return api.commander.Output("git", "describe", "--tags")
}

func (api GitAPI) PushTag(tag string) (err error) {
	return api.commander.Run("git", "push", "origin", tag)
}

func (api GitAPI) SetConfig(key string, value string) (err error) {
	return api.commander.Run("git", "config", key, value)
}

func (api GitAPI) SetConfigIfNotSet(key string, value string) (err error) {
	if _, err = api.GetConfig(key); err != nil {
		err = api.SetConfig(key, value)
	}

	return err
}
