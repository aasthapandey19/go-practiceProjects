package main

import "testing"

func TestNewDeck(t *testing.T) {

	testDeck := newDeck()

	if len(testDeck) != 52 {
		t.Errorf("Expected deck length of 52, recieved %d ", len(testDeck))
	}
}
