package piscine

// Max returns the maximum value in the input integer slice 'arr'.
func Max(arr []int) int {
	// Initialize the maximum value to be the first element in the slice
	max := arr[0]
	// Iterate over the remaining elements in the slice and update the maximum value if necessary
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	// Return the maximum value
	return max
}
