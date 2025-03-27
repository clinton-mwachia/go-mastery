## Memoization

Memoization is an optimization technique that enhances the performance of recursive functions by storing the results of expensive function calls and returning the cached result when the same inputs occur again. This approach is particularly beneficial in dynamic programming, where solutions to subproblems are reused multiple times.

**Implementing Memoization in Go**

In Go, memoization can be implemented using a map to cache computed results. Below is a detailed example demonstrating how to apply memoization to compute Fibonacci numbers.

**Example: Fibonacci Sequence with Memoization**

The Fibonacci sequence is a classic example to illustrate memoization. A naive recursive implementation recalculates the same values multiple times, leading to exponential time complexity. By employing memoization, we can store already computed values and retrieve them when needed, significantly improving performance.

```go
package main

import (
	"fmt"
)

// fibonacci function calculates the nth Fibonacci number using memoization.
func fibonacci(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	// Check if the result is already computed
	if val, exists := memo[n]; exists {
		return val
	}
	// Compute the result and store it in the memoization map
	memo[n] = fibonacci(n-1, memo) + fibonacci(n-2, memo)
	return memo[n]
}

func main() {
	n := 40
	memo := make(map[int]int)
	result := fibonacci(n, memo)
	fmt.Printf("Fibonacci number at position %d is %d\n", n, result)
}
```

**Explanation:**

1. **Base Case:** If `n` is 0 or 1, the function returns `n` directly, as these are the first two numbers in the Fibonacci sequence.

2. **Memoization Check:** Before computing the Fibonacci number for `n`, the function checks if the result already exists in the `memo` map. If it does, the cached result is returned, avoiding redundant calculations.

3. **Recursive Calculation:** If the result is not in the cache, the function recursively calculates the Fibonacci numbers for `n-1` and `n-2`, stores the computed result in the `memo` map, and then returns it.

4. **Initialization:** In the `main` function, a `memo` map is initialized and passed to the `fibonacci` function to store computed values.

This implementation reduces the time complexity from exponential to linear, as each Fibonacci number is computed only once and stored for future reference.

**Advanced Memoization with Structs**

For more complex scenarios, encapsulating the memoization logic within a struct can provide better organization and reusability.

```go
package main

import (
	"fmt"
)

// Memoizer struct encapsulates the memoization map.
type Memoizer struct {
	cache map[int]int
}

// NewMemoizer creates a new instance of Memoizer.
func NewMemoizer() *Memoizer {
	return &Memoizer{cache: make(map[int]int)}
}

// Fibonacci calculates the nth Fibonacci number using memoization.
func (m *Memoizer) Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	if val, exists := m.cache[n]; exists {
		return val
	}
	m.cache[n] = m.Fibonacci(n-1) + m.Fibonacci(n-2)
	return m.cache[n]
}

func main() {
	n := 40
	memoizer := NewMemoizer()
	result := memoizer.Fibonacci(n)
	fmt.Printf("Fibonacci number at position %d is %d\n", n, result)
}
```

**Explanation:**

- **Memoizer Struct:** Contains a `cache` map to store computed Fibonacci numbers.

- **NewMemoizer Function:** Initializes and returns a new `Memoizer` instance with an empty cache.

- **Fibonacci Method:** Performs the memoized calculation, similar to the previous example, but encapsulated within the `Memoizer` struct.

This approach promotes better organization, especially when dealing with multiple functions requiring memoization.

**Considerations:**

- **Concurrency:** The above implementations are not safe for concurrent use. If multiple goroutines access the memoization map simultaneously, data races may occur. To handle concurrency, consider using synchronization mechanisms like mutexes or the `sync.Map` type.

- **Memory Usage:** Memoization trades increased memory usage for faster computation. Ensure that the memory overhead is acceptable for your application's requirements.

By integrating memoization into your Go programs, you can significantly enhance the efficiency of recursive functions and optimize performance in scenarios involving repeated computations.
