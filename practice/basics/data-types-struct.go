package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Person1", age: 30}
	fmt.Printf("%+v\n", p)
	p.updateName("New Person")
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newName string) {
	p.name = newName
}

func (p person) updateName2(newName string) {
	p.name = newName
}
