package main

import (
	"fmt"
	"time"
)

func sendMessage(name string, channel chan string) {
	fmt.Printf("Sending message to %s...\n>>> ", name)
	fmt.Println("Hello", name)
	time.Sleep(1 * time.Second)
	channel <- fmt.Sprintf("Message sent to %s.\n", name)
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
		fmt.Println("Status: ", <-channel)
	}

}
