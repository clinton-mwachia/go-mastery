package main

import (
	"errors"
	"fmt"
)

// Stack represents a stack data structure
type Stack struct {
	items []int
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	topIndex := len(s.items) - 1
	item := s.items[topIndex]
	s.items = s.items[:topIndex]
	return item, nil
}

// Peek returns the top item without removing it
func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	stack := &Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Stack size:", stack.Size()) // Output: Stack size: 3

	topItem, _ := stack.Peek()
	fmt.Println("Top item:", topItem) // Output: Top item: 30

	poppedItem, _ := stack.Pop()
	fmt.Println("Popped item:", poppedItem) // Output: Popped item: 30

	fmt.Println("Stack size after pop:", stack.Size()) // Output: Stack size after pop: 2
}
