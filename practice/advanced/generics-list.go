package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	node := &List[int]{next: nil, val: 20}
	nodeHead := node
	node.push(30)
	nodeHead.print()
	node.push(40)
	nodeHead.print()
}

func (l *List[T]) print() {
	for l != nil {
		fmt.Printf("%v (%v)->", l.val, l)
		l = l.next
	}
	fmt.Println("nil")
}

func (l *List[T]) push(val T) {
	currentNode := l
	node := &List[T]{next: nil, val: val}
	currentNode.next = node
	l = currentNode
}
