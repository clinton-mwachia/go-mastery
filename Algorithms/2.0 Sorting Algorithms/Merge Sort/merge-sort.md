# **3. Merge Sort**

_Description_: Merge Sort is a divide-and-conquer algorithm that splits the array into halves, recursively sorts them, and then merges the sorted halves back together.

_Implementation_:

```go
package main

import "fmt"

// MergeSort sorts a slice of integers using the Merge Sort algorithm.
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Split the array into two halves
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Merge the sorted halves
	return merge(left, right)
}

// merge combines two sorted slices into a single sorted slice.
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	// Compare elements and merge
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	data := []int{38, 27, 43, 3, 9, 82, 10}
	sortedData := MergeSort(data)
	fmt.Println("Sorted array:", sortedData)
}
```

_Explanation_:

- The `MergeSort` function recursively divides the slice into halves until each sub-slice contains a single element.
- The `merge` function then combines these sorted sub-slices back together in order, resulting in a fully sorted slice.

## Time Complexity

- Best: O(n log n)​
- Average: O(n log n)​
- Worst: O(n log n)​

## Space Complexity

> O(n)​

## Use Case

Consistent performance; suitable for large datasets; requires additional memory for merging.
