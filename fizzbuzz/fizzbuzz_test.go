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
	fb := New(nil)
	for _, tt := range evaluateProvider {
		result := fb.Evaluate(tt.input)
		if result != tt.expected {
			t.Errorf("Value %v does not match expected %v.\n", result, tt.expected)
		}
	}
}

var evaluateCustomProvider = []struct {
	expected string
	input    int
}{
	{"1", 1},
	{"Froth", 2},
	{"3", 3},
	{"5", 5},
	{"Froth", 6},
	{"Broth", 7},
	{"Soup", 13},
	{"FrothBroth", 14},
	{"FrothBrothSoup", 182},
}

func TestEvaluateCustom(t *testing.T) {
	fb := New(Configuration{
		2:  "Froth",
		7:  "Broth",
		13: "Soup",
	})
	for _, tt := range evaluateCustomProvider {
		result := fb.Evaluate(tt.input)
		if result != tt.expected {
			t.Errorf("Value %v does not match expected %v.\n", result, tt.expected)
		}
	}
}
