package internal

import "github.com/restechnica/semverbot/pkg/modes"

const (
	// DefaultConfigFilePath the default relative filepath to the config file.
	DefaultConfigFilePath = ".semverbot.toml"

	// DefaultGitBranchDelimiters the default delimiters used by the git-branch mode.
	DefaultGitBranchDelimiters = "/"

	// DefaultGitCommitDelimiters the default delimiters used by the git-commit mode.
	DefaultGitCommitDelimiters = "[]/"

	// DefaultGitTagsPrefix the default prefix prepended to git tags.
	DefaultGitTagsPrefix = "v"

	// DefaultGitTagsSuffix the default prefix prepended to git tags.
	DefaultGitTagsSuffix = ""

	// DefaultMode the default mode for incrementing versions.
	DefaultMode = modes.Auto

	// DefaultVersion the default version when no other version can be found.
	DefaultVersion = "0.0.0"
)
