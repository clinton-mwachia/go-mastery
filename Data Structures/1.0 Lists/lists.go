package main

import "fmt"

// Node represents an element in the linked list
type Node struct {
	value int
	next  *Node
}

// LinkedList represents the linked list itself
type LinkedList struct {
	head *Node
}

// Append adds a new node with the given value at the end of the list
func (list *LinkedList) Append(value int) {
	newNode := &Node{value: value}
	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Prepend adds a new node with the given value at the beginning of the list
func (list *LinkedList) Prepend(value int) {
	newNode := &Node{value: value, next: list.head}
	list.head = newNode
}

// Delete removes the first occurrence of the node with the given value
func (list *LinkedList) Delete(value int) {
	if list.head == nil {
		return
	}
	if list.head.value == value {
		list.head = list.head.next
		return
	}
	current := list.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}
}

// Display prints out the values in the list
func (list *LinkedList) Display() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := &LinkedList{}
	list.Append(10)
	list.Append(20)
	list.Prepend(5)
	list.Display() // Output: 5 -> 10 -> 20 -> nil

	list.Delete(10)
	list.Display() // Output: 5 -> 20 -> nil
}
