package piscine

// IsUpper returns true if all characters in 'str' are uppercase letters (A-Z),
// and false otherwise.
func IsUpper(str string) bool {
	runeCount := arrayCount(str) // Count the number of runes in 'str'
	count := 0                   // Initialize a variable to count the number of uppercase letters in 'str'
	for _, i := range str {      // Iterate over each rune in 'str'
		if i >= 'A' && i <= 'Z' { // If the rune is an uppercase letter, increment the count
			count++
		}
	}
	return count == runeCount // Return true if the number of uppercase letters is equal to the number of runes in 'str'
}
