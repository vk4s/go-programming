package main

func main() {
	cards := newDeck()
	cards.shuffle()

	// cards.print()

	hand, _ := deal(cards, 4)
	hand.print()
	hand.shuffle()

	hand.print()

	// remaining.print()

	// fmt.Println(deal(cards, 2))
	// err := hand.saveToFile("data/hand.bin")
	// if err != nil {
	// 	panic(err)
	// }

	// newCards, err := newDeckFromFile("data/hand.bin")
	// if err != nil {
	// 	panic(err)
	// }

	// newCards.print()
}
