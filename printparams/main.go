package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get all command-line arguments provided to the program
	arg := os.Args

	// Iterate over all arguments except for the first one (which is the name of the program)
	for i, s := range arg {
		if i != 0 {
			// Iterate over each character in the current argument and print it using z01
			for _, a := range s {
				z01.PrintRune(a)
			}
			// Print a newline character at the end of each argument
			z01.PrintRune('\n')
		}
	}
}
