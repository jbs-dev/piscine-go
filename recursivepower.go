package piscine

// RecursivePower that takes two integers, nb and power,
// as input and recursively calculates nb raised to the power of power
func RecursivePower(nb int, power int) int {
	// Check if power is 0
	if power == 0 {
		// If power is 0, the result is 1 (by definition, any number raised to the power of 0 is equal to 1)
		return 1
	}
	// Check if power is negative
	if power < 0 {
		// If power is negative, the result is 0 (since any number raised to a negative power is equal to 1 divided by that number raised to the absolute value of the power)
		return 0
	} else {
		// If power is positive, recursively calculate nb * nb^(power-1)
		return nb * RecursivePower(nb, power-1)
	}
}
