package main

import (
	"fmt"
)

func pivotIndex(nums []int) int {
	l := len(nums) - 1
	if len(nums) == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	if nums[l]-nums[0] == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if nums[l]-nums[i] == nums[i-1] {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(arr))
}
