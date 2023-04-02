package piscine

// Compare compares two strings and returns an integer based on their lexicographical order
func Compare(a, b string) int {
	if a == b { // if the two strings are equal, return 0
		return 0
	} else if a < b { // if string a is less than string b, return -1
		return -1
	} else { // if string a is greater than string b, return 1
		return 1
	}
}
