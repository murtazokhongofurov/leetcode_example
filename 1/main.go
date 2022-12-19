package main

import "fmt"

func twoSum(nums []int, target int) []int {
	l := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				l = append(l, i)
				l = append(l, j)
			}
		}
	}
	return l
}
func main() {
	t := 9
	arr := []int{2, 7, 11, 15}
	fmt.Println(twoSum(arr, t))
}
