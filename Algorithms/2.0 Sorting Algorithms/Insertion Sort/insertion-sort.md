## **2. Insertion Sort**

_Description_: Insertion Sort builds the final sorted array one item at a time, with each new item being inserted into its correct position among the previously sorted items.

_Implementation_:

```go
package main

import "fmt"

// InsertionSort sorts a slice of integers using the Insertion Sort algorithm.
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		// Move elements of arr[0..i-1], that are greater than key,
		// to one position ahead of their current position
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func main() {
	data := []int{12, 11, 13, 5, 6}
	InsertionSort(data)
	fmt.Println("Sorted array:", data)
}
```

_Explanation_:

- The `InsertionSort` function iterates over the slice, inserting each element into its correct position relative to the already sorted portion of the slice.
- This process continues until the entire slice is sorted.

## Time Complexity

- Best: O(n)​
- Average: O(n²)​
- Worst: O(n²)​

## Space Complexity

> O(1)​

## Use Case

Efficient for small datasets or nearly sorted data.
