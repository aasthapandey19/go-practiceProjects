package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	alex := person{"Alex", "Anderson", contactInfo{"em", 99}}
	//checo := person{firstName: "Sergio", lastName: "Perez", contact: contactInfo{"", 0}}
	//fmt.Println(alex)
	//fmt.Println(checo.lastName)
	//var max person
	//max.firstName = "Max"
	//max.lastName = "Verstappen"
	//max.contact.email = "jdasdj"
	//max.contact.zip = 98
	//fmt.Println(max)
	alex.updateName("Charles")

	alex.print()
}

func (p *person) updateName(updationOnName string) {
	p.firstName = updationOnName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
