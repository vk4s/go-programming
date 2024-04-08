package main

import "fmt"

func main() {
	number := 12

	fmt.Println(number)
	addOne(number)
	fmt.Println(number)

}

func addOne(number int) {
	number = number + 1
}
