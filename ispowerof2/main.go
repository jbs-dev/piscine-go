package main

import (
	"os"
	"strconv"
)

// The main function expects exactly one argument.
// If there is no argument or more than one argument, the program exits with status code 0.
// The argument must be a valid integer, otherwise the program panics with an error.
// The function then checks if the integer is a power of 2 or not and writes the result to stdout.
func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	// To check if a number is a power of 2, we use a bitwise AND operation between the number and (number-1).
	// If the result is 0, then the number is a power of 2.
	isPowerOf2 := number&(number-1) == 0

	if isPowerOf2 {
		os.Stdout.Write([]byte("true\n"))
	} else {
		os.Stdout.Write([]byte("false\n"))
	}
}
