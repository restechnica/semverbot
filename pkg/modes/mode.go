package modes

// Mode interface which increments a specific semver level.
type Mode interface {
	Increment(prefix string, targetVersion string) (nextVersion string, err error)
}
