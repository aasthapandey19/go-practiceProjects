package main

import "fmt"

type englishBot struct{}

type spanishBot struct{}

func main() {
	var ebs englishBot
	fmt.Println(ebs.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "hi there"
}
func (sb spanishBot) getGreeting() string {
	return "hola!"
}
