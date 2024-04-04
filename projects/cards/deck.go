package main

import (
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
	// write []byte -> file

	data := []byte(d.toString())

	err := os.WriteFile(file, data, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (d deck) toString() string {
	cards := strings.Join(d, ";")
	return cards
}
