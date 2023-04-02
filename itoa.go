package piscine

// Itoa converts an integer 'n' to a string and returns the resulting string.
func Itoa(n int) string {
	// Initialize an empty string to hold the result
	var result string
	// Determine the sign of the number and make it positive for simplicity
	sign := 1
	if n < 0 {
		sign = -1
		n = -n
	}
	// If the number is 0, return "0"
	if n == 0 {
		return "0"
	}
	// Convert the number to a string by repeatedly taking the last digit and adding it to the result string
	for n > 0 {
		digit := n % 10
		result = string(rune('0'+digit)) + result
		n /= 10
	}
	// If the number was originally negative, add a "-" sign to the beginning of the result string
	if sign == -1 {
		result = "-" + result
	}
	// Return the result string
	return result
}
