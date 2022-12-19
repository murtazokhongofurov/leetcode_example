package main

import "fmt"

func isPalindrome(x int) bool {
	n := ReverseNew(x)
	if n == x {
		return true
	}
	return false
}

func ReverseNew(num int) int {
	res := 0
	for num > 0 {
		res = res*10 + num%10
		num /= 10
	}
	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(isPalindrome(n))
}
