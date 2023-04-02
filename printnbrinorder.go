package piscine

import "github.com/01-edu/z01"

// PrintNbrInOrder that takes an integer as input and prints its digits in ascending order
func PrintNbrInOrder(n int) {
	// If n is 0, simply print 0 and return
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// If n is positive, extract its digits and store them in an array
	if n > 0 {
		var array []int
		eachValue := 0

		// Repeatedly divide n by 10 and take the remainder to extract each digit
		for n != 0 {
			eachValue = n % 10
			n /= 10
			array = append(array, eachValue)
		}

		// Implement bubble sort to sort the extracted digits in ascending order
		for i := 0; i < len(array)-1; i++ {
			for j := 0; j < len(array)-i-1; j++ {
				if array[j] > array[j+1] {
					// Swap adjacent elements if they are in the wrong order
					array[j], array[j+1] = array[j+1], array[j]
				}
			}
		}

		// Print the sorted digits
		for i := 0; i < len(array); i++ {
			z01.PrintRune(rune(array[i] + 48)) // Convert each digit to a rune by adding 48 to its integer value
		}
	}
}

// Bubble sort is a simple sorting algorithm that repeatedly steps through the list,
// compares adjacent elements and swaps them if they are in the wrong order.
// 	The pass through the list is repeated until the list is sorted.
//		The algorithm gets its name from the way smaller elements "bubble" to the top of the list.
//			In this implementation, we start with an unsorted list of digits extracted from the input number,
//				and we repeatedly swap adjacent elements until the list is sorted in ascending order.
