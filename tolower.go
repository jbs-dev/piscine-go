package piscine

// ToLower takes a string as input and returns a new string with all uppercase characters converted to lowercase.
func ToLower(s string) string {
	// Convert the input string to a slice of runes
	// containing multibyte characters.
	runeArray := []rune(s)

	// Iterate through the rune slice.
	for i := range runeArray {
		// If the current rune is an uppercase letter (between 'A' and 'Z'), convert it to lowercase
		// by adding 32 to its Unicode code point. The difference between the code points
		// of 'A' and 'a' (or 'Z' and 'z') is 32.
		if runeArray[i] >= 'A' && runeArray[i] <= 'Z' {
			runeArray[i] = runeArray[i] + 32
		}
	}

	// Convert the modified rune slice back to a string and return it.
	return string(runeArray)
}
