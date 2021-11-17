package semver

// Mode interface which increments a specific semver level.
type Mode interface {
	Increment(targetVersion string) (nextVersion string, err error)
}
