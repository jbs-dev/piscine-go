package piscine

// DivMod calculates the integer division and remainder of two integers and stores the results in two pointers
func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b // calculate the integer division of a and b and store the result in the first pointer
	*mod = a % b // calculate the remainder of a and b and store the result in the second pointer
}
