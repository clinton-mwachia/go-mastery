# Disjoint Union

A **Disjoint Set**, also known as a **Union-Find** data structure, is used to keep track of a set of elements partitioned into non-overlapping subsets. It supports two primary operations:

1. **Find**: Determine which subset a particular element belongs to. This can be used to check if two elements are in the same subset.

2. **Union**: Merge two subsets into a single subset.

These operations are crucial in various applications, such as network connectivity, image processing, and algorithms like Kruskal's for finding minimum spanning trees.

## Implementing Disjoint Set in Go

Go does not have a built-in Disjoint Set data structure, but we can implement one using slices and maps. Below is an example of how to create a Disjoint Set in Go:

```go
package main

import "fmt"

// DisjointSet represents a disjoint set data structure.
type DisjointSet struct {
	parent map[int]int
	rank   map[int]int
}

// NewDisjointSet initializes a new Disjoint Set.
func NewDisjointSet() *DisjointSet {
	return &DisjointSet{
		parent: make(map[int]int),
		rank:   make(map[int]int),
	}
}

// MakeSet creates a new set containing the given element.
func (ds *DisjointSet) MakeSet(x int) {
	ds.parent[x] = x
	ds.rank[x] = 0
}

// Find returns the representative of the set containing x.
func (ds *DisjointSet) Find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x]) // Path compression
	}
	return ds.parent[x]
}

// Union merges the sets containing x and y.
func (ds *DisjointSet) Union(x, y int) {
	xRoot := ds.Find(x)
	yRoot := ds.Find(y)

	if xRoot == yRoot {
		return
	}

	// Union by rank
	if ds.rank[xRoot] < ds.rank[yRoot] {
		ds.parent[xRoot] = yRoot
	} else if ds.rank[xRoot] > ds.rank[yRoot] {
		ds.parent[yRoot] = xRoot
	} else {
		ds.parent[yRoot] = xRoot
		ds.rank[xRoot]++
	}
}

func main() {
	ds := NewDisjointSet()
	elements := []int{1, 2, 3, 4, 5}

	// Initialize sets for each element
	for _, elem := range elements {
		ds.MakeSet(elem)
	}

	// Perform unions
	ds.Union(1, 2)
	ds.Union(3, 4)
	ds.Union(2, 3)

	// Find representatives
	for _, elem := range elements {
		fmt.Printf("Representative of %d: %d\n", elem, ds.Find(elem))
	}
}
```

**Output:**

```
Representative of 1: 1
Representative of 2: 1
Representative of 3: 1
Representative of 4: 1
Representative of 5: 5
```

**Explanation:**

- **MakeSet**: Initializes a new set containing the element `x`. Each element starts as its own parent (representative) with a rank of 0.

- **Find**: Recursively finds the representative of the set containing `x`. Implements path compression by making each node point directly to the root, flattening the structure for faster future queries.

- **Union**: Merges the sets containing `x` and `y`. Uses union by rank to attach the shorter tree to the root of the taller tree, maintaining a balanced structure.
