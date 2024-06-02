package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ErrDequeEmpty - ошибка при чтении из пустой очереди
var ErrDequeEmpty = fmt.Errorf("Deque is empty")

type Node struct {
	value int
	next  *Node
	prev  *Node
}
type Deque struct {
	head *Node
	tail *Node
	size int
}

func NewDeque() *Deque {
	return &Deque{}
}

// IsEmpty - для проверки или список пуст
func (d *Deque) IsEmpty() bool {
	return d.size == 0
}

// Size - размер списка
func (d *Deque) Size() int {
	return d.size
}

// AddFront - добавления елемента в начала списка
func (d *Deque) AddFront(item int) {
	newNode := &Node{value: item}
	if d.IsEmpty() {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.size++
}

// AddRear - добавления елемента в конец списка
func (d *Deque) AddRear(item int) {
	newNode := &Node{value: item}
	if d.IsEmpty() {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.prev = d.tail
		d.tail.next = newNode
		d.tail = newNode
	}
	d.size++
}

// RemoveFront - удаления елемента c начала списка
func (d *Deque) RemoveFront() (int, error) {
	if d.IsEmpty() {
		return 0, ErrDequeEmpty
	}
	frontValue := d.head.value
	if d.size == 1 {
		d.head = nil
		d.tail = nil
	} else {
		d.head = d.head.next
		d.head.prev = nil
	}
	d.size--

	return frontValue, nil
}

// RemoveRear - удаления елемента c конца списка
func (d *Deque) RemoveRear() (int, error) {
	if d.IsEmpty() {
		return 0, ErrDequeEmpty
	}
	frontValue := d.tail.value
	if d.size == 1 {
		d.head = nil
		d.tail = nil
	} else {
		d.tail = d.tail.prev
		d.tail.next = nil
	}
	d.size--

	return frontValue, nil
}

// Remove - удаления заказа из очереди
func (d *Deque) Remove(item int) bool {
	curr := d.head
	for curr != nil {
		if curr.value == item {
			if curr.prev != nil {
				curr.prev.next = curr.next
			} else {
				d.head = curr.next
			}
			if curr.next != nil {
				curr.next.prev = curr.prev
			} else {
				d.tail = curr.prev
			}
			d.size--
			return true
		}
		curr = curr.next
	}

	return false
}

// Print - вывод списка
func (d *Deque) Print() {
	if d.IsEmpty() {
		fmt.Println(ErrDequeEmpty)
	} else {
		elements := make([]int, 0, d.size)
		curr := d.head
		for curr != nil {
			elements = append(elements, curr.value)
			curr = curr.next
		}
		fmt.Println(elements)
	}
}

func praseInput(input string) []int {
	values := strings.Fields(input)
	orders := make([]int, len(values))
	for i, v := range values {
		val, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Ошибка преобразования значенияЖ %s\n", v)
			os.Exit(1)
		}
		orders[i] = val
	}

	return orders
}

func main() {
	fmt.Println("создаем список")
	deque := NewDeque()

	fmt.Println("Введите номера заказов через пробел (например 1 2 3): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	orders := praseInput(input)
	for _, order := range orders {
		deque.AddRear(order)
	}

	fmt.Println("добавленные заказы: ")
	deque.Print()

	//Удаления одно заказа
	fmt.Println("Введите номер заказа для удаления: ")
	input, _ = reader.ReadString('\n')
	orderToRemove, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Некорректный номер заказа")
		return
	}
	deleted := deque.Remove(orderToRemove)
	if deleted {
		fmt.Println("Заказ успешно удален")
		fmt.Println("Оставшиеся заказы: ")
		deque.Print()
	} else {
		fmt.Println("Заказ не найден")
	}

}
