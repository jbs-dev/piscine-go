package piscine

// IsPrintable returns true if all characters in 'str' are printable,
// and false otherwise.
func IsPrintable(str string) bool {
	// Iterate over each rune in 'str'
	for _, i := range str {
		// If the rune is not printable, return false
		if charIsPrintable(i) {
			return false
		}
	}
	// If all runes are printable, return true
	return true
}

// charIsPrintable returns true if the given rune is a non-printable ASCII character,
// and false otherwise.
func charIsPrintable(r rune) bool {
	// Non-printable ASCII characters have code points from 0 to 31
	if r >= 0 && r <= 31 {
		return true
	}
	return false
}
