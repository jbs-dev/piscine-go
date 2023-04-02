package piscine

// SplitWhiteSpaces takes a string and returns a slice of its words,
// separated by one or more whitespace characters.
func SplitWhiteSpaces(str string) []string {
	// Remove unnecessary whitespace characters from the beginning, end, and
	// between words, and store the cleaned string.
	str = clearWord(str)

	// Calculate the number of words in the cleaned string.
	len := getWordLen(str)

	// Create a new slice with the same length as the number of words.
	arr := make([]string, len)

	tmp := "" // Temporary string to hold a word while iterating through the cleaned string.
	j := 0    // Index for adding words to the arr slice.

	// Check if the first character is not a whitespace, then add it to tmp.
	if !spaceCheck(str[0]) {
		tmp += string(str[0])
	}

	// Iterate through the rest of the cleaned string.
	for i := 1; i < strLen(str); i++ {
		// If the previous character is a whitespace, store the word (stored in tmp)
		// in the arr slice and reset tmp with the current character.
		if spaceCheck(str[i-1]) {
			arr[j] = tmp
			tmp = string(str[i])
			j++
		} else if !spaceCheck(str[i]) {
			// If the current character is not a whitespace, add it to tmp to form a word.
			tmp += string(str[i])
		}
	}
	// Store the last word in the arr slice.
	arr[j] = tmp

	return arr
}

// clearWord takes a string and removes unnecessary whitespace characters.
// It trims leading and trailing whitespaces and collapses multiple consecutive
// whitespaces to a single space between words.
func clearWord(s string) string {
	result := ""
	for i, v := range s {
		// Add the first character to the result string if it's not a whitespace.
		if i == 0 {
			if !spaceCheck(byte(v)) {
				result += string(v)
			}
		} else {
			// Add the character to the result string if the previous and current characters
			// are not both whitespaces.
			if !(spaceCheck(s[i-1]) && spaceCheck(s[i])) {
				result += string(v)
			}
		}
	}
	// Remove trailing whitespace from the result string, if any.
	len := strLen(result)
	if spaceCheck(result[len-1]) {
		result = result[:len-1]
	}
	return result
}

// spaceCheck checks if a byte represents a whitespace character.
// It returns true if the byte is a space (' '), newline ('\n'), or tab ('\t').
func spaceCheck(s byte) bool {
	return (s == ' ' || s == '\n' || s == '\t')
}

// getWordLen returns the number of words in a cleaned string,
// where words are separated by single whitespace characters.
func getWordLen(s string) int {
	if strLen(s) == 0 {
		return 0
	}
	count := 1 // Initialize count to 1, assuming there's at least one word in the cleaned string.
	for _, v := range s {
		// If the current character is a whitespace, increment the count of words.
		if spaceCheck(byte(v)) {
			count++
		}
	}
	return count
}

// strLen returns the length of a string by iterating through its characters
// and counting them.
func strLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
