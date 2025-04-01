## String Comparison

In Go, comparing strings is a fundamental operation that can be achieved using various methods. Below are the primary techniques for string comparison, along with code examples and explanations:

**1. Using Comparison Operators**

Go supports the standard comparison operators for strings: `==`, `!=`, `<`, `>`, `<=`, and `>=`. These operators perform lexicographical comparisons based on Unicode code points.

_Example:_

```go
package main

import "fmt"

func main() {
    str1 := "apple"
    str2 := "banana"

    // Equality check
    if str1 == str2 {
        fmt.Println("str1 and str2 are equal.")
    } else {
        fmt.Println("str1 and str2 are not equal.")
    }

    // Lexicographical comparison
    if str1 < str2 {
        fmt.Println("str1 comes before str2.")
    } else {
        fmt.Println("str1 comes after str2.")
    }
}
```

_Output:_

```
str1 and str2 are not equal.
str1 comes before str2.
```

**2. Using `strings.Compare` Function**

The `strings` package provides the `Compare` function, which performs a three-way comparison between two strings. It returns:

- `0` if `a == b`
- `-1` if `a < b`
- `+1` if `a > b`

_Example:_

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    result := strings.Compare("apple", "banana")
    fmt.Println(result) // Output: -1

    result = strings.Compare("grape", "grape")
    fmt.Println(result) // Output: 0

    result = strings.Compare("orange", "grape")
    fmt.Println(result) // Output: 1
}
```

_Note:_ While `strings.Compare` is available, it's generally clearer and more idiomatic to use the built-in comparison operators (`==`, `<`, `>`, etc.) for readability and performance.

**3. Case-Insensitive Comparison with `strings.EqualFold`**

For case-insensitive comparisons, Go provides the `strings.EqualFold` function. This function compares two strings and returns `true` if they are equal, irrespective of case.

_Example:_

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    str1 := "GoLang"
    str2 := "golang"

    if strings.EqualFold(str1, str2) {
        fmt.Println("str1 and str2 are equal (case-insensitive).")
    } else {
        fmt.Println("str1 and str2 are not equal (case-insensitive).")
    }
}
```

_Output:_

```
str1 and str2 are equal (case-insensitive).
```

**4. Handling Whitespace and Newlines**

When comparing strings that may contain leading or trailing whitespace (including newlines), it's advisable to use `strings.TrimSpace` to remove any extraneous whitespace before performing the comparison.

_Example:_

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    input := "hello\n"
    target := "hello"

    if strings.TrimSpace(input) == target {
        fmt.Println("The strings are equal after trimming.")
    } else {
        fmt.Println("The strings are not equal after trimming.")
    }
}
```

_Output:_

```
The strings are equal after trimming.
```

By utilizing these methods, you can effectively compare strings in Go, tailoring the approach to the specific requirements of your application.
