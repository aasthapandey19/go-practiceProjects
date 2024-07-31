package main

import (
	"fmt"
	"os"
	"strings"
)

// new type to extend string data typee
type deck []string

func newDeck() deck {
	freshCards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			createdCard := value + " of " + suite
			freshCards = append(freshCards, createdCard)
		}
	}
	return freshCards
}

func (d deck) print() {
	for i, x := range d {
		fmt.Println(i, x)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, "\n")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}
