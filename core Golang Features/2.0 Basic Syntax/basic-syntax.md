# Basic Syntax in Go

```go
// Package declaration: Every Go program must belong to a package.
// The "main" package is the entry point for an executable program.
package main

// Import statement: Used to include standard and custom packages.
import "fmt"

// The main function: The execution starts from the main function in the main package.
func main() {
    // Printing to the console using the fmt package.
    fmt.Println("Welcome to Go!")

    // Declaring a variable using the 'var' keyword.
    var message string = "Go is awesome!"
    fmt.Println(message)
}

```

---

## Explanation

### **1. Package Declaration**

- Every Go program must belong to a **package**.
- The `main` package is special because it defines an executable program.

### **2. Importing Packages**

- `import "fmt"` includes the **fmt** package, which provides functions for formatted I/O.

### **3. The `main` Function**

- Execution starts from the `main()` function in the `main` package.

### **4. Defining your Variable/function or any valid Go item**

- Variables can be declared using `var` or the shorthand `:=`.
- Constants use the `const` keyword and cannot be changed.
- Functions using `func`.
- Many more.

---

## Running the Code

1. Save the file as `basic_syntax.go`.
2. Open a terminal and navigate to the file's location.
3. Run:

   ```sh
   go run basic_syntax.go
   ```
