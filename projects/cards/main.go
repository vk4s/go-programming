package main

import "fmt"

func main() {
	cards := newDeck()
	// cards.print()

	hand, remaining := deal(cards, 4)
	hand.print()
	fmt.Println()
	remaining.print()

	// fmt.Println(deal(cards, 2))
	hand.saveToFile("data/hand.bin")

}
