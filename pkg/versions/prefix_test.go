package versions

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddPrefix(t *testing.T) {
	type Test struct {
		Name    string
		Version string
		Prefix  string
	}

	var tests = []Test{
		{Name: "HappyPath", Version: "0.0.0", Prefix: "v"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var got = AddPrefix(test.Version, test.Prefix)
			assert.True(t, strings.HasPrefix(got, test.Prefix))
		})
	}
}
