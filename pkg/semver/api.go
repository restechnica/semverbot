package semver

// ModeAPI an API to work with different modes.
type ModeAPI struct {
	GitBranchMode GitBranchMode
	GitCommitMode GitCommitMode
}

// NewModeAPI creates a new semver mode API with a mode detector to pass
// it on to the different modes that require it.
// Returns the new ModeAPI.
func NewModeAPI(detector ModeDetector) ModeAPI {
	return ModeAPI{
		GitBranchMode: NewGitBranchMode(detector),
		GitCommitMode: NewGitCommitMode(detector),
	}
}

// SelectMode selects the mode corresponding to the mode string.
// Returns the corresponding mode.
func (api ModeAPI) SelectMode(mode string) Mode {
	switch mode {
	case Auto:
		return NewAutoMode([]Mode{api.GitBranchMode, api.GitCommitMode})
	case GitCommit:
		return api.GitCommitMode
	case GitBranch:
		return api.GitBranchMode
	case Patch:
		return NewPatchMode()
	case Minor:
		return NewMinorMode()
	case Major:
		return NewMajorMode()
	default:
		return NewPatchMode()
	}
}
