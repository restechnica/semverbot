package config

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYAMLLoader_Load(t *testing.T) {
	type Test struct {
		File string
		Name string
		Want Root
	}

	var tests = []Test{
		{Name: "LoadEmptyYAML", File: "../../../test/configs/empty.yaml", Want: Root{}},
		{Name: "LoadFullYAML", File: "../../../test/configs/full.yaml", Want: Root{
			Env: Env{
				Files:   []string{"./some-path/some-file"},
				Scripts: []EnvScript{{Bin: "node", Path: "./some-path/some-script.js"}},
				Vars:    map[string]string{"SOME_VAR": "some_value", "ANOTHER_VAR": "another_value"},
			},
			Git: Git{
				Config:    map[string]string{"user.email": "github-actions@github.com", "user.name": "github-actions"},
				Unshallow: true,
			},
			Semver: Semver{
				Bin:      "python3",
				Matches:  map[string]string{"fix/": "patch", "feature/": "minor", "release/": "major"},
				Path:     "./some-path/some-script.py",
				Strategy: "git-commit",
			},
		}},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var want = test.Want

			var data, _ = ioutil.ReadFile(test.File)

			var loader = NewYAMLLoader()
			loader.readFile = func(path string) ([]byte, error) {
				return data, nil
			}

			var got, err = loader.Load("fake-path")

			assert.NoError(t, err)
			assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
		})
	}

	type ErrorTest struct {
		Name           string
		ReadFileError  error
		UnmarshalError error
	}

	var errorTests = []ErrorTest{
		{Name: "ReturnErrorOnReadFileError", ReadFileError: fmt.Errorf("some-error"), UnmarshalError: nil},
		{Name: "ReturnErrorOnUnmarshalError", ReadFileError: nil, UnmarshalError: fmt.Errorf("some-error")},
	}

	for _, test := range errorTests {
		t.Run(test.Name, func(t *testing.T) {
			var loader = NewYAMLLoader()
			loader.readFile = func(path string) ([]byte, error) {
				return nil, test.ReadFileError
			}
			loader.unmarshal = func(data []byte, config *Root) (err error) {
				return test.UnmarshalError
			}

			var _, got = loader.Load("fake-path")
			assert.Error(t, got)
		})
	}
}

func TestYAMLLoader_Overload(t *testing.T) {
	type Test struct {
		File     string
		Name     string
		Original Root
		Want     Root
	}

	var original = Root{Git: Git{Unshallow: true}, Env: Env{Files: []string{"some-script"}}}

	var tests = []Test{
		{Name: "OverloadEmptyYAML", File: "../../../test/configs/empty.yaml", Original: original, Want: Root{}},
		{Name: "OverloadYAML", File: "../../../test/configs/overload.yaml", Want: Root{
			Env:    Env{Files: []string{"./some-path/some-file"}},
			Git:    Git{Unshallow: false},
			Semver: Semver{Strategy: "patch"},
		}},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var want = test.Want

			var data, _ = ioutil.ReadFile(test.File)

			var loader = NewYAMLLoader()
			loader.readFile = func(path string) ([]byte, error) {
				return data, nil
			}

			var got, err = loader.Overload("fake-path", test.Original)

			assert.NoError(t, err)
			assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
		})
	}

	type ErrorTest struct {
		Name           string
		ReadFileError  error
		UnmarshalError error
	}

	var errorTests = []ErrorTest{
		{Name: "ReturnErrorOnReadFileError", ReadFileError: fmt.Errorf("some-error"), UnmarshalError: nil},
		{Name: "ReturnErrorOnUnmarshalError", ReadFileError: nil, UnmarshalError: fmt.Errorf("some-error")},
	}

	for _, test := range errorTests {
		t.Run(test.Name, func(t *testing.T) {
			var loader = NewYAMLLoader()
			loader.readFile = func(path string) ([]byte, error) {
				return nil, test.ReadFileError
			}
			loader.unmarshal = func(data []byte, config *Root) (err error) {
				return test.UnmarshalError
			}

			var _, got = loader.Overload("fake-path", Root{})
			assert.Error(t, got)
		})
	}
}
