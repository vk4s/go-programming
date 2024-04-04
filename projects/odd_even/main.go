package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		switch number % 2 {
		case 0:
			fmt.Println("Even")
		case 1:
			fmt.Println("Odd")
		default:
		}
	}
}
