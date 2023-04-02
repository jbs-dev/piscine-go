package piscine

// NRune returns the nth rune (Unicode code point) in the input string 's'.
// If the value of 'n' is out of range or negative, it returns the null rune.
func NRune(s string, n int) rune {
	// Convert the input string 's' to a rune slice
	runeArray := []rune(s)
	// Count the number of runes in the slice
	arrayCount := 0
	for i := range s {
		arrayCount = i + 1
	}

	// If 'n' is within range, return the nth rune (subtracting 1 to account for 0-based indexing)
	if n <= arrayCount && n-1 >= 0 {
		return runeArray[n-1]
	} else {
		// Otherwise, return the null rune (U+0000)
		return 0
	}
}
