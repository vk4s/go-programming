package main

import (
	"fmt"
	"time"
)

func main() {
	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"epsilon",
		"zeta",
	}

	for i, v := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, v))
	}
	time.Sleep(1 * time.Second)
	printSomething("Second go routine")
}

func printSomething(s string) {
	fmt.Println(s)
}
