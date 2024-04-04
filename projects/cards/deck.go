package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
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

func newDeckFromFile(file string) (deck, error) {
	cards := deck{}
	data, err := os.ReadFile(file)
	decodedText := decodeB64(string(data))
	cards.fromString(decodedText)
	return cards, err
}

func (d deck) print() {
	for idx, element := range d {
		fmt.Println(idx, element)
	}
	fmt.Println()
}

func deal(d deck, handSize int) (deck, deck) {
	// create a hand of 'n' cards
	hand := d[:handSize]
	remaining := d[handSize:]
	return hand, remaining
}

func (d deck) shuffle() {
	for i, _ := range d {
		n := rand.Intn(len(d) - 1)
		d[i], d[n] = d[n], d[i]
	}
}

func (d deck) saveToFile(file string) error {
	// convert []string -> 'string' -> []byte
	// write []byte -> base64 encode -> []byte -> file
	encodedText := encodeB64(d.toString())
	return os.WriteFile(file, []byte(encodedText), 0666)
}

func (d deck) toString() string {
	return strings.Join([]string(d), ";")
}

func (d *deck) fromString(cardsText string) {
	cards := strings.Split(cardsText, ";")
	cardsDeck := deck{}
	for _, card := range cards {
		cardsDeck = append(cardsDeck, card)
	}
	*d = cardsDeck
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
