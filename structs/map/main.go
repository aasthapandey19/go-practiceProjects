package main

import "fmt"

func main() {
	fmt.Println("Map implmntation")

	colors := map[string]string{
		"red":  "ferrari",
		"blue": "redbull",
	}
	fmt.Println(colors["red"])
}
