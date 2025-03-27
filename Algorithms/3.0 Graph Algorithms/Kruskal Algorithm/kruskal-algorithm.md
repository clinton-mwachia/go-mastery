## Kruskal's Algorithm

Kruskal's Algorithm is a fundamental approach in graph theory for determining the Minimum Spanning Tree (MST) of a weighted, undirected graph. The MST is a subset of the graph's edges that connects all vertices with the minimal possible total edge weight, ensuring no cycles are formed. Kruskal's Algorithm achieves this by following a greedy strategy: it repeatedly selects the shortest edge that doesn't form a cycle until all vertices are connected.

**Algorithm Steps:**

1. **Sort All Edges**: Arrange the graph's edges in ascending order based on their weights.

2. **Initialize Disjoint Sets**: Create a disjoint-set (union-find) data structure to keep track of which vertices are connected. Initially, each vertex is in its own set.

3. **Construct the MST**:
   - Iterate through the sorted edges.
   - For each edge, determine the root parents of the vertices it connects using the union-find structure.
   - If the vertices are in different sets (i.e., adding the edge won't form a cycle), include this edge in the MST and merge the sets.
   - Continue this process until the MST contains exactly \( V-1 \) edges, where \( V \) is the number of vertices in the graph.

**Go Implementation:**

Below is a Go implementation of Kruskal's Algorithm:

```go
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
```

**Explanation:**

- **Edge and Graph Structures**: The `Edge` struct represents a connection between two vertices with a specific weight. The `Graph` struct contains the total number of vertices and a slice of `Edge` instances.

- **Union-Find Data Structure**: The `UnionFind` struct manages disjoint sets, supporting efficient union and find operations. It employs path compression in the `Find` method and union by rank in the `Union` method to optimize performance.

- **Kruskal's Algorithm Implementation**:

  - Edges are sorted in ascending order based on their weights.
  - The algorithm iterates through the sorted edges, adding each edge to the MST if it doesn't form a cycle, as determined by the union-find structure.
  - This process continues until the MST contains \( V-1 \) edges.

- **Main Function**: An example graph is defined, and Kruskal's Algorithm is applied to find and print the edges of the MST.

**Performance Considerations:**

- **Time Complexity**: The algorithm's time complexity is \( O(E \log E) \), where \( E \) is the number of edges. This complexity arises from the initial sorting of the edges and the subsequent union-find operations.

- **Space Complexity**: The space complexity is \( O(V + E) \), accounting for the storage of the graph's edges and the union-find data structures.
