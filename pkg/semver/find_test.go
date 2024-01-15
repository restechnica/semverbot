package semver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	type Test struct {
		Name      string
		Prefix    string
		Versions  []string
		WantIndex int
	}

	var tests = []Test{
		{Name: "FindVersionIfValid", Prefix: "v", Versions: []string{"v1.0.1", "v0.1.1", "v0.1.0"}, WantIndex: 0},
		{Name: "FindVersionWithCustomPrefixIfValid", Prefix: "test-", Versions: []string{"test-1.0.1", "test-0.1.1", "test-0.1.0"}, WantIndex: 0},
		{Name: "SkipVersionIfInvalid", Prefix: "v", Versions: []string{"invalid1", "invalid2", "v0.1.0"}, WantIndex: 2},
		{Name: "FindVersionWhenDifferentOrder", Prefix: "v", Versions: []string{"v1.3.1", "v0.2.0", "v2.3.0"}, WantIndex: 2},
		{Name: "FindVersionWhenMultiplePrefixes", Prefix: "v", Versions: []string{"v1.3.1", "v0.2.0", "2.3.0"}, WantIndex: 2},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var versions = test.Versions
			var want = versions[test.WantIndex]
			var got, err = Find(test.Prefix, versions)

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
			var _, got = Find("v", test.Versions)
			assert.Error(t, got)
		})
	}
}
