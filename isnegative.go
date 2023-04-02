package piscine

import "github.com/01-edu/z01"

// IsNegative prints 'T' if 'nb' is negative, and 'F' otherwise.
// It then prints a newline character.
func IsNegative(nb int) {
	// If 'nb' is non-negative, print 'F'
	if nb >= 0 {
		z01.PrintRune('F')
		// Otherwise, print 'T'
	} else {
		z01.PrintRune('T')
	}
	// Print a newline character
	z01.PrintRune('\n')
}
