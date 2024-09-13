package modes

import (
	blangsemver "github.com/blang/semver/v4"

	"github.com/restechnica/semverbot/pkg/semver"
)

// Minor semver version level for minor
const Minor = "minor"

// MinorMode implementation of the Mode interface.
// It makes use of the minor level of semver versions.
type MinorMode struct{}

// NewMinorMode creates a new MinorMode.
// Returns the new MinorMode.
func NewMinorMode() MinorMode {
	return MinorMode{}
}

// Increment increments a given version using the MinorMode.
// Returns the incremented version.
func (mode MinorMode) Increment(prefix string, suffix string, targetVersion string) (nextVersion string, err error) {
	var version blangsemver.Version

	if version, err = semver.Parse(prefix, suffix, targetVersion); err != nil {
		return
	}

	// at point of writing IncrementMinor always returns a nil value error
	_ = version.IncrementMinor()

	return version.FinalizeVersion(), err
}

// String returns a string representation of an instance.
func (mode MinorMode) String() string {
	return Minor
}
