package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Push(val T) {
	if l.next == nil {
		l.val = val
		l.next = &List[T]{}
	} else {
		l = l.next
		l.Push(val)
	}
}

func (l *List[T]) Println() {
	fmt.Print("[ ")
	for {
		val := l.val
		if l.next == nil {
			break
		}
		fmt.Printf("%v ", val)
		l = l.next
	}
	fmt.Println("]")
}

func ten() {
	l := List[float64]{}
	l.Println()
	l.Push(1.0)
	l.Println()
	l.Push(2.0)
	l.Println()

	l2 := List[string]{}
	l2.Println()
	l2.Push("hello")
	l2.Println()
	l2.Push("world")
	l2.Println()
}
