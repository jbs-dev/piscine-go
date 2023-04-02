package piscine

// Join concatenates the input string slice 'strs' into a single string, separated by the input separator 'sep',
// and returns the resulting string.
func Join(strs []string, sep string) string {
	// Initialize an empty string to hold the result
	result := ""
	// Count the number of strings in 'strs'
	count := 0
	for i := range strs {
		count = i + 1
	}
	// Concatenate each string in 'strs' to the result string, separated by 'sep'
	for i, str := range strs {
		// If it's the last string, don't add the separator
		if i == count-1 {
			result += str
		} else {
			result += str + sep
		}
	}
	// Return the result string
	return result
}
