package semver

import "fmt"

// Find finds a valid semver version in a slice of strings.
// Returns a version when found, otherwise an error stating no valid semver version has been found.
func Find(versions []string) (found string, err error) {
	for _, version := range versions {
		if _, err = Parse(version); err == nil {
			return version, err
		}
	}

	return found, fmt.Errorf("could not find a valid semver version")
}
