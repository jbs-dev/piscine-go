package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lens := 0            // Initialize a counter variable to count the number of command-line arguments.
	argum := os.Args[1:] // Get command-line arguments, excluding the program name.

	// Increment the counter for each command-line argument.
	for range argum {
		lens++
	}

	// For each command-line argument, try to read the content of the corresponding file.
	for i := 1; i <= lens; i++ {
		// Read the file specified by the current argument.
		file, err := os.ReadFile(os.Args[i])

		// Convert the file content to a string.
		strFile := string(file)

		// If an error occurred while reading the file, print an error message and exit with a non-zero status code.
		if err != nil {
			string := "ERROR: " + "open " + os.Args[i] + ": no such file or directory"

			// Print the error message one rune at a time.
			for _, v := range string {
				z01.PrintRune(v)
			}

			// Print a newline character after the error message.
			z01.PrintRune('\n')

			// Exit the program with a non-zero status code to indicate an error.
			os.Exit(1)
		} else {
			// If no error occurred, print the file content one rune at a time.
			for _, v := range strFile {
				z01.PrintRune(v)
			}
		}
	}

	// If no arguments were provided, read from standard input and print the content.
	if lens == 0 {
		// Keep reading from standard input until EOF is reached or the read content is empty.
		for {
			// Read from standard input using io.ReadAll instead of ioutil.ReadAll.
			text, err := io.ReadAll(os.Stdin)

			// If an EOF (End Of File) error occurs or the read content is empty, break the loop.
			if err == io.EOF || len(text) == 0 {
				break
			}

			// Print the content read from standard input one rune at a time.
			for _, v := range string(text) {
				z01.PrintRune(v)
			}
		}
	}
}
