package main

import (
	"fmt"
	"math"
)

// Input: nums = [-1,2,1,-4], target = 1
// Output: 2
// Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

func threeSumClosest(nums []int, target int) int {
	curr := nums[0]
	for i := 0; i < len(nums)-2; i++ {
		for j := i+1; j < len(nums)-1; j++ {
			for k := 0; k < len(nums); k++ {
				if math.Abs(float64(target)-(float64(nums[i])+float64(nums[j])+float64(nums[k]))) < float64(target)-float64(curr) {
					curr = nums[i]+nums[j] + nums[k]
				}
			}
		}
	}
	return curr
}

func main() {
	arr := []int{1,2,0}
	var target int
	fmt.Scan(&target)
	fmt.Println(threeSumClosest(arr, target))
}
