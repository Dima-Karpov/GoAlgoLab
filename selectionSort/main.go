package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100
	}

	selectionSort(ar)
	fmt.Println(ar)
	selectionSortByMax(ar)
	fmt.Println(ar)
	bidirectionalSelectionSort(ar)
	fmt.Println(ar)
}

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}

	return arr
}

func selectionSortByMax(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		maxIndex := i
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[maxIndex] {
				maxIndex = j
			}
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}
	}

	return arr
}

func bidirectionalSelectionSort(arr []int) []int {
	left := 0
	right := len(arr) - 1

	for left < right {
		minIndex := left
		maxIndex := right

		for i := left; i <= right; i++ {
			if arr[i] < arr[minIndex] {
				minIndex = i
			}
			if arr[i] > arr[maxIndex] {
				maxIndex = i
			}
		}
		if minIndex != left {
			arr[left], arr[minIndex] = arr[minIndex], arr[left]
		}
		if maxIndex == left {
			maxIndex = minIndex
		}
		if maxIndex != right {
			arr[right], arr[maxIndex] = arr[maxIndex], arr[right]
		}

		left++
		right--
	}

	return arr
}
