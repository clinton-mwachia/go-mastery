# Stacks with Generics

- **Generics:** With the introduction of generics in Go 1.18, you can create a stack that holds any data type. Here's an example:

  ```go
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
  ```

In this example, the `Stack` struct is defined with a type parameter `T`, allowing it to store any data type. The `Push` and `Pop` methods operate on this generic type.

- **Concurrency:** If multiple goroutines will access the stack simultaneously, consider adding synchronization mechanisms like mutexes to prevent race conditions.

This implementation provides a basic stack with essential operations, suitable for various applications requiring LIFO behavior.
