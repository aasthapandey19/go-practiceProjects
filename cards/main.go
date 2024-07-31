package main

import "fmt"

//import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
}
