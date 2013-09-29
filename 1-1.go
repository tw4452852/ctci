package ctci

// Implement an algorithm to determine if a string has all unique characters.
// What if you can not use additional data structures?

// IsUnique determine if a string has all unique characters
// Assume a string only contain the ascii characters
func IsUnique(s string) bool {
	seen := make([]int, 8) // every bit represent a character
	for _, c := range s {
		if isSet(byte(c), seen) {
			return false
		}
		set(byte(c), seen)
	}
	return true
}

func isSet(c byte, seen []int) bool {
	index := c / 32
	offset := c % 32
	return (seen[index] & (1 << offset)) != 0
}

func set(c byte, seen []int) {
	index := c / 32
	offset := c % 32
	seen[index] |= (1 << offset)
}
