package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	a1 := []rune(s1)
	a2 := []rune(s2)
	for i := 0; i < len(a2)-1; i++ {
		for j := 0; j < len(a1)-1; j++ {
			if a1[i] == a2[j] && a1[i+1] == a2[j+1] {
				return true
			}
		}
	}
	return false
}

func main() {
	var str1, str2 string
	print("str1-> ")
	fmt.Scan(&str1)
	print("str2-> ")
	fmt.Scan(&str2)
	fmt.Println(checkInclusion(str1, str2))
}
