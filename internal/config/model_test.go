package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEnv(t *testing.T) {
	t.Run("CheckDefaultValues", func(t *testing.T) {
		var want = Env{
			Vars:    map[string]string{},
			Scripts: []EnvScript{},
			Files:   []string{},
		}
		var got = NewEnv()

		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})
}

func TestNewGit(t *testing.T) {
	t.Run("CheckDefaultValues", func(t *testing.T) {
		var want = Git{
			Config:    map[string]string{},
			Unshallow: true,
		}
		var got = NewGit()
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})
}

func TestNewSemver(t *testing.T) {
	t.Run("CheckDefaultValues", func(t *testing.T) {
		var want = Semver{
			Matches:  map[string]string{},
			Strategy: "auto",
		}
		var got = NewSemver()
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})
}

func TestNewRoot(t *testing.T) {
	t.Run("CheckDefaultValues", func(t *testing.T) {
		var want = Root{
			Env:    Env{Vars: map[string]string{}, Scripts: []EnvScript{}, Files: []string{}},
			Git:    Git{Config: map[string]string{}, Unshallow: true},
			Semver: Semver{Matches: map[string]string{}, Strategy: "auto"},
		}
		var got = NewRoot()
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})
}
