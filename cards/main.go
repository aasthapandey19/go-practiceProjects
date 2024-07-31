package main

//import "fmt"

func main() {
	//cards := newDeck()
	//cards.saveToFile("myDeckOfCards")
	//fmt.Println(cards.toString())
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
