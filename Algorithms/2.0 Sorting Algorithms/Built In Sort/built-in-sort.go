package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sorting integers
	ints := []int{5, 3, 4, 1, 2}
	sort.Ints(ints)
	fmt.Println("Sorted integers:", ints)

	// Sorting strings
	strs := []string{"banana", "apple", "cherry"}
	sort.Strings(strs)
	fmt.Println("Sorted strings:", strs)
}