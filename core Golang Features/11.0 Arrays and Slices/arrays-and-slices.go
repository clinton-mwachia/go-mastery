package main

import "fmt"

func main() {
	// Declaring an array of 5 integers
	var arr [5]int
	fmt.Println("Default array:", arr) // [0 0 0 0 0]

	// Initializing an array with specific values
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Initialized array:", arr2) // [1 2 3 4 5]

	// Array with inferred length
	arr3 := [...]int{10, 20, 30}
	fmt.Println("Array with inferred length:", arr3) // [10 20 30]

	fmt.Println("-------------------slices------------------------")

	// Creating a slice using a slice literal
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice) // [1 2 3 4 5]

	// Creating a slice from an array
	arr_slice := [5]int{10, 20, 30, 40, 50}
	slice2 := arr_slice[1:4]
	fmt.Println("Slice from array:", slice2) // [20 30 40]

	// Using the make function to create a slice
	slice3 := make([]int, 3, 5)
	fmt.Println("Slice with make:", slice3) // [0 0 0]

	// appending to slices
	slice = append(slice, 4, 5)
	fmt.Println("Appended slice:", slice) // [1 2 3 4 5]

	// copying slices
	src := []int{1, 2, 3}
	dest := make([]int, len(src))
	copy(dest, src)
	fmt.Println("Copied slice:", dest)
}
