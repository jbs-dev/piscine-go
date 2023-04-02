package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	// First digit loop
	for a := '0'; a <= '9'; a = a + 1 {
		// Second digit loop
		for b := '0'; b <= '9'; b = b + 1 {
			// Initialize c and d for each pair of digits
			d := b + 1
			// Third digit loop
			for c := a; c <= '9'; c = c + 1 {
				// Fourth digit loop
				for ; d <= '9'; d = d + 1 {
					// Print the combination
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(' ')
					z01.PrintRune(c)
					z01.PrintRune(d)
					// Add a comma and space if it's not the last combination
					if a < '9' || b < '8' || c < '9' || d < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				// Reset d for the next iteration
				d = '0'
			}
		}
	}
	// Add a newline at the end of the output
	z01.PrintRune('\n')
}
