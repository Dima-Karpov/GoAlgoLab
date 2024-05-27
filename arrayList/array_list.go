package main

import "fmt"

// получения размера списка
func Size(list []int) int {
	return len(list)
}

// добавления нового элемента в конец списка
func Add(list []int, el int) []int {
	return append(list, el)
}

// добавления нового элемента в произвольную позицию index
func Insert(list []int, el, index int) []int {
	list = append(list, el)
	for key := Size(list) - 1; key > index; key-- {
		list[key] = list[key-1]
	}
	list[index] = el

	return list
}

// удаления элемента в позиции индекса
func Remove(list []int, index int) []int {
	for key := index + 1; key < Size(list); key++ {
		list[key-1] = list[key]
	}

	return list[:Size(list)-1]
}

func main() {
	list := []int{1, 2, 3}
	fmt.Println("Size: ", Size(list))
	list = Add(list, 4)
	fmt.Println("Add: ", list)
	list = Insert(list, 0, 0)
	fmt.Println("Insert: ", list)
	list = Remove(list, 2)
	fmt.Println("Remove: ", list)
}
