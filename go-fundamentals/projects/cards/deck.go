package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Deck []string

func newDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "King", "Queen", "Jack", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func newDeckFromFile(file string) (Deck, error) {
	cards := Deck{}
	data, err := os.ReadFile(file)
	decodedText := decodeB64(string(data))
	cards.fromString(decodedText)
	return cards, err
}

func (d Deck) print() {
	for idx, element := range d {
		fmt.Println(idx, element)
	}
	fmt.Println()
}

func deal(d Deck, handSize int) (Deck, Deck) {
	hand := d[:handSize]
	remaining := d[handSize:]
	return hand, remaining
}

func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)

	for i := range d {
		newPosition := randomGenerator.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d Deck) saveToFile(file string) error {
	// convert []string -> 'string' -> []byte
	// write []byte -> base64 encode -> []byte -> file
	encodedText := encodeB64(d.toString())
	return os.WriteFile(file, []byte(encodedText), 0666)
}

func (d Deck) toString() string {
	return strings.Join([]string(d), ";")
}

func (d *Deck) fromString(cardsText string) {
	cards := strings.Split(cardsText, ";")
	cardsDeck := Deck{}
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
