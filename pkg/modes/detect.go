package modes

import (
	"fmt"

	"github.com/restechnica/semverbot/internal/util"
	"github.com/restechnica/semverbot/pkg/semver"
)

func DetectModeFromString(str string, semverMap semver.Map, delimiters string) (detected Mode, err error) {
	var modes []Mode

	if modes, err = DetectModesFromString(str, semverMap, delimiters); err != nil {
		return nil, err
	}

	if len(modes) == 0 {
		return nil, fmt.Errorf(`failed to detect mode from string '%s' with delimiters '%s'`, str, delimiters)
	}

	var priority = map[string]int{
		Patch: 1,
		Minor: 2,
		Major: 3,
	}

	for _, mode := range modes {
		if detected == nil {
			detected = mode
		}

		if priority[mode.String()] > priority[detected.String()] {
			detected = mode
		}
	}

	return detected, err
}

// DetectModesFromString detects multiple modes based on a string.
// Mode detection is limited to PatchMode, MinorMode, MajorMode.
// The order of a detected modes is relative to their position in the string.
// Returns a slice of the detected modes.
func DetectModesFromString(str string, semverMap semver.Map, delimiters string) (detected []Mode, err error) {
	var substrings = util.SplitByDelimiterString(str, delimiters)

	for _, substring := range substrings {
		for level, values := range semverMap {
			if util.SliceContainsString(values, substring) {
				switch level {
				case Patch:
					detected = append(detected, NewPatchMode())
				case Minor:
					detected = append(detected, NewMinorMode())
				case Major:
					detected = append(detected, NewMajorMode())
				default:
					return nil, fmt.Errorf("failed to detect mode due to unsupported semver level: '%s'", level)
				}
			}
		}
	}

	return detected, err
}
