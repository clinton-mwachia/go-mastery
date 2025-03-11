package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// method 1
	go printNumbers()

	// method 2
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}() // Launching printNumbers as a goroutine

	// Keep the main function alive to allow the goroutine to complete
	time.Sleep(600 * time.Millisecond)
	fmt.Println("Main function ends")
}
