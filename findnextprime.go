package piscine

// IsPrime2 returns a boolean value indicating whether the given integer is prime or not
func IsPrime2(nb int) bool {
	if nb <= 1 { // If nb is less than or equal to 1, it's not prime
		return false
	}
	for i := 2; i <= nb/2; i++ { // Check whether nb is divisible by any number between 2 and nb/2
		if nb%i == 0 { // If nb is divisible by any number in that range, it's not prime
			return false
		}
	}
	return true // If nb is not divisible by any number in the range, it's prime
}

// FindNextPrime returns the next prime number after the given integer
func FindNextPrime(nb int) int {
	if IsPrime2(nb) { // If nb is already prime, return it
		return nb
	} else { // Otherwise, find the next prime number after nb
		i := nb + 1
		for !(IsPrime2(i)) { // Increment i until the next prime number is found
			i++
		}
		return i // Return the next prime number
	}
}

/* package piscine

func FindNextPrime(nb int) int {
	for a := nb; a >= nb; a++ {
		if IsPrime(a) {
			return a
		}
	}
	return 1
}*/
