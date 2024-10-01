package main

import "fmt"

type Node struct {
	data string
	prev *Node
	next *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func initial() {
	head.data = "San Francisco"
	head.prev = nil
	head.next = nil

	var nodeOakland *Node = &Node{data: "Oakland", prev: head, next: nil}

	head.next = nodeOakland

	var nodeBerkeley *Node = &Node{data: "Berkeley", prev: nodeOakland, next: nil}
	nodeOakland.next = nodeBerkeley

	tail.data = "Fremont"
	tail.prev = nodeBerkeley
	tail.next = nil
	nodeBerkeley.next = tail
}

func add(data string) {
	var newNode *Node = new(Node)
	newNode.data = data
	newNode.next = nil
	tail.next = newNode
	newNode.prev = tail
	tail = newNode
}

func insert(insertPosition int, data string) {
	var p = head
	var i = 0
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
	newNode.prev = p
	newNode.next.prev = newNode
}

func removeNode(position int) {
	var p = head
	var i = 0
	// Move the node to the previous node position that was deleted
	for {
		if p.next == nil || i >= position-1 {
			break
		}
		p = p.next
		i++
	}

	var temp = p.next    // Save the node you want to delete
	p.next = p.next.next // Previous node next points to next of delete the node
	p.next.prev = p
	temp.next = nil // Set the delete node next to null
	temp.prev = nil // Set the delete node prev to null
}

func output(node *Node) {
	var p = node
	var end *Node = nil
	for {
		if p == nil {
			break
		}
		fmt.Printf("%s -> ", p.data)
		end = p
		p = p.next
	}
	fmt.Printf("End\n")

	p = end
	for {
		if p == nil {
			break
		}

		fmt.Printf("%s -> ", p.data)
		p = p.prev
	}
	fmt.Printf("Start\n")
}

func main() {
	initial()

	output(head)
	fmt.Print("\n")

	fmt.Printf("Add a new node Walnut: \n")
	add("Walnut")
	output(head)
	fmt.Print("\n")

	fmt.Printf("Insert a new node Atlanto at index 2: \n")
	insert(2, "Atlanto")
	output(head)
	fmt.Print("\n")

	fmt.Printf("Delete node index 2: \n")
	removeNode(2)
	output(head)
	fmt.Print("\n")
}
