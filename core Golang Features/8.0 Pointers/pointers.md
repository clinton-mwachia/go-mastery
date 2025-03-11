# Pointers in Go

## Introduction

Pointers in Go allow us to work with memory addresses, enabling efficient function calls, avoiding unnecessary copying of large data structures, and modifying variables directly.

---

## Code

```go
package main

import "fmt"

// Function that modifies a variable using a pointer
func modifyValue(val *int) {
    *val = *val * 2
}

// Function demonstrating pointer swapping
func swap(a, b *int) {
    temp := *a
    *a = *b
    *b = temp
}

func main() {
    // Declaring a pointer
    var ptr *int
    num := 10
    ptr = &num // Storing the address of num

    fmt.Println("Value of num:", num)
    fmt.Println("Pointer Address:", ptr)
    fmt.Println("Value at Pointer Address:", *ptr) // Dereferencing pointer

    // Modifying variable through pointer
    fmt.Println("\nModifying Value through Pointer:")
    modifyValue(&num)
    fmt.Println("Modified num:", num)

    // Pointer Swapping Example
    fmt.Println("\nPointer Swapping Example:")
    x, y := 5, 10
    fmt.Println("Before swap: x =", x, ", y =", y)
    swap(&x, &y)
    fmt.Println("After swap: x =", x, ", y =", y)
}
```

---

## Explanation

### **1. What is a Pointer?**

A **pointer** is a variable that stores the **memory address** of another variable. In Go, a pointer is defined using the `*` (asterisk) symbol.

```go
var ptr *int // Declaring a pointer to an int
num := 10
ptr = &num // Storing the address of num in ptr
```

### **2. Accessing Pointer Values (Dereferencing)**

To **access the value stored at the memory address** a pointer refers to, we use the `*` operator. This is called **dereferencing**.

```go
fmt.Println(*ptr) // Prints value stored at the address in ptr
```

---

### **3. Passing Pointers to Functions**

Passing variables by **reference** (using pointers) allows functions to **modify the original variable**.

Example:

```go
func modifyValue(val *int) {
    *val = *val * 2
}
```

Usage:

```go
num := 10
modifyValue(&num)
fmt.Println(num) // Output: 20
```

---

### **4. Swapping Two Values Using Pointers**

Since Go does not support **pass-by-reference** directly, pointers are used to swap values in a function.

```go
func swap(a, b *int) {
    temp := *a
    *a = *b
    *b = temp
}
```

Usage:

```go
x, y := 5, 10
swap(&x, &y)
fmt.Println(x, y) // Output: 10 5
```

---

## Running the Code

1. Save the file as `pointers.go`.
2. Open a terminal and navigate to the file's location.
3. Run:

   ```sh
   go run pointers.go
   ```

4. Expected output:

   ```go
   Value of num: 10
   Pointer Address: 0xc0000140a8
   Value at Pointer Address: 10

   Modifying Value through Pointer:
   Modified num: 20

   Pointer Swapping Example:
   Before swap: x = 5 , y = 10
   After swap: x = 10 , y = 5
   ```
