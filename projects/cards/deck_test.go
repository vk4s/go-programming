package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	deckSize := 52
	if len(d) != deckSize {
		t.Errorf("Expected deck length of '%v', but got '%v'", deckSize, len(d))
	}

	firstCard := "Ace of Spades"
	if d[0] != firstCard {
		t.Errorf("Expected first card of '%v', but got '%v'", firstCard, d[0])
	}

	lastCard := "Ten of Clubs"
	if d[len(d)-1] != lastCard {
		t.Errorf("Expected last card of '%v', but got '%v'", lastCard, d[0])
	}
}

func TestSaveDeckToFileAndLoadDeckFromFile(t *testing.T) {
	filename := "_decktesting.bin"

	// delete any existing _dectesting.bin file from current dir
	err := os.Remove(filename)
	if err == nil {
		// file doesn't exist
		fmt.Printf("File %v deleted..\n", filename)
	}

	// create a deck
	cards := newDeck()

	// save the deck to a file named _decktesting.bin
	err = cards.saveToFile(filename)
	if err != nil {
		t.Error(err)
	}

	// load from file
	newCards, err := newDeckFromFile(filename)
	if err != nil {
		t.Error(err)
	}

	// assert the length
	if len(cards) != len(newCards) {
		t.Errorf("Length of both decks not equal. %v != %v", len(cards), len(newCards))
	}

	// delete the _dectesting.bin file from current dir
	err = os.Remove(filename)
	if err != nil {
		// file doesn't exist
		fmt.Printf("Error deleting %v file.\n", filename)
	}
}

func TestShuffle(t *testing.T) {
	cards := newDeck()

	originalFirstCard, originalLastCard := cards[0], cards[len(cards)-1]

	cards.shuffle()

	shuffledFirstCard, shuffledLastCard := cards[0], cards[len(cards)-1]

	if originalFirstCard == shuffledFirstCard {
		t.Errorf("First cards are same after shuffle, expected different. %v==%v", originalFirstCard, shuffledFirstCard)
	}

	if originalLastCard == shuffledLastCard {
		t.Errorf("Last cards are same after shuffle, expected different. %v==%v", originalLastCard, shuffledLastCard)
	}
}

func TestDeal(t *testing.T) {
	cards := newDeck()

	handSize := 5

	hand, remaining := deal(cards, handSize)

	if len(hand) != handSize {
		t.Errorf("Wrong number of cards in hand after deal. %v != %v",
			len(hand), handSize)
	}

	if len(remaining) != len(cards)-len(hand) {
		t.Errorf("Cards remaining are not equal to 'total cards' - 'cards in hand'. %v != %v",
			len(remaining), len(cards)-len(hand))
	}

	if len(remaining)+len(hand) != len(cards) {
		t.Errorf("Remaining number of cards not equal to 'cards in hand'+'cards remaining'. %v != %v",
			len(remaining)+len(hand), len(cards))
	}
}

func TestEncodingUtilityFunctions(t *testing.T) {
	originalText := "Game of Cards"

	encodedText := encodeB64(originalText)
	decodedText := decodeB64(encodedText)

	if originalText != decodedText {
		t.Errorf("Decoded text is not equal to orignal text. '%v'!='%v'",
			originalText, decodedText)
	}
}
