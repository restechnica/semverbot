package semver

import blangsemver "github.com/blang/semver/v4"

func Parse(version string) (blangsemver.Version, error) {
	return blangsemver.ParseTolerant(version)
}
