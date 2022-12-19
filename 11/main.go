package main

import "fmt"

func maxArea(height []int) int {
	s := 0
	i, j := 0, len(height)-1
	for i < j {
		if height[i] < height[j] {
			s = max(s, height[i]*(j-i))
			i++
		} else {
			s = max(s, height[j]*(j-i))
			j--
		}
	}
	return s
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(arr))
}
