package piscine

// MakeRange returns a slice of integers from 'min' to 'max', inclusive.
// If 'max' is less than 'min', the function returns an empty slice.
func MakeRange(min, max int) []int {
	var array []int
	// Check if 'max' is greater than 'min'
	if max > min {
		// Allocate a slice of length (max-min) to hold the range of integers
		array = make([]int, max-min)
		// Populate the slice with the range of integers
		for i := range array {
			array[i] = i + min
		}
	}
	// Return the resulting slice
	return array
}
