package main

import "fmt"

func tribonacci(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	if n == 2 {
		return 1
	}
	var (
		a, b, c = 0, 1, 1
	)
	for i := 2; i < n; i++ {
		a, b, c = b, c, a+b+c
	}

	return c
}

func main() {
	var n int
	print("n -> ")
	fmt.Scan(&n)
	fmt.Println(tribonacci(n))

}
