package piscine

// CollatzCountdown calculates the number of steps it takes to reach 1 in the Collatz sequence
func CollatzCountdown(start int) int {
	if start <= 0 { // if the input is negative or zero, return -1 (as per the function requirements)
		return -1
	}

	steps := 0 // initialize a variable to hold the number of steps required to reach 1

	for start != 1 { // iterate over the sequence until the value of start reaches 1
		if start%2 == 0 { // if start is even, divide it by 2
			start = start / 2
		} else { // if start is odd, multiply it by 3 and add 1
			start = 3*start + 1
		}
		steps++ // increment the number of steps required to reach 1
	}
	return steps // return the number of steps required to reach 1
}
