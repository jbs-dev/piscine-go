package piscine

// Rot14 that takes a string as input and rotates each letter in the string by 14 positions forward
// in the ASCII table (with wraparound, so that 'Z' becomes 'N' and 'z' becomes 'n')
func Rot14(str string) string {
	// Initialize an empty string to store the result
	result := ""

	// Iterate over each rune in the input string
	for _, res := range str {
		// Rotate the current rune by 14 positions and append the result to the output string
		result += string(rot14(res))
	}

	// Return the rotated string
	return result
}

func rot14(a rune) rune {
	// Check if the rune is a capital letter or lowercase letter between A and M or a and m
	if a >= 'A' && a < 'M' || a >= 'a' && a < 'm' {
		// If so, rotate the rune by 14 positions forward in the ASCII table
		return a + 14
	}
	// Check if the rune is a capital letter or lowercase letter between N and Z or n and z
	if a >= 'M' && a <= 'Z' || a >= 'm' && a <= 'z' {
		// If so, rotate the rune by 12 positions backward in the ASCII table
		return a - 12
	}
	// If the rune is not a letter between A and Z or a and z, return it unchanged
	return a
}
