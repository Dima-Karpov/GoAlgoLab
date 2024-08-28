package main

import "fmt"

func sort(arr []int, length int) {
	for i := 0; i < length; i++ {
		// Take unsorted new elements
		var insertElement = arr[i]
		// Inserted position
		var insertPosition = i
		for j := insertPosition - 1; j >= 0; j-- {
			// if the new element is smaller than the sorted element, it is shifted to the right
			if insertElement < arr[j] {
				arr[j+1] = arr[j]
				insertPosition--
			}
		}
		// Insert the new element
		arr[insertPosition] = insertElement
	}
}

func main() {
	// index start form 0
	var scores = []int{90, 70, 50, 80, 60, 85}
	var length = len(scores)

	sort(scores, length)

	for i := 0; i < length; i++ {
		fmt.Printf("%d\t", scores[i])
	}
}
