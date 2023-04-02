package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get the name of the current executable from the first command-line argument.
	name := []rune(os.Args[0])[2:]
	// Iterate over each character in the name and print it to the console.
	for _, val := range name {
		z01.PrintRune(val)
	}
	// Print a newline character to the console to move to the next line.
	z01.PrintRune('\n')
}
