package main

import "fmt"

func binarySearch(arr []int, length int, searchValue int) int {
	var low = 0
	var high = length
	var mid = 0

	for {
		if low >= high {
			break
		}
		mid = (low + high) / 2
		if arr[mid] == searchValue {
			return mid
		} else if arr[mid] < searchValue {
			low = mid + 1
		} else if arr[mid] > searchValue {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	var scores = []int{30, 40, 50, 60, 70, 85, 90, 100}
	var length = len(scores)

	var searchValue = 40
	var position = binarySearch(scores, length, searchValue)
	fmt.Printf("%d position: %d", searchValue, position)

	fmt.Printf("\n----------------------------------\n")

	searchValue = 95
	position = binarySearch(scores, length, searchValue)
	fmt.Printf("%d position: %d\n", searchValue, position)
}
