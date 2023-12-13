package day05

import (
	"fmt"
)

type ConverterNode struct {
	Value *Converter
	Next  *ConverterNode
}

// LinkedList represents a linked list
type LinkedList struct {
	Head *ConverterNode
	Tail *ConverterNode
}

// Append adds a new node with the given data to the end of the linked list
func (list *LinkedList) Append(data *Converter) {
	newNode := &ConverterNode{Value: data, Next: nil}

	if list.Head == nil && list.Tail == nil {
		// If the list is empty, set the new node as the head
		list.Head = newNode
		list.Tail = newNode
		return
	}

	// Traverse the list to find the last node
	current := list.Tail
	for current.Next != nil {
		current = current.Next
	}

	// Append the new node to the end
	current.Next = newNode
}

// Display prints the elements of the linked list
func (list *LinkedList) Display() {
	current := list.Head
	for current != nil {
		//fmt.Printf("%d -> ", current.Con)
		current = current.Next
	}
	fmt.Println("nil")
}
