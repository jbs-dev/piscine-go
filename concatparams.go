package piscine

// ConcatParams concatenates all strings in a slice, separated by a newline character, and returns the result
func ConcatParams(args []string) string {
	result := ""     // initialize a variable to hold the resulting string
	count := 0       // initialize a variable to hold the number of elements in the slice
	for range args { // iterate over the slice to count the number of elements
		count++
	}
	for i := range args { // iterate over the slice to concatenate the elements
		if i == count-1 { // if the current element is the last element, concatenate it without a newline character
			result += args[i]
		} else { // if the current element is not the last element, concatenate it with a newline character
			result += args[i] + "\n"
		}
	}
	return result // return the resulting string
}
