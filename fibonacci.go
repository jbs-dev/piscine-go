package piscine

func Fibonacci(index int) int {
	if index < 0 { // If the index is negative, return -1
		return -1
	}
	if index == 0 { // If the index is 0, return 0
		return 0
	}
	if index == 1 { // If the index is 1, return 1
		return 1
	} else { // For all other cases, calculate the Fibonacci number recursively
		return Fibonacci(index-1) + Fibonacci(index-2)
	}
}
