package main

import "fmt"

// ErrQueueOverFlow - ошибка переполнения очереди
var ErrQueueOverFlow = fmt.Errorf("Queue over flow")

// ErrQueueEmpty - ошибка при чтении из пустой очереди
var ErrQueueEmpty = fmt.Errorf("Queue is empty")

// IntQueue - структура очереди
type IntQueue struct {
	data      []int
	headIndex int
	tailIndex int
	size      int
}

// NewIntQueue - конструктор очереди
func NewINtQueue(size int) *IntQueue {
	return &IntQueue{make([]int, size, size), 0, 0, 0}
}

// Size - получения размеров очереди
func (q IntQueue) Size() int {
	return q.size
}

// MaxSize - получения максимального размера очереди
func (q IntQueue) MaxSize() int {
	return len(q.data)
}

// Tail - получения хвоста очереди
func (q IntQueue) Tail() (int, error) {
	if q.headIndex == q.tailIndex {
		return 0, ErrQueueEmpty
	}
	return q.data[q.tailIndex], nil
}

// Head - получения головы (начало) очереди
func (q IntQueue) Head() (int, error) {
	if q.headIndex == q.tailIndex {
		return 0, ErrQueueEmpty
	}
	return q.data[q.headIndex], nil
}

// Queue - добавление элемента в очередь
func (q *IntQueue) Queue(el int) error {
	if q.tailIndex == q.MaxSize() {
		return ErrQueueOverFlow
	}
	q.data[q.tailIndex] = el
	q.size++
	q.tailIndex++

	return nil
}

// Dequeue - извлечение елемента из очереди
func (q *IntQueue) Dequeue() (int, error) {
	if q.headIndex == q.tailIndex {
		return 0, ErrQueueEmpty
	}
	head := q.data[q.headIndex]
	for key := q.headIndex + 1; key < q.tailIndex; key++ {
		q.data[key-1] = q.data[key]
	}
	q.tailIndex--
	q.size--

	return head, nil
}

// Print - печать очереди
func (q IntQueue) Print() {
	if q.headIndex != q.tailIndex {
		for i := q.headIndex; i < q.tailIndex; i++ {
			fmt.Printf("%d\t", q.data[i])
		}
		fmt.Printf("\n")
	} else {
		fmt.Println("%s\n", ErrQueueEmpty.Error())
	}
}

func main() {
	fmt.Println("Создаем очередь")
	q := NewINtQueue(5)
	q.Print()

	fmt.Println("Получения хвоста очереди")
	_, err := q.Head()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Добавляем три элемента '0 1 2'")
	q.Queue(0)
	q.Queue(1)
	q.Queue(2)
	fmt.Println(q.Size())
	q.Print()

	fmt.Println("Добавляем еще два элемента '3 4'")
	q.Queue(3)
	q.Queue(4)
	fmt.Println("Добавляем и еще один элемент '5'")
	err = q.Queue(5)
	if err != nil {
		fmt.Println(err.Error())
	}
	q.Print()

	fmt.Println("Извлекаем первый елемент из очереди '0'")
	el, err := q.Dequeue()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Наш извлеченный елемент из очереди")
	fmt.Println(el)
	fmt.Println(q.Size())
	q.Print()

}
