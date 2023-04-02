package piscine

// BasicJoin concatenates a slice of strings into a single string
func BasicJoin(strs []string) string {
	result := ""               // initialize a variable to hold the final concatenated string
	for _, str := range strs { // iterate over the slice of strings
		result += str // concatenate each string to the final result
	}
	return result // return the final concatenated string
}
