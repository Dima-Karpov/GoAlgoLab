package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

var root *Node = nil

func createNewNode(newData int) *Node {
	var newNode *Node = new(Node)
	newNode.data = newData
	newNode.left = nil
	newNode.right = nil

	return newNode
}

// In-order traversal binary search tree
func inOrder(root *Node) {
	if root == nil {
		return
	}
	inOrder(root.left) // Traversing the left subtree
	fmt.Printf("%d ", root.data)
	inOrder(root.right) // Traversing the right subtree
}

func insert(node *Node, newData int) {
	if root == nil {
		root = &Node{
			data:  newData,
			left:  nil,
			right: nil,
		}
		return
	}

	var compareValue = newData - node.data
	// Recursive left subtree, continue to find the insertion position
	if compareValue < 0 {
		if node.left == nil {
			node.left = createNewNode(newData)
		} else {
			insert(node.left, newData)
		}
	} else if compareValue > 0 {
		// Recursive right subtree, continue to find the insertion position
		if node.right == nil {
			node.right = createNewNode(newData)
		} else {
			insert(node.right, newData)
		}
	}
}

func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.data)
	preOrder(root.left)
	preOrder(root.right)
}

func postOrder(root *Node) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Printf("%d ", root.data)
}

func searchMinValue(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.left == nil {
		return node
	}
	return searchMinValue(node.left)
}

func searchMaxValue(node *Node) *Node {
	if node == nil || node.data == 0 {
		return nil
	}
	if node.right == nil {
		return node
	}
	return searchMaxValue(node.right)
}

func removeNode(node *Node, newData int) *Node {
	if node == nil {
		return node
	}
	var compareValue = newData - node.data
	if compareValue > 0 {
		node.right = removeNode(node.right, newData)
	} else if compareValue < 0 {
		node.left = removeNode(node.left, newData)
	} else if node.left != nil && node.right != nil {
		// Find the minimum node of the right subtree to replace the current node
		node.data = searchMinValue(node.right).data
		node.right = removeNode(node.right, node.data)
	} else {
		if node.left != nil {
			node = node.left
		} else {
			node = node.right
		}
	}

	return node
}

func main() {
	insert(root, 60)
	insert(root, 40)
	insert(root, 20)
	insert(root, 10)
	insert(root, 30)
	insert(root, 50)
	insert(root, 80)
	insert(root, 70)
	insert(root, 90)

	fmt.Printf("\nIn-order traversal binary search tree \n")
	inOrder(root)

	fmt.Printf("\nPre-order traversal binary search tree \n")
	preOrder(root)

	fmt.Printf("\nPost-order traversal binary search tree \n")
	postOrder(root)

	fmt.Printf("\nMin value \n")
	var minNode = searchMinValue(root)
	fmt.Printf("%d\n", minNode.data)

	fmt.Printf("\nMax value \n")
	var maxNode = searchMaxValue(root)
	fmt.Printf("%d\n", maxNode.data)

	fmt.Printf("\ndelete node is: 10\n")
	removeNode(root, 10)

	fmt.Printf("\nIn-order traversal binary search tree \n")
	inOrder(root)

	fmt.Printf("\ndelete node is: 20\n")
	removeNode(root, 20)

	fmt.Printf("\nIn-order traversal binary search tree \n")
	inOrder(root)

	fmt.Printf("\ndelete node is: 40\n")
	removeNode(root, 40)

	fmt.Printf("\nIn-order traversal binary search tree \n")
	inOrder(root)

}
