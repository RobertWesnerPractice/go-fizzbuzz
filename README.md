# Go FizzBuzz

An unnecessarily customizable FizzBuzz implementation with tests.

## Installation

```bash
go get github.com/RobertWesnerPractice/go-fizzbuzz/fizzbuzz
```

## Running with default configuration

```go
fb := fizzbuzz.New(nil)
for i := 1; i <= 100; i++ {
    fmt.Println(fb.Evaluate(i))
}
```

## Running with custom configuration

```go
fb := fizzbuzz.New(fizzbuzz.Configuration{
    2:  "Froth",
    7:  "Broth",
    13: "Soup",
})
for i := 1; i <= 182; i++ {
    fmt.Println(fb.Evaluate(i))
}
```