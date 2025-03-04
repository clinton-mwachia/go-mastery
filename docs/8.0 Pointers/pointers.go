package main

import "fmt"

// Function that modifies a variable using a pointer
func modifyValue(val *int) {
	*val = *val * 2
}

// Function demonstrating pointer swapping
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	// Declaring a pointer
	var ptr *int
	num := 10
	ptr = &num // Storing the address of num

	fmt.Println("Value of num:", num)
	fmt.Println("Pointer Address:", ptr)
	fmt.Println("Value at Pointer Address:", *ptr) // Dereferencing pointer

	// Modifying variable through pointer
	fmt.Println("\nModifying Value through Pointer:")
	modifyValue(&num)
	fmt.Println("Modified num:", num)

	// Pointer Swapping Example
	fmt.Println("\nPointer Swapping Example:")
	x, y := 5, 10
	fmt.Println("Before swap: x =", x, ", y =", y)
	swap(&x, &y)
	fmt.Println("After swap: x =", x, ", y =", y)
}
