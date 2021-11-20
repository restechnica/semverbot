package internal

const (
	// DefaultConfig the default config.
	DefaultConfig = `mode = "auto"

[git]

[git.config]
email = "semverbot@github.com"
name = "semverbot"

[git.tags]
prefix = "v"

[semver]
patch = ["fix", "bug"]
minor = ["feature"]
major = ["release"]

[modes]

[modes.git-branch]
delimiters = "/"

[modes.git-commit]
delimiters = "[]"

`

	// DefaultConfigFilePath the default relative filepath to the config file.
	DefaultConfigFilePath = ".semverbot.toml"

	// DefaultGitBranchDelimiters the default delimiters used by the git-branch mode.
	DefaultGitBranchDelimiters = "/"

	// DefaultGitCommitDelimiters the default delimiters used by the git-commit mode.
	DefaultGitCommitDelimiters = "[]"

	// DefaultVersion the default version when no other version can be found.
	DefaultVersion = "0.0.0"
)
