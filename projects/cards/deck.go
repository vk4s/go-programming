package main

import "fmt"

// create a new type of "deck", which is slice of strings
type deck []string

func print(cards deck) {
	for idx, element := range cards {
		fmt.Println(idx, element)
	}
}
