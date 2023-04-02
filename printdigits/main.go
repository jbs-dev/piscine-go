package main

import "github.com/01-edu/z01"

func main() {
	// Set i to the Unicode code point of '0'
	i := '0'
	// Iterate from the Unicode code point of '0' to the Unicode code point of '9'
	for i <= '9' {
		// Print the current rune
		z01.PrintRune(i)
		// Increment i to the next Unicode code point
		i++
	}
	// Print a newline character at the end of the string
	z01.PrintRune('\n')
}
