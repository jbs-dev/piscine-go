package piscine

// Sqrt returns the integer square root of a number nb.
// If nb does not have an integer square root, the function returns 0.
func Sqrt(nb int) int {
	// If the input number is 1, its square root is 1.
	if nb == 1 {
		return 1
	}
	// If the input number is 2, it doesn't have an integer square root, so return 0.
	if nb == 2 {
		return 0
	}
	// Check if the input number is positive.
	if nb > 0 {
		result := 1 // Variable to store the square of the current integer a.
		sqrt := 0   // Variable to store the integer square root value.
		// Iterate through positive integers, starting from 1.
		for a := 1; result <= nb; a++ {
			result = a * a // Calculate the square of the current integer a.
			sqrt++         // Increment the integer square root value.

			// If the square of the current integer a is equal to the input number,
			// return the integer square root value.
			if result == nb {
				return sqrt
			}
		}
		// If the loop finishes without finding an integer square root, return 0.
		return 0
	} else {
		// If the input number is not positive, it doesn't have an integer square root, so return 0.
		return 0
	}
}
