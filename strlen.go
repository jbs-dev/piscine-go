package piscine

// StrLen returns the length of a string, which is the count of characters in the string.
func StrLen(s string) int {
	// Convert the input string to a slice of runes. A rune represents a Unicode code point,
	// which can be a single character or a sequence of bytes for multibyte characters.
	// This conversion ensures that the length calculation will be accurate for strings
	// containing multibyte characters.
	runeSlice := []rune(s)

	// Use the built-in len function to determine the length of the rune slice, which
	// corresponds to the number of characters in the input string.
	return len(runeSlice)
}
