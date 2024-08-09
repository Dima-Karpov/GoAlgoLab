package main

import "fmt"

func max(arr []int, length int) int {
	for i := 0; i < length-1; i++ {
		if arr[i] > arr[i+1] { // swap
			var temp = arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = temp
		}
	}
	var maxValue = arr[length-1]
	return maxValue
}

func main() {
	var scores = []int{60, 50, 95, 80, 70}
	var length = len(scores)
	var maxValue = max(scores, length)
	fmt.Printf("Max Value = %d\n", maxValue)
}
