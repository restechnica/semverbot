package semver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	type Test struct {
		Name      string
		Prefix    string
		Suffix    string
		Versions  []string
		WantIndex int
	}

	var tests = []Test{
		{Name: "FindVersionIfValid", Prefix: "v", Suffix: "", Versions: []string{"v1.0.1", "v0.1.1", "v0.1.0"}, WantIndex: 0},
		{Name: "FindVersionWithCustomPrefixIfValid", Prefix: "test-", Suffix: "", Versions: []string{"test-1.0.1", "test-0.1.1", "test-0.1.0"}, WantIndex: 0},
		{Name: "FindVersionWithCustomSuffixIfValid", Prefix: "v", Suffix: "a", Versions: []string{"v1.0.1a", "v0.1.1a", "v0.1.0a"}, WantIndex: 0},
		{Name: "SkipVersionIfInvalid", Prefix: "v", Suffix: "", Versions: []string{"invalid1", "invalid2", "v0.1.0"}, WantIndex: 2},
		{Name: "FindVersionWhenDifferentOrder", Prefix: "v", Suffix: "", Versions: []string{"v1.3.1", "v0.2.0", "v2.3.0"}, WantIndex: 2},
		{Name: "FindVersionWhenMultiplePrefixes", Prefix: "v", Suffix: "", Versions: []string{"v1.3.1", "v0.2.0", "2.3.0"}, WantIndex: 2},
		{Name: "FindVersionWhenMultipleSuffixes", Prefix: "v", Suffix: "n", Versions: []string{"v1.3.1a", "v0.2.0-alt", "v2.3.0n"}, WantIndex: 2},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var versions = test.Versions
			var want = versions[test.WantIndex]
			var got, err = Find(test.Prefix, test.Suffix, versions)

			assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
			assert.NoError(t, err)
		})
	}

	type ErrorTest struct {
		Name     string
		Versions []string
	}

	var errorTests = []ErrorTest{
		{Name: "ReturnErrorOnInvalidVersions", Versions: []string{"invalid", "semver", "versions"}},
		{Name: "ReturnErrorOnNoVersions", Versions: []string{}},
	}

	for _, test := range errorTests {
		t.Run(test.Name, func(t *testing.T) {
			var _, got = Find("v", "a", test.Versions)
			assert.Error(t, got)
		})
	}
}
