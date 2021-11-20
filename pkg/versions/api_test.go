package versions

import (
	"fmt"
	"github.com/restechnica/semverbot/pkg/cli"
	"testing"

	"github.com/restechnica/semverbot/internal/mocks"
	"github.com/restechnica/semverbot/pkg/git"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAPI_GetVersion(t *testing.T) {
	type Test struct {
		Name    string
		Version string
	}

	var tests = []Test{
		{Name: "ReturnVersion", Version: "0.0.0"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var cmder = mocks.NewMockCommander()
			cmder.On("Output", mock.Anything, mock.Anything).Return(test.Version, nil)

			var gitAPI = git.API{Commander: cmder}
			var versionAPI = API{GitAPI: gitAPI}

			var got, err = versionAPI.GetVersion()

			assert.NoError(t, err)
			assert.Equal(t, test.Version, got, `want: "%s, got: "%s"`, test.Version, got)
		})
	}

	type ErrorTest struct {
		Error error
		Name  string
	}

	var errorTests = []ErrorTest{
		{Name: "ReturnErrorOnGitError", Error: fmt.Errorf("some-error")},
	}

	for _, test := range errorTests {
		t.Run(test.Name, func(t *testing.T) {
			var cmder = mocks.NewMockCommander()
			cmder.On("Output", mock.Anything, mock.Anything).Return("", test.Error)

			var gitAPI = git.API{Commander: cmder}
			var versionAPI = API{GitAPI: gitAPI}

			var _, got = versionAPI.GetVersion()
			assert.Error(t, got)
		})
	}
}

func TestAPI_GetVersionOrDefault(t *testing.T) {
	type Test struct {
		Name    string
		Version string
	}

	var tests = []Test{
		{Name: "ReturnVersionWithoutError", Version: "0.0.0"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var cmder = mocks.NewMockCommander()
			cmder.On("Output", mock.Anything, mock.Anything).Return(test.Version, nil)

			var gitAPI = git.API{Commander: cmder}
			var versionAPI = API{GitAPI: gitAPI}

			var got, err = versionAPI.GetVersion()

			assert.NoError(t, err)
			assert.Equal(t, test.Version, got, `want: "%s, got: "%s"`, test.Version, got)
		})
	}

	type ErrorTest struct {
		Error error
		Name  string
	}

	var errorTests = []ErrorTest{
		{Name: "ReturnDefaultVersionOnGitError", Error: fmt.Errorf("some-error")},
	}

	for _, test := range errorTests {
		t.Run(test.Name, func(t *testing.T) {
			var cmder = mocks.NewMockCommander()
			cmder.On("Output", mock.Anything, mock.Anything).Return("", test.Error)

			var gitAPI = git.API{Commander: cmder}
			var versionAPI = API{GitAPI: gitAPI}

			var got = versionAPI.GetVersionOrDefault(cli.DefaultVersion)
			assert.Equal(t, cli.DefaultVersion, got, `want: "%s, got: "%s"`, cli.DefaultVersion, got)
		})
	}
}

func TestAPI_PredictVersion(t *testing.T) {

}

func TestAPI_PushVersion(t *testing.T) {

}

func TestAPI_ReleaseVersion(t *testing.T) {

}

func TestAPI_UpdateVersion(t *testing.T) {

}

func TestNewAPI(t *testing.T) {
	t.Run("ValidateState", func(t *testing.T) {
		var api = NewAPI()
		assert.NotNil(t, api.GitAPI)
	})
}
