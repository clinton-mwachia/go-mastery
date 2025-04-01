## String Indexing

In Go, the `strings` package provides several functions to locate the position of substrings within a string. Two primary functions for this purpose are `Index` and `LastIndex`.

**1. `strings.Index` Function**

The `Index` function returns the index of the first occurrence of a specified substring within a string. If the substring is not found, it returns `-1`.

_Syntax_:

```go
func Index(s, substr string) int
```

_Parameters_:

- `s`: The string to be searched.
- `substr`: The substring to find within `s`.

_Returns_:

- The index of the first occurrence of `substr` in `s`, or `-1` if `substr` is not present.

_Example_:

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "Hello, World!"
    substr := "World"
    index := strings.Index(str, substr)
    if index != -1 {
        fmt.Printf("The first occurrence of %q starts at index %d.\n", substr, index)
    } else {
        fmt.Printf("%q not found in %q.\n", substr, str)
    }
}
```

_Output_:

```
The first occurrence of "World" starts at index 7.
```

**2. `strings.LastIndex` Function**

The `LastIndex` function returns the index of the last occurrence of a specified substring within a string. If the substring is not found, it returns `-1`.

_Syntax_:

```go
func LastIndex(s, substr string) int
```

_Parameters_:

- `s`: The string to be searched.
- `substr`: The substring to find within `s`.

_Returns_:

- The index of the last occurrence of `substr` in `s`, or `-1` if `substr` is not present.

_Example_:

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "Go gopher, go!"
    substr := "go"
    index := strings.LastIndex(str, substr)
    if index != -1 {
        fmt.Printf("The last occurrence of %q starts at index %d.\n", substr, index)
    } else {
        fmt.Printf("%q not found in %q.\n", substr, str)
    }
}
```

_Output_:

```
The last occurrence of "go" starts at index 12.
```

**Additional Indexing Functions**

Go's `strings` package also offers other indexing functions:

- `IndexAny`: Returns the index of the first occurrence of any Unicode code point from a set of characters.
- `LastIndexAny`: Returns the index of the last occurrence of any Unicode code point from a set of characters.
- `IndexByte`: Returns the index of the first occurrence of a specific byte.
- `LastIndexByte`: Returns the index of the last occurrence of a specific byte.
- `IndexFunc`: Returns the index of the first rune satisfying a given function.
- `LastIndexFunc`: Returns the index of the last rune satisfying a given function.

These functions provide flexibility for various string searching needs. For more detailed information and examples, refer to the [Go `strings` package documentation](https://pkg.go.dev/strings).
