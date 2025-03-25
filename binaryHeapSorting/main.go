package main

import "fmt"

// Adjustment heap
func adjustHeap(array []int, currentIndex int, maxLength int) {
	var noLeafValue = array[currentIndex] // Current non-leaf node

	//2 * currentIndex + 1 Current left subtree subscript
	for j := 2*currentIndex + 1; j <= maxLength; j = currentIndex*2 + 1 {
		if j < maxLength && array[j] < array[j+1] {
			j++ // j Large subscript
		}

		if noLeafValue >= array[j] {
			break
		}

		array[currentIndex] = array[j] // Move up to the parent node
		currentIndex = j
	}

	array[currentIndex] = noLeafValue // To put in the position
}

// Initialize the heap
func createHeap(array []int, length int) {
	// Build a heap, (length - 1) / 2 scan half of the nodes with child nodes
	for i := (length - 1) / 2; i >= 0; i-- {
		adjustHeap(array, i, length-1)
	}
}

func heapSort(array []int, length int) {
	for i := length - 1; i > 0; i-- {
		var temp = array[0]
		array[0] = array[i]
		array[i] = temp
		adjustHeap(array, 0, i-1)
	}
}

func main() {
	var scores = []int{10, 90, 20, 80, 30, 70, 40, 60, 50}
	var length = len(scores)

	fmt.Print("Before building a heap: \n")
	for i := 0; i < length; i++ {
		fmt.Printf("%d, ", scores[i])
	}
	fmt.Print("\n\n")

	fmt.Print("After building a heap: \n")
	createHeap(scores, length)
	for i := 0; i < length; i++ {
		fmt.Printf("%d, ", scores[i])
	}
	fmt.Print("\n\n")

	fmt.Print("After heap sorting: \n")
	heapSort(scores, length)
	for i := 0; i < length; i++ {
		fmt.Printf("%d, ", scores[i])
	}
}
