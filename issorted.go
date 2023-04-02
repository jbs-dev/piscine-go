package piscine

// IsSorted returns true if the given integer slice 'tab' is sorted in ascending or descending order,
// as determined by the function 'f', and false otherwise.
func IsSorted(f func(a, b int) int, tab []int) bool {
	// Count the length of 'tab'
	length := 0
	for i := range tab {
		length = i + 1
	}

	// Initialize variables to track whether the slice is sorted in ascending or descending order
	asciending := true
	descending := true

	// Check if 'tab' is sorted in descending order
	for i := 1; i < length; i++ {
		if !(f(tab[i-1], tab[i]) >= 0) {
			descending = false
		}
	}

	// Check if 'tab' is sorted in ascending order
	for i := 1; i < length; i++ {
		if !(f(tab[i-1], tab[i]) <= 0) {
			asciending = false
		}
	}

	// Return true if 'tab' is sorted in ascending or descending order, and false otherwise
	return asciending || descending
}
