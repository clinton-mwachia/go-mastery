# Control Flow in Go

## Introduction

Control flow structures allow the execution of code to follow different paths based on conditions, loops, or branches. Go provides several control flow mechanisms, including **if-else statements, switch cases, loops (for), and defer, panic, and recover** for handling errors.

---

## Code

```go
package main

import "fmt"

func main() {
    // If-Else Statement
    age := 18
    fmt.Println("If-Else Statement:")
    if age >= 18 {
        fmt.Println("You are an adult.")
    } else {
        fmt.Println("You are a minor.")
    }

    // If with Short Statement
    fmt.Println("\nIf with Short Statement:")
    if num := 10; num%2 == 0 {
        fmt.Println(num, "is even.")
    } else {
        fmt.Println(num, "is odd.")
    }

    // Switch Statement
    fmt.Println("\nSwitch Statement:")
    day := 3
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    default:
        fmt.Println("Other day")
    }

    // Switch without Condition (Equivalent to Multiple If-Else)
    fmt.Println("\nSwitch without Condition:")
    score := 85
    switch {
    case score >= 90:
        fmt.Println("Grade: A")
    case score >= 80:
        fmt.Println("Grade: B")
    case score >= 70:
        fmt.Println("Grade: C")
    default:
        fmt.Println("Grade: F")
    }

    // For Loop
    fmt.Println("\nFor Loop:")
    for i := 1; i <= 5; i++ {
        fmt.Println("Iteration:", i)
    }

    // For Loop with range
	fmt.Println("\nFor Loop with range:")
	nums := []int{2, 3, 4}
	for i, num := range nums {
		fmt.Println("Index:", i, "Num: ", num)
	}

    // While-like Loop
    fmt.Println("\nWhile-like Loop:")
    j := 1
    for j <= 5 {
        fmt.Println("Value of j:", j)
        j++
    }

    // Infinite Loop with Break
    fmt.Println("\nInfinite Loop with Break:")
    k := 1
    for {
        if k > 3 {
            break
        }
        fmt.Println("k:", k)
        k++
    }

    // Loop with Continue
    fmt.Println("\nLoop with Continue:")
    for i := 1; i <= 5; i++ {
        if i == 3 {
            continue // Skip iteration when i == 3
        }
        fmt.Println("i:", i)
    }

    // Defer Statement
    fmt.Println("\nDefer Statement:")
    defer fmt.Println("This will execute at the end of main function.")
    fmt.Println("Hello, World!")

    // Panic and Recover
    fmt.Println("\nPanic and Recover:")
    safeFunction()
}

// Function demonstrating panic and recover
func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    fmt.Println("Executing safeFunction...")
    panic("Something went wrong!") // Causes a panic
}
```

---

## Explanation

### **1. If-Else Statements**

Used for conditional branching.

| Statement                 | Description                                    |
| ------------------------- | ---------------------------------------------- |
| `if`                      | Executes a block if the condition is `true`.   |
| `else`                    | Executes if the `if` condition is `false`.     |
| `if-else if-else`         | Multiple conditions checked sequentially.      |
| `if` with short statement | Declares a variable inside the `if` condition. |

Example:

```go
if x > 10 {
    fmt.Println("x is greater than 10")
} else {
    fmt.Println("x is 10 or less")
}
```

---

### **2. Switch Case**

Simplifies multiple conditional checks.

| Feature                    | Description                        |
| -------------------------- | ---------------------------------- |
| `switch case`              | Checks for equality of a variable. |
| `switch without condition` | Behaves like multiple `if-else`.   |
| `default`                  | Executes when no case matches.     |

Example:

```go
switch value {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
default:
    fmt.Println("Other number")
}
```

---

### **3. Loops (`for`)**

Go only has the `for` loop for iterations.

| Type                             | Description                              |
| -------------------------------- | ---------------------------------------- |
| `for init; condition; increment` | Traditional loop (like `for` in C/Java). |
| `for condition`                  | While-like loop.                         |
| `for {}`                         | Infinite loop.                           |
| `continue`                       | Skips current iteration.                 |
| `break`                          | Exits the loop early.                    |

Example:

```go
for i := 0; i < 5; i++ {
    fmt.Println("Iteration:", i)
}
```

---

### **4. Defer Statement**

Defers execution of a function until the surrounding function exits.

Example:

```go
func example() {
    defer fmt.Println("This runs last.")
    fmt.Println("This runs first.")
}
```

Output:

```
This runs first.
This runs last.
```

---

### **5. Panic and Recover**

- `panic`: Stops normal execution and starts unwinding the stack.
- `recover`: Catches a panic and allows the program to continue running.

Example:

```go
func handleError() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    panic("Unexpected error!")
}
```

---

## Running the Code

1. Save the file as `control-flow.go`.
2. Open a terminal and navigate to the file's location.
3. Run:

   ```sh
   go run control-flow.go
   ```

4. Expected output:

   ```
   If-Else Statement:
   You are an adult.

   If with Short Statement:
   10 is even.

   Switch Statement:
   Wednesday

   Switch without Condition:
   Grade: B

   For Loop:
   Iteration: 1
   Iteration: 2
   ...

   Defer Statement:
   Hello, World!
   This will execute at the end of main function.

   Panic and Recover:
   Executing safeFunction...
   Recovered from panic: Something went wrong!
   ```
