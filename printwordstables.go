package piscine

import "github.com/01-edu/z01"

// PrintWordsTables that takes a slice of strings as input and prints each string to the console,
// one character at a time, followed by a newline character
func PrintWordsTables(table []string) {
	// Iterate over each element in the input slice of strings
	for _, v := range table {
		// Iterate over each rune in the current string
		for _, v2 := range v {
			// Print the current rune using the z01.PrintRune() function
			z01.PrintRune(v2)
		}
		// Print a newline character after each string is printed
		z01.PrintRune('\n')
	}
}
