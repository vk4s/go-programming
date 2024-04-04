package main

// import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "King of Spades")
	print(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
