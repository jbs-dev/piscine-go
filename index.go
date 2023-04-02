package piscine

// Index returns the index of the first occurrence of 'toFind' in 's',
// or -1 if 'toFind' is not found.
func Index(s string, toFind string) int {
	// Count the number of characters in 'toFind'
	toFindCount := arrayCount(toFind)
	// Initialize a variable to keep track of the number of characters in 'toFind' that have been matched so far
	secondIndex := 0
	// Iterate over each character in 'toFind'
	for _, j := range toFind {
		// Iterate over each character in 's'
		for i1, i2 := range s {
			// Check if the current character in 's' matches the current character in 'toFind'
			if i2 == j {
				// If 'toFind' has only one character, return its index
				if toFindCount == 1 {
					return i1
					// If 'toFind' has more than one character
				} else if toFindCount > 1 {
					// Iterate over each character in 'toFind' to check if they match the corresponding character in 's'
					for a := 0; a < toFindCount; a++ {
						if s[i1+a] == toFind[a] {
							secondIndex++
						} else {
							return -1
						}
					}
					// If all characters in 'toFind' match, return the index of the first character in 'toFind'
					if secondIndex == toFindCount {
						return i1
					}
				} else {
					// If 'toFind' has no characters, return -1
					return -1
				}
			}
		}
		// If no characters in 'toFind' have been matched so far, return -1
		if secondIndex <= 0 {
			return -1
		}
	}
	// If 'toFind' has no characters, return the count of characters in 'toFind'
	return toFindCount
}

// arrayCount returns the number of characters in 's'
func arrayCount(s string) int {
	count := 0
	for i := range s {
		count = i + 1
	}
	return count
}
