package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode uint
}

type Person struct {
	firstName string
	lastName  string
	age       uint8
	contact   ContactInfo
}

func main() {

	john := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       24,
		contact: ContactInfo{
			email:   "john@jd.example",
			zipCode: 411983,
		},
	}

	john.print()

	var alex Person

	alex.print()

	alex = john

	alex.print()

	alex.updateAge(28)
	alex.print()
}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *Person) updateAge(age uint8) {
	p.age = age
}
