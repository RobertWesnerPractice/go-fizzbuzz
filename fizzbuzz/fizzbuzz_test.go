package fizzbuzz

import (
	"testing"
)

var evaluateProvider = []struct {
	expected string
	input    int
}{
	{"1", 1},
	{"2", 2},
	{"Fizz", 3},
	{"Buzz", 5},
	{"Fizz", 6},
	{"Buzz", 10},
	{"FizzBuzz", 15},
	{"Buzz", 25},
	{"26", 26},
	{"FizzBuzz", 75},
	{"Fizz", 333},
}

func TestEvaluate(t *testing.T) {
	for _, tt := range evaluateProvider {
		result := Evaluate(tt.input)
		if result != tt.expected {
			t.Errorf("Value %v does not match expected %v.\n", result, tt.expected)
		}
	}
}
