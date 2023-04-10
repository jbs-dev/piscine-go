package piscine

// StrLen returns the length of a string, which is the count of characters in the string.
func StrLen(s string) int {
	runeSlice := []rune(s) // Convert the input string to a slice of runes.
	return len(runeSlice)
}

// Use the built-in len function to determine the length of the rune slice, which
// corresponds to the number of characters in the input string.
