package main

import "fmt"

func reverse(arr []int, length int) {
	var middle = length / 2

	for i := 0; i <= middle; i++ {
		var temp = arr[i]
		arr[i] = arr[length-i-1]
		arr[length-i-1] = temp
	}
}

func main() {
	var scores = []int{50, 60, 70, 80, 90}
	var length = len(scores)

	reverse(scores, length)

	for i := 0; i < length; i++ {
		fmt.Printf("%d\t", scores[i])
	}
}
