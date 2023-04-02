package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// main function reads command line arguments, checks for errors and calls ztail function with specified arguments.
func main() {
	// read command line arguments
	args := os.Args[1:]

	// if no arguments are given, print the program name
	if len(args) == 0 {
		fmt.Println(os.Args[0])
	} else {
		// convert the first argument to an integer
		num, err := strconv.Atoi(os.Args[1])

		// if there is an error converting the argument to an integer, print the error message
		if err != nil {
			fmt.Println(err.Error())
		} else {
			// check if there is at least one filename argument provided
			if len(os.Args[2:]) == 0 {
				fmt.Println("File name missing")
			} else {
				// call the ztail function with the provided arguments
				for _, res := range os.Args[2:] {
					ztail(res, num)
				}
			}
		}
	}
}

// ztail function prints the last `numBytes` bytes of the file with the given filename `s`.
func ztail(s string, numBytes int) {
	// open the file
	file, err := os.Open(s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close() // make sure to close the file

	// get the size of the file
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fileSize := fileInfo.Size()

	// calculate the number of bytes to read from the end of the file
	offset := fileSize - int64(numBytes)
	if offset < 0 {
		offset = 0
	}

	// seek to the offset position from the end of the file
	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// create a buffer to read the last `numBytes` of the file
	buffer := make([]byte, numBytes)
	_, err = io.ReadFull(file, buffer)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// print the buffer to the console
	fmt.Printf("%s\n", buffer)
}
