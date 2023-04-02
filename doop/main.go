package main

import "os"

func main() {
	// Check if there are exactly 4 command-line arguments (including the program name).
	// It's important to have exactly two operands and one operator for the program to function correctly.
	if len(os.Args) == 4 {
		// Check if both the operands are valid numbers.
		// This function ensures that the input strings can be safely converted to integers.
		if validArgs(os.Args[1], os.Args[3]) {
			// Convert the operands to integers using the atoi function.
			a := atoi(os.Args[1])
			operator := os.Args[2]
			b := atoi(os.Args[3])

			// Perform the operation specified by the operator and output the result.
			// The switch statement checks the operator and performs the corresponding operation.
			switch operator {
			case "+":
				answer := a + b
				// Check if the addition did not result in an integer overflow.
				if answer >= a {
					os.Stdout.WriteString(itoa(answer) + "\n")
				}
			case "-":
				answer := a - b
				// Check if the subtraction did not result in an integer underflow.
				if answer <= a {
					os.Stdout.WriteString(itoa(answer) + "\n")
				}
			case "/":
				// Division by zero is undefined, so we check if the divisor is zero.
				if b == 0 {
					os.Stdout.WriteString("No division by 0\n")
				} else {
					answer := a / b
					os.Stdout.WriteString(itoa(answer) + "\n")
				}
			case "*":
				answer := a * b
				// Check if the multiplication did not result in an integer overflow.
				if answer/a == b {
					os.Stdout.WriteString(itoa(answer) + "\n")
				}
			case "%":
				// Modulo by zero is undefined, so we check if the divisor is zero.
				if b == 0 {
					os.Stdout.WriteString("No modulo by 0\n")
				} else {
					answer := a % b
					os.Stdout.WriteString(itoa(answer) + "\n")
				}
			}
		}
	}
}

// validArgs checks if the given strings are valid integer numbers.
// It checks if both input strings consist of digits or a leading '-' sign for negative numbers.
func validArgs(value1, value2 string) bool {
	// Return false if any of the input strings is too large for a valid int64.
	if value1 == "-9223372036854775809" || value2 == "-9223372036854775809" {
		return false
	}
	// Check if all characters in the first input string are either digits or a leading '-' sign.
	for _, v := range value1 {
		if !(v >= '0' && v <= '9') && v != '-' {
			return false
		}
	}
	// Check if all characters in the second input string are either digits or a leading '-' sign.
	for _, v := range value2 {
		if !(v >= '0' && v <= '9') && v != '-' {
			return false
		}
	}
	return true
}

// atoi converts a string to an integer.
// It handles positive and negative numbers, ensuring correct sign and magnitude.
func atoi(n string) int {
	isPositive := false
	num := ""
	result := 0

	// Check if the number is positive or negative.
	// If positive, set isPositive to true and use the input string as is.
	// If negative, remove the '-' sign and use the rest of the input string.
	if n[0] != '-' {
		isPositive = true
		num = n
	} else {
		num = n[1:]
	}

	// Iterate through each character of the input string (excluding the sign).
	// Calculate the integer value of the character and accumulate the result.
	for i, v := range num {
		if i != len(num)-1 {
			result += int(v - '0') // Subtract '0' to get the integer value of the character.
			result *= 10           // Shift the result left by one digit to make room for the next digit.
		} else {
			result += int(v - '0')
		}
	}

	// If the number was positive, return the result as is.
	// If the number was negative, return the result negated.
	if isPositive {
		return result
	}
	return -result
}

// itoa converts an integer to a string.
// It handles positive and negative numbers, ensuring correct sign and digit sequence.
func itoa(n int) string {
	isNegative := false
	result := ""

	// Check if the number is positive or negative.
	// If positive, negate the input number and set isNegative to true.
	// This allows us to handle positive and negative numbers using the same loop.
	if n > 0 {
		n = n * -1
		isNegative = true
	}

	// Build the output string in reverse order, starting from the least significant digit.
	for n <= -10 {
		mod := -n % 10                          // Calculate the least significant digit.
		result = string(rune('0'+mod)) + result // Add the digit as a character to the output string.
		n = n / 10                              // Remove the least significant digit from the input number.
	}
	result = string(rune('0'+(n*-1))) + result // Add the last digit to the output string.

	// If the number was negative, add a '-' sign to the output string.
	if !isNegative {
		if result != "0" {
			result = "-" + result
			return result
		}
	}
	return result
}
