package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		if len(nums)%2 == 1 {
			return float64(nums[len(nums)/2]) / 1
		}
	}
	return float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / 2.0
}

func main() {
	var arr1, arr2 = []int{1, 3}, []int{2}
	fmt.Println(findMedianSortedArrays(arr1, arr2))

}
