# **4. Heap Sort**

_Description_: Heap Sort builds a max-heap from the input data and then repeatedly extracts the maximum element from the heap, reconstructing the heap each time, until all elements are sorted.

_Implementation_:

```go
package main

import "fmt"

// HeapSort sorts a slice of integers using the Heap Sort algorithm.
func HeapSort(arr []int) {
	n := len(arr)

	// Build a max-heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from the heap
	for i := n - 1; i >= 0; i-- {
		// Move current root to end
		arr[0], arr[i] = arr[i], arr[0]
		// Call max heapify on the reduced heap
		heapify(arr, i, 0)
	}
}

// heapify ensures the max-heap property for a subtree rooted at index i.
func heapify(arr []int, n, i int) {
	largest := i     // Initialize largest as root
	left := 2*i + 1  // Left child
	right := 2*i + 2 // Right child

	// If left child is larger than root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// If right child is larger than largest so far
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		// Recursively heapify the affected sub-tree
		heapify(arr, n, largest)
	}
}

func main() {
	data := []int{12, 11, 13, 5, 6, 7}
	HeapSort(data)
	fmt.Println("Sorted array:", data)
}
```

_Explanation_:

- The `HeapSort` function first transforms the slice into a max-heap.

## Time Complexity

- Best: O(n log n)​
- Average: O(n log n)​
- Worst: O(n log n)​

## Space Complexity

> O(1)​

## Use Case

Efficient and in-place; suitable for large datasets
