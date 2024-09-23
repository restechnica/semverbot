package modes

// Mode interface which increments a specific semver level.
type Mode interface {
	Increment(prefix string, suffix string, targetVersion string) (nextVersion string, err error)
	String() string
}
