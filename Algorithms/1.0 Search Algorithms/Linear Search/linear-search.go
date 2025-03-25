package main

import "fmt"

// LinearSearch function searches for a target value in a slice.
// It returns the index of the target and a boolean indicating if the target was found.
func LinearSearch(slice []int, target int) (int, bool) {
	for i, v := range slice {
		if v == target {
			return i, true
		}
	}
	return -1, false
}

func main() {
	data := []int{10, 20, 30, 40, 50}
	target := 30

	index, found := LinearSearch(data, target)
	if found {
		fmt.Printf("Found %d at index %d\n", target, index)
	} else {
		fmt.Printf("%d not found in the slice\n", target)
	}
}
