package fizzbuzz

import (
	"sort"
	"strconv"
)

type Configuration map[int]string

type FizzBuzz struct {
	config Configuration
}

func (f FizzBuzz) Evaluate(value int) string {
	var result = ""

	// make sure all keys are in ascending order, not random!
	configKeys := make([]int, 0, len(f.config))
	for key := range f.config {
		configKeys = append(configKeys, key)
	}
	sort.Ints(configKeys)

	for _, key := range configKeys {
		if value%key == 0 {
			result += f.config[key]
		}
	}

	if len(result) == 0 {
		return strconv.Itoa(value)
	}

	return result
}

func New(config Configuration) FizzBuzz {
	if config == nil {
		config = Configuration{
			3: "Fizz",
			5: "Buzz",
		}
	}

	return FizzBuzz{config}
}
