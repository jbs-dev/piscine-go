package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

// PrintInfo prints the instructions for how to use the program.
func PrintInfo() {
	fmt.Print("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.")
}

// Sort sorts the string passed to it and prints the sorted string to the console.
func Sort(s string) {
	// Create an array of 1000 integers to store the frequency of each character in the string.
	var array [1000]int
	// Iterate over each character in the string and update the corresponding frequency in the array.
	for _, char := range s {
		array[int(char)]++
	}
	// Iterate over each index in the array.
	for index, char := range array {
		// Iterate over each occurrence of the character at the current index.
		for char > 0 {
			// Print the character to the console.
			z01.PrintRune(rune(index))
			// Decrement the count of the current character.
			char--
		}
	}
	// Print a newline character at the end of the string.
	z01.PrintRune('\n')
}

func main() {
	// Get the command-line arguments.
	argArray := os.Args[1:]
	// Initialize variables.
	str := ""
	check := false
	SortIt := false
	// Iterate over each argument.
	for _, element := range argArray {
		// Set check to true.
		check = true
		// If the argument is "-h" or "--help", print the instructions and exit the loop.
		if element == "-h" || element == "--help" {
			PrintInfo()
			break
		}
		// Get the length of the current argument.
		len := 0
		for i := range element {
			len = i + 1
		}
		// If the argument is not empty:
		if len > 0 {
			// If the argument is a flag, update str or SortIt as appropriate.
			if element[0] == '-' {
				if len > 2 && element[2] == 'i' {
					if len > 8 {
						str += element[9:]
					}
				}
				if element[1] == 'i' {
					if len > 3 {
						str += element[3:]
					}
				}
				// If the argument is not a flag, add it to str.
			} else {
				str = element + str
			}
		}
		// If the argument is "-o" or "--order", set SortIt to true.
		if element == "-o" || element == "--order" {
			SortIt = true
		}
	}
	// If no arguments were provided, print the instructions.
	if !check {
		PrintInfo()
	}
	// If the "-o" or "--order" flag was provided, sort the string and print it to the console.
	if SortIt {
		Sort(str)
		// If the "-o" or "--order" flag was not provided, print the string to the console.
	} else {
		fmt.Println(str)
	}
}
