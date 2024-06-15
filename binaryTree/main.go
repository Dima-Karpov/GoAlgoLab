package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Insert добавляет новое значение в дерево
func (t *TreeNode) Insert(value int) {
	if value <= t.Value {
		if t.Left == nil {
			t.Left = &TreeNode{Value: value}
		} else {
			t.Left.Insert(value)
		}
	} else {
		if t.Right == nil {
			t.Right = &TreeNode{Value: value}
		} else {
			t.Right.Insert(value)
		}
	}
}

// Search ищет значение в дереве
func (t *TreeNode) Search(value int) bool {
	if t == nil {
		return false
	} else if value == t.Value {
		return true
	} else if value < t.Value {
		return t.Left.Search(value)
	}

	return t.Right.Search(value)
}

// InOrder выполняет обход дерева в порядке
func (t *TreeNode) InOrder() {
	if t == nil {
		return
	}
	t.Left.InOrder()
	fmt.Println(t.Value)
	t.Right.InOrder()
}

// Delete удаление элемента из дерева
func (t *TreeNode) Delete(value int) *TreeNode {
	if t == nil {
		return nil
	}

	if value < t.Value {
		t.Left = t.Left.Delete(value)
	} else if value > t.Value {
		t.Right = t.Right.Delete(value)
	} else {
		if t.Left == nil && t.Right == nil {
			return nil
		} else if t.Left == nil {
			return t.Right
		} else if t.Right == nil {
			return t.Left
		}

		// У узла два потомка
		minRight := t.Right.MinValueNode()
		t.Value = minRight.Value
		t.Right = t.Right.Delete(minRight.Value)
	}
	return t
}

// MinValueNode находит узел с минимальным значением в дереве
func (t *TreeNode) MinValueNode() *TreeNode {
	current := t
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func main() {
	root := &TreeNode{Value: 10}

	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	fmt.Println("Search 7: ", root.Search(7))
	fmt.Println("Search 20: ", root.Search(20))

	fmt.Println("InOrder traversal: ")
	root.InOrder()

	root = root.Delete(7)
	root = root.Delete(10)

	fmt.Println("InOrder traversal after deletion: ")
	root.InOrder()
}
