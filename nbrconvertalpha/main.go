package main

import (
	"os"

	"github.com/01-edu/z01"
)

// main function
func main() {
	// get command-line arguments
	my_arr := os.Args[1:]
	if len(my_arr) >= 1 {
		// check if the flag "--upper" is present
		if my_arr[0] == "--upper" {
			// iterate over the rest of the arguments
			for _, numStr := range my_arr[1:] {
				num := 0
				// convert the string representation of a number to an integer
				for _, digit := range numStr {
					num = num*10 + int(digit-'0')
				}
				// if the number is between 1 and 26, print the corresponding uppercase letter
				if num >= 1 && num <= 26 {
					z01.PrintRune('A' + rune(num-1))
				} else {
					// otherwise, print a space
					z01.PrintRune(' ')
				}
			}
		} else {
			// iterate over the arguments
			for _, numStr := range my_arr {
				num := 0
				// convert the string representation of a number to an integer
				for _, digit := range numStr {
					num = num*10 + int(digit-'0')
				}
				// if the number is between 1 and 26, print the corresponding lowercase letter
				if num >= 1 && num <= 26 {
					z01.PrintRune('a' + rune(num-1))
				} else {
					// otherwise, print a space
					z01.PrintRune(' ')
				}
			}
		}
		// print a newline character at the end
		z01.PrintRune('\n')
	}
}
