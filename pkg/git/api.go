package git

// API interface to interact with git.
type API interface {
	CreateAnnotatedTag(tag string) (err error)
	FetchTags() (err error)
	FetchUnshallow() (err error)
	GetConfig(key string) (value string, err error)
	GetLatestAnnotatedTag() (tag string, err error)
	GetLatestCommitMessage() (message string, err error)
	GetMergedBranchName() (name string, err error)
	PushTag(tag string) (err error)
	SetConfig(key string, value string) (err error)
	SetConfigIfNotSet(key string, value string) (err error)
}
