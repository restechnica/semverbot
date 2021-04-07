package semver

import (
	"fmt"
	"strings"
)

type ModeDetector struct {
	ModeDetectionMap map[string][]string
}

func NewModeDetector(modeDetectionMap map[string][]string) ModeDetector {
	return ModeDetector{ModeDetectionMap: modeDetectionMap}
}

// DetectMode detects which semver level mode (patch, minor, major) based
// on a string and the ModeDetectionMap.
// Returns the detected mode or an error if no mode was detected.
func (detector ModeDetector) DetectMode(target string) (detected Mode, err error) {
	for mode, substrings := range detector.ModeDetectionMap {
		for _, substring := range substrings {
			if strings.Contains(target, substring) {
				switch mode {
				case Patch:
					detected = NewPatchMode()
				case Minor:
					detected = NewMinorMode()
				case Major:
					detected = NewMajorMode()
				}
				return detected, err
			}
		}
	}

	return detected, fmt.Errorf(`failed to detect mode from "%s"`, target)
}
