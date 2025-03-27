## 3. Using Go's Standard Library (`sort` Package)

Go's standard library provides the `sort` package, which offers efficient sorting functions for slices of built-in types.

**Implementation:**

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sorting integers
	ints := []int{5, 3, 4, 1, 2}
	sort.Ints(ints)
	fmt.Println("Sorted integers:", ints)

	// Sorting strings
	strs := []string{"banana", "apple", "cherry"}
	sort.Strings(strs)
	fmt.Println("Sorted strings:", strs)
}
```

**Explanation:**

- The `sort.Ints` function sorts a slice of integers in ascending order.
- The `sort.Strings` function sorts a slice of strings in lexicographical order.

**Usage:**

In the `main` function, we define slices of integers and strings. We use `sort.Ints` and `sort.Strings` to sort these slices and then print the sorted results.

For more advanced sorting, such as sorting a slice of custom structs, you can implement the `sort.Interface` by defining methods `Len()`, `Less(i, j int) bool`, and `Swap(i, j int)`. This allows for sorting based on custom criteria.
