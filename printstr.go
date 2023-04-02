package piscine

import "github.com/01-edu/z01"

// PrintStr that takes a string as input and prints it to the console one rune at a time
func PrintStr(s string) {
	// Iterate over each rune in the input string s
	for _, r := range s {
		// Print the current rune using the z01.PrintRune() function
		z01.PrintRune(rune(r))
	}
}
