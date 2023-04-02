package piscine

// IsAlpha returns true if all characters in 'str' are alphabetic or numeric,
// and false otherwise.
func IsAlpha(str string) bool {
	// Count the number of runes in 'str'
	runeCount := arrayCount(str)
	// Initialize a variable to count the number of alphabetic or numeric characters in 'str'
	count := 0
	// Iterate over each rune in 'str'
	for _, char := range str {
		// If the rune is alphabetic or numeric, increment the count
		if isNumberOrAlph(char) {
			count++
		}
	}
	// Return true if the number of alphabetic or numeric characters is equal to the number of runes in 'str',
	// and false otherwise
	return count == runeCount
}

// isNumberOrAlph returns true if the given rune is alphabetic or numeric,
// and false otherwise.
func isNumberOrAlph(r rune) bool {
	if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
		return true
	}
	return false
}
