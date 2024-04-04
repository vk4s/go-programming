package main

import "testing"

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
