package main

import "fmt"

func main() {
	card := []string{"Ace of Diamonds", newCard()}
	card = append(card, "Joker")
	fmt.Println(card)
}

func newCard() string {
	return "five of Diamonds"
}
