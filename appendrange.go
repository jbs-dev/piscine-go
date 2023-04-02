package piscine

// AppendRange returns an array of integers containing all values between min (inclusive) and max (exclusive)
func AppendRange(min, max int) []int {
	var array []int // initialize an empty integer array
	if max > min {  // only execute the loop if max is greater than min
		for i := min; i < max; i++ { // iterate over all integers between min and max
			array = append(array, i) // append the current integer to the array
		}
	}
	return array // return the array of integers
}
