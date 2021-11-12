package morestrings

func ReverseString(s string) string {
	// convert to slice of runes
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		//swap runes
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
