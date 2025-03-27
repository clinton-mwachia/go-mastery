package main

import (
	"fmt"
	"sort"
)

// Edge represents a weighted connection between two vertices.
type Edge struct {
	Source, Destination, Weight int
}

// Graph represents a graph with vertices and edges.
type Graph struct {
	Vertices int
	Edges    []Edge
}

// UnionFind is a data structure to manage disjoint sets.
type UnionFind struct {
	Parent, Rank []int
}

// NewUnionFind initializes a UnionFind structure for 'n' elements.
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{Parent: parent, Rank: rank}
}

// Find returns the root parent of a given node 'i'.
func (uf *UnionFind) Find(i int) int {
	if uf.Parent[i] != i {
		uf.Parent[i] = uf.Find(uf.Parent[i]) // Path compression
	}
	return uf.Parent[i]
}

// Union merges two subsets 'x' and 'y' using union by rank.
func (uf *UnionFind) Union(x, y int) {
	xRoot := uf.Find(x)
	yRoot := uf.Find(y)
	if xRoot == yRoot {
		return
	}
	if uf.Rank[xRoot] < uf.Rank[yRoot] {
		uf.Parent[xRoot] = yRoot
	} else if uf.Rank[xRoot] > uf.Rank[yRoot] {
		uf.Parent[yRoot] = xRoot
	} else {
		uf.Parent[yRoot] = xRoot
		uf.Rank[xRoot]++
	}
}

// KruskalMST applies Kruskal's Algorithm to find the MST of the graph.
func (g *Graph) KruskalMST() []Edge {
	// Sort edges by ascending weight
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	uf := NewUnionFind(g.Vertices)
	mst := []Edge{}

	for _, edge := range g.Edges {
		x := uf.Find(edge.Source)
		y := uf.Find(edge.Destination)
		if x != y {
			mst = append(mst, edge)
			uf.Union(x, y)
		}
	}

	return mst
}

func main() {
	// Example usage
	graph := Graph{
		Vertices: 4,
		Edges: []Edge{
			{0, 1, 10},
			{0, 2, 6},
			{0, 3, 5},
			{1, 3, 15},
			{2, 3, 4},
		},
	}

	mst := graph.KruskalMST()
	fmt.Println("Edges in the Minimum Spanning Tree:")
	for _, edge := range mst {
		fmt.Printf("%d -- %d == %d\n", edge.Source, edge.Destination, edge.Weight)
	}
}
