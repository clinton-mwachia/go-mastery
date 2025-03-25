# Search Algorithms

In Go, implementing search algorithms like Linear Search and Binary Search can be achieved with straightforward code. Below are examples of both, along with explanations of their functionality.

### Linear Search

Linear Search, also known as sequential search, is a simple method that checks each element in a slice sequentially until the target element is found or the slice ends. It's suitable for unsorted or small datasets.

**Implementation:**

```go
package main

import "fmt"

// LinearSearch function searches for a target value in a slice.
// It returns the index of the target and a boolean indicating if the target was found.
func LinearSearch(slice []int, target int) (int, bool) {
    for i, v := range slice {
        if v == target {
            return i, true
        }
    }
    return -1, false
}

func main() {
    data := []int{10, 20, 30, 40, 50}
    target := 30

    index, found := LinearSearch(data, target)
    if found {
        fmt.Printf("Found %d at index %d\n", target, index)
    } else {
        fmt.Printf("%d not found in the slice\n", target)
    }
}
```

**Explanation:**

- The `LinearSearch` function iterates over each element in the `slice`.
- If it finds the `target`, it returns the index and `true`.
- If the loop completes without finding the `target`, it returns `-1` and `false`.

**Usage:**

In the `main` function, we define a slice `data` and a `target` value. We call `LinearSearch` with these parameters and print the result based on whether the target was found.
