package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get the command-line arguments.
	arg := os.Args

	// Get the number of arguments passed.
	count := 0
	for s := range arg {
		count = s + 1
	}

	// Loop over the arguments in reverse order.
	for j := count - 1; j >= 1; j-- {
		// Print each argument and add a newline character.
		for _, element := range arg[j] {
			z01.PrintRune(element)
		}
		z01.PrintRune('\n')
	}
}
