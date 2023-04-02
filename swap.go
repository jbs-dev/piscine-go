package piscine

// Swap takes two integer pointers as input and swaps the values they point to.
func Swap(a *int, b *int) {
	// In this line, the value pointed to by 'a' is assigned to a temporary variable,
	// then the value pointed to by 'b' is assigned to the value pointed to by 'a',
	// and finally, the value in the temporary variable is assigned to the value
	// pointed to by 'b'. This effectively swaps the values pointed to by 'a' and 'b'.
	*a, *b = *b, *a
}
