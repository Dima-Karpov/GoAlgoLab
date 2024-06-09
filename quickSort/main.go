package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	ar = quickSort(ar)

	fmt.Println(ar)
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left := []int{}
	right := []int{}
	middle := []int{}

	for _, val := range arr {
		if val < pivot {
			left = append(left, val)
		} else if val == pivot {
			middle = append(middle, val)
		} else {
			right = append(right, val)
		}
	}
	return append(append(quickSort(left), middle...), quickSort(right)...)
}
