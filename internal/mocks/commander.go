package mocks

import "github.com/stretchr/testify/mock"

// MockCommander a commander interface mock implementation.
type MockCommander struct {
	mock.Mock
}

// NewMockCommander creates a new MockCommander.
// Returns the new MockCommander.
func NewMockCommander() *MockCommander {
	return &MockCommander{}
}

// Output runs a mock command.
// Returns mocked output or a mocked error.
func (mock *MockCommander) Output(name string, arg ...string) (string, error) {
	args := mock.Called(name, arg)
	return args.String(0), args.Error(1)
}

// Run runs a mock command.
// Returns a mocked error.
func (mock *MockCommander) Run(name string, arg ...string) error {
	args := mock.Called(name, arg)
	return args.Error(0)
}
