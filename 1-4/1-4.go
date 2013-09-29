package ctci

// Write a method to decide if two strings are anagrams or not.

// Assume two string only contain ascii charactors
func IsAnagram(a, b []byte) bool {
	counts := make([]int, 256)
	for _, ab := range a {
		counts[ab]++
	}
	for _, bb := range b {
		counts[bb]--
	}
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}
