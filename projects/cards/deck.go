package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "King", "Queen", "Jack", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func newDeckFromFile(file string) deck {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	decodedText := decodeB64(string(data))

	cards := deck{}
	cards.fromString(decodedText)
	return cards
}

// 'd' here is a 'this' of C++/JS and 'self' of Python
func (d deck) print() {
	for idx, element := range d {
		fmt.Println(idx, element)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	// create a hand of 'n' cards

	hand := d[:handSize]
	remaining := d[handSize:]

	return hand, remaining
}

func (d deck) saveToFile(file string) {
	// convert []string -> 'string' -> []byte
	// write []byte -> base64 encode -> []byte -> file

	cardsString := d.toString()
	encodedText := encodeB64(cardsString)

	err := os.WriteFile(file, []byte(encodedText), 0666)
	if err != nil {
		panic(err)
	}
}

func encodeB64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func decodeB64(b64EncodedText string) string {

	rawDecodedText, err := base64.StdEncoding.DecodeString(b64EncodedText)
	if err != nil {
		panic(err)
	}

	return string(rawDecodedText)
}

func (d deck) toString() string {
	cards := strings.Join(d, ";")
	return cards
}

func (d *deck) fromString(cardsText string) {
	cards := strings.Split(cardsText, ";")
	cardsDeck := deck{}
	for _, card := range cards {
		cardsDeck = append(cardsDeck, card)
	}
	*d = cardsDeck
}
