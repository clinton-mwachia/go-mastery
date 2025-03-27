# **1. Depth-First Search (DFS)**

_Description_: Depth-First Search explores a graph by traversing as far along a branch as possible before backtracking. It's useful for tasks like detecting cycles and finding connected components.

_Implementation_:

```go
package main

import "fmt"

// Graph represents an adjacency list graph.
type Graph struct {
	vertices map[string][]string
}

// NewGraph creates a new Graph.
func NewGraph() *Graph {
	return &Graph{vertices: make(map[string][]string)}
}

// AddEdge adds an edge to the graph.
func (g *Graph) AddEdge(from, to string) {
	g.vertices[from] = append(g.vertices[from], to)
}

// DFS performs a depth-first search starting from the given vertex.
func (g *Graph) DFS(start string, visited map[string]bool) {
	if visited[start] {
		return
	}
	visited[start] = true
	fmt.Println(start)
	for _, neighbor := range g.vertices[start] {
		g.DFS(neighbor, visited)
	}
}

func main() {
	g := NewGraph()
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "E")

	visited := make(map[string]bool)
	g.DFS("A", visited)
}
```

_Explanation_:

- The `Graph` struct represents a graph using an adjacency list.
- The `AddEdge` method adds a directed edge from one vertex to another.
- The `DFS` method recursively visits each vertex, marking it as visited and printing its value.
