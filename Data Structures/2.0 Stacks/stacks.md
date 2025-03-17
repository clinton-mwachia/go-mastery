# Stacks

A **stack** is a fundamental data structure that operates on a Last-In-First-Out (LIFO) principle, meaning the last element added is the first to be removed. In Go, stacks can be implemented using slices, which provide dynamic resizing and efficient element access.

**Implementing a Stack Using Slices**

Below is an example of a stack implementation using a slice of integers:

```go
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
```

**Explanation:**

- **Stack Structure:** The `Stack` struct contains a slice of integers, `items`, which holds the elements of the stack.

- **Push Method:** Adds a new item to the top of the stack by appending it to the `items` slice.

- **Pop Method:** Removes and returns the top item from the stack. It checks if the stack is empty to prevent errors and then slices off the top element.

- **Peek Method:** Returns the top item without removing it, allowing inspection of the stack's top element.

- **IsEmpty Method:** Checks if the stack is empty by verifying if the length of `items` is zero.

- **Size Method:** Returns the current number of items in the stack.

**Using the Stack:**

In the `main` function, we create a new `Stack` instance and perform various operations:

- **Push:** Add items `10`, `20`, and `30` to the stack.

- **Size:** Retrieve the number of items in the stack, which is `3` after the pushes.

- **Peek:** Check the top item without removing it; the top item is `30`.

- **Pop:** Remove the top item (`30`) from the stack.

- **Size after Pop:** Verify the stack size after popping, which is now `2`.
