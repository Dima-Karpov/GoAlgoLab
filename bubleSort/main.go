package main

import (
	"fmt"
	"math/rand"
	"time"
)

// init - необходимо для того, чтобы рандом был похож на рандомный
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	arr := make([]int, 50)
	for i := range arr {
		arr[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100, 100]
	}

	bubbleSort(arr)
	fmt.Println(arr)
	bubbleSortRecursive(arr)
	fmt.Println(arr)
	bubbleSortReversed(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	n := len(arr)
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
		n--

		if !swapped {
			break
		}
	}
}
func bubbleSortReversed(arr []int) {
	n := len(arr)
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] < arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
		n--

		if !swapped {
			break
		}
	}
}
func bubbleSortRecursive(arr []int) {
	n := len(arr)
	if n == 1 {
		return
	}

	for i := 1; i < n; i++ {
		if arr[i-1] > arr[i] {
			arr[i-1], arr[i] = arr[i], arr[i-1]
		}
	}

	bubbleSortRecursive(arr[:n-1])

}
