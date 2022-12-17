package main

import "fmt"

func moveZeroes(nums []int) []int {
	var n int
	for _, v := range nums {
		if v != 0 {
			nums[n] = v
			n++
		}
	}
	for n < len(nums) {
		nums[n] = 0
		n++
	}
	return nums
}

func main() {
	nums := []int{0, 1, 0, 3, 0, 12}
	fmt.Println(moveZeroes(nums))
}
