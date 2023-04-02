package piscine

// Capitalize capitalises the first letter of each word in a string
func Capitalize(s string) string {
	runeArray := []rune(s)           // convert the string to an array of runes (Unicode characters)
	for i, char := range runeArray { // iterate over the array of runes
		if isNumberOrAlph(char) { // check if the current rune is a letter or a number
			if i == 0 || !isNumberOrAlph(runeArray[i-1]) { // check if the current rune is the first rune in a word
				if runeArray[i] >= 'a' && runeArray[i] <= 'z' { // capitalize the current letter if it is lowercase
					runeArray[i] = char - 32
				}
			} else { // if the current rune is not the first rune in a word
				if runeArray[i] >= 'A' && runeArray[i] <= 'Z' { // convert the current letter to lowercase if it is uppercase
					runeArray[i] = char + 32
				}
			}
		}
	}
	return string(runeArray) // return the final string
}

// isNumberOrAlph checks if a rune is a letter or a number
/* func isNumberOrAlph(r rune) bool {
	if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
		return true
	}
	return false
} */
