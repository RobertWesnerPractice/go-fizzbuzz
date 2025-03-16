package fizzbuzz

import (
	"strconv"
)

func Evaluate(value int) string {
	var result = ""

	if value%3 == 0 {
		result += "Fizz"
	}

	if value%5 == 0 {
		result += "Buzz"
	}

	if len(result) == 0 {
		return strconv.Itoa(value)
	}

	return result
}
