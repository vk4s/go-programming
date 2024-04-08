package main

import "fmt"

type bot interface {
	// any function with same signature
	// (name, parameters and return type; receiver type is not considered)
	// automatically becomes an honorary member of this `bot` interface
	// e.g. getGreeting() function in the englishBot and spanishBot types
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// [OPTIONAL] if we are not using the receiver value we can remote that. i.e. `eb`
func (englishBot) getGreeting() string {
	// VERY custom logic to generating English greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic to generating Spanish greeting
	return "Hola!"
}
