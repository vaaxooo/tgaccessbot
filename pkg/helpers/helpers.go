package helpers

import "strings"

// SplitValue splits a string by a separator
func SplitValue(value string, separator string) []string {
	if value == "" {
		return []string{}
	}

	return strings.Split(value, separator)
}

// The function checks if a given item is present in a slice of strings and returns true if it is
// found, otherwise false.
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
