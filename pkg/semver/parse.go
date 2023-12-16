package semver

import (
	blangsemver "github.com/blang/semver/v4"
	"strings"
)

// Parse parses a version string into a semver version struct.
// It tolerates certain version specifications that do not strictly adhere to semver specs.
// See the library documentation for more information.
// Returns the parsed blang/semver/v4 Version.
func Parse(prefix string, version string) (blangsemver.Version, error) {
	var mappedVersion = strings.Replace(version, prefix, "v", 1)
	return blangsemver.ParseTolerant(mappedVersion)
}
