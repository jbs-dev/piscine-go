package piscine

// Any returns true if at least one element in the given string array satisfies the given function f
func Any(f func(string) bool, arr []string) bool {
	for _, s := range arr { // iterate over each element in the string array
		if f(s) { // if the given function f returns true for the current element
			return true // return true to indicate that at least one element satisfies the function
		}
	}
	return false // if the function f returns false for all elements, return false to indicate that none satisfy the function
}
