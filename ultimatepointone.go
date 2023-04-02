package piscine

// UltimatePointOne takes a pointer to a pointer to a pointer to an integer (n) as input.
// It sets the value of the integer pointed to by the triple pointer to 1.
func UltimatePointOne(n ***int) {
	***n = 1 // Dereference the triple pointer three times to access the integer it points to, and set its value to 1.
}
