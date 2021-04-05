package semver

import (
	blangsemver "github.com/blang/semver/v4"
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

	if version, err = blangsemver.Parse(targetVersion); err != nil {
		return
	}

	// at point of writing IncrementPatch always returns a nil value error
	_ = version.IncrementPatch()

	return version.FinalizeVersion(), err
}
