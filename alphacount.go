package piscine

// AlphaCount returns the number of alphabetical characters in the given string
func AlphaCount(str string) int {
	count := 0              // initialize a variable to keep track of the count of alphabetical characters
	for _, s := range str { // iterate over each character in the string
		if isAlpha(s) { // check if the character is alphabetical
			count++ // increment the count variable
		}
	}
	return count // return the count of alphabetical characters in the string
}

// isAlpha returns true if the given rune is an alphabetical character, false otherwise
func isAlpha(alpha rune) bool {
	// iterate over the lowercase alphabet
	for a := 'a'; a <= 'z'; a++ {
		if alpha == a { // if the given rune matches the current lowercase letter
			return true // return true to indicate that it is an alphabetical character
		}
	}
	// iterate over the uppercase alphabet
	for a := 'A'; a <= 'Z'; a++ {
		if alpha == a { // if the given rune matches the current uppercase letter
			return true // return true to indicate that it is an alphabetical character
		}
	}
	return false // if the given rune does not match any lowercase or uppercase letter, return false
}
