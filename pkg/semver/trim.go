package semver

import (
	"strings"

	blangsemver "github.com/blang/semver/v4"
)

// Trim trims a semver version string of anything but major.minor.patch information.
// Returns the trimmed semver version.
func Trim(prefix string, suffix string, version string) (string, error) {
	var semverVersion blangsemver.Version
	var err error

	var versionWithoutPrefix = strings.Replace(version, prefix, prefix, 1)
	var versionWithoutPrefixOrSuffix = strings.Replace(versionWithoutPrefix, suffix, suffix, 1)

	if semverVersion, err = Parse(prefix, suffix, versionWithoutPrefixOrSuffix); err != nil {
		return version, err
	}

	return semverVersion.FinalizeVersion(), err
}
