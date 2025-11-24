package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Add(value T) {
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{nil, value}
}

func (l *List[T]) String() string {
	return fmt.Sprintf("%v", l.val)
}

func main() {
	list := &List[int]{nil, 1}
	list.Add(2)
	list.Add(3)

	for n := list; n != nil; n = n.next {
		fmt.Println(n)
	}
}
