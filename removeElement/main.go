package main

import "fmt"

func main() {
	res := removeElement([]int{3, 2, 2, 3}, 3)
	if res != 2 {
		fmt.Printf("неправильный результав %v\n", res)
		return
	} else {
		fmt.Printf("правильный результав %v\n", res)
	}

}

func removeElement(nums []int, val int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}
