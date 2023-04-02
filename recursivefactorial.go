package piscine

// RecursiveFactorial that takes an integer nb as input and recursively calculates the factorial of nb.
func RecursiveFactorial(nb int) int {
	// Check if nb is less than 25
	if nb < 25 {
		// If nb is 0, return 1 (by definition, 0! is equal to 1)
		if nb == 0 {
			return 1
		}
		// If nb is negative, return 0 (since factorials are not defined for negative numbers)
		if nb < 0 {
			return 0
		} else {
			// If nb is positive, recursively calculate nb * (nb-1)!
			return nb * RecursiveFactorial(nb-1)
		}
	}
	// If nb is greater than or equal to 25, return 0 (since factorials of such large numbers can cause integer overflow)
	return 0
}
