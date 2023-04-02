package piscine

// ForEach applies the given function f to each element of the input slice arr.
func ForEach(f func(int), arr []int) {
	for i := range arr { // Iterate over the indices of the input slice arr
		f(arr[i]) // Apply the function f to the i-th element of arr
	}
}
