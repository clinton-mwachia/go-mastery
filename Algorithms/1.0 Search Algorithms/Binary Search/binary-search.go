package main

import "fmt"

// BinarySearch function searches for a target value in a sorted slice.
// It returns the index of the target and a boolean indicating if the target was found.
func BinarySearch(slice []int, target int) (int, bool) {
	low, high := 0, len(slice)-1

	for low <= high {
		mid := low + (high-low)/2
		if slice[mid] == target {
			return mid, true
		}
		if slice[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, false
}

func main() {
	data := []int{10, 20, 30, 40, 50} // Must be sorted
	target := 30

	index, found := BinarySearch(data, target)
	if found {
		fmt.Printf("Found %d at index %d\n", target, index)
	} else {
		fmt.Printf("%d not found in the slice\n", target)
	}
}
