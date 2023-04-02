package piscine

// IsPrime returns true if 'nb' is a prime number, and false otherwise.
func IsPrime(nb int) bool {
	// Check if 'nb' is greater than zero
	if nb > 0 {
		// If 'nb' is 1, it is not prime
		if nb == 1 {
			return false
		}
		// If 'nb' is 2 or 3, it is prime
		if nb == 2 || nb == 3 {
			return true
		}
		// If 'nb' is divisible by 2 or 3, it is not prime
		if nb%3 == 0 || nb%2 == 0 {
			return false
		}
		// Iterate over all integers from 2 to 'nb'-1
		for a := 2; a < nb; a++ {
			// If 'nb' is divisible by any integer between 2 and 'nb'-1, it is not prime
			if nb%a == 0 {
				return false
			}
		}
		// If 'nb' is not divisible by any integer between 2 and 'nb'-1, it is prime
		return true
	} else {
		// If 'nb' is less than or equal to zero, it is not prime
		return false
	}
}
