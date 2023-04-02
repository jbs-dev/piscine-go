package piscine

// Unmatch takes an array of integers (arr) as input and returns the first integer
// that appears an odd number of times in the array. If no such integer is found, it returns -1.
func Unmatch(arr []int) int {
	// Iterate through the array elements.
	for _, res := range arr {
		fois := 0 // Initialize a counter variable (fois) to count the occurrences of the current integer (res) in the array.

		// Iterate through the array again to count the occurrences of the current integer (res).
		for _, el := range arr {
			// If the current element (el) matches the integer (res), increment the counter variable (fois).
			if el == res {
				fois++
			}
		}

		// If the integer (res) appears an odd number of times (fois is 1 or an odd number),
		// return the integer as the result.
		if fois == 1 || fois%2 == 1 {
			return res
		}
	}
	// If no integer appears an odd number of times in the array, return -1.
	return -1
}
