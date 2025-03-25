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
