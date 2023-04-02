package main

import (
	"fmt"
	"os"
)

func main() {
	// Get command-line arguments
	arr := os.Args[1:]

	// Check for errors in the arguments
	if checkErrors(arr) {
		fmt.Println("Error")
	} else {
		// Optimize the sudoku input and solve it
		sudoku := optimizeSudoku(arr)
		if solveSudoku(&sudoku, len(sudoku)) {
			// Print the solved sudoku
			printSudoku(sudoku)
		} else {
			fmt.Println("Error")
		}
	}
}

// checkErrors checks for errors in the sudoku input.
// It returns true if an error is found, false otherwise.
func checkErrors(arr []string) bool {
	if len(arr) == 0 {
		return true
	}

	// Check for invalid characters or empty cells
	for i, str := range arr {
		if len(str) == 0 {
			return true
		}
		for _, ch := range arr[i] {
			if !(ch >= '1' && ch <= '9' || ch == '.') {
				return true
			}
		}
	}

	// Check for duplicate numbers in same column
	for column := 0; column < 9; column++ {
		var seenNumbers [10]bool // to keep track of numbers seen in the column
		for row := 0; row < 9; row++ {
			if arr[row][column] != '.' {
				number := runeToInt(rune(arr[row][column]))
				if seenNumbers[number] {
					return true
				}
				seenNumbers[number] = true
			}
		}
	}

	return false
}

// solveSudoku recursively solves the sudoku using backtracking.
// It returns true if the sudoku is solved, false otherwise.
func solveSudoku(arr *[][]int, length int) bool {
	// Check if there are any empty cells
	isEmpty := true
	row := -1
	column := -1
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if (*arr)[i][j] == 0 {
				row = i
				column = j
				isEmpty = false
				break
			}
		}
	}
	if isEmpty {
		return true
	}

	// Try all possible numbers in the empty cell
	for number := 1; number <= 9; number++ {
		if isCorrect(*arr, row, column, number) {
			(*arr)[row][column] = number
			if solveSudoku(arr, length) {
				return true
			} else {
				(*arr)[row][column] = 0
			}
		}
	}

	return false
}

// optimizeSudoku optimizes the input sudoku string.
func optimizeSudoku(arr []string) [][]int {
	sudoku := make([][]int, 9)
	for i := range sudoku {
		sudoku[i] = make([]int, 9)
	}
	for i, str := range arr {
		for j, ch := range str {
			sudoku[i][j] = runeToInt(ch)
		}
	}
	return sudoku
}

// printSudoku prints the solved sudoku.
func printSudoku(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Print(arr[i][j])
			fmt.Print(" ")
		}
	}
}

func isCorrect(arr [][]int, row int, column int, num int) bool {
	// Returns true if the given number is correct in the given row, column and sub-sudoku, otherwise false
	return !checkRow(arr, row, num) && !checkColumn(arr, column, num) && !checkSubSudoku(arr, row-(row%3), column-(column%3), num)
}

func checkRow(arr [][]int, row int, num int) bool {
	// Checks if the given number exists in the given row of the sudoku
	for column := 0; column < len(arr); column++ {
		if arr[row][column] == num {
			return true
		}
	}
	return false
}

func checkColumn(arr [][]int, column int, num int) bool {
	// Checks if the given number exists in the given column of the sudoku
	for row := 0; row < len(arr); row++ {
		if arr[row][column] == num {
			return true
		}
	}
	return false
}

func checkSubSudoku(arr [][]int, rowIndex int, columnIndex int, num int) bool {
	// Checks if the given number exists in the sub-sudoku of the given row and column indices
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			if arr[rowIndex+row][columnIndex+column] == num {
				return true
			}
		}
	}
	return false
}

func runeToInt(number rune) int {
	// Converts a given rune character representing a digit to an integer
	count := 0
	for i := '1'; i <= number; i++ {
		count++
	}
	return count
}
