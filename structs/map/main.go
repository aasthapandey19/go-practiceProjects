package main

import "fmt"

func main() {
	fmt.Println("Map implmntation")

	colors := map[string]string{
		"red":   "ferrari",
		"blue":  "redbull",
		"green": "alpine",
	}
	fmt.Println(colors["red"])
	delete(colors, "blue")
	printMap(colors)
}

func printMap(c map[string]string) {

	for color, team := range c {
		fmt.Println(color, " is the color of team ", team)
	}
}
