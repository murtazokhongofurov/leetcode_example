package main

import (
	"fmt"
	"strings"
)

func firstUniqChar(s string) int {
	for i := 0; i < len(s); i++ {
		if strings.Count(s, string(s[i])) == 1 {
			return strings.Index(s, string(s[i]))
		}
	}
	return -1
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(firstUniqChar(s))
}
