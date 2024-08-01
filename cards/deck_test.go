package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	testDeck := newDeck()

	if len(testDeck) != 52 {
		t.Errorf("Expected deck length of 52, recieved %d ", len(testDeck))
	}

	if testDeck[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, got %s", testDeck[0])
	}
	if testDeck[len(testDeck)-1] != "King of Hearts" {
		t.Errorf("Expected King of Hearts, got %s", testDeck[len(testDeck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")
	if len(loadedDeck) != 1 {
		t.Errorf("Expected all 52 cards, received %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
