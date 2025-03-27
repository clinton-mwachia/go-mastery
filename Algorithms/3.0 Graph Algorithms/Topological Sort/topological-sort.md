## Topological Sort

Topological sorting is a fundamental algorithm used to linearly order the vertices of a Directed Acyclic Graph (DAG) such that for every directed edge `u → v`, vertex `u` precedes vertex `v` in the ordering. This is particularly useful in scenarios like task scheduling, resolving symbol dependencies in linkers, and determining the compilation order in build systems.

**Implementing Topological Sort in Go**

In Go, topological sorting can be implemented using Depth-First Search (DFS). Below is a comprehensive example demonstrating this approach:

```go
package main

import (
	"fmt"
)

// Graph represents a directed graph using an adjacency list.
type Graph struct {
	vertices map[string][]string
}

// NewGraph creates a new Graph instance.
func NewGraph() *Graph {
	return &Graph{vertices: make(map[string][]string)}
}

// AddEdge adds a directed edge from vertex u to vertex v.
func (g *Graph) AddEdge(u, v string) {
	g.vertices[u] = append(g.vertices[u], v)
}

// TopologicalSort performs a topological sort on the graph.
func (g *Graph) TopologicalSort() ([]string, error) {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)
	var result []string

	// Helper function for DFS
	var dfs func(v string) error
	dfs = func(v string) error {
		if recStack[v] {
			return fmt.Errorf("cycle detected: %s", v)
		}
		if visited[v] {
			return nil
		}
		visited[v] = true
		recStack[v] = true
		for _, neighbor := range g.vertices[v] {
			if err := dfs(neighbor); err != nil {
				return err
			}
		}
		recStack[v] = false
		result = append([]string{v}, result...)
		return nil
	}

	// Perform DFS from each vertex
	for vertex := range g.vertices {
		if !visited[vertex] {
			if err := dfs(vertex); err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}

func main() {
	g := NewGraph()
	g.AddEdge("A", "C")
	g.AddEdge("B", "C")
	g.AddEdge("C", "D")
	g.AddEdge("D", "E")

	order, err := g.TopologicalSort()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Topological Sort Order:", order)
}
```

**Explanation:**

- **Graph Representation:** The `Graph` struct uses a map to represent the adjacency list of the graph, where each key is a vertex, and the corresponding value is a slice of adjacent vertices.

- **AddEdge Method:** Adds a directed edge from vertex `u` to vertex `v` by appending `v` to the adjacency list of `u`.

- **TopologicalSort Method:** Performs the topological sort using a DFS approach:

  - Utilizes two maps, `visited` to track visited vertices and `recStack` to detect cycles.
  - Defines a recursive `dfs` function that explores each vertex's neighbors. If a cycle is detected (i.e., a vertex is revisited in the current recursion stack), an error is returned.
  - After visiting all neighbors of a vertex, it prepends the vertex to the `result` slice, ensuring the correct topological order.

- **Cycle Detection:** The algorithm detects cycles by maintaining a recursion stack (`recStack`). If a vertex is encountered that is already in the recursion stack, it indicates a cycle, and an error is returned.

- **Usage:** In the `main` function, a graph is created, edges are added, and the topological sort is performed. The resulting order is then printed.

**Output:**

```
Topological Sort Order: [A B C D E]
```
