package main

func main() {
	cards := newDeck()
	// cards.print()

	hand, _ := deal(cards, 5)
	hand.print()
	// fmt.Println()
	// remaining.print()

	// fmt.Println(deal(cards, 2))
	err := hand.saveToFile("data/hand.bin")
	if err != nil {
		panic(err)
	}

	newCards, err := newDeckFromFile("data/hand.bin")
	if err != nil {
		panic(err)
	}

	newCards.print()
}
