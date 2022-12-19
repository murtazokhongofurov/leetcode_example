package main

import "fmt"

func longestPalindrome(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		for j := i; j <= i+1; j++ {
			p, q := i, j
			if q >= len(s) {
				continue
			}

			for p >= 0 && q < len(s) {
				if s[p] == s[q] {
					p--
					q++
				} else {
					break
				}
			}

			if len(result) < q-p-1 {
				result = s[p+1 : q]
			}
		}
	}
	return result
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(longestPalindrome(str))
}
