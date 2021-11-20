package cli

import "github.com/restechnica/semverbot/internal"

var (
	// DefaultConfig the default config.
	DefaultConfig = internal.DefaultConfig

	// DefaultConfigFilePath the default relative filepath to the config file.
	DefaultConfigFilePath = internal.DefaultConfigFilePath

	// DefaultGitBranchDelimiters the default delimiters used by the git-branch mode.
	DefaultGitBranchDelimiters = internal.DefaultGitBranchDelimiters

	// DefaultGitCommitDelimiters the default delimiters used by the git-commit mode.
	DefaultGitCommitDelimiters = internal.DefaultGitCommitDelimiters

	// DefaultVersion the default version when no other version can be found.
	DefaultVersion = internal.DefaultVersion
)
