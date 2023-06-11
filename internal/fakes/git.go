package fakes

import "fmt"

// FakeGitAPI a git.API interface fake implementation.
type FakeGitAPI struct {
	Config     map[string]string
	LocalTags  []string
	PushedTags []string
}

// NewFakeGitAPI creates a new FakeGitAPI.
// Returns the new FakeGitAPI.
func NewFakeGitAPI() *FakeGitAPI {
	return &FakeGitAPI{
		Config:     map[string]string{},
		LocalTags:  []string{},
		PushedTags: []string{},
	}
}

// CreateAnnotatedTag creates a fake tag.
func (fake *FakeGitAPI) CreateAnnotatedTag(tag string) (err error) {
	fake.LocalTags = append(fake.LocalTags, tag)
	return err
}

// FetchTags does nothing.
func (fake *FakeGitAPI) FetchTags() (err error) {
	return err
}

// FetchUnshallow does nothing.
func (fake *FakeGitAPI) FetchUnshallow() (err error) {
	return err
}

// GetConfig returns a fake config.
func (fake *FakeGitAPI) GetConfig(key string) (value string, err error) {
	var config, exists = fake.Config[key]

	if exists {
		return config, nil
	}

	return "", fmt.Errorf("config does not exist")
}

// GetLatestAnnotatedTag returns a fake tag.
func (fake *FakeGitAPI) GetLatestAnnotatedTag() (tag string, err error) {
	if len(fake.LocalTags) == 0 {
		return tag, fmt.Errorf("no tags found")
	}
	return fake.LocalTags[len(fake.LocalTags)-1], nil
}

// GetLatestCommitMessage does nothing.
func (fake *FakeGitAPI) GetLatestCommitMessage() (message string, err error) {
	return message, err
}

// GetMergedBranchName does nothing.
func (fake *FakeGitAPI) GetMergedBranchName() (name string, err error) {
	return name, err
}

// GetTags does nothing
func (fake *FakeGitAPI) GetTags() (tags string, err error) {
	return tags, err
}

// PushTag pushes a fake tag.
func (fake *FakeGitAPI) PushTag(tag string) (err error) {
	fake.PushedTags = append(fake.PushedTags, tag)
	return err
}

// SetConfig sets a fake config.
func (fake *FakeGitAPI) SetConfig(key string, value string) (err error) {
	fake.Config[key] = value
	return err
}

// SetConfigIfNotSet sets a fake config if it does not exist.
func (fake *FakeGitAPI) SetConfigIfNotSet(key string, value string) (actual string, err error) {
	if actual, err = fake.GetConfig(key); err != nil {
		err = fake.SetConfig(key, value)
		actual = value
	}

	return actual, err
}
