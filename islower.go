package piscine

// IsLower returns true if all characters in 'str' are lowercase letters,
// and false otherwise.
func IsLower(str string) bool {
	// Count the number of runes in 'str'
	runeCount := arrayCount(str)
	// Initialize a variable to count the number of lowercase letters in 'str'
	count := 0
	// Iterate over each rune in 'str'
	for _, i := range str {
		// If the rune is a lowercase letter, increment the count
		if i >= 'a' && i <= 'z' {
			count++
		}
	}
	// Return true if the number of lowercase letters is equal to the number of runes in 'str',
	// and false otherwise
	return count == runeCount
}
