package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type rectangle struct {
	length, width float64
}

type triangle struct {
	base, height float64
}

type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return math.Pow(s.side, 2)
}

func (r rectangle) getArea() float64 {
	return r.length * r.width
}
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	sq := square{side: 2}
	rect := rectangle{length: 2, width: 3}
	tr := triangle{base: 5, height: 3}

	printArea(sq)
	printArea(rect)
	printArea(tr)

}
