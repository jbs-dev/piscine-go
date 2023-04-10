package piscine

// StrRev reverses the characters in a given string and returns the reversed string.
func StrRev(s string) string {
	runes := []rune(s)     // Convert the input string to a slice of runes.
	length := 0            // Initialize a length variable to store the length of the rune slice.
	for i := range runes { // Iterate through the rune slice, updating the length variable to the last index.
		length = i
	}
	// Reverse the characters in the rune slice using a two-pointer approach.
	// One pointer (i) starts at the beginning of the slice and moves forward,
	// while the other pointer (j) starts at the end of the slice and moves backward.
	for i, j := 0, length; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Swap the characters at positions i and j.
	}
	return string(runes) // Convert the reversed rune slice back to a string and return it.
}
