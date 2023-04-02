package piscine

import "github.com/01-edu/z01"

// PrintComb prints all the possible three-digit numbers in ascending order where the three digits are distinct.
// Each number is separated by a comma and a space, except for the last number, which is followed by a newline character.
func PrintComb() {
	// Loop over all possible values of the first digit
	for a := '0'; a <= '7'; a++ {
		// Loop over all possible values of the second digit
		for b := a + 1; b <= '8'; b++ {
			// Loop over all possible values of the third digit
			for c := b + 1; c <= '9'; c++ {
				// If the three digits are distinct and in ascending order, print the number
				if a < b && b < c {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
					// If this is the last number, print a newline character
					if a == '7' && b == '8' && c == '9' {
						z01.PrintRune(10)
					} else {
						// Otherwise, print a comma and a space
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}
