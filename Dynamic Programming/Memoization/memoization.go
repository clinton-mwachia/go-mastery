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
