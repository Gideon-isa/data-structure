package main

import (
	"fmt"
	"golangtutorial/structure"
)

func main() {
	myStack := structure.Stack{}
	myStack.Items = []int{2, 3, 4, 5, 6}

	myStack.Push(50)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)

	ages := make(map[int]string)

	ages[0] = "first"
	ages[1] = "second"
	ages[2] = "third"

	fmt.Println(ages)

	myQueue := structure.Queue{
		Items: []int{20, 30, 40},
	}

	fmt.Println(myQueue)

	myQueue.Enqueue(45, 50, 60)
	fmt.Println(myQueue)
	myQueue.Dequeue()

	fmt.Println(myQueue)

	myList := structure.LinkedList{}
	node1 := structure.Node{Data: 48}
	node2 := structure.Node{Data: 18}
	node3 := structure.Node{Data: 16}
	node4 := structure.Node{Data: 48}
	myList.Prepend(&node1)
	myList.Prepend(&node2)
	myList.Prepend(&node3)
	myList.Prepend(&node4)

	fmt.Println(myList)
	myList.PrintList()
}
