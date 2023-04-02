package piscine

// IsNumeric returns true if all characters in 'str' are digits (0-9),
// and false otherwise.
func IsNumeric(str string) bool {
	// Count the number of runes in 'str'
	runeCount := arrayCount(str)
	// Initialize a variable to count the number of digits in 'str'
	count := 0
	// Iterate over each rune in 'str'
	for _, char := range str {
		// If the rune is a digit, increment the count
		if isDigit(char) {
			count++
		}
	}
	// Return true if the number of digits is equal to the number of runes in 'str',
	// and false otherwise
	return count == runeCount
}

// isDigit returns true if the given rune is a digit (0-9),
// and false otherwise.
func isDigit(digit rune) bool {
	// Iterate over the range of ASCII values for digits (0-9)
	for a := '0'; a <= '9'; a++ {
		// If the given rune matches the current digit, return true
		if digit == a {
			return true
		}
	}
	// If the given rune does not match any digit, return false
	return false
}
