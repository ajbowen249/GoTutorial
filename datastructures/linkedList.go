// Package dataStructures implements an integer binary tree
// and an integer linked list.
package dataStructures

import (
	"fmt"
)

// IntLinkedList is an integer linked list.
type IntLinkedList struct {
	Head *IntLLNode // Head is the head node of the list.
	Tail *IntLLNode // Tail is the head node of the list.
}

// IntLLNode is a node in an integer linked list.
type IntLLNode struct {
	Value    int
	Previous *IntLLNode
	Next     *IntLLNode
}

// Add creates a new node with the given value and inserts it into the linked list.
func (ll *IntLinkedList) Add(newValue int) {
	newNode := new(IntLLNode)
	newNode.Value = newValue

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
		return
	}

	ll.Tail.Next = newNode
	newNode.Previous = ll.Tail
	ll.Tail = newNode
}

// Get Seeks through the linked list and returns the node at the given index.
func (ll *IntLinkedList) Get(index int) int {
	current := ll.Head

	for index > 0 {
		if current == nil {
			return 0
		}

		current = current.Next
		index--
	}

	return current.Value
}

// Print prints the list to standard output.
func (ll *IntLinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}
