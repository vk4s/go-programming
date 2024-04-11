package main

import "fmt"

func main() {
	var name string = "Vikash"
	name2 := "Poppins"

	// Can't declare multiple variables with DIFFERENT types
	// var age int, address string, is_citizen bool
	// fmt.Println(age, address, is_citizen)

	// Can declare multiple variables with SAME type (default value is 'falsy')
	var python, java, c, golang bool
	var age int
	fmt.Println(python, java, c, golang, age)

	fmt.Printf("Hello %v and %v\n", name, name2)
	fmt.Println("Hello world!!")
}
