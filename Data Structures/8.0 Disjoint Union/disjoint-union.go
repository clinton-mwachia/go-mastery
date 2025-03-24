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
