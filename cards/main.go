package main

//import "fmt"

func main() {
	cards := newDeck()
	dealtCards, remainingCards := deal(cards, 5)
	dealtCards.print()
	remainingCards.print()
}
