package main

import (
	"fmt"
	"time"
)

func main() {
	// Create an unbuffered channel
	ch := make(chan string)

	// Start a goroutine
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello from goroutine"
	}()

	// Receive value from channel
	msg := <-ch
	fmt.Println(msg)

	// Create a buffered channel with capacity 2
	ch2 := make(chan int, 2)

	// Send values to the channel
	ch2 <- 1
	ch2 <- 2

	// Receive values from the channel
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	// closing channels
	close(ch)

	// Range over the channel to receive values until it's closed
	for value := range ch {
		fmt.Println(value)
	}
}
