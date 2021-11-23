package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	type Test struct {
		Delimiters   string
		Name         string
		TargetString string
		Value        string
		Want         bool
	}

	var tests = []Test{
		{Name: "EmptyDelimitersReturnsFalse", Delimiters: "", TargetString: "some-value", Value: "value", Want: false},
		{Name: "EmptyTargetStringWithNonEmptyValueReturnsFalse", Delimiters: "-", TargetString: "", Value: "value", Want: false},
		{Name: "EmptyTargetStringWithEmptyValueReturnsFalse", Delimiters: "-", TargetString: "", Value: "", Want: false},
		{Name: "EmptyValueReturnsFalse", Delimiters: "/", TargetString: "some/string", Value: "", Want: false},
		{Name: "IncorrectDelimitersReturnsFalse", Delimiters: "/", TargetString: "some-value", Value: "value", Want: false},
		{Name: "OneDelimiterButWrongLocationReturnsFalse", Delimiters: "/", TargetString: "some/feature [test]", Value: "feature", Want: false},
		{Name: "OneDelimiterButCorrectLocationReturnsTrue", Delimiters: "/", TargetString: "some/feature/ [test]", Value: "feature", Want: true},
		{Name: "MultipleDelimitersReturnsTrue", Delimiters: "/[]", TargetString: "some/feature [test]", Value: "test", Want: true},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var got = Contains(test.TargetString, test.Value, test.Delimiters)
			assert.Equal(t, test.Want, got, `want: "%s, got: "%s"`, test.Want, got)
		})
	}
}

func TestSplitByDelimiterString(t *testing.T) {
	type Test struct {
		Delimiters   string
		Name         string
		TargetString string
		Want         []string
	}

	var tests = []Test{
		{Name: "EmptyDelimiters", Delimiters: "", TargetString: "some-string", Want: []string{"some-string"}},
		{Name: "EmptyTargetString", Delimiters: "/", TargetString: "", Want: []string{}},
		{Name: "IncorrectDelimiters", Delimiters: "/", TargetString: "some-string", Want: []string{"some-string"}},
		{Name: "OneDelimiter", Delimiters: "/", TargetString: "some/feature [test]", Want: []string{"some", "feature [test]"}},
		{Name: "MultipleDelimiters", Delimiters: "/[]", TargetString: "some/feature [test]", Want: []string{"some", "feature ", "test"}},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var got = SplitByDelimiterString(test.TargetString, test.Delimiters)
			assert.Equal(t, test.Want, got, `want: "%s, got: "%s"`, test.Want, got)
		})
	}
}

func TestSliceContainsString(t *testing.T) {
	type Test struct {
		Name  string
		Slice []string
		Value string
		Want  bool
	}

	var tests = []Test{
		{Name: "EmptySlice", Slice: []string{}, Value: "test", Want: false},
		{Name: "EmptyValue", Slice: []string{"test"}, Value: "", Want: false},
		{Name: "OneSliceElementWithValue", Slice: []string{"test"}, Value: "test", Want: true},
		{Name: "MultipleSliceElementsWithValue", Slice: []string{"one", "test", "two"}, Value: "test", Want: true},
		{Name: "OneSliceElementWithoutValue", Slice: []string{"without"}, Value: "test", Want: false},
		{Name: "MultipleElementWithoutValue", Slice: []string{"without", "with", "out"}, Value: "test", Want: false},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var got = SliceContainsString(test.Slice, test.Value)
			assert.Equal(t, test.Want, got, `want: "%s, got: "%s"`, test.Want, got)
		})
	}
}
