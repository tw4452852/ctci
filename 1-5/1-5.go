package ctci

// Write a method to replace all spaces in a string with "%20".

func ReplaceSpace(s []byte) []byte {
	spaceCount := 0
	for _, sb := range s {
		if sb == ' ' {
			spaceCount++
		}
	}
	if spaceCount == 0 {
		return s
	}
	r := make([]byte, len(s)+2*spaceCount)
	ri := 0
	for _, sb := range s {
		if sb == ' ' {
			r[ri] = '%'
			r[ri+1] = '2'
			r[ri+2] = '0'
			ri += 3
		} else {
			r[ri] = sb
			ri++
		}
	}
	return r
}
