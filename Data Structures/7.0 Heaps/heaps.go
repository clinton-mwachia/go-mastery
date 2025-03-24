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
