package main

import "fmt"

// SelectionSort sorts a slice of integers using the Selection Sort algorithm.
func SelectionSort(arr []int) {
	n := len(arr)
	for i := range n - 1 {
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
	data := []int{64, 25, 20, 1, 12, 22, 11}
	SelectionSort(data)
	fmt.Println("Sorted array:", data)
}
