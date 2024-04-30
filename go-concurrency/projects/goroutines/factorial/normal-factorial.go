package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// the number for which we need to calculate the factorial
	var number uint64 = 25

	fact := Factorial(number)

	fmt.Printf("Factorial of %d is %d\n", number, fact)

	elapsed := time.Since(start)
	fmt.Printf("Time taken to run the code: %s\n", elapsed)
}

func Factorial(number uint64) uint64 {
	if number <= 1 {
		return 1
	}
	return number * Factorial(number-1)
}
