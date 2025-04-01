## Cutting

In Go's `strings` package, the "cutting" functions are designed to split strings based on specified separators or to remove specific prefixes and suffixes. These functions include `Cut`, `CutPrefix`, and `CutSuffix`. Below is a detailed explanation and examples for each:

**1. `strings.Cut`**

The `Cut` function slices a string into two parts based on the first occurrence of a specified separator. It returns the text before the separator, the text after the separator, and a boolean indicating whether the separator was found. This function simplifies common operations that previously required combinations of `Index`, `IndexByte`, `IndexRune`, and `SplitN`.

**Syntax:**

```go
func Cut(s, sep string) (before, after string, found bool)
```

**Example:**

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "username:password"
    before, after, found := strings.Cut(str, ":")
    if found {
        fmt.Printf("Before: %s\nAfter: %s\n", before, after)
    } else {
        fmt.Println("Separator not found")
    }
}
```

**Output:**

```
Before: username
After: password
```

In this example, the `Cut` function splits the string `str` at the first occurrence of the colon (`:`), resulting in `before` containing "username" and `after` containing "password".

**2. `strings.CutPrefix`**

The `CutPrefix` function checks if a string starts with a specified prefix. If it does, the function returns the string without the prefix and a boolean value of `true`. If the string does not start with the prefix, it returns the original string and `false`. citeturn0search5

**Syntax:**

```go
func CutPrefix(s, prefix string) (after string, found bool)
```

**Example:**

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "Hello, World!"
    prefix := "Hello, "
    if after, found := strings.CutPrefix(str, prefix); found {
        fmt.Printf("String without prefix: %s\n", after)
    } else {
        fmt.Println("Prefix not found")
    }
}
```

**Output:**

```
String without prefix: World!
```

Here, `CutPrefix` removes the "Hello, " prefix from the string, leaving "World!".

**3. `strings.CutSuffix`**

The `CutSuffix` function operates similarly to `CutPrefix` but works on suffixes. It checks if a string ends with a specified suffix and, if so, returns the string without the suffix and `true`. Otherwise, it returns the original string and `false`.

**Syntax:**

```go
func CutSuffix(s, suffix string) (before string, found bool)
```

**Example:**

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "example.txt"
    suffix := ".txt"
    if before, found := strings.CutSuffix(str, suffix); found {
        fmt.Printf("String without suffix: %s\n", before)
    } else {
        fmt.Println("Suffix not found")
    }
}
```

**Output:**

```
String without suffix: example
```

In this case, `CutSuffix` removes the ".txt" suffix from "example.txt", resulting in "example".

These functions provide efficient and readable ways to manipulate strings by cutting them based on specific prefixes, suffixes, or separators. They enhance code clarity and reduce the need for more verbose string manipulation techniques.
