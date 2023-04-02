package piscine

import "github.com/01-edu/z01"

// Prints an integer to the console
func PrintNbr(n int) {
	// If n is negative, print a '-' character
	if n < 0 {
		z01.PrintRune('-')
		// Convert n to its positive value to handle negative numbers
		n = -n
	}

	// If n is zero, print '0'
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Handle each digit of n separately using the check() function
	for n > 0 {
		check(n % 10)
		n /= 10
	}
}

// Converts an integer to its string representation and prints it to the console
func check(r int) {
	c := '0'

	// If r is zero, print '0' and return
	if r == 0 {
		z01.PrintRune(c)
		return
	}

	// Increment c for each digit in r%10
	for i := 1; i <= r%10; i++ {
		c++
	}

	// Decrement c for each digit less than r%10 (to handle negative numbers)
	for i := -1; i >= r%10; i-- {
		c++
	}

	// If r/10 is not zero, continue recursively with r/10
	if r/10 != 0 {
		check(r / 10)
	}

	// Print the resulting rune c
	z01.PrintRune(c)
}
