package piscine

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

/* func EvenOrOdd(number int) string {
  return map[bool]string{true: "Even", false: "Odd"}[number%2 == 0]
} */
