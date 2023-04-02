package piscine

import "github.com/01-edu/z01" // Import the z01 library for printing

// Converts an integer to its string representation
func intToString(n int) string {
	if n == 0 {
		return "0"
	} else if n == 1 {
		return "1"
	} else if n == 2 {
		return "2"
	} else if n == 3 {
		return "3"
	} else if n == 4 {
		return "4"
	} else if n == 5 {
		return "5"
	} else if n == 6 {
		return "6"
	} else if n == 7 {
		return "7"
	} else if n == 8 {
		return "8"
	} else {
		return "9"
	}
}

// Converts an integer to its string representation recursively
func itoa(n int) string {
	var result string

	// If n is zero, return "0"
	if n == 0 {
		return "0"
	}

	// While n is greater than zero
	for n > 0 {
		// Get the last digit of n and append it to the result string
		result = intToString(n%10) + result
		// Divide n by 10 to remove the last digit
		n /= 10
	}

	// Return the resulting string
	return result
}

// Prints a string to the console using the z01 library
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

// Recursively generates and prints all combinations of n digits that start with prev
func printComb(n int, prev int, result string, count *int) {
	for i := 0; i < 10; i++ {
		// If i is greater than prev
		if prev < i {
			// If n is 1, print the current combination with a comma separator
			if n == 1 {
				if *count > 0 {
					printString(", ")
				}
				printString(result + itoa(i))
				*count++
				// If n is greater than 1, continue generating combinations recursively
			} else {
				printComb(n-1, i, result+itoa(i), count)
			}
		}
	}
}

// Generates and prints all combinations of n digits that don't have repeated digits
func PrintCombN(n int) {
	var count int = 0
	for i := 0; i < 10; i++ {
		// If n is greater than 1, generate combinations recursively
		if n > 1 {
			printComb(n-1, i, itoa(i), &count)
			// If n is 1, print the current digit with a comma separator
		} else {
			if count > 0 {
				printString(", ")
			}
			printString(itoa(i))
			count++
		}
	}

	// Print a newline character at the end
	printString("\n")
}
