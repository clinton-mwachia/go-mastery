## Sorting Algorithms

In Go, implementing sorting algorithms can be achieved through both custom implementations and utilizing the standard library's `sort` package. Below are examples of common sorting algorithms implemented in Go, along with explanations of their functionality.

### 1. Bubble Sort

Bubble Sort is a simple comparison-based algorithm that repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order. This process continues until the list is sorted.

**Implementation:**

```go
package main

import "fmt"

// BubbleSort sorts a slice of integers using the Bubble Sort algorithm.
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap elements
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	data := []int{64, 34, 25, 12, 22, 11, 90}
	BubbleSort(data)
	fmt.Println("Sorted array:", data)
}
```

**Explanation:**

- The `BubbleSort` function iterates over the slice multiple times.
- During each pass, it compares adjacent elements and swaps them if they are in the wrong order.
- The largest unsorted element "bubbles" to its correct position at the end of the slice in each pass.
- This process repeats until no swaps are needed, indicating that the slice is sorted.

**Usage:**

In the `main` function, we define a slice `data` with unsorted integers. We call `BubbleSort` to sort the slice and then print the sorted array.

## Time Complexity

- Best: O(n)​
- Average: O(n²)​
- Worst: O(n²)​
- Space Complexity: O(1)​

## Use Case

Educational purposes; not suitable for large datasets due to inefficiency.
