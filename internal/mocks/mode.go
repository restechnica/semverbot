package mocks

import "github.com/stretchr/testify/mock"

// MockMode a semver mode interface mock implementation.
type MockMode struct {
	mock.Mock
}

// NewMockMode creates a new MockMode.
// Returns the new MockMode.
func NewMockMode() *MockMode {
	return &MockMode{}
}

// Increment mock increments a version.
// Returns an incremented mock version.
func (mock *MockMode) Increment(targetVersion string) (nextVersion string, err error) {
	args := mock.Called(targetVersion)
	return args.String(0), args.Error(1)
}
