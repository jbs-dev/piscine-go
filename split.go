package piscine

// Split that takes a string str and a string charset as input,
//	and returns a slice of strings that contains all the substrings of str
//		that are separated by any of the characters in charse

func Split(str, charset string) []string {
	// Determine the length of the input string and the separator string
	lenStr := 0
	lenCharset := 0
	for range str {
		lenStr++
	}
	for range charset {
		lenCharset++
	}

	// Count the number of words in the input string
	word := 0
	for i := 0; i < lenStr-lenCharset; i++ {
		if str[i:i+lenCharset] == charset {
			word++
		}
	}

	// Initialize a new string slice with length word+1
	newStr := make([]string, word+1)

	// Initialize variables for keeping track of the current word index, current word length, and current position in str
	j := -1
	k := 0
	for i := 0; i < lenStr-lenCharset; i++ {
		// If the current substring of str is equal to the separator string, store the current word in newStr and reset k and i to the appropriate values
		if str[i:i+lenCharset] == charset {
			j++
			newStr[j] = str[i-k : i]
			k = 0
			i = i + lenCharset - 1
			// If the current substring of str is not equal to the separator string, increment k to keep track of the current word length
		} else {
			k++
		}
	}

	// Calculate the length of the final word and store it in q
	q := 0
	for _, v := range newStr {
		for range v {
			q++
		}
		q += lenCharset
	}
	q -= lenCharset

	// Store the final word in newStr
	j++
	newStr[j] = str[q:]

	// Return the new string slice
	return newStr
}
