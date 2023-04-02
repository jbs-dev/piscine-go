package piscine

// Map applies the input function 'f' to each element in the input integer slice 'arr',
// and returns a new slice of booleans where each element corresponds to the result of 'f' applied to the corresponding element in 'arr'.
func Map(f func(int) bool, arr []int) []bool {
	// Calculate the length of 'arr'
	lens := 0
	for range arr {
		lens++
	}
	// Initialize an empty boolean slice to hold the results
	index := 0
	array := make([]bool, lens)
	// Apply 'f' to each element in 'arr', and add the result to the 'array' slice
	for _, v := range arr {
		array[index] = f(v)
		index++
	}
	// Return the resulting slice
	return array
}
