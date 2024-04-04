package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "King of Spades")

	for idx, element := range cards {
		fmt.Println(idx, element)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
