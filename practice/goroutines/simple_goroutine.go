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

func sendMessage(name string, channel chan string) {
	fmt.Printf("Sending message to %s...\n>>> ", name)
	fmt.Println("Hello", name)

	elapsedTime := time.Duration(getRandomNumber(5) * int(time.Second))
	time.Sleep(elapsedTime)

	channel <- fmt.Sprintf("Message sent to %s. Elapsed time: %dms", name, elapsedTime.Milliseconds())
}

func main() {
	names := []string{
		"Vik",
		"John",
		"Jane",
		"Krishna",
		"Priya",
		"Dave",
		"Karan",
		"Mohit",
		"Rahul",
	}

	// Create a channel to control go routines' execution
	channel := make(chan string)

	for _, name := range names {
		go sendMessage(name, channel)
	}

	// This waits for the go routines to finish
	// and prints the messages one-by-one as they finish
	for range len(names) {
		fmt.Println("Status:", <-channel)
	}
}
