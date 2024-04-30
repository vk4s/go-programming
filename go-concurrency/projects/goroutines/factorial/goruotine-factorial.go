package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	// the number for which we need to calculate the factorial
	var number uint64 = 25

	var wg sync.WaitGroup

	wg.Add(1)

	fact := FactorialGo(number, &wg)

	wg.Wait()

	fmt.Printf("Factorial of %d is %d\n", number, fact)

	elapsed := time.Since(start)
	fmt.Printf("Time taken to run the code: %s\n", elapsed)
}

func FactorialGo(number uint64, wg *sync.WaitGroup) uint64 {
	// fmt.Println("Checking factorial for: ", number)
	if number <= 1 {
		wg.Done()
		return 1
	}
	wg.Add(1)
	fact := number * FactorialGo(number-1, wg)
	wg.Done()
	return fact
}
