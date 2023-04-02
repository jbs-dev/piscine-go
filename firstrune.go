package piscine

// FirstRune is a function that takes a string as input and returns the first rune of that string.
func FirstRune(s string) rune {
	// Convert the input string to a slice of runes
	sRune := []rune(s)
	// Return the first rune of the slice
	return sRune[0]
}
