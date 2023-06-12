package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockGitAPI a git.API interface mock implementation.
type MockGitAPI struct {
	mock.Mock
}

// NewMockGitAPI creates a new MockGitAPI.
// Returns the new MockGitAPI.
func NewMockGitAPI() *MockGitAPI {
	return &MockGitAPI{}
}

// CreateAnnotatedTag mocks creating a tag.
// Returns a mocked error.
func (mock *MockGitAPI) CreateAnnotatedTag(tag string) (err error) {
	args := mock.Called(tag)
	return args.Error(0)
}

// FetchTags mocks fetching tags.
// Returns a mocked error.
func (mock *MockGitAPI) FetchTags() (output string, err error) {
	args := mock.Called()
	return args.String(0), args.Error(1)
}

// FetchUnshallow mocks changing to an unshallow repo.
// Returns a mocked error.
func (mock *MockGitAPI) FetchUnshallow() (output string, err error) {
	args := mock.Called()
	return args.String(0), args.Error(0)
}

// GetConfig mocks getting a config.
// Returns a mocked config or a mocked error.
func (mock *MockGitAPI) GetConfig(key string) (value string, err error) {
	args := mock.Called(key)
	return args.String(0), args.Error(1)
}

// GetLatestAnnotatedTag mocks getting the latest annotated tag.
// Returns a mocked tag or a mocked error.
func (mock *MockGitAPI) GetLatestAnnotatedTag() (tag string, err error) {
	args := mock.Called()
	return args.String(0), args.Error(1)
}

// GetLatestCommitMessage mocks getting the latest commit message.
// Returns a mocked commit message or a mocked error.
func (mock *MockGitAPI) GetLatestCommitMessage() (message string, err error) {
	args := mock.Called()
	return args.String(0), args.Error(1)
}

// GetMergedBranchName mocks getting a merged branch name.
// Returns a mocked merged branch name or a mocked error.
func (mock *MockGitAPI) GetMergedBranchName() (name string, err error) {
	args := mock.Called()
	return args.String(0), args.Error(1)
}

// GetTags mocks getting all tags.
// Returns a mocked string of tags or a mocked error.
func (mock *MockGitAPI) GetTags() (tags string, err error) {
	args := mock.Called()
	return args.String(0), args.Error(1)
}

// PushTag pushes a fake tag.
// Returns a mocked error.
func (mock *MockGitAPI) PushTag(tag string) (err error) {
	args := mock.Called(tag)
	return args.Error(0)
}

// SetConfig mocks setting a config.
// Returns a mocked error.
func (mock *MockGitAPI) SetConfig(key string, value string) (err error) {
	args := mock.Called(key, value)
	return args.Error(0)
}

// SetConfigIfNotSet mocks setting a config if not set.
// Returns a mocked error.
func (mock *MockGitAPI) SetConfigIfNotSet(key string, value string) (actual string, err error) {
	args := mock.Called(key, value)
	return args.String(0), args.Error(1)
}
