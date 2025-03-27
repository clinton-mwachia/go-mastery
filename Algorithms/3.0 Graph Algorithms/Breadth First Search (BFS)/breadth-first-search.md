# **2. Breadth First Search (BFS)**

_Description_: Breadth-First Search traverses a graph level by level, visiting all neighbors of a node before moving to the next level. It's commonly used for finding the shortest path in unweighted graphs.

_Implementation_:

```go
package main

import (
	"fmt"
)

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

// BFS performs a breadth-first search starting from the given vertex.
func (g *Graph) BFS(start string) {
	visited := make(map[string]bool)
	queue := []string{start}

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		if visited[vertex] {
			continue
		}
		fmt.Println(vertex)
		visited[vertex] = true
		for _, neighbor := range g.vertices[vertex] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
}

func main() {
	g := NewGraph()
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "E")

	g.BFS("A")
}
```

_Explanation_:

- The `Graph` struct and `AddEdge` method function similarly to the DFS example.
- The `BFS` method uses a queue to explore each level of the graph, printing each vertex as it's visited.
