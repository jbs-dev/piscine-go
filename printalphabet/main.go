package main

import "github.com/01-edu/z01"

func main() {
	// Set the variable i to the Unicode value of lowercase 'a'.
	i := 97
	// Loop through the ASCII values for lowercase letters 'a' to 'z'.
	for i <= 122 {
		// Convert the current ASCII value to a rune and print it to the console using z01.
		z01.PrintRune(rune(i))
		// Increment i to move to the next ASCII value.
		i++
	}
	// Print a newline character at the end of the output.
	z01.PrintRune('\n')
}
