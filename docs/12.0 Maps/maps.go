package main

import "fmt"

func main() {
	// Initialize a map with predefined key-value pairs
	m := map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 3,
	}

	// Retrieve a value
	appleCount := m["apple"]
	fmt.Println(appleCount) // Output: 2

	// add key-value to the map
	m["pineapple"] = 10

	// delete item
	delete(m, "apple")
	// confirmiv value is deleted
	value, exists := m["apple"]
	if exists {
		fmt.Println("apple count:", value)
	} else {
		fmt.Println("apple not found")
	}

	// Iterate over the map
	for fruit, count := range m {
		fmt.Printf("%s: %d\n", fruit, count)
	}
}
