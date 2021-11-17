package cli

const (
	// DefaultConfig the default config.
	DefaultConfig = `[git]

[git.config]
email = "semverbot@github.com"
name = "semverbot"

[git.tags]
prefix = "v"

[semver]
mode = "auto"

[semver.match]
patch = ["fix/", "[fix]"]
minor = ["feature/", "[feature]"]
major = ["release/", "[release]"]

`

	// DefaultConfigFilePath the default relative filepath to the config file.
	DefaultConfigFilePath = ".semverbot.toml"

	// DefaultVersion the default version when no other version can be found.
	DefaultVersion = "0.0.0"
)
