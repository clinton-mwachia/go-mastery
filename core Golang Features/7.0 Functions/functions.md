# Functions in Go

## Introduction

Functions in Go are building blocks of modular code, allowing code reusability and organization. Go supports **normal functions, multiple return values, variadic functions, anonymous functions, closures, and recursion**.

---

## Code

```go
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
```

---

## Explanation

### **1. Defining Functions**

Functions in Go are defined using the `func` keyword.

```go
func functionName(parameters) returnType {
    // function body
}
```

Example:

```go
func greet(name string) string {
    return "Hello, " + name
}
```

---

### **2. Multiple Return Values**

Go allows returning multiple values from a function.

```go
func divide(a, b int) (int, string) {
    if b == 0 {
        return 0, "Error: Division by zero"
    }
    return a / b, "Success"
}
```

Usage:

```go
result, msg := divide(10, 2)
fmt.Println(result, msg) // Output: 5 Success
```

---

### **3. Variadic Functions**

Variadic functions accept a variable number of arguments.

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

Usage:

```go
fmt.Println(sum(1, 2, 3, 4)) // Output: 10
```

---

### **4. Anonymous Functions**

Functions without names can be assigned to variables.

```go
multiply := func(a, b int) int {
    return a * b
}
```

Usage:

```go
fmt.Println(multiply(4, 6)) // Output: 24
```

---

### **5. Closures (Function Returning a Function)**

A function can return another function.

```go
func multiplier(factor int) func(int) int {
    return func(n int) int {
        return n * factor
    }
}
```

Usage:

```go
timesThree := multiplier(3)
fmt.Println(timesThree(5)) // Output: 15
```

---

### **6. Recursive Functions**

A function that calls itself is **recursive**.

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```

Usage:

```go
fmt.Println(factorial(5)) // Output: 120
```

---

## Running the Code

1. Save the file as `functions.go`.
2. Open a terminal and navigate to the file's location.
3. Run:

   ```sh
   go run functions.go
   ```

4. Expected output:

   ```
   Function with parameters and return:
   3 + 5 = 8

   Function with multiple return values:
   10 / 2 = 5 Message: Success

   Variadic Function:
   Sum of 1, 2, 3: 6

   Anonymous Function:
   4 * 6 = 24

   Closure Function:
   3 * 5 = 15

   Recursive Function:
   Factorial of 5: 120
   ```

```

```
