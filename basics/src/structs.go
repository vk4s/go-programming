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

	fmt.Println(john)

	var alex Person

	fmt.Printf("%+v\n", alex)

	alex = john

	fmt.Printf("%+v\n", alex)

}
