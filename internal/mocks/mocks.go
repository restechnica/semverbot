package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockCommander a commander interface mock implementation.
type MockCommander struct {
	mock.Mock
}

// NewMockCommander creates a new MockCommander.
// returns the new MockCommander.
func NewMockCommander() *MockCommander {
	return &MockCommander{}
}

// Output runs a mock command.
// returns mocked output or a mocked error.
func (mock *MockCommander) Output(name string, arg ...string) (string, error) {
	args := mock.Called(name, arg)
	return args.String(0), args.Error(1)
}

// Run runs a mock command.
// returns a mocked error.
func (mock *MockCommander) Run(name string, arg ...string) error {
	args := mock.Called(name, arg)
	return args.Error(0)
}

// MockSemverMode a semver mode interface mock implementation.
type MockSemverMode struct {
	mock.Mock
}

// NewMockSemverMode creates a new MockSemverMode.
// returns the new MockSemverMode.
func NewMockSemverMode() *MockSemverMode {
	return &MockSemverMode{}
}

// Increment mock increments a version.
// returns an incremented mock version.
func (mock *MockSemverMode) Increment(targetVersion string) (nextVersion string, err error) {
	args := mock.Called(targetVersion)
	return args.String(0), args.Error(1)
}
