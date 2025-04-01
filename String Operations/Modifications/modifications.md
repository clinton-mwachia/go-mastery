## String Modifications

In Go, the `strings` package offers a variety of functions for string manipulation, including `Map` and other modification functions. Here's an overview of these functions along with code examples:

**1. `strings.Map` Function**

The `strings.Map` function applies a given mapping function to each character in a string, returning a new string with the modified characters. If the mapping function returns a negative value, the character is omitted from the resulting string.

_Function Signature:_

```go
func Map(mapping func(rune) rune, s string) string
```

_Example Usage:_

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    // Define a mapping function that converts letters to uppercase
    toUpper := func(r rune) rune {
        if r >= 'a' && r <= 'z' {
            return r - 'a' + 'A'
        }
        return r
    }

    original := "Hello, Gophers!"
    modified := strings.Map(toUpper, original)
    fmt.Println(modified) // Output: "HELLO, GOPHERS!"
}
```

In this example, the `toUpper` function converts lowercase letters to uppercase, and `strings.Map` applies this function to each character in the `original` string.

**2. `strings.Replace` and `strings.ReplaceAll` Functions**

These functions replace occurrences of a substring within a string.

_Function Signatures:_

```go
func Replace(s, old, new string, n int) string
func ReplaceAll(s, old, new string) string
```

_Example Usage:_

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    original := "oink oink oink"
    modified := strings.ReplaceAll(original, "oink", "moo")
    fmt.Println(modified) // Output: "moo moo moo"
}
```

Here, `strings.ReplaceAll` replaces all occurrences of "oink" with "moo" in the `original` string.

**3. `strings.Repeat` Function**

The `strings.Repeat` function returns a new string consisting of a specified number of copies of the original string.

_Function Signature:_

```go
func Repeat(s string, count int) string
```

_Example Usage:_

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    repeated := strings.Repeat("na", 2)
    fmt.Println("ba" + repeated) // Output: "banana"
}
```

In this example, `strings.Repeat` creates a string by repeating "na" twice, resulting in "banana" when concatenated with "ba".

**4. `strings.Trim` and Related Functions**

These functions remove leading and trailing characters from a string.

_Function Signatures:_

```go
func Trim(s, cutset string) string
func TrimFunc(s string, f func(rune) bool) string
func TrimLeft(s, cutset string) string
func TrimLeftFunc(s string, f func(rune) bool) string
func TrimRight(s, cutset string) string
func TrimRightFunc(s string, f func(rune) bool) string
```

_Example Usage:_

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    original := "¡¡¡Hello, Gophers!!!"
    trimmed := strings.Trim(original, "!¡")
    fmt.Println(trimmed) // Output: "Hello, GOPHERS"
}
```

Here, `strings.Trim` removes all leading and trailing characters specified in the `cutset` from the `original` string.

These functions, along with others in the `strings` package, provide powerful tools for string manipulation in Go. For more detailed information, refer to the official Go documentation.
