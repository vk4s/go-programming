package main

// import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "King of Spades")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
