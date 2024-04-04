package main

import "fmt"

// create a new type of "deck", which is slice of strings
type deck []string

// 'd' here is a 'this' of C++/JS and 'self' of Python
func (d deck) print() {
	for idx, element := range d {
		fmt.Println(idx, element)
	}
}
