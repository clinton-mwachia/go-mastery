package main

import "fmt"

// BubbleSort sorts a slice of integers using the Bubble Sort algorithm.
func BubbleSort(arr []int) {
	n := len(arr)
	for i := range n - 1 {
		for j := range n - i - 1 {
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
