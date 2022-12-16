package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	nums1 := []int{}
	for i := 0; i < len(nums); i++ {
		nums1 = append(nums1, nums[i]*nums[i])
	}
	return nums1
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(nums)
}
