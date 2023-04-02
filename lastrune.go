package piscine

// LastRune returns the last rune (Unicode code point) in the input string 's'.
func LastRune(s string) rune {
	// Convert the input string to a rune slice
	sRune := []rune(s)
	// Return the last rune in the slice
	return sRune[len(sRune)-1]
}
