package semver

import (
	"strings"

	blangsemver "github.com/blang/semver/v4"
)

// Trim trims a semver version string of anything but major.minor.patch information.
// Returns the trimmed semver version.
func Trim(prefix string, version string) (string, error) {
	var semverVersion blangsemver.Version
	var err error

	var mappedVersion = strings.Replace(version, prefix, prefix, 1)

	if semverVersion, err = Parse(prefix, mappedVersion); err != nil {
		return version, err
	}

	return semverVersion.FinalizeVersion(), err
}
