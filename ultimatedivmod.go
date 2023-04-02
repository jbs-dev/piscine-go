package piscine

// UltimateDivMod takes two pointers to integers, a and b, as input.
// It divides the integer pointed by a by the integer pointed by b,
// and stores the quotient in the location pointed by a.
// It also calculates the remainder of this division and stores it in the location pointed by b.
func UltimateDivMod(a *int, b *int) {
	y := *a      // Store the original value of the integer pointed by a in the variable y.
	*a = *a / *b // Divide the integer pointed by a by the integer pointed by b, and update the value at the location pointed by a with the quotient.
	*b = y % *b  // Calculate the remainder of the division (original value of a divided by b) and update the value at the location pointed by b with the remainder.
}
