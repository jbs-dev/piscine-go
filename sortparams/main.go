package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get the command line arguments.
	arg := os.Args

	// Count the number of arguments.
	count := 0
	for s := range arg {
		count = s + 1
	}

	// Sort the arguments.
	for i := 1; i < count; i++ {
		for j := i + 1; j < count; j++ {
			// Compare the current argument with the next one and swap if necessary.
			if arg[i] > arg[j] {
				arg[i], arg[j] = arg[j], arg[i]
			}
		}
	}

	// Print the sorted arguments.
	for j := 1; j <= count-1; j++ {
		// Print each character in the argument followed by a newline.
		for _, element := range arg[j] {
			z01.PrintRune(element)
		}
		z01.PrintRune('\n')
	}
}
