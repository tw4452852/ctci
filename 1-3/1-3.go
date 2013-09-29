package ctci

// Design an algorithm and write code to remove the duplicate
// characters in a string without using any additional buffer.
// NOTE: One or two additional variables are fine.
// An extra copy of the array is not.
// FOLLOW UP
// Write the test cases for this method.

func RemoveDuplicate(s []byte) []byte {
	length := len(s) - 1
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if s[i] == s[j] {
				for k := j; k != length-1; k++ {
					s[k] = s[k+1]
				}
				length--
				j--
			}
		}
	}
	return append(s[:length], 0)
}
