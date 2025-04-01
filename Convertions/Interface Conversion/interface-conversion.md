## 6. Interface Type Conversions

In Go, interfaces can hold values of any type that implements the interface. Type assertions and type switches are used to extract the underlying value or determine its type at runtime.

_Type Assertion Example:_

```go
package main

import "fmt"

func main() {
    var i interface{} = "hello"

    s, ok := i.(string)
    if ok {
        fmt.Println(s) // Successfully asserted i holds a string
    } else {
        fmt.Println("i does not hold a string")
    }
}
```

In this example, a type assertion checks if the interface `i` holds a string and extracts it if true. citeturn0search6

_Type Switch Example:_

```go
package main

import "fmt"

func main() {
    var i interface{} = 42

    switch v := i.(type) {
    case int:
        fmt.Printf("i is an int: %d\n", v)
    case string:
        fmt.Printf("i is a string: %s\n", v)
    default:
        fmt.Println("i is of an unknown type")
    }
}
```

This code uses a type switch to determine the type of the value held by the interface `i` and handles each type accordingly. citeturn0search6
