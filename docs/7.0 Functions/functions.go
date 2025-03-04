package main

import "fmt"

// Function with parameters and return value
func add(a int, b int) int {
	return a + b
}

// Function with multiple return values
func divide(numerator, denominator float64) (float64, string) {
	if denominator == 0 {
		return 0, "Error: Division by zero"
	}
	return numerator / denominator, "Success"
}

// Variadic function (accepts multiple arguments)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Anonymous function assigned to a variable
var multiply = func(a, b int) int {
	return a * b
}

// Function that returns another function (Closure)
func multiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// Recursive function (Factorial)
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("Function with parameters and return:")
	fmt.Println("3 + 5 =", add(3, 5))

	fmt.Println("\nFunction with multiple return values:")
	result, message := divide(10, 2)
	fmt.Println("10 / 2 =", result, "Message:", message)

	fmt.Println("\nVariadic Function:")
	fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))

	fmt.Println("\nAnonymous Function:")
	fmt.Println("4 * 6 =", multiply(4, 6))

	fmt.Println("\nClosure Function:")
	timesThree := multiplier(3)
	fmt.Println("3 * 5 =", timesThree(5))

	fmt.Println("\nRecursive Function:")
	fmt.Println("Factorial of 5:", factorial(5))
}
