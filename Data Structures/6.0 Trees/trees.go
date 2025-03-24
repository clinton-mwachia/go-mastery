package main

import "fmt"

// Node represents a node in the binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert adds a new key to the BST
func (n *Node) Insert(key int) {
	if n == nil {
		return
	} else if key < n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	}
}

// Search checks if a key exists in the BST
func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}
	if key < n.Key {
		return n.Left.Search(key)
	} else if key > n.Key {
		return n.Right.Search(key)
	}
	return true
}

// InOrder traverses the BST in ascending order
func (n *Node) InOrder() {
	if n != nil {
		n.Left.InOrder()
		fmt.Printf("%d ", n.Key)
		n.Right.InOrder()
	}
}

func main() {
	root := &Node{Key: 10}
	keys := []int{5, 15, 3, 7, 12, 18}

	for _, key := range keys {
		root.Insert(key)
	}

	fmt.Println("In-order traversal:")
	root.InOrder() // Output: 3 5 7 10 12 15 18

	fmt.Println("\nSearch for key 7:", root.Search(7)) // Output: true
	fmt.Println("Search for key 9:", root.Search(9))   // Output: false
}
