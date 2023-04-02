package piscine

// Compact removes all empty strings from a slice and returns the length of the resulting slice
func Compact(ptr *[]string) int {
	length := 0              // initialize a variable to hold the length of the resulting slice
	for _, s := range *ptr { // iterate over the slice
		if s != "" { // check if the current string is not empty
			length++ // if the string is not empty, increment the length of the resulting slice
		}
	}

	count := 0                      // initialize a variable to hold the current index of the resulting slice
	array := make([]string, length) // create a new slice with the same length as the resulting slice
	for _, s := range *ptr {        // iterate over the slice
		if s != "" { // check if the current string is not empty
			array[count] = s // add the non-empty string to the resulting slice
			count++          // increment the current index of the resulting slice
		}
	}

	*ptr = array  // update the input slice to the resulting slice
	return length // return the length of the resulting slice
}
