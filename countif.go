package piscine

// CountIf counts the number of elements in a slice for which a given function returns true
func CountIf(f func(string) bool, tab []string) int {
	count := 0              // initialize a variable to hold the count of elements for which the given function returns true
	for _, s := range tab { // iterate over the slice
		if f(s) { // if the given function returns true for the current element, increment the count variable
			count++
		}
	}
	return count // return the count of elements for which the given function returns true
}
