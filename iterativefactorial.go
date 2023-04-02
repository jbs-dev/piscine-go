package piscine

func IterativeFactorial(nb int) int {
	// Check if nb is less than 25 (to avoid integer overflow)
	if nb < 25 {
		// Check if nb is 0, return 1
		if nb == 0 {
			return 1
		}
		// Check if nb is negative, return 0
		if nb < 0 {
			return 0
		}
		// Otherwise, calculate the factorial using a for loop
		b := 1
		for a := nb; a >= 1; a-- {
			b *= a
		}
		return b
	}
	// If nb is greater than or equal to 25, return 0 (to avoid integer overflow)
	return 0
}
