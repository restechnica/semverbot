package semver

import (
	"strings"

	blangsemver "github.com/blang/semver/v4"
)

// Parse parses a version string into a semver version struct.
// It tolerates certain version specifications that do not strictly adhere to semver specs.
// See the library documentation for more information.
// Returns the parsed blang/semver/v4 Version.
func Parse(prefix string, suffix string, version string) (blangsemver.Version, error) {
	var versionWithoutPrefix = strings.Replace(version, prefix, "v", 1)
	var versionWithoutPrefixAndSuffix = strings.Replace(versionWithoutPrefix, suffix, "", 1)
	return blangsemver.ParseTolerant(versionWithoutPrefixAndSuffix)
}
