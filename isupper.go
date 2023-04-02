package piscine

// IsUpper returns true if all characters in 'str' are uppercase letters (A-Z),
// and false otherwise.
func IsUpper(str string) bool {
	// Count the number of runes in 'str'
	runeCount := arrayCount(str)
	// Initialize a variable to count the number of uppercase letters in 'str'
	count := 0
	// Iterate over each rune in 'str'
	for _, i := range str {
		// If the rune is an uppercase letter, increment the count
		if i >= 'A' && i <= 'Z' {
			count++
		}
	}
	// Return true if the number of uppercase letters is equal to the number of runes in 'str',
	// and false otherwise
	return count == runeCount
}
