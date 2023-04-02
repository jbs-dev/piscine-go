package piscine

import (
	"github.com/01-edu/z01"
)

// PrintNbrBase prints an integer in a given base.
func PrintNbrBase(nbr int, base string) {
	answer := ""           // the final answer string
	numDigits := len(base) // number of digits in the given base
	maxPower := numDigits  // maximum power of the base
	if nbr < 0 {           // if the number is negative, add a negative sign to the answer and make numDigits negative
		answer = "-"
		maxPower *= -1
	}
	if numDigits > 1 { // if the base has more than one digit
		for nbr/maxPower >= numDigits { // while nbr divided by maxPower is greater than or equal to the base
			maxPower *= numDigits // multiply maxPower by the base
		}
		for maxPower != 0 { // while maxPower is not zero
			digitIndex := nbr / maxPower               // calculate the index of the current digit in the base
			answer = answer + string(base[digitIndex]) // add the character at the digitIndex to the answer
			nbr = nbr - (nbr/maxPower)*maxPower        // subtract the product of nbr divided by maxPower and maxPower from nbr
			maxPower /= numDigits                      // divide maxPower by the base
		}
		seenDigits := map[rune]bool{} // keep track of digits seen in the base
		for _, digit := range base {  // loop through each digit in the base
			if digit == '+' || digit == '-' { // if the digit is a plus or minus sign, set the answer to "NV"
				answer = "NV"
				break
			}
			if !seenDigits[digit] { // if the digit has not been seen yet
				seenDigits[digit] = true // mark it as seen
			} else { // if the digit has already been seen
				answer = "NV" // set the answer to "NV"
				break
			}
		}
	} else { // if the base has only one digit
		answer = "NV" // set the answer to "NV"
	}
	for _, c := range answer { // loop through each character in the answer string
		z01.PrintRune(c) // print the character
	}
}
