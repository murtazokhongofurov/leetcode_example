package main

import (
	"fmt"
)

func main() {
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	s, e := 0, prices[0]
	for i := 0; i < len(prices); i++ {
		s = Max(s, prices[i]-e)
		e = Min(e, prices[i])
	}
	return s
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
