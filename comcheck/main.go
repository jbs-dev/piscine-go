package main

import (
	"fmt"
	"os"
)

func main() {
	// Define a slice of strings that we want to check for in the command-line arguments.
	words := []string{"01", "galaxy", "galaxy 01"}

	count := 0          // Initialize a counter variable to count the number of matches found.
	args := os.Args[1:] // Get command-line arguments, excluding the program name.

	// Iterate through the command-line arguments.
	for i := range args {
		// For each argument, iterate through the words slice to check if the argument matches any word.
		for _, s := range words {
			// If the current argument matches a word in the words slice, increment the counter.
			if args[i] == s {
				count++
			}
		}
	}

	// If at least one match was found (count >= 1), print an alert message.
	if count >= 1 {
		fmt.Println("Alert!!!")
	}
}
