package modes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/restechnica/semverbot/pkg/semver"
)

func TestDetectModeFromString(t *testing.T) {
	var semverMap = semver.Map{
		Patch: {"fix", "bug", "patch"},
		Minor: {"feature", "feat", "minor"},
		Major: {"release", "major"},
	}

	type Test struct {
		String     string
		Delimiters string
		Name       string
		SemverMap  semver.Map
		Want       Mode
	}

	var tests = []Test{
		{Name: "DetectPatchModeWithSlash", String: "fix/some-bug", Delimiters: "/", SemverMap: semverMap, Want: NewPatchMode()},
		{Name: "DetectPatchModeWithSquareBrackets", String: "[bug] some fix", Delimiters: "[]", SemverMap: semverMap, Want: NewPatchMode()},
		{Name: "DetectMinorModeWithSlash", String: "feature/some-bug", Delimiters: "/", SemverMap: semverMap, Want: NewMinorMode()},
		{Name: "DetectMinorModeWithRoundBrackets", String: "feat(subject): some changes", Delimiters: "():", SemverMap: semverMap, Want: NewMinorMode()},
		{Name: "DetectMinorModeWithSquareBrackets", String: "[feature] some changes", Delimiters: "[]", SemverMap: semverMap, Want: NewMinorMode()},
		{Name: "DetectMajorModeWithSlash", String: "release/some-bug", Delimiters: "/", SemverMap: semverMap, Want: NewMajorMode()},
		{Name: "DetectMinorWithMultipleModes", String: "some [fix] and release/feat(subject)", Delimiters: "()[]/", SemverMap: semverMap, Want: NewMinorMode()},
		{Name: "DetectMajorWithMultipleModes0", String: "[fix] some [feature] and [release]", Delimiters: "[]", SemverMap: semverMap, Want: NewMajorMode()},
		{Name: "DetectMajorWithMultipleModes1", String: "[release] some [fix] test [feature]", Delimiters: "[]", SemverMap: semverMap, Want: NewMajorMode()},
		{Name: "DetectMajorWithMultipleModes2", String: "[feature] some [release] test [fix]", Delimiters: "[]", SemverMap: semverMap, Want: NewMajorMode()},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var got, err = DetectModeFromString(test.String, test.SemverMap, test.Delimiters)

			assert.NoError(t, err)
			assert.IsType(t, test.Want, got, `want: '%s, got: '%s'`, test.Want, got)
		})
	}

	type ErrorTest struct {
		String     string
		Delimiters string
		Error      error
		Name       string
		SemverMap  semver.Map
	}

	var errorTests = []ErrorTest{
		{
			Name:       "DetectNothingWithEmptySemverMap",
			String:     "[feature] some changes",
			Delimiters: "[]",
			Error:      fmt.Errorf(`failed to detect mode from string '[feature] some changes' with delimiters '[]'`),
			SemverMap:  semver.Map{},
		},
		{
			Name:       "DetectNothingWithEmptyDelimiters",
			String:     "[feature] some changes",
			Delimiters: "",
			Error:      fmt.Errorf(`failed to detect mode from string '[feature] some changes' with delimiters ''`),
			SemverMap:  semverMap,
		},
		{
			Name:       "DetectNothingWithEmptyCommitMessage",
			String:     "",
			Delimiters: "/",
			Error:      fmt.Errorf(`failed to detect mode from string '' with delimiters '/'`),
			SemverMap:  semverMap,
		},
		{
			Name:       "DetectNothingWithFaultySemverMap",
			String:     "[feature] some changes",
			Delimiters: "[]",
			Error:      fmt.Errorf(`failed to detect mode due to unsupported semver level: 'mnr'`),
			SemverMap: semver.Map{
				"mnr": {"feature"},
			},
		},
	}

	for _, test := range errorTests {
		t.Run(test.Name, func(t *testing.T) {
			var _, got = DetectModeFromString(test.String, test.SemverMap, test.Delimiters)

			assert.Error(t, got)
			assert.Equal(t, test.Error, got, `want: '%s', got: '%s'`, test.Error, got)
		})
	}
}

func TestDetectModesFromString(t *testing.T) {
	var semverMap = semver.Map{
		Patch: {"fix", "bug", "patch"},
		Minor: {"feature", "feat", "minor"},
		Major: {"release", "major"},
	}

	type Test struct {
		String     string
		Delimiters string
		Name       string
		SemverMap  semver.Map
		Want       []Mode
	}

	var tests = []Test{
		{Name: "DetectPatchModeWithSlash", String: "fix/some-bug", Delimiters: "/", SemverMap: semverMap, Want: []Mode{NewPatchMode()}},
		{Name: "DetectPatchModeWithSquareBrackets", String: "[bug] some fix", Delimiters: "[]", SemverMap: semverMap, Want: []Mode{NewPatchMode()}},
		{Name: "DetectMinorModeWithSlash", String: "feature/some-bug", Delimiters: "/", SemverMap: semverMap, Want: []Mode{NewMinorMode()}},
		{Name: "DetectMinorModeWithRoundBrackets", String: "feat(subject): some changes", Delimiters: "():", SemverMap: semverMap, Want: []Mode{NewMinorMode()}},
		{Name: "DetectMinorModeWithSquareBrackets", String: "[feature] some changes", Delimiters: "[]", SemverMap: semverMap, Want: []Mode{NewMinorMode()}},
		{Name: "DetectMajorModeWithSlash", String: "release/some-bug", Delimiters: "/", SemverMap: semverMap, Want: []Mode{NewMajorMode()}},
		{Name: "DetectMultipleModes0", String: "[fix] some [feature] and [release]", Delimiters: "[]", SemverMap: semverMap, Want: []Mode{NewPatchMode(), NewMinorMode(), NewMajorMode()}},
		{Name: "DetectMultipleModes1", String: "some [feature] and release/test and fix(subject)", Delimiters: "()[]/", SemverMap: semverMap, Want: []Mode{NewMinorMode()}},
		{Name: "DetectMultipleModes2", String: "some [fix] and release/test and feat(subject)", Delimiters: "()[]/", SemverMap: semverMap, Want: []Mode{NewPatchMode()}},
		{Name: "DetectMultipleModes3", String: "some [fix] and release/feat(subject)", Delimiters: "()[]/", SemverMap: semverMap, Want: []Mode{NewPatchMode(), NewMinorMode()}},
		{Name: "DetectMultipleModesInOrder0", String: "[release] some [fix] test [feature]", Delimiters: "[]", SemverMap: semverMap, Want: []Mode{NewMajorMode(), NewPatchMode(), NewMinorMode()}},
		{Name: "DetectMultipleModesInOrder1", String: "[feature] some [release] test [fix]", Delimiters: "[]", SemverMap: semverMap, Want: []Mode{NewMinorMode(), NewMajorMode(), NewPatchMode()}},
		{Name: "DetectNothingWithEmptySemverMap", String: "feature/some-feature", Delimiters: "/", SemverMap: semver.Map{}, Want: []Mode{}},
		{Name: "DetectNothingWithEmptyDelimiters", String: "feature/some-feature", Delimiters: "", SemverMap: semverMap, Want: []Mode{}},
		{Name: "DetectNothingWithEmptyString", String: "", Delimiters: "/", SemverMap: semverMap, Want: []Mode{}},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var got, err = DetectModesFromString(test.String, test.SemverMap, test.Delimiters)

			assert.NoError(t, err)

			assert.Equal(t, len(test.Want), len(got), `want: '%s, got: '%s'`, test.Want, got)

			for i, mode := range got {
				assert.IsType(t, test.Want[i], mode, `want: '%s, got: '%s'`, test.Want, got)
			}
		})
	}

	type ErrorTest struct {
		String     string
		Delimiters string
		Error      error
		Name       string
		SemverMap  semver.Map
	}

	var errorTests = []ErrorTest{
		{
			Name:       "DetectNothingWithFaultySemverMap",
			String:     "feature/some-feature",
			Delimiters: "/",
			Error:      fmt.Errorf(`failed to detect mode due to unsupported semver level: 'mnr'`),
			SemverMap: semver.Map{
				"mnr": {"feature"},
			},
		},
	}

	for _, test := range errorTests {
		t.Run(test.Name, func(t *testing.T) {
			var _, got = DetectModesFromString(test.String, test.SemverMap, test.Delimiters)

			assert.Error(t, got)
			assert.Equal(t, test.Error, got, `want: '%s, got: '%s'`, test.Error, got)
		})
	}
}
