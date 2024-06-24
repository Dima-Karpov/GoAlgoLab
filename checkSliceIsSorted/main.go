package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func checkSliceIsSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			return false
		}
	}
	return true
}

func main() {
	array1 := []int{0, 1, 2, 3, 4, 5}
	array2 := []int{9, 7, 4, 1, 3, 5}
	array3 := []int{0}
	array4 := []int{}
	array5 := []int{1, 1}
	array6 := []int{3, 2, 1}
	array7 := []int{5, 15, 2, 13, 7, 16, 10, 2}
	array8 := []int{1, 9, 7, 4, 6, 2, 1, 13, 22, -3, 12, 76}

	arrayOfArrays := [][]int{
		array1,
		array2,
		array3,
		array4,
		array5,
		array6,
		array7,
		array8,
	}

	for i := 0; i < len(arrayOfArrays); i++ {
		fmt.Println("перед сортировкой", checkSliceIsSorted(arrayOfArrays[i]))
	}
	for i := 0; i < len(arrayOfArrays); i++ {
		arrayOfArrays[i] = mergeSort(arrayOfArrays[i])

		fmt.Println(mergeSort(arrayOfArrays[i]))
	}
	for i := 0; i < len(arrayOfArrays); i++ {
		fmt.Println("после сортировки", checkSliceIsSorted(arrayOfArrays[i]))
	}

}
