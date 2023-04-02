package piscine

// IterativePower returns the result of raising 'nb' to the power of 'power' using an iterative approach.
// If 'power' is 0, the function returns 1. If 'power' is negative, the function returns 0.
func IterativePower(nb int, power int) int {
	// If 'power' is 0, the result is always 1
	if power == 0 {
		return 1
	}
	// If 'power' is negative, the result is always 0
	if power < 0 {
		return 0
	} else {
		// Initialize 'result' to 1, then multiply it by 'nb' 'power' times
		result := 1
		for a := 1; a <= power; a++ {
			result *= nb
		}
		// Return the result
		return result
	}
}
