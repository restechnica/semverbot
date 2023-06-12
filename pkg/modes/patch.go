package modes

import (
	blangsemver "github.com/blang/semver/v4"

	"github.com/restechnica/semverbot/pkg/semver"
)

// Patch semver version level for patch
const Patch = "patch"

// PatchMode implementation of the Mode interface.
// It makes use of the patch level of semver versions.
type PatchMode struct{}

// NewPatchMode creates a new PatchMode.
// Returns the new PatchMode.
func NewPatchMode() PatchMode {
	return PatchMode{}
}

// Increment increments a given version using the PatchMode.
// Returns the incremented version.
func (mode PatchMode) Increment(targetVersion string) (nextVersion string, err error) {
	var version blangsemver.Version

	if version, err = semver.Parse(targetVersion); err != nil {
		return
	}

	// at point of writing IncrementPatch always returns a nil value error
	_ = version.IncrementPatch()

	return version.FinalizeVersion(), err
}

// String returns a string representation of an instance.
func (mode PatchMode) String() string {
	return Patch
}
