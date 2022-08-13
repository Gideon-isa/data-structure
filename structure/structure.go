package structure

import "fmt"

type Stack struct {
	Items []int
}

// Queue represents a queue that holds a slice
type Queue struct {
	Items []int
}

type Node struct {
	Data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

// Add ints to the Linked list
func (l *LinkedList) Prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

// Printing the List
func (l LinkedList) PrintList() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

// Push
func (s *Stack) Push(i ...int) {
	s.Items = append(s.Items, i...)
}

// Pop
func (s *Stack) Pop() int {
	l := len(s.Items) - 1
	r := s.Items[l]
	s.Items = s.Items[:l]
	return r
}

// Queue adds a value at the end
func (q *Queue) Enqueue(i ...int) {
	q.Items = append(q.Items, i...)
}

// Dequeue
func (q *Queue) Dequeue() int {
	toRemove := q.Items[0]
	q.Items = q.Items[1:]
	return toRemove
}
