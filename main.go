package main

import (
	"fmt"
	"github.com/RobertWesnerPractice/go-fizzbuzz/fizzbuzz"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Println(fizzbuzz.Evaluate(i))
	}
}
