# Heaps

A **heap** is a specialized tree-based data structure that satisfies the **heap property**:

- In a **max heap**, for any given node _C_, if _P_ is the parent of _C_, then the key (the value) of _P_ is greater than or equal to the key of _C_. This ensures that the largest key is always at the root.

- In a **min heap**, the key of _P_ is less than or equal to the key of _C_, ensuring that the smallest key is at the root.

Heaps are commonly used to implement priority queues, where the highest (or lowest) priority element is always efficiently accessible.

## Implementing Heaps in Go

Go provides the `container/heap` package, which offers heap operations for any type that implements the `heap.Interface`. This interface requires the following methods:

- `Len() int`: Returns the number of elements in the collection.

- `Less(i, j int) bool`: Reports whether the element with index `i` should sort before the element with index `j`.

- `Swap(i, j int)`: Swaps the elements with indexes `i` and `j`.

- `Push(x interface{})`: Adds an element `x` to the collection.

- `Pop() interface{}`: Removes and returns the last element from the collection.

By implementing these methods, any data type can utilize heap operations such as insertion and removal while maintaining the heap property.

### Example: Min-Heap Implementation

Below is an example of implementing a min-heap in Go using the `container/heap` package:

```go
package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of integers.
type IntHeap []int

// Len returns the number of elements in the collection.
func (h IntHeap) Len() int { return len(h) }

// Less reports whether the element with index i should sort before the element with index j.
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

// Swap swaps the elements with indexes i and j.
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push adds an element x to the heap.
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the last element from the heap.
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	// Initialize a slice with some integers.
	ints := &IntHeap{2, 1, 5}
	heap.Init(ints) // Establish the heap invariants.

	// Add a new integer to the heap.
	heap.Push(ints, 3)

	// Remove and print each element from the heap in ascending order.
	for ints.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(ints))
	}
}
```

**Output:**

```
1 2 3 5
```

**Explanation:**

1. **Define the `IntHeap` type**: This is a slice of integers that will implement the `heap.Interface`.

2. **Implement the `heap.Interface` methods**:

   - `Len()`: Returns the number of elements in the heap.

   - `Less(i, j int) bool`: Defines the ordering; for a min-heap, the element at index `i` should be less than the element at index `j`.

   - `Swap(i, j int)`: Swaps the elements at indexes `i` and `j`.

   - `Push(x interface{})`: Appends the new element `x` to the heap.

   - `Pop() interface{}`: Removes and returns the last element from the heap.

3. **Initialize and use the heap**:

   - Create an instance of `IntHeap` with some initial values.

   - Use `heap.Init` to establish the heap invariants.

   - Add a new element using `heap.Push`.

   - Remove elements using `heap.Pop`, which returns the smallest element due to the min-heap property.

This implementation leverages the `container/heap` package to manage the heap operations, ensuring that the smallest element is always at the root, and insertion and removal operations maintain the heap property.

For more information and examples, refer to the official Go documentation for the `container/heap` package. [container/heap](https://pkg.go.dev/container/heap)
