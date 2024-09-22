package versions

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSuffix(t *testing.T) {
	type Test struct {
		Name    string
		Version string
		Suffix  string
	}

	var tests = []Test{
		{Name: "HappyPath", Version: "0.0.0", Suffix: "a"},
		{Name: "HappyPathAlt", Version: "0.0.0", Suffix: "-alt"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var got = AddSuffix(test.Version, test.Suffix)
			assert.True(t, strings.HasSuffix(got, test.Suffix))
		})
	}
}
