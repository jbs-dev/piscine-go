package piscine

// SortWordArr that takes a slice of strings as input and sorts it in lexicographic order
// /(i.e., in alphabetical order, ignoring case).
func SortWordArr(array []string) {
	// Initialize an empty string to store temporary results
	result := ""

	// Use nested for loops to iterate over every pair of elements in the input slice
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			// Convert the current pair of elements to byte arrays
			a := []byte(array[j])
			b := []byte(array[i])

			// Compare the current pair of byte arrays using the comparearrayascii function
			if comparearrayascii(a, b) {
				// If the second element is less than the first element, swap them
				result = array[i]
				array[i] = array[j]
				array[j] = result
			}
		}
	}
}

// comparearrayascii function takes two byte arrays as input and compares them in ASCII order
func comparearrayascii(a, b []byte) bool {
	// Determine the minimum length of the two byte arrays
	n := 0
	if len(a) <= len(b) {
		n = len(a)
	} else {
		n = len(b)
	}

	// Compare the two byte arrays, element by element, in ASCII order
	for i := 0; i < n; i++ {
		// If the current element in the second byte array is less than the corresponding element in the first byte array,
		// the second byte array is considered "less than" the first byte array
		if a[i] < b[i] {
			return true
			// If the current element in the second byte array is greater than the corresponding element in the first byte array,
			// the second byte array is considered "greater than" the first byte array
		} else if a[i] > b[i] {
			return false
		}
	}

	// If the two byte arrays are equal up to the minimum length, the shorter array is considered "less than" the longer array
	return len(a) < len(b)
}
