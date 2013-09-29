package ctci

import (
	"strings"
)

// Assume you have a method isSubstring which checks if one word is a substring of another.
// Given two strings, s1 and s2,
// write code to check if s2 is a rotation of s1 using only one call to isSubstring
// (i.e.,"waterbottler" is a rotation of "erbottlewat").

func IsRotation(src, dst string) bool {
	if len(src) != len(dst) {
		return false
	}
	return strings.Contains(src+src, dst)
}
