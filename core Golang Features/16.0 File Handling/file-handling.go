package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a new file or truncate if it exists
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when the function exits
}
