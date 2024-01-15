package modes

import (
	blangsemver "github.com/blang/semver/v4"
	"github.com/restechnica/semverbot/pkg/semver"
)

// Major semver version level for major
const Major = "major"

// MajorMode implementation of the Mode interface.
// It makes use of the major level of semver versions.
type MajorMode struct{}

// NewMajorMode creates a new MajorMode.
// Returns the new MajorMode.
func NewMajorMode() MajorMode {
	return MajorMode{}
}

// Increment increments a given version using the MajorMode.
// Returns the incremented version.
func (mode MajorMode) Increment(prefix string, targetVersion string) (nextVersion string, err error) {
	var version blangsemver.Version

	if version, err = semver.Parse(prefix, targetVersion); err != nil {
		return
	}

	// at point of writing IncrementMajor always returns a nil value error
	_ = version.IncrementMajor()

	return version.FinalizeVersion(), err
}

// String returns a string representation of an instance.
func (mode MajorMode) String() string {
	return Major
}
