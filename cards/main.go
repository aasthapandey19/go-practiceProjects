package main

import "fmt"

func main() {
	card := []string{"Ace of Diamonds", newCard()}
	card = append(card, "Joker")
	for i, car := range card {
		fmt.Println(i, car)
	}
}

func newCard() string {
	return "five of Diamonds"
}
