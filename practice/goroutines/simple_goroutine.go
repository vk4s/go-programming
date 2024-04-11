package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	fmt.Printf("Sending message to %s...\n>>> ", name)
	fmt.Println("Hello", name)
	time.Sleep(2 * time.Second)
	fmt.Printf("Message sent.\n\n")
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

	for _, name := range names {
		sayHello(name)
	}

}
