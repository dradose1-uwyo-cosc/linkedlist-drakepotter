//[Drake Potter]
//COSC 3750
//[02/06/26]

package main

import (
	"fmt"
	"hw-linked/ds"
)

func main() {
	fmt.Println("Only here so the import doesn't leave an error")

	linkedlist := &ds.LinkedList{}
	linkedlist.InsertAt(0, "first")
	linkedlist.Insert("first")
	linkedlist.Insert("first")
	linkedlist.Insert("second")
	linkedlist.Insert("third")
	linkedlist.Insert("fourth")
	linkedlist.Insert("fifth")
	linkedlist.RemoveAt(4)
	linkedlist.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist.GetSize())
	fmt.Println("-------------")
	linkedlist.RemoveAll("first")
	linkedlist.PrintList()
	fmt.Println("-------------")
	linkedlist.Reverse()
	linkedlist.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist.GetSize())
	fmt.Println("-------------")

	stack := &ds.Stack{}
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")
	data, _ := stack.Pop()
	println("Popped from stack:", data)

	queue := &ds.Queue{}
	queue.Push("first")
	queue.Push("second")
	queue.Push("third")
	data, _ = queue.Pop()
	println("Popped from queue:", data)

	//LinkedList tests
	l := ds.LinkedList{}
	l.Insert("Donkey")
	l.PrintList()

	l.InsertAt(0, "Shrek")
	l.PrintList()

	l.Remove("Donkey")
	l.PrintList()

	if l.GetSize() == 1 {
		fmt.Println("Size is correct")
	}

	for i := 0; i < 4; i++ {
		l.Insert("Farquaad")
	}
	l.PrintList()

	l.Reverse()
	l.PrintList()

	l.RemoveAt(4)
	l.PrintList()

	l.RemoveAll("Farquaad")
	if l.IsEmpty() {
		fmt.Println("Empty")
	}

	//Stack tests
	var stack2 ds.Stack
	stack2.Push("Bob")
	stack2.Push("Pat")
	fmt.Println(stack2.Pop())

	//Queue tests
	var queue2 ds.Queue
	queue2.Push("Lester")
	queue2.Push("Woody")
	fmt.Println(queue2.Pop())
}
