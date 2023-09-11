package semver

import (
	"fmt"
	"strings"

	blangsemver "github.com/blang/semver/v4"
)

// Find finds the biggest valid semver version in a slice of strings.
// The initial order of the versions does not matter.
// Returns the biggest valid semver version if found, otherwise an error stating no valid semver version has been found.
func Find(versions []string) (found string, err error) {
	var parsedVersions blangsemver.Versions
	var parsedVersion blangsemver.Version

	for _, version := range versions {
		if parsedVersion, err = Parse(version); err != nil {
			continue
		}

		parsedVersions = append(parsedVersions, parsedVersion)
	}

	if len(parsedVersions) == 0 {
		return found, fmt.Errorf("could not find a valid semver version")
	}

	blangsemver.Sort(parsedVersions)

	var targetVersion = parsedVersions[len(parsedVersions)-1]

	// necessary because blangsemver's Version.String() strips any prefix
	for _, version := range versions {
		if strings.Contains(version, targetVersion.String()) {
			found = version
			break
		}
	}

	return found, nil
}
