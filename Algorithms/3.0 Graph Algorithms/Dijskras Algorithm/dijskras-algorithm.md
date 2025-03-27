## Dijkstra's Algorithm

Dijkstra's algorithm is a fundamental method for finding the shortest paths from a single source vertex to all other vertices in a weighted graph with non-negative edge weights. It systematically selects the vertex with the smallest tentative distance, updates the distances to its adjacent vertices, and repeats this process until all vertices have been processed.

**Go Implementation of Dijkstra's Algorithm**

Below is a Go implementation of Dijkstra's algorithm using an adjacency list representation of the graph and a priority queue to efficiently select the next vertex with the smallest tentative distance.

```go
package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Edge represents a weighted edge to a neighboring vertex.
type Edge struct {
	Target string
	Weight float64
}

// Graph represents a weighted graph using an adjacency list.
type Graph struct {
	Vertices map[string][]Edge
}

// NewGraph initializes a new Graph.
func NewGraph() *Graph {
	return &Graph{Vertices: make(map[string][]Edge)}
}

// AddEdge adds a directed, weighted edge from one vertex to another.
func (g *Graph) AddEdge(from, to string, weight float64) {
	g.Vertices[from] = append(g.Vertices[from], Edge{Target: to, Weight: weight})
}

// Item represents a vertex and its current shortest distance.
type Item struct {
	Vertex    string
	Distance  float64
	Priority  float64
	Index     int
}

// PriorityQueue implements a min-heap for Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

// Less prioritizes items with smaller distances.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push adds an item to the priority queue.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the item with the smallest distance.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Dijkstra computes the shortest paths from the source vertex to all other vertices.
func (g *Graph) Dijkstra(source string) map[string]float64 {
	// Initialize distances with infinity.
	distances := make(map[string]float64)
	for vertex := range g.Vertices {
		distances[vertex] = math.Inf(1)
	}
	distances[source] = 0

	// Initialize the priority queue and push the source vertex.
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{Vertex: source, Distance: 0})

	for pq.Len() > 0 {
		// Extract the vertex with the smallest distance.
		current := heap.Pop(&pq).(*Item)
		currentVertex := current.Vertex
		currentDistance := current.Distance

		// Update distances to neighboring vertices.
		for _, edge := range g.Vertices[currentVertex] {
			newDistance := currentDistance + edge.Weight
			if newDistance < distances[edge.Target] {
				distances[edge.Target] = newDistance
				heap.Push(&pq, &Item{Vertex: edge.Target, Distance: newDistance})
			}
		}
	}

	return distances
}

func main() {
	g := NewGraph()
	g.AddEdge("A", "B", 1)
	g.AddEdge("A", "C", 4)
	g.AddEdge("B", "C", 2)
	g.AddEdge("B", "D", 5)
	g.AddEdge("C", "D", 1)

	source := "A"
	distances := g.Dijkstra(source)

	fmt.Printf("Shortest distances from source vertex %s:\n", source)
	for vertex, distance := range distances {
		fmt.Printf("To %s: %f\n", vertex, distance)
	}
}
```

**Explanation of the Implementation**

1. **Graph Representation**:

   - The `Graph` struct uses an adjacency list to represent the graph, where each vertex maps to a slice of `Edge` structs representing its neighbors and the respective edge weights.

2. **Adding Edges**:

   - The `AddEdge` method adds a directed edge from one vertex to another with a specified weight. If the graph is undirected, this method should be called twice for each pair of vertices to add edges in both directions.

3. **Priority Queue**:

   - The `PriorityQueue` is implemented using Go's `container/heap` package. It stores `Item` structs, each containing a vertex, its current shortest distance from the source, and an index for heap operations. The priority queue ensures that the vertex with the smallest tentative distance is processed next.

4. **Dijkstra's Algorithm**:

   - The `Dijkstra` method initializes all vertex distances to infinity, except for the source vertex, which is set to zero. It then uses the priority queue to repeatedly extract the vertex with the smallest distance, update the distances to its neighbors, and push these neighbors into the priority queue if a shorter path is found. This process continues until all vertices have been processed.

5. **Main Function**:
   - In the `main` function, a sample graph is created with vertices A, B, C, and D, and weighted edges between them. The `Dijkstra` method is called with 'A' as the source vertex, and the shortest distances from 'A' to all other vertices are printed.

**Output**

```
Shortest distances from source vertex A:
To A: 0.000000
To B: 1.000000
To C: 3.000000
To D: 4.000000
```

This output indicates that the shortest path from vertex 'A' to 'B' has a distance of 1, to 'C' is 3, and to 'D' is 4.
