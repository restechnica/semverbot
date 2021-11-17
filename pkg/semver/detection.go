package semver

import (
	"fmt"
	"strings"
)

// ModeDetector detects which mode should be applied.
type ModeDetector struct {
	SemverMatchMap map[string][]string
}

// NewModeDetector creates a new ModeDetector with a detection configuration.
// Returns the new ModeDetector.
func NewModeDetector(semverMatchMap map[string][]string) ModeDetector {
	return ModeDetector{SemverMatchMap: semverMatchMap}
}

// DetectMode detects the semver level mode (patch, minor, major) based
// on a string and the SemverMatchMap.
// Returns the detected mode or an error if no mode was detected.
func (detector ModeDetector) DetectMode(target string) (detected Mode, err error) {
	for mode, substrings := range detector.SemverMatchMap {
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
