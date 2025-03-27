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
	Vertex   string
	Distance float64
	Priority float64
	Index    int
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
