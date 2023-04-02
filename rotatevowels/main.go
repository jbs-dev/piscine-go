package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get the arguments passed to the program, ignoring the first argument (which is the name of the program itself).
	argruments := os.Args[1:]

	// Determine the number of arguments passed to the program.
	length := 0
	for i := range argruments {
		length = i + 1
	}

	// If there are arguments, proceed with the program.
	if length != 0 {
		// Create an empty string to store the concatenated arguments and a boolean to keep track of whether the current argument is the first one.
		str := ""
		first := false

		// Concatenate the arguments into a single string with spaces between them.
		for _, arg := range argruments {
			if first {
				str += " "
			}
			first = true
			str += arg
		}

		// Convert the string to an array of runes and create two empty slices: one to store the positions of vowels in the string, and another to store the vowels themselves.
		runes := []rune(str)
		var pos []int
		var vow []rune

		// Iterate over the array of runes and find the positions and values of all the vowels.
		for i, r := range runes {
			if r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' || r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
				pos = append(pos, i)
				vow = append(vow, runes[i])
			}
		}

		// Reverse the order of the vowels.
		swap(vow)

		// Replace the vowels in the string with the reversed vowels.
		for i := range pos {
			runes[pos[i]] = vow[i]
		}

		// Print the modified string.
		for _, r := range runes {
			z01.PrintRune(r)
		}
	}
	// Print a newline character at the end of the program.
	z01.PrintRune('\n')
}

// Reverse the order of elements in the slice.
func swap(vow []rune) {
	// Get the length of the slice.
	len := 0
	for i := range vow {
		len = i + 1
	}
	// Iterate over the slice and swap the positions of the elements.
	for i, j := 0, len-1; i < j; i, j = i+1, j-1 {
		temp := vow[i]
		vow[i] = vow[j]
		vow[j] = temp
	}
}
