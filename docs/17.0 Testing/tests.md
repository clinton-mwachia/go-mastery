## Testing

In Go, testing is an integral part of the development process, facilitated by the language's built-in `testing` package and the `go test` command. This design encourages developers to write and run tests efficiently, promoting code reliability and maintainability.

**Creating a Basic Test:**

To illustrate testing in Go, consider a simple function that adds two integers:

**Create a go module**

```go
go mod init testing
```

**Create a `maths` directory and add the files below**

```go
// math.go
package math

// Add returns the sum of two integers.
func Add(a, b int) int {
    return a + b
}
```

To test this function, follow these steps:

1. **Create a Test File:**

   In the same directory as `math.go`, create a file named `math_test.go`. By convention, test files in Go end with `_test.go`.

2. **Write the Test Function:**

   In `math_test.go`, implement the test:

   ```go
   // math_test.go
   package math

   import "testing"

   // TestAdd tests the Add function.
   func TestAdd(t *testing.T) {
       result := Add(2, 3)
       expected := 5
       if result != expected {
           t.Errorf("Add(2, 3) = %d; want %d", result, expected)
       }
   }
   ```

   Here, `TestAdd` calls the `Add` function and checks if the result matches the expected value. If not, it reports an error.

3. **Run the Test:**

   Execute the test using the `go test` command:

   ```bash
   $ go test
   ```

   If the test passes, you'll see output indicating success.

**Table-Driven Tests:**

For functions that require multiple test cases, table-driven tests are a common pattern:

```go
// math_test.go
package math

import "testing"

// TestAdd tests the Add function using table-driven tests.
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"zero and positive", 0, 4, 4},
        {"negative and positive", -1, 1, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

In this approach, a slice of test cases is defined, and each case is run using `t.Run`, which provides better test reporting.

**Benchmark Tests:**

Go also supports benchmarking to measure performance:

```go
// math_test.go
package math

import "testing"

// BenchmarkAdd benchmarks the Add function.
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

Run benchmarks with:

```bash
$ go test -bench=.
```

**Code Coverage:**

To assess test coverage, use the `-cover` flag:

```bash
$ go test -cover
```

This command reports the percentage of code covered by tests.

By leveraging Go's testing capabilities, you can ensure your code is both functional and efficient, leading to more reliable software development.
