# Search Algorithms

In Go, implementing search algorithms like Linear Search and Binary Search can be achieved with straightforward code. Below are examples of both, along with explanations of their functionality.

### Binary Search

Binary Search is an efficient algorithm for finding a target value within a sorted slice. It repeatedly divides the search interval in half, reducing the time complexity to O(log n).

**Implementation:**

```go
package main

import "fmt"

// BinarySearch function searches for a target value in a sorted slice.
// It returns the index of the target and a boolean indicating if the target was found.
func BinarySearch(slice []int, target int) (int, bool) {
    low, high := 0, len(slice)-1

    for low <= high {
        mid := low + (high-low)/2
        if slice[mid] == target {
            return mid, true
        }
        if slice[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1, false
}

func main() {
    data := []int{10, 20, 30, 40, 50} // Must be sorted
    target := 30

    index, found := BinarySearch(data, target)
    if found {
        fmt.Printf("Found %d at index %d\n", target, index)
    } else {
        fmt.Printf("%d not found in the slice\n", target)
    }
}
```

**Explanation:**

- The `BinarySearch` function initializes `low` and `high` pointers to the start and end of the slice.
- It calculates the `mid` index and compares the `target` with the middle element.
- If the middle element matches the `target`, it returns the index and `true`.
- If the `target` is greater, it adjusts the `low` pointer; if less, it adjusts the `high` pointer.
- This process continues until the `target` is found or the pointers converge.

**Usage:**

In the `main` function, we define a sorted slice `data` and a `target` value. We call `BinarySearch` and print the result based on whether the target was found.

**Note:** Binary Search requires the input slice to be sorted. Ensure the data is sorted before applying this algorithm.
