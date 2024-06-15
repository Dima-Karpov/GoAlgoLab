package main

// возвращает индекс и значение максимального элемента
func findMaxInSlice(slice []int) (int, int) {
	var index, val int
	for k, v := range slice {
		if v > val {
			index, val = k, v
		}
	}

	return index, val
}

// возвращает только индекс
func findInSlice(slice []int, value int) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			return i
		}
	}
	return -1
}
