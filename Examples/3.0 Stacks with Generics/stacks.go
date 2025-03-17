package main

import (
	"errors"
	"fmt"
)

// Stack represents a generic stack data structure
type Stack[T any] struct {
	items []T
}

// Push adds an item to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	topIndex := len(s.items) - 1
	item := s.items[topIndex]
	s.items = s.items[:topIndex]
	return item, nil
}

func main() {
	intStack := &Stack[int]{}
	intStack.Push(100)
	intStack.Push(200)
	fmt.Println(intStack.Pop()) // Output: 200, <nil>

	stringStack := &Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")
	fmt.Println(stringStack.Pop()) // Output: world, <nil>
}
