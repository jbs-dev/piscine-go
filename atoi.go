package piscine

// Atoi converts a string to an integer
func Atoi(s string) int {
	runes := []rune(s) // convert the string to an array of runes (Unicode characters)

	num := 0               // initialize a variable to hold the final integer value
	length := 0            // initialize a variable to hold the length of the string
	for i := range runes { // iterate over the array of runes
		length = i // set the length variable to the index of the final rune
	}
	if length == 0 { // if the length of the string is 0, return 0
		return 0
	}

	ten := 1                      // initialize a variable to hold the value of the current place (ones, tens, hundreds, etc.)
	for j := 0; j < length; j++ { // iterate over the array of runes (excluding the last rune)
		if runes[j] == '+' || runes[j] == '-' { // if the current rune is a plus or minus sign, skip it
			continue
		}
		ten *= 10 // multiply the value of the current place by 10 for each non-sign rune in the string
	}

	for i := range runes { // iterate over the array of runes
		if (runes[0] == '+' || runes[0] == '-') && i == 0 { // if the first rune is a plus or minus sign, skip it
			continue
		}
		if runes[i] < '0' || runes[i] > '9' { // if the current rune is not a digit, return 0
			return 0
		}
		temp := 0                         // initialize a variable to hold the value of the current digit
		for k := '0'; k < runes[i]; k++ { // iterate over all digits less than the current digit
			temp++ // increment the value of the current digit for each digit less than it
		}
		num += temp * ten // add the value of the current digit multiplied by the value of the current place to the final integer value
		ten /= 10         // divide the value of the current place by 10 to move to the next place
	}

	if runes[0] == '-' { // if the first rune is a minus sign, make the final integer value negative
		num *= -1
	}
	return num // return the final integer value
}
