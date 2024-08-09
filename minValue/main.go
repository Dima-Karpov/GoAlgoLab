package main

import "fmt"

func min(arr []int, length int) int {
	var minIndex = 0
	for j := 1; j < length; j++ {
		if arr[minIndex] > arr[j] {
			minIndex = j
		}
	}
	return arr[minIndex]
}

func main() {
	var scores = []int{60, 50, 95, 80, 70}
	var length = len(scores)
	var minValue = min(scores, length)
	fmt.Printf("Min Value = %d\n", minValue)
}
