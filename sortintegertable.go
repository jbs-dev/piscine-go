package piscine

// SortIntegerTable that takes a slice of integers as input and sorts it in ascending order
func SortIntegerTable(table []int) {
	// Initialize a variable i to 1
	i := 1
	// Use a while loop to repeatedly iterate over the input slice until it is sorted
	for i < len(table) {
		// If the current element is less than the previous element, swap them and reset i to 1
		if table[i-1] > table[i] {
			tmp := table[i]
			table[i] = table[i-1]
			table[i-1] = tmp
			i = 1
			// If the current element is greater than or equal to the previous element, increment i by 1
		} else {
			i++
		}
	}
}
