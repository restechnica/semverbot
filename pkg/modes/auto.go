package modes

// Auto mode name for AutoMode.
const Auto = "auto"

// AutoMode implementation of the Mode interface.
// It makes use of several modes and defaults to PatchMode as a last resort.
type AutoMode struct {
	Modes []Mode
}

// NewAutoMode creates a new AutoMode.
// The order of modes in the modes slices is important and determines in which order the modes are applied in AutoMode.Increment.
// Returns the new AutoMode.
func NewAutoMode(modes []Mode) AutoMode {
	return AutoMode{Modes: modes}
}

// Increment increments a given version using AutoMode.
// It will attempt to increment the target version with its internal modes and defaults to PatchMode as a last resort.
// Returns the incremented version or an error if anything went wrong.
func (autoMode AutoMode) Increment(targetVersion string) (nextVersion string, err error) {
	for _, mode := range autoMode.Modes {
		if nextVersion, err = mode.Increment(targetVersion); err == nil {
			return nextVersion, err
		}
	}

	return PatchMode{}.Increment(targetVersion)
}
