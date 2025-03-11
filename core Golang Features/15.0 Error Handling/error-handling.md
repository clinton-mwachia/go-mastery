## Error Handling

In Go, **error handling** is managed through explicit, separate return values, differing from exception-based approaches in languages like Java or Python. This design encourages developers to handle errors proactively, ensuring robust and predictable code.

**Defining the `error` Type:**

The `error` type in Go is a built-in interface:

```go
type error interface {
    Error() string
}
```

Any type that implements the `Error()` method with a string return satisfies the `error` interface.

**Basic Error Handling:**

Functions in Go often return an error as the last return value. It's idiomatic to check this value to determine if the function executed successfully.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()
    // Proceed with file operations
}
```

In this example, `os.Open` attempts to open a file and returns a file pointer and an error. If an error occurs, it's handled immediately, and the function exits gracefully.

**Creating Custom Errors:**

Developers can define custom error messages using the `errors.New` function from the `errors` package.

```go
package main

import (
    "errors"
    "fmt"
)

func validateAge(age int) error {
    if age < 18 {
        return errors.New("age must be 18 or older")
    }
    return nil
}

func main() {
    err := validateAge(16)
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

Here, `validateAge` returns an error if the age is less than 18. The main function checks this error and handles it accordingly.

**Custom Error Types:**

For more detailed error information, custom error types can be implemented by defining a struct that satisfies the `error` interface.

```go
package main

import (
    "fmt"
)

type DivisionError struct {
    Dividend float64
    Divider  float64
}

func (e *DivisionError) Error() string {
    return fmt.Sprintf("cannot divide %.2f by %.2f", e.Dividend, e.Divider)
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, &DivisionError{a, b}
    }
    return a / b, nil
}

func main() {
    result, err := divide(4, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
```

In this code, `DivisionError` provides context about the division operation that caused the error. When a division by zero is attempted, the custom error offers specific details.

**Wrapping Errors:**

Error wrapping, allowing errors to be annotated with additional context while preserving the original error. The `fmt.Errorf` function supports this with the `%w` verb.

```go
package main

import (
    "fmt"
    "errors"
)

func readFile(filename string) error {
    err := errors.New("file not found")
    if err != nil {
        return fmt.Errorf("readFile: %w", err)
    }
    return nil
}

func main() {
    err := readFile("missing.txt")
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

Here, `readFile` wraps the original error with additional context. This approach aids in understanding the error's origin when it's propagated up the call stack.

**Unwrapping and Inspecting Errors:**

To inspect wrapped errors, Go provides the `errors.Is` and `errors.As` functions.

```go
package main

import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("not found")

func findItem(id int) error {
    return fmt.Errorf("findItem: %w", ErrNotFound)
}

func main() {
    err := findItem(123)
    if errors.Is(err, ErrNotFound) {
        fmt.Println("Item not found")
    } else {
        fmt.Println("Error:", err)
    }
}
```

In this example, `errors.Is` checks if the error or any wrapped error matches `ErrNotFound`. This method enables precise error handling based on specific error types.

**Panic and Recover:**

While Go emphasizes explicit error handling, it also provides `panic` and `recover` for handling unexpected conditions. Use these sparingly, as they can lead to less predictable code behavior.

```go
package main

import (
    "fmt"
)

func mayPanic() {
    panic("something went wrong")
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    mayPanic()
    fmt.Println("After mayPanic")
}
```

Here, `mayPanic` triggers a panic, and the deferred function with `recover` handles it, allowing the program to continue running.

**Best Practices:**

- **Handle Errors Immediately:** Check and handle errors as soon as they occur to maintain code clarity and prevent unintended behavior.
- **Provide Context:** When returning errors, add context to aid in debugging. Wrapping errors with additional information is a common practice.
- **Avoid Silent Failures:** Always handle errors explicitly. Ignoring errors can lead to unpredictable program states.
- **Use Panic Sparingly:** Reserve `panic` for unrecoverable situations. For expected errors, use the conventional error handling approach.
