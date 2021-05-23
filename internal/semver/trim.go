package semver

import blangsemver "github.com/blang/semver/v4"

func Trim(version string) (string, error) {
	var semverVersion blangsemver.Version
	var err error

	if semverVersion, err = Parse(version); err != nil {
		return version, err
	}

	return semverVersion.FinalizeVersion(), err
}
