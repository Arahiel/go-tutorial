package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expectedDeckSize := 16
	if len(d) != expectedDeckSize {
		t.Errorf("Expected deck length of %v, but got %v", expectedDeckSize, len(d))
	}

	expectedFirstItem := "Ace of Spades"
	expectedLastItem := "Four of Clubs"

	if d[0] != expectedFirstItem {
		t.Errorf("Expected: %v, got: %v", expectedFirstItem, d[0])
	}

	if d[len(d)-1] != expectedLastItem {
		t.Errorf("Expected: %v, got: %v", expectedLastItem, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	deckFileName := "_decktesting"
	os.Remove(deckFileName)

	d := newDeck()

	err := d.saveToFile(deckFileName)

	if err != nil {
		t.Errorf("Cannot save to file! %v", err)
	}

	loadedDeck := newDeckFromFile(deckFileName)

	if loadedDeck.toString() != d.toString() {
		t.Errorf("Expected: %v, got: %v", d.toString(), loadedDeck.toString())
	}

	os.Remove(deckFileName)
}
