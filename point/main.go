package main

import "github.com/01-edu/z01"

// Define a point struct with integer x and y coordinates.
type point struct {
	x int
	y int
}

// Set the x and y coordinates of a point.
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	// Create a pointer to a point.
	points := &point{}

	// Set the coordinates of the point using the setPoint function.
	setPoint(points)

	// Print the x coordinate.
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	IntoRune(points.x)
	// Print a comma and a space.
	z01.PrintRune(',')
	z01.PrintRune(' ')
	// Print the y coordinate.
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	IntoRune(points.y)
	// Print a newline character.
	z01.PrintRune('\n')
}

// Recursive function that converts an integer to a rune.
func check(r int) {
	// Start with '0' as the character to be printed.
	c := '0'
	// If r is 0, print '0' and return.
	if r == 0 {
		z01.PrintRune(c)
		return
	}
	// If r is negative, print '-' and set r to its absolute value.
	if r < 0 {
		z01.PrintRune('-')
		r = -r
	}
	// For each digit in r, increment c to the appropriate character.
	for i := 1; i <= r%10; i++ {
		c++
	}
	// If the digit is negative, decrement c to the appropriate character.
	for i := -1; i >= r%10; i-- {
		c++
	}
	// If there are more digits, recurse with the remaining digits of r.
	if r/10 != 0 {
		check(r / 10)
	}
	// Print the character for the current digit.
	z01.PrintRune(c)
}

// Convert an integer to a rune and print it.
func IntoRune(n int) {
	// If n is negative, print '-' and convert the absolute value of n.
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	// Convert n to a rune using the check function.
	check(n)
}
