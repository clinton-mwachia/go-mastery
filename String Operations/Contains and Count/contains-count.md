## Contains and Count

In Go, the `strings` package provides the `Contains` and `Count` functions for efficient string manipulation. Below are detailed explanations and examples for each:

**1. `strings.Contains` Function**

The `strings.Contains` function checks if a specified substring exists within a given string. It returns a boolean value: `true` if the substring is found, and `false` otherwise.

**Syntax:**

```go
func Contains(s, substr string) bool
```

- `s`: The string to be searched.
- `substr`: The substring to search for.

**Example:**

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    mainStr := "Hello, World!"
    subStr1 := "World"
    subStr2 := "Gopher"

    fmt.Printf("Does \"%s\" contain \"%s\"? %t\n", mainStr, subStr1, strings.Contains(mainStr, subStr1))
    fmt.Printf("Does \"%s\" contain \"%s\"? %t\n", mainStr, subStr2, strings.Contains(mainStr, subStr2))
}
```

**Output:**

```
Does "Hello, World!" contain "World"? true
Does "Hello, World!" contain "Gopher"? false
```

In this example, `strings.Contains` checks for the presence of "World" and "Gopher" in "Hello, World!", returning `true` and `false` respectively.

**2. `strings.Count` Function**

The `strings.Count` function counts the number of non-overlapping instances of a specified substring within a given string.

**Syntax:**

```go
func Count(s, substr string) int
```

- `s`: The string to be searched.
- `substr`: The substring to count.

**Example:**

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    mainStr := "banana"
    subStr := "ana"
    char := "a"

    fmt.Printf("The substring \"%s\" appears %d times in \"%s\".\n", subStr, strings.Count(mainStr, subStr), mainStr)
    fmt.Printf("The character \"%s\" appears %d times in \"%s\".\n", char, strings.Count(mainStr, char), mainStr)
}
```

**Output:**

```
The substring "ana" appears 1 times in "banana".
The character "a" appears 3 times in "banana".
```

In this example, `strings.Count` determines that the substring "ana" appears once in "banana", and the character "a" appears three times.

**Note:** `strings.Count` counts non-overlapping instances. For example, in the string "banana", the substring "ana" is counted once, not twice, because the second "ana" overlaps with the first.

For more information and additional string functions, refer to the [Go `strings` package documentation](https://pkg.go.dev/strings).
