package piscine

import "github.com/01-edu/z01"

// PrintSolution prints the resulting 8-queens solution to the console
func PrintSolution(result [8]int) {
	for _, val := range result {
		z01.PrintRune(rune(val) + '0') // convert each integer value in the result array to a rune and print it to the console
	}
	z01.PrintRune('\n') // print a newline character to the console after printing the result array
}

// PlaceQueen recursively places the queens on the chess board and finds the 8-queens solution
func PlaceQueen(horiz [8]bool, fDiag [15]bool, bDiag [15]bool, result [8]int, col int) {
	if col == 8 {
		return // if all 8 queens have been placed, return from the function
	}
	for row := 0; row <= 7; row++ { // iterate over all rows in the current column
		if horiz[row] || bDiag[row+col] || fDiag[7-row+col] {
			// if there is already a queen in the same row, or if there is already a queen in the same diagonal (forward or backward), skip this row
			continue
		}
		// if there is no queen in the same row or diagonal, continue with the recursive call to PlaceQueen with updated parameters
		horizTmp := horiz  // create a temporary copy of the horiz array
		fDiagTmp := fDiag  // create a temporary copy of the fDiag array
		bDiagTmp := bDiag  // create a temporary copy of the bDiag array
		resulTmp := result // create a temporary copy of the result array

		horizTmp[row] = true       // mark the current row as occupied by a queen
		fDiagTmp[7-row+col] = true // mark the current forward diagonal as occupied by a queen
		bDiagTmp[row+col] = true   // mark the current backward diagonal as occupied by a queen
		resulTmp[col] = row + 1    // add the current row number to the result array (add 1 to convert from 0-indexed to 1-indexed)

		if col == 7 {
			PrintSolution(resulTmp) // if all 8 queens have been placed, print the resulting solution to the console
			return
		}
		// continue with the recursive call to PlaceQueen with updated parameters
		PlaceQueen(horizTmp, fDiagTmp, bDiagTmp, resulTmp, col+1)
	}
}

// EightQueens finds and prints the solution for the 8-queens problem
func EightQueens() {
	var horiz [8]bool                         // initialize an array to keep track of the occupied rows
	var fDiag [15]bool                        // initialize an array to keep track of the occupied forward diagonals
	var bDiag [15]bool                        // initialize an array to keep track of the occupied backward diagonals
	var table [8]int                          // initialize an array to store the resulting solution
	PlaceQueen(horiz, fDiag, bDiag, table, 0) // start the recursive process of placing the queens on the chess board and finding the solution
}
