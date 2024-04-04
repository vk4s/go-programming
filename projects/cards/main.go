package main

import "fmt"

func main() {
	cards := newDeck()
	// cards.print()

	hand, remaining := deal(cards, 3)
	hand.print()
	fmt.Println()
	remaining.print()

	// fmt.Println(deal(cards, 2))

}
