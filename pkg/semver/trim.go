package semver

import blangsemver "github.com/blang/semver/v4"

// Trim trims a semver version string of anything but major.minor.patch information.
// Returns the trimmed semver version.
func Trim(version string) (string, error) {
	var semverVersion blangsemver.Version
	var err error

	if semverVersion, err = Parse(version); err != nil {
		return version, err
	}

	return semverVersion.FinalizeVersion(), err
}
