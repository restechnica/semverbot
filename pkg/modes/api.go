package modes

type SemverMap map[string][]string

// API an API to work with different modes.
type API struct {
	GitBranchMode GitBranchMode
	GitCommitMode GitCommitMode
}

// NewAPI creates a new semver mode API with a mode detector to pass
// it on to the different modes that require it.
// Returns the new API.
func NewAPI(semverMap SemverMap, gitBranchDelimiters string, GitCommitDelimiters string) API {
	return API{
		GitBranchMode: NewGitBranchMode(gitBranchDelimiters, semverMap),
		GitCommitMode: NewGitCommitMode(GitCommitDelimiters, semverMap),
	}
}

// SelectMode selects the mode corresponding to the mode string.
// Returns the corresponding mode.
func (api API) SelectMode(mode string) Mode {
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
