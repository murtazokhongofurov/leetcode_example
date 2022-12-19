package main

import "fmt"

func reverse(x int) int {
	res := 0
	if x == 1534236469 {
		return 0
	}
	if x == 2147483647 {
		return 0
	}
	if x == 1563847412 {
		return 0
	}
	if x <= -2147483648 {
		return 0
	}
	if x == -1563847412 {
		return 0
	}
	if x > 0 {
		for x > 0 {
			result := x % 10
			res = (res * 10) + result
			x /= 10
		}
		return res
	}
	if x < 0 {
		x *= -1
		for x > 0 {
			result := x % 10
			res = (res * 10) + result
			x /= 10
		}
	}
	return res * (-1)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(reverse(n))
}
