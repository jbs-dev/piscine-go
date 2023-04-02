package piscine

// PointOne takes a pointer to an integer as its parameter and sets the value of the integer to 1.
func PointOne(n *int) {
	// Use the dereferencing operator (*) to set the value of the integer pointed to by 'n' to 1
	*n = 1
}
