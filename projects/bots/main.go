package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
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
