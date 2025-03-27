## Sorting Algorithms

In Go, implementing sorting algorithms can be achieved through both custom implementations and utilizing the standard library's `sort` package. Below are examples of common sorting algorithms implemented in Go, along with explanations of their functionality.

### 1. Quick Sort

Quick Sort is a highly efficient sorting algorithm that uses a divide-and-conquer approach. It selects a 'pivot' element and partitions the other elements into two sub-arrays according to whether they are less than or greater than the pivot. The sub-arrays are then sorted recursively.

**Implementation:**

```go
package main

import "fmt"

// QuickSort sorts a slice of integers using the Quick Sort algorithm.
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	// Choose a pivot
	pivot := len(arr) / 2

	// Move the pivot to the right
	arr[pivot], arr[right] = arr[right], arr[pivot]

	// Partitioning
	for i := range arr {
		if arr[i] < arr[right] {
			// Swap elements
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Swap the pivot back to its correct place
	arr[left], arr[right] = arr[right], arr[left]

	// Recursively apply QuickSort to the sub-arrays
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])

	return arr
}

func main() {
	data := []int{10, 7, 8, 9, 1, 5}
	sortedData := QuickSort(data)
	fmt.Println("Sorted array:", sortedData)
}
```

**Explanation:**

- The `QuickSort` function checks if the length of the slice is less than 2; if so, it returns the slice as it's already sorted.
- It selects a pivot element and partitions the slice into elements less than and greater than the pivot.
- The pivot is moved to its correct position, and the function recursively sorts the sub-arrays on either side of the pivot.

**Usage:**

In the `main` function, we define a slice `data` with unsorted integers. We call `QuickSort` to sort the slice and then print the sorted array.

## Time Complexity

- Best: O(n log n)​
- Average: O(n log n)​
- Worst: O(n²)​

## Space Complexity

> O(log n)​

## Use Case

Generally fast; preferred for large datasets; performance can degrade with poorly chosen pivots.
