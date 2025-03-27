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
