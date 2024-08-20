package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	var length = len(scores)

	sort(scores, length)

	for i := 0; i < length; i++ {
		fmt.Printf("%d,", scores[i])
	}
}

func sort(arr []int, length int) {
	var minIndex int // Save the index of the selected minimum

	for i := 0; i < length; i++ {
		minIndex = i
		// Save the minimum value of each loop as the first element
		var minValue = arr[minIndex]
		for j := i; j < length-1; j++ {
			if minValue > arr[j+1] {
				minValue = arr[j+1]
				minIndex = j + 1
			}
		}

		// if minIndex changed, current minimum is exchanged with minIndex
		if i != minIndex {
			var temp = arr[i]
			arr[i] = arr[minIndex]
			arr[minIndex] = temp
		}
	}
}
