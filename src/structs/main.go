package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 94000,
		},
	}
	alex.print()
	alexPointer := &alex
	alexPointer.changeName("Rubierto")
	alex.print()
}

func (p *person) changeName(name string) {
	(*p).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p) //%+v prints all parameters with their values
}
