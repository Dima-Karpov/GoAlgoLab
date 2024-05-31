package main

import "fmt"

// ErrorWrongListIndex - ошибка при передачи несущесвующего в списке индекса
var ErrorWrongListIndex = fmt.Errorf("Wrong list index")

// IntNode - описание типа Узел списка
type IntNode struct {
	Value int
	Next  *IntNode
}

// NewIntNode - создание нового узал списка
func NewIntNode(value int) *IntNode {
	return &IntNode{value, nil}
}

// IntList - описание типа Список целых чисел
type IntList struct {
	size int
	Head *IntNode
}

// NewIntList - создание нового списка целых чисел
func NewIntList() *IntList {
	return &IntList{0, nil}
}

// Size - получения размера списка
func (l *IntList) Size() int {
	return l.size
}

// Get - обновление произвольного элемента списка
func (l *IntList) Get(index int) (*IntNode, error) {
	if index < 0 || index >= l.size {
		return nil, ErrorWrongListIndex
	}

	node := l.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}

	return node, nil
}

// Set - обновление произвольного элемента списка
func (l *IntList) Set(el int, index int) error {
	if index < 0 || index >= l.size {
		return ErrorWrongListIndex
	}

	node, err := l.Get(index)
	if err != nil {
		return err
	}
	node.Value = el

	return nil
}

// Add - добавление нового элемента в начало строки
func (l *IntList) Add(el int) {
	newNode := NewIntNode(el)
	if l.Head == nil {
		l.Head = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
	l.size++
}

// Insert - вставка нового эдемента в произвольную позицию
func (l *IntList) Insert(el int, index int) error {
	if index < 0 || index >= l.Size() {
		return ErrorWrongListIndex
	}

	newNode := NewIntNode(el)
	if index == 0 {
		l.Add(el)

		return nil
	}
	node, err := l.Get(index - 1)
	if err != nil {
		return err
	}
	newNode.Next = node.Next
	node.Next = newNode
	l.size++

	return nil
}

// Remove - удаление элемента из произвольной позиции
func (l *IntList) Remove(index int) error {
	if index < 0 || index >= l.size {
		return ErrorWrongListIndex
	}

	node, err := l.Get(index - 1)
	if err != nil {
		return err
	}
	node.Next = node.Next.Next
	l.size--

	return nil
}

// Print - печать списка
func (l *IntList) Print() {
	node := l.Head
	if node != nil {
		for node != nil {
			fmt.Printf("%d\t", node.Value)
			node = node.Next
		}
		fmt.Printf("\n")
	} else {
		fmt.Println("List is empty.")
	}
}

func main() {
	list := NewIntList()
	fmt.Println("Create list")
	list.Print()

	// добавление
	fmt.Println("Add values in list")
	list.Add(2)
	list.Add(1)
	list.Add(0)
	list.Print()

	// вставка нового эдемента в произвольную позицию
	fmt.Println("Inserting a new element at an arbitrary position")
	list.Insert(0, 1)
	list.Print()

	// удаление элемента из произвольной позиции
	fmt.Println("Removing an element from an arbitrary position")
	list.Remove(2)
	list.Print()

	// получения размера списка
	fmt.Printf("List size: %d\n", list.Size())

}
