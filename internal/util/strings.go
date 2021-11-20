package util

import "strings"

// SplitByDelimiterString splits a string by multiple delimiters.
// Returns the resulting slice of strings.
func SplitByDelimiterString(target string, delimiters string) []string {
	var splitDelimiters = strings.Split(delimiters, "")

	return strings.FieldsFunc(target, func(r rune) bool {
		for _, delimiter := range splitDelimiters {
			if delimiter == string(r) {
				return true
			}
		}
		return false
	})
}

// SliceContainsString returns true if a string equals an element in the slice.
func SliceContainsString(container []string, value string) bool {
	for _, contained := range container {
		if contained == value {
			return true
		}
	}
	return false
}

// Contains returns true if a string, split by delimiters, contains another string.
func Contains(target string, value string, delimiters string) bool {
	var slice = SplitByDelimiterString(target, delimiters)
	return SliceContainsString(slice, value)
}
