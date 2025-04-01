## Titles and Cases

In Go, the `strings` package provides functions for case conversion, including converting strings to uppercase, lowercase, and title case. Additionally, the `golang.org/x/text/cases` package offers more advanced title casing functionality.

**1. ToUpper and ToLower Functions**

The `strings` package includes `ToUpper` and `ToLower` functions to convert strings to uppercase and lowercase, respectively.

_Example:_

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    original := "Hello, World!"
    upper := strings.ToUpper(original)
    lower := strings.ToLower(original)
    fmt.Println("Original:", original)
    fmt.Println("Uppercase:", upper)
    fmt.Println("Lowercase:", lower)
}
```

_Output:_

```
Original: Hello, World!
Uppercase: HELLO, WORLD!
Lowercase: hello, world!
```

**2. ToTitle Function**

The `strings` package provides the `ToTitle` function, which converts all Unicode letters in a string to their title case. However, it's important to note that `strings.Title` has been deprecated due to its inconsistent handling of word boundaries.

_Example:_

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    original := "hello, world!"
    title := strings.ToTitle(original)
    fmt.Println("Original:", original)
    fmt.Println("Title Case:", title)
}
```

_Output:_

```
Original: hello, world!
Title Case: HELLO, WORLD!
```

**3. Advanced Title Case Conversion with `golang.org/x/text/cases`**

For more accurate title case conversion that respects word boundaries, you can use the `cases` package from `golang.org/x/text`. This package provides the `Title` function, which offers improved handling of title casing.

_Initialize the module_

```go
go mod init case-titles
```

_First, install the package:_

```bash
go get golang.org/x/text/cases
go get golang.org/x/text/language
```

_Example:_

```go
package main

import (
    "fmt"
    "golang.org/x/text/cases"
    "golang.org/x/text/language"
)

func main() {
    original := "hello, world!"
    titleCaser := cases.Title(language.English)
    title := titleCaser.String(original)
    fmt.Println("Original:", original)
    fmt.Println("Title Case:", title)
}
```

_Output:_

```
Original: hello, world!
Title Case: Hello, World!
```

This approach provides more accurate title casing by considering language-specific rules and word boundaries.

**Documentation References:**

- [strings package - GoDoc](https://pkg.go.dev/strings)
- [text/cases package](https://pkg.go.dev/golang.org/x/text/cases)

By utilizing these functions, you can effectively perform case conversions in your Go programs, ensuring that strings are formatted as needed for your application's requirements.
