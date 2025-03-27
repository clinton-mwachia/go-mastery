package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Node represents a point in the grid.
type Node struct {
	x, y    int
	g, h, f float64
	parent  *Node
	index   int // index in the priority queue
}

// PriorityQueue implements heap.Interface and holds Nodes.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].f < pq[j].f
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}

// heuristic estimates the cost from current node to the goal.
func heuristic(a, b *Node) float64 {
	return math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y))
}

// neighbors returns the adjacent nodes of the current node.
func neighbors(node *Node, grid [][]int) []*Node {
	dirs := []struct{ dx, dy int }{
		{0, -1}, {-1, 0}, {1, 0}, {0, 1},
	}
	var result []*Node
	for _, dir := range dirs {
		x, y := node.x+dir.dx, node.y+dir.dy
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 0 {
			result = append(result, &Node{x: x, y: y})
		}
	}
	return result
}

// reconstructPath builds the path from start to goal.
func reconstructPath(node *Node) []*Node {
	var path []*Node
	for current := node; current != nil; current = current.parent {
		path = append([]*Node{current}, path...)
	}
	return path
}

// aStar finds the shortest path between start and goal nodes in a grid.
func aStar(start, goal *Node, grid [][]int) []*Node {
	openSet := &PriorityQueue{}
	heap.Init(openSet)
	heap.Push(openSet, start)

	cameFrom := make(map[Node]*Node)

	gScore := make(map[Node]float64)
	gScore[*start] = 0

	fScore := make(map[Node]float64)
	fScore[*start] = heuristic(start, goal)

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*Node)

		if current.x == goal.x && current.y == goal.y {
			return reconstructPath(current)
		}

		for _, neighbor := range neighbors(current, grid) {
			tentativeGScore := gScore[*current] + 1

			if g, exists := gScore[*neighbor]; !exists || tentativeGScore < g {
				cameFrom[*neighbor] = current
				gScore[*neighbor] = tentativeGScore
				fScore[*neighbor] = tentativeGScore + heuristic(neighbor, goal)
				heap.Push(openSet, neighbor)
			}
		}
	}

	return nil // no path found
}

func main() {
	// Define a simple grid: 0 = walkable, 1 = obstacle
	grid := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}

	start := &Node{x: 0, y: 0}
	goal := &Node{x: 4, y: 4}

	path := aStar(start, goal, grid)
	if path != nil {
		fmt.Println("Path found:")
		for _, node := range path {
			fmt.Printf("(%d, %d) -> ", node.x, node.y)
		}
		fmt.Println("Goal")
	} else {
		fmt.Println("No path found.")
	}
}
