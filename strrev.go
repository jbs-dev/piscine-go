package piscine

// StrRev reverses the characters in a given string and returns the reversed string.
func StrRev(s string) string {
	// Convert the input string to a slice of runes. A rune represents a Unicode code point,
	// which can be a single character or a sequence of bytes for multibyte characters.
	// This conversion ensures that the reversal will be accurate for strings
	// containing multibyte characters.
	runes := []rune(s)

	// Initialize a length variable to store the length of the rune slice.
	length := 0

	// Iterate through the rune slice, updating the length variable to the last index.
	for i := range runes {
		length = i
	}

	// Reverse the characters in the rune slice using a two-pointer approach.
	// One pointer (i) starts at the beginning of the slice and moves forward,
	// while the other pointer (j) starts at the end of the slice and moves backward.
	for i, j := 0, length; i < j; i, j = i+1, j-1 {
		// Swap the characters at positions i and j.
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the reversed rune slice back to a string and return it.
	return string(runes)
}
