package main

import "fmt"

func main() {
	card := []string{"Ace of Diamonds", newCard()}
	card = append(card, "Joker")
	for i, card := range card {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "five of Diamonds"
}
