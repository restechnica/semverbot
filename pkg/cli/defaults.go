package cli

import (
	"fmt"

	"github.com/restechnica/semverbot/internal"
)

var (
	// DefaultAdditionalConfigFilePaths additional default relative filepaths to the config file.
	DefaultAdditionalConfigFilePaths = []string{".sbot.toml", ".semverbot/config.toml", ".sbot/config.toml"}

	// DefaultConfigFilePath the default relative filepath to the config file.
	DefaultConfigFilePath = internal.DefaultConfigFilePath

	// DefaultGitBranchDelimiters the default delimiters used by the git-branch mode.
	DefaultGitBranchDelimiters = internal.DefaultGitBranchDelimiters

	// DefaultGitCommitDelimiters the default delimiters used by the git-commit mode.
	DefaultGitCommitDelimiters = internal.DefaultGitCommitDelimiters

	// DefaultGitTagsPrefix the default prefix prepended to git tags.
	DefaultGitTagsPrefix = internal.DefaultGitTagsPrefix

	// DefaultMode the default mode for incrementing versions.
	DefaultMode = internal.DefaultMode

	// DefaultVersion the default version when no other version can be found.
	DefaultVersion = internal.DefaultVersion
)

func GetDefaultConfig() string {
	const template = `mode = "%s"

[git]

[git.config]
email = "semverbot@github.com"
name = "semverbot"

[git.tags]
prefix = "%s"

[semver]
patch = ["fix", "bug"]
minor = ["feature"]
major = ["release"]

[modes]

[modes.git-branch]
delimiters = "%s"

[modes.git-commit]
delimiters = "%s"
`

	return fmt.Sprintf(
		template,
		DefaultMode,
		DefaultGitTagsPrefix,
		DefaultGitBranchDelimiters,
		DefaultGitCommitDelimiters,
	)
}
