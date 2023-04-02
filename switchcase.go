package piscine

// SwitchCase takes a string as input and returns a new string with the case of each character switched.
// If a character is uppercase, it will be converted to lowercase, and vice versa.
func SwitchCase(s string) string {
	// Convert the input string to a slice of runes
	runeArray := []rune(s)

	// Iterate through the rune slice.
	for i := range runeArray {
		// If the current rune is a lowercase letter (between 'a' and 'z'), convert it to uppercase
		// by subtracting 32 from its Unicode code point. The difference between the code points
		// of 'A' and 'a' (or 'Z' and 'z') is 32.
		if runeArray[i] >= 'a' && runeArray[i] <= 'z' {
			runeArray[i] = runeArray[i] - 32
		} else {
			// If the current rune is not a lowercase letter, it is assumed to be an uppercase letter.
			// Convert it to lowercase by adding 32 to its Unicode code point.
			runeArray[i] = runeArray[i] + 32
		}
	}

	// Convert the modified rune slice back to a string and return it.
	return string(runeArray)
}
