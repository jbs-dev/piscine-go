package piscine

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	// Start a nested loop to iterate through all possible pairs of elements in the array
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			// If the result of the function 'f' is 1 for the current pair of elements,
			// swap the elements using the 'modify' function
			if f(array[i], array[j]) == 1 {
				temp := array[i]
				modify(array, i, array[j])
				modify(array, j, temp)
			}
		}
	}
}

func modify(array []string, ind int, val string) {
	// Modify the element at the given index with the new value
	array[ind] = val
}
