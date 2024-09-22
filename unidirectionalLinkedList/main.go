package main

import "fmt"

type Node struct {
	data string
	next *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func initial() {
	head.data = "San Francisco"
	head.next = nil

	var nodeOakland *Node = &Node{data: "Oakland", next: nil}
	head.next = nodeOakland

	var nodeBerkeley *Node = &Node{data: "Berkeley", next: nil}
	nodeOakland.next = nodeBerkeley

	tail.data = "Fremont"
	tail.next = nil
	nodeBerkeley.next = tail
}

func insert(insertPosition int, data string) {
	var p = head
	var i = 0

	// Move the node to the insertion position
	for {
		if p.next == nil || i >= insertPosition-1 {
			break
		}
		p = p.next
		i++
	}
	var newNode *Node = new(Node)
	newNode.data = data
	newNode.next = p.next // newNode next point to next node
	p.next = newNode      // current next point to newNode
}

func removeNode(removePosition int) {
	var p = head
	var i = 0

	// Move the node to the previous node position that was deleted
	for {
		if p.next == nil || i >= removePosition-1 {
			break
		}
		p = p.next
		i++
	}
	var temp = p.next // Save the node you want to delete
	// Previous node next points to next of delete the node
	p.next = p.next.next
	temp.next = nil
}

func output(node *Node) {
	var p = node

	for {
		if p == nil {
			break
		}
		fmt.Printf("%s -> ", p.data)
		p = p.next
	}
	fmt.Printf("End\n\n")
}

func main() {
	initial()

	output(head)

	fmt.Printf("Insert a new node Walnut at index = 2: \n")
	insert(2, "Walnut")
	output(head)

	fmt.Printf("Delete a new node Walnut at index = 2:\n")
	removeNode(2)
	output(head)

}
