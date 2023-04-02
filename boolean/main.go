package main

import (
	"os"

	"github.com/01-edu/z01"
)

// printStr takes a string (s) as input and prints it to the standard output followed by a newline character.
func printStr(s string) {
	// Iterate through the string's characters (runes).
	for _, r := range s {
		// Print each character (rune) using the PrintRune function from the z01 package.
		z01.PrintRune(r)
	}
	// Print a newline character at the end of the string.
	z01.PrintRune('\n')
}

// isEven takes an integer (nbr) as input and returns true if it is even, and false if it is odd.
func isEven(nbr int) bool {
	if (nbr)%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	args := os.Args[1:] // Get command-line arguments, excluding the program name.
	count := 0          // Initialize a counter variable to count the number of arguments.

	// Iterate through the arguments to count them.
	for i := range args {
		count = i + 1
	}

	// Call isEven with the number of arguments (count) to determine if there are an even or odd number of arguments.
	// Depending on the result, call printStr with the corresponding message.
	if isEven(count) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
