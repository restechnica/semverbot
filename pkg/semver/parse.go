package semver

import blangsemver "github.com/blang/semver/v4"

// Parse parses a version string into a semver version struct.
// It tolerates certain version specifications that do not strictly adhere to semver specs.
// See the library documentation for more information.
// Returns the parsed blang/semver/v4 Version.
func Parse(version string) (blangsemver.Version, error) {
	return blangsemver.ParseTolerant(version)
}
