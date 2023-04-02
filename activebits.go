package piscine

func ActiveBits(n int) int {
	// Initialize a variable 'count' to 0
	count := 0
	// Start a loop that runs as long as 'n' is greater than 0
	for n > 0 {
		// Check if the least significant bit of 'n' is 1
		if n&1 == 1 {
			// If it is, increment 'count' by 1
			count++
		}
		// Right shift 'n' by 1 bit
		n >>= 1
	}
	// Return the value of 'count'
	return count
}
