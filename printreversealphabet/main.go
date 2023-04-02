package main

import "github.com/01-edu/z01"

func main() {
	// Start with the ASCII value of 'z'.
	i := 122
	// Loop from 'z' down to 'a'.
	for i >= 97 {
		// Print the corresponding character for the current ASCII value.
		z01.PrintRune(rune(i))
		// Decrement the ASCII value by 1.
		i--
	}
	// Print a newline character at the end of the output.
	z01.PrintRune('\n')
}
