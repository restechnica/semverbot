package git

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/restechnica/semverbot/internal/mocks"
)

func TestCLI_CreateAnnotatedTag(t *testing.T) {
	t.Run("ReturnErrorOnCommanderError", func(t *testing.T) {
		var want = fmt.Errorf("some-error")

		var cmder = mocks.NewMockCommander()
		cmder.On("Run", mock.Anything, mock.Anything).Return(want)

		var gitCLI = CLI{Commander: cmder}
		var got = gitCLI.CreateAnnotatedTag("0.0.0")

		assert.Error(t, got)
		assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
	})
}

func TestCLI_FetchTags(t *testing.T) {
	t.Run("ReturnErrorOnCommanderError", func(t *testing.T) {
		var want = fmt.Errorf("some-error")

		var cmder = mocks.NewMockCommander()
		cmder.On("Run", mock.Anything, mock.Anything).Return(want)

		var gitCLI = CLI{Commander: cmder}
		var got = gitCLI.FetchTags()

		assert.Error(t, got)
		assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
	})
}

func TestCLI_FetchUnshallow(t *testing.T) {
	t.Run("ReturnErrorOnCommanderError", func(t *testing.T) {
		var want = fmt.Errorf("some-error")

		var cmder = mocks.NewMockCommander()
		cmder.On("Run", mock.Anything, mock.Anything).Return(want)

		var gitCLI = CLI{Commander: cmder}
		var got = gitCLI.FetchUnshallow()

		assert.Error(t, got)
		assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
	})
}

func TestCLI_GetConfig(t *testing.T) {
	t.Run("ReturnErrorOnCommanderError", func(t *testing.T) {
		var want = fmt.Errorf("some-error")

		var cmder = mocks.NewMockCommander()
		cmder.On("Output", mock.Anything, mock.Anything).Return("value", want)

		var gitCLI = CLI{Commander: cmder}
		var _, got = gitCLI.GetConfig("key")

		assert.Error(t, got)
		assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
	})
}

func TestCLI_GetLatestAnnotatedTag(t *testing.T) {
	t.Run("ReturnErrorOnCommanderError", func(t *testing.T) {
		var want = fmt.Errorf("some-error")

		var cmder = mocks.NewMockCommander()
		cmder.On("Output", mock.Anything, mock.Anything).Return("value", want)

		var gitCLI = CLI{Commander: cmder}
		var _, got = gitCLI.GetLatestAnnotatedTag()

		assert.Error(t, got)
		assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
	})
}

func TestCLI_GetLatestCommitMessage(t *testing.T) {
	t.Run("ReturnErrorOnCommanderError", func(t *testing.T) {
		var want = fmt.Errorf("some-error")

		var cmder = mocks.NewMockCommander()
		cmder.On("Output", mock.Anything, mock.Anything).Return("value", want)

		var gitCLI = CLI{Commander: cmder}
		var _, got = gitCLI.GetLatestCommitMessage()

		assert.Error(t, got)
		assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
	})
}

func TestCLI_GetMergedBranchName(t *testing.T) {
	t.Run("ReturnErrorOnCommanderError", func(t *testing.T) {
		var want = fmt.Errorf("some-error")

		var cmder = mocks.NewMockCommander()
		cmder.On("Output", mock.Anything, mock.Anything).Return("value", want)

		var gitCLI = CLI{Commander: cmder}
		var _, got = gitCLI.GetMergedBranchName()

		assert.Error(t, got)
		assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
	})
}

func TestCLI_PushTag(t *testing.T) {
	t.Run("ReturnErrorOnCommanderError", func(t *testing.T) {
		var want = fmt.Errorf("some-error")

		var cmder = mocks.NewMockCommander()
		cmder.On("Run", mock.Anything, mock.Anything).Return(want)

		var gitCLI = CLI{Commander: cmder}
		var got = gitCLI.PushTag("tag")

		assert.Error(t, got)
		assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
	})
}

func TestCLI_SetConfig(t *testing.T) {
	t.Run("ReturnErrorOnCommanderError", func(t *testing.T) {
		var want = fmt.Errorf("some-error")

		var cmder = mocks.NewMockCommander()
		cmder.On("Run", mock.Anything, mock.Anything).Return(want)

		var gitCLI = CLI{Commander: cmder}
		var got = gitCLI.SetConfig("key", "value")

		assert.Error(t, got)
		assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
	})
}

func TestCLI_SetConfigIfNotSet(t *testing.T) {
	t.Run("DoNotSetConfigIfConfigExists", func(t *testing.T) {
		var want = "initial"

		var cmder = mocks.NewMockCommander()
		cmder.On("Output", mock.Anything, mock.Anything).Return(want, nil)

		var gitCLI = CLI{Commander: cmder}
		var got, err = gitCLI.SetConfigIfNotSet("key", "value")

		cmder.AssertCalled(t, "Output", mock.Anything, mock.Anything)
		cmder.AssertNotCalled(t, "Run", mock.Anything, mock.Anything)

		assert.NoError(t, err)
		assert.Equal(t, nil, err, `want: "%s, got: "%s"`, nil, err)
		assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
	})

	t.Run("SetConfigIfConfigDoesNotExist", func(t *testing.T) {
		var initial = "initial"
		var want = "value"

		var cmder = mocks.NewMockCommander()
		cmder.On("Output", mock.Anything, mock.Anything).Return(initial, fmt.Errorf("some-error"))
		cmder.On("Run", mock.Anything, mock.Anything).Return(nil)

		var gitCLI = CLI{Commander: cmder}
		var got, err = gitCLI.SetConfigIfNotSet("key", want)

		cmder.AssertCalled(t, "Output", mock.Anything, mock.Anything)
		cmder.AssertCalled(t, "Run", mock.Anything, mock.Anything)

		assert.NoError(t, err)
		assert.Equal(t, nil, err, `want: "%s, got: "%s"`, nil, err)
		assert.Equal(t, nil, err, `want: "%s, got: "%s"`, nil, err)

		assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
	})

	t.Run("ReturnErrorOnSetConfigError", func(t *testing.T) {
		var want = fmt.Errorf("some-error")

		var cmder = mocks.NewMockCommander()
		cmder.On("Output", mock.Anything, mock.Anything).Return("value", fmt.Errorf("some-error"))
		cmder.On("Run", mock.Anything, mock.Anything).Return(fmt.Errorf("some-error"))

		var gitCLI = CLI{Commander: cmder}
		var _, got = gitCLI.SetConfigIfNotSet("key", "value")

		assert.Error(t, got)
		assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
	})
}

func TestNewCLI(t *testing.T) {
	t.Run("ValidateState", func(t *testing.T) {
		var cli = NewCLI()
		assert.NotNil(t, cli)
		assert.NotNil(t, cli.Commander)
	})
}
