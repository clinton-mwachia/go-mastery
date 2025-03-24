## Trees

A **tree** is a hierarchical data structure consisting of nodes, where each node has a value and references to its child nodes. A common type of tree is the **Binary Search Tree (BST)**, where each node has at most two children, and for any given node, all values in its left subtree are less than the node's value, and all values in its right subtree are greater.

**1. Define the Node Structure**

Each node in the BST will contain an integer key and pointers to its left and right child nodes.

```go
package main

import "fmt"

// Node represents a node in the binary search tree
type Node struct {
    Key   int
    Left  *Node
    Right *Node
}
```

**2. Implement the Insert Function**

The `Insert` function adds a new key to the BST while maintaining its properties.

```go
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
```

**3. Implement the Search Function**

The `Search` function checks whether a given key exists in the BST.

```go
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
```

**4. Implement the InOrder Traversal Function**

In-order traversal visits nodes in ascending order of their keys.

```go
// InOrder traverses the BST in ascending order
func (n *Node) InOrder() {
    if n != nil {
        n.Left.InOrder()
        fmt.Printf("%d ", n.Key)
        n.Right.InOrder()
    }
}
```

**5. Example Usage**

Here's how you can create a BST, insert keys, search for a key, and perform in-order traversal:

```go
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
```

**Considerations:**

- **Balancing:** The provided implementation does not balance the tree, which can lead to inefficient operations if the tree becomes skewed. For balanced trees, consider using data structures like AVL trees or Red-Black trees.

- **Deletion:** Implementing a deletion operation requires handling three cases: deleting a leaf node, deleting a node with one child, and deleting a node with two children.

- **Generics:** Go 1.18 introduced support for generics, allowing for more flexible and type-safe implementations of data structures like BSTs.
