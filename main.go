package main

import (
	"fmt"
	"github.com/RobertWesnerPractice/go-fizzbuzz/fizzbuzz"
)

func main() {
	fb := fizzbuzz.New(nil)
	for i := 1; i <= 100; i++ {
		fmt.Println(fb.Evaluate(i))
	}
}
