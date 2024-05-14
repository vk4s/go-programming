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
}
