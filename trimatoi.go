package piscine

// TrimAtoi takes a string as input and returns its integer representation,
// ignoring any non-digit characters before the first digit and after the last digit.
// If a minus sign is present immediately before the first digit, the returned integer will be negative.
func TrimAtoi(s string) int {
	var array []int      // Array to store the digits found in the string.
	result := 0          // Final integer result.
	minusIndex := -1     // Index of the last minus sign found.
	firstDigitIndex := 0 // Index of the first digit found.
	index := 0           // Current index in the input string.
	arrayCount := 0      // Count of digits found in the input string.

	// Iterate through the input string.
	for _, rune := range s {
		// If the current character is a minus sign, store its index.
		if rune == '-' {
			minusIndex = index
		}
		// If the current character is a digit, store the digit in the array.
		if isDigit(rune) {
			// If this is the first digit found, store its index.
			if firstDigitIndex == 0 {
				firstDigitIndex = index
			}
			// Append the digit to the array.
			array = append(array, int(rune-'0'))
		}
		// Increment the index for the next iteration.
		index++
	}

	// Calculate the count of digits found in the input string.
	// The range keyword is used to iterate over the array and count the number of elements.
	for count := range array {
		arrayCount = count + 1
	}

	// Convert the array of digits to the final integer result.
	// Iterate through the array, multiplying the current result by 10 and adding the current digit.
	// This way, the digits are combined into a single integer.
	for i := 0; i < arrayCount; i++ {
		result = result*10 + array[i]
	}

	// If a minus sign is found before the first digit, negate the result.
	// The condition (minusIndex != -1) ensures that a minus sign was actually found in the string.
	if minusIndex < firstDigitIndex && minusIndex != -1 {
		result = result * -1
	}
	return result
}

/* isDigit checks if a given rune is a digit (between '0' and '9') and returns true if it is, false otherwise.
// This helper function is used in the main TrimAtoi function to identify digits in the input string.
func isDigit(digit rune) bool {
	// Iterate through the range of digit characters from '0' to '9'.
	for a := '0'; a <= '9'; a++ {
		// If the input rune matches the current digit character, return true.
		if digit == a {
			return true
		}
	}
	// If the input rune did not match any digit characters, return false.
	return false
} */
