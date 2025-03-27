**1. Selection Sort**

_Description_: Selection Sort divides the array into a sorted and an unsorted region. It repeatedly selects the smallest element from the unsorted region and moves it to the end of the sorted region.

_Implementation_:

```go
package main

import "fmt"

// SelectionSort sorts a slice of integers using the Selection Sort algorithm.
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Swap the found minimum element with the first element
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	data := []int{64, 25, 12, 22, 11}
	SelectionSort(data)
	fmt.Println("Sorted array:", data)
}
```

_Explanation_:

- The `SelectionSort` function iterates over the slice, selecting the smallest element in the unsorted region and swapping it with the first unsorted element.
- This process continues until the entire slice is sorted.
