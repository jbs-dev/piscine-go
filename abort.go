package piscine

func Abort(a, b, c, d, e int) int {
	// Create an integer array with the input values
	array := []int{a, b, c, d, e}
	// Set a variable 'count' to the length of the array (5 in this case)
	count := 5
	// Set a temporary variable to 0
	temp := 0
	// Start a nested for loop
	for i := 0; i < count-1; i++ {
		// Set the value of 'temp' to 'i'
		temp = i
		// Start a second loop to iterate through the remaining elements of the array
		for j := i + 1; j < count; j++ {
			// If the value at index 'j' is less than the value at index 'temp'
			if array[j] < array[temp] {
				// Set 'temp' to 'j'
				temp = j
			}
		}
		// Swap the values at indices 'i' and 'temp'
		array[i], array[temp] = array[temp], array[i]
	}
	// Return the value at the middle index of the sorted array (index 2 in this case)
	return array[2]
}
