package piscine

// BasicAtoi converts a string to an integer using decimal (base 10) representation
func BasicAtoi(s string) int {
	runes := []rune(s) // convert the string to an array of runes (Unicode characters)

	num := 0               // initialize a variable to hold the final integer value
	length := 0            // initialize a variable to hold the length of the string
	for i := range runes { // iterate over the array of runes
		length = i // set the length variable to the index of the final rune
	}

	ten := 1                      // initialize a variable to hold the value of the current place (ones, tens, hundreds, etc.)
	for j := 0; j < length; j++ { // iterate over the array of runes (excluding the last rune)
		ten *= 10 // multiply the value of the current place by 10 for each non-sign rune in the string
	}

	for i := range runes { // iterate over the array of runes
		temp := 0                         // initialize a variable to hold the value of the current digit
		for k := '0'; k < runes[i]; k++ { // iterate over all digits less than the current digit
			temp++ // increment the value of the current digit for each digit less than it
		}
		num += temp * ten // add the value of the current digit multiplied by the value of the current place to the final integer value
		ten /= 10         // divide the value of the current place by 10 to move to the next place
	}

	return num // return the final integer value
}
