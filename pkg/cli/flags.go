package cli

var (
	// ConfigFlag a flag which configures the config file location.
	ConfigFlag string

	// DebugFlag a flag which sets the log level verbosity to Debug if true
	DebugFlag bool

	// ModeFlag a flag which indicates the semver mode to increment the current version with.
	ModeFlag string

	// VerboseFlag a flag which increases log level verbosity to Info if true
	VerboseFlag bool
)
