package main

//import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("myDeckOfCards")
	//fmt.Println(cards.toString())
}
