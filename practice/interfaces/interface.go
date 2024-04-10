package main

import "fmt"

type shape interface {
	area() int
}

type square struct {
	side int
}

type rectangle struct {
	width  int
	length int
}

func (s square) area() int {
	return s.side * s.side
}

func (r rectangle) area() int {
	return r.length * r.width
}

func printArea(s shape) {
	fmt.Println(s.area())
}

func main() {
	mySquare := square{side: 2}
	printArea(mySquare)

	myRect := rectangle{length: 2, width: 5}
	printArea(myRect)
}
