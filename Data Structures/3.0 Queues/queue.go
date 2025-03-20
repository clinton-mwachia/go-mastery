package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 3)

	queue <- "task1"
	queue <- "task2"
	queue <- "task3"

	close(queue)

	for item := range queue {
		fmt.Println("Dequeued:", item)
	}
}
