package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	alex := person{"Alex", "Anderson"}
	checo := person{firstName: "Sergio", lastName: "Perez"}
	fmt.Println(alex)
	fmt.Println(checo.lastName)
}
