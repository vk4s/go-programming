package main

import "fmt"

func main() {
	welcome_message := "-- Welcome to the Game of Cards --"

	fmt.Printf("%v\n%v\n", welcome_message,
		separator_line("*", len(welcome_message)))

	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "King of Spades")

	fmt.Println("The participants are:")

	for idx, element := range cards {
		fmt.Printf("\t%v -> %v\n", idx, element)
	}
}

func newCard() string {
	return "Five of Diamonds"
}

func separator_line(character string, times int) string {
	// repeat a character multiple time
	str := ""
	for ; times >= 0; times-- {
		str += character
	}
	return str
}
