package main

func main() {
	cards := newDeck()
	// cards.print()

	hand, _ := deal(cards, 5)
	hand.print()
	// fmt.Println()
	// remaining.print()

	// fmt.Println(deal(cards, 2))
	hand.saveToFile("data/hand.bin")
	newCards := newDeckFromFile("data/hand.bin")
	newCards.print()
}
