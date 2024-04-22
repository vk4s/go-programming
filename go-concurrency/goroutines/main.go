package main

import (
	"fmt"
)

func main() {
	go printSomething("First go routine")
	printSomething("Second go routine")
}

func printSomething(s string) {
	fmt.Println(s)
}
