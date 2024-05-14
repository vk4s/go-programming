package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomNumber(max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)
	return randomGenerator.Intn(max - 1)
}

func main() {
	var numbers []int

	for range 100 {
		numbers = append(numbers, getRandomNumber(100))
	}

	fmt.Println(numbers)
	fmt.Println(len(numbers))

	for _, n := range numbers {
		// if n in divisible by 2, print "Two"
		// if n in divisible by 3, print "Three"
		// if the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.
		// Otherwise, print “Never mind”
		fmt.Print(n, " ")
		if n%2 == 0 {
			fmt.Println("Two")
		} else if n%3 == 0 {
			fmt.Println("Three")
		} else if n%2 == 0 && n%3 == 0 {
			fmt.Println("Six!")
		} else {
			fmt.Println("Never mind")
		}

	}

}
