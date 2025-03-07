In Go, a **goroutine** is a lightweight thread managed by the Go runtime, enabling concurrent execution of functions. Goroutines are fundamental to Go's concurrency model, allowing developers to perform multiple tasks simultaneously without the complexity of traditional threading mechanisms.

**Key Characteristics of Goroutines:**

- **Lightweight:** Goroutines consume minimal memory, allowing the creation of thousands of concurrent tasks within a single program.
- **Dynamic Growth:** The stack size of a goroutine starts small and grows as needed, optimizing resource usage.
- **Concurrent Execution:** Goroutines run concurrently with other goroutines, facilitating efficient multitasking.

**Creating and Using Goroutines:**

To start a new goroutine, prepend the `go` keyword to a function call. Here's an example:

```go
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
```

**Explanation:**

- The `printNumbers` function prints numbers from 1 to 5, pausing for 100 milliseconds between each print.
- In the `main` function, `printNumbers` is invoked as a goroutine using the `go` keyword.
- A `time.Sleep` call in the `main` function ensures it remains active long enough for the goroutine to complete its execution.

**Output:**

```
1
1
2
2
3
3
4
4
5
5
Main function ends
```
