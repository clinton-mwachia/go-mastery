package main

import "fmt"

// Graph represents a graph using an adjacency list
type Graph struct {
	adjacencyList map[string][]string
}

// NewGraph creates a new graph
func NewGraph() *Graph {
	return &Graph{adjacencyList: make(map[string][]string)}
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(vertex string) {
	if _, exists := g.adjacencyList[vertex]; !exists {
		g.adjacencyList[vertex] = []string{}
	}
}

// AddEdge adds an edge between two vertices
func (g *Graph) AddEdge(vertex1, vertex2 string) {
	g.AddVertex(vertex1)
	g.AddVertex(vertex2)
	g.adjacencyList[vertex1] = append(g.adjacencyList[vertex1], vertex2)
	g.adjacencyList[vertex2] = append(g.adjacencyList[vertex2], vertex1)
}

// Display prints the adjacency list of the graph
func (g *Graph) Display() {
	for vertex, neighbors := range g.adjacencyList {
		fmt.Printf("%s -> %v\n", vertex, neighbors)
	}
}

func main() {
	graph := NewGraph()

	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "D")
	graph.AddEdge("D", "E")

	graph.Display()
}
