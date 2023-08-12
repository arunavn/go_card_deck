package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected 16 cards but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card as Ace of Spades but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Club" {
		t.Errorf("Expected last card as Ace of Spades but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	deck_loaded := newDeckFromFile("_decktesting")
	if len(deck_loaded) != len(deck) {
		t.Errorf("Wrote %v cards, but loaded %v cards", len(deck), len(deck_loaded))
	}
}
