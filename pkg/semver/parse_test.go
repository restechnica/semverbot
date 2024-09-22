package semver

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	type Test struct {
		Name     string
		Major    string
		Minor    string
		Patch    string
		Prebuild string
		Prefix   string
		Suffix   string
	}

	var tests = []Test{
		{Name: "Default", Major: "0", Minor: "0", Patch: "0"},
		{Name: "Patch", Major: "0", Minor: "0", Patch: "1"},
		{Name: "Minor", Major: "0", Minor: "2", Patch: "0"},
		{Name: "Major", Major: "3", Minor: "0", Patch: "0"},
		{Name: "DiscardPrefix", Major: "1", Minor: "0", Patch: "0", Prefix: "v"},
		{Name: "DiscardSuffix", Major: "2", Minor: "0", Patch: "0", Suffix: "a"},
		{Name: "DiscardSuffixAlt", Major: "2", Minor: "0", Patch: "0", Suffix: "-alt"},
		{Name: "KeepPrebuild", Major: "2", Minor: "0", Patch: "0", Suffix: "", Prebuild: "-pre+001"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var version = fmt.Sprintf(`%s%s.%s.%s%s%s`, test.Prefix, test.Major, test.Minor, test.Patch,
				test.Suffix, test.Prebuild)

			var got, err = Parse(test.Prefix, test.Suffix, version)

			assert.Equal(t, test.Major, fmt.Sprint(got.Major), `want: "%s", got: "%d"`, test.Major, got.Major)
			assert.Equal(t, test.Minor, fmt.Sprint(got.Minor), `want: "%s", got: "%s"`, test.Minor, got.Minor)
			assert.Equal(t, test.Patch, fmt.Sprint(got.Patch), `want: "%s", got: "%s"`, test.Patch, got.Patch)

			if test.Prefix != "" {
				assert.False(t, strings.HasPrefix(got.String(), test.Prefix))
			}

			if test.Suffix != "" {
				assert.False(t, strings.HasSuffix(got.String(), test.Suffix))
			}

			if test.Prebuild != "" {
				assert.True(t, strings.HasSuffix(got.String(), test.Prebuild))
			}

			assert.NoError(t, err)
		})
	}
}
