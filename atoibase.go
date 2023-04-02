package piscine

// AtoiBase converts a string to an integer using a given base represented by a string of unique digits
func AtoiBase(s string, t string) int {
	res := 0               // initialize a variable to hold the final integer value
	lens := 0              // initialize a variable to hold the length of the base string
	arr := map[rune]bool{} // initialize a map to hold the unique digits in the base string
	for _, c := range t {  // iterate over the runes in the base string
		if arr[c] || c == '-' || c == '+' { // if the current rune is a duplicate, a plus sign, or a minus sign, return 0
			return res
		}
		arr[c] = true // add the current rune to the map of unique digits
		lens++        // increment the length variable
	}
	if lens > 1 { // only execute the loop if the length of the base string is greater than 1
		for _, v := range s { // iterate over the runes in the input string
			count := 0  // initialize a variable to hold the value of the current digit
			if arr[v] { // if the current rune is a valid digit in the base string
				for _, j := range t { // iterate over the runes in the base string
					if j == v { // if the current rune matches the current digit in the input string
						break // stop iterating over the base string
					}
					count++ // increment the value of the current digit for each rune in the base string that does not match the current digit
				}
				res = res*lens + count // update the final integer value by adding the value of the current digit multiplied by the current power of the base
			}
		}
	}
	return res // return the final integer value
}
