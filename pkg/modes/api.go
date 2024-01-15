package modes

import "github.com/rs/zerolog/log"

// API an API to work with different modes.
type API struct {
	GitBranchMode GitBranchMode
	GitCommitMode GitCommitMode
}

// NewAPI creates a new semver mode API.
// Returns the new API.
func NewAPI(gitBranchMode GitBranchMode, gitCommitMode GitCommitMode) API {
	return API{
		GitBranchMode: gitBranchMode,
		GitCommitMode: gitCommitMode,
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
		log.Warn().Msg("mode invalid, falling back to patch mode")
		return NewPatchMode()
	}
}
