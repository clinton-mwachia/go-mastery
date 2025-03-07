# Maps In Go

In Go, a **map** is a built-in data type that associates keys with values, allowing for efficient data retrieval based on unique identifiers. Maps are commonly used for tasks like counting occurrences, grouping data, or any scenario where a key-value relationship is appropriate.

**Key Characteristics of Maps:**

- **Key-Value Pairs:** Each key is unique and maps to a specific value.
- **Dynamic Size:** Maps can grow or shrink as elements are added or removed.
- **Reference Type:** Maps are reference types, meaning they point to the underlying data and can be `nil` if not initialized.

**Creating and Initializing Maps:**

There are several ways to create and initialize maps in Go:

1. **Using the `make` Function:**

   ```go
   // Declare and initialize an empty map with string keys and int values
   m := make(map[string]int)
   ```

2. **Using a Map Literal:**

   ```go
   // Initialize a map with predefined key-value pairs
   m := map[string]int{
       "apple":  2,
       "banana": 5,
       "orange": 3,
   }
   ```

**Adding and Updating Elements:**

To add or update elements in a map, use the following syntax:

```go
m["key"] = value
```

**Example:**

```go
package main

import "fmt"

func main() {
    // Initialize a map
    m := make(map[string]int)

    // Add key-value pairs
    m["apple"] = 2
    m["banana"] = 5

    // Update a value
    m["apple"] = 4

    fmt.Println(m) // Output: map[apple:4 banana:5]
}
```

**Retrieving Elements:**

To retrieve a value from a map, use the key:

```go
value := m["key"]
```

**Example:**

```go
package main

import "fmt"

func main() {
    // Initialize a map
    m := map[string]int{
        "apple":  2,
        "banana": 5,
    }

    // Retrieve a value
    appleCount := m["apple"]
    fmt.Println(appleCount) // Output: 2
}
```

**Checking for Key Existence:**

To determine if a key exists in a map, use the two-value assignment:

```go
value, exists := m["key"]
```

If the key exists, `exists` will be `true`; otherwise, it will be `false`.

**Example:**

```go
package main

import "fmt"

func main() {
    // Initialize a map
    m := map[string]int{
        "apple":  2,
        "banana": 5,
    }

    // Check if a key exists
    value, exists := m["orange"]
    if exists {
        fmt.Println("Orange count:", value)
    } else {
        fmt.Println("Orange not found")
    }
    // Output: Orange not found
}
```

**Deleting Elements:**

To remove a key-value pair from a map, use the `delete` function:

```go
delete(m, "key")
```

**Example:**

```go
package main

import "fmt"

func main() {
    // Initialize a map
    m := map[string]int{
        "apple":  2,
        "banana": 5,
    }

    // Delete a key-value pair
    delete(m, "apple")

    fmt.Println(m) // Output: map[banana:5]
}
```

**Iterating Over a Map:**

Use a `for` loop with the `range` keyword to iterate over all key-value pairs in a map:

```go
for key, value := range m {
    // Process key and value
}
```

**Example:**

```go
package main

import "fmt"

func main() {
    // Initialize a map
    m := map[string]int{
        "apple":  2,
        "banana": 5,
    }

    // Iterate over the map
    for fruit, count := range m {
        fmt.Printf("%s: %d\n", fruit, count)
    }
    // Output:
    // apple: 2
    // banana: 5
}
```

**Important Considerations:**

- **Key Types:** The key type must be comparable, meaning it supports the `==` and `!=` operators. Valid key types include numbers, strings, pointers, and structs with comparable fields. Slices, maps, and functions cannot be used as map keys.

- **Zero Value:** The zero value of a map is `nil`. A `nil` map behaves like an empty map when reading but causes a runtime panic if you attempt to add or modify elements. Always initialize maps before use.

- **Concurrency:** Maps are not safe for concurrent use. If multiple goroutines access a map simultaneously without synchronization, the program may exhibit unpredictable behavior or crash. Use synchronization primitives like `sync.Mutex` or the concurrent `sync.Map` type for safe concurrent access.

### Running the code

## Full Code Example

```go
package main

import "fmt"

func main() {
	// Initialize a map with predefined key-value pairs
	m := map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 3,
	}

	// Retrieve a value
	appleCount := m["apple"]
	fmt.Println(appleCount) // Output: 2

	// add key-value to the map
	m["pineapple"] = 10

	// delete item
	delete(m, "apple")
	// confirmiv value is deleted
	value, exists := m["apple"]
	if exists {
		fmt.Println("apple count:", value)
	} else {
		fmt.Println("apple not found")
	}

	// Iterate over the map
	for fruit, count := range m {
		fmt.Printf("%s: %d\n", fruit, count)
	}
}
```

1. Save the code in a file named `maps.go`.
2. Open a terminal and navigate to the directory containing `maps.go`.
3. Run the code using the command:

   ```sh
   go run maps.go
   ```

4. The expected output will be:

```go
2
apple not found
banana: 5
orange: 3
pineapple: 10
```
