package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	for {
		if len(str1) < len(str2) {
			str1, str2 = str2, str1
		}

		l2 := len(str2)
		if len(str2) == 0 {
			return str1
		}
		for len(str1) >= l2 {
			if str1[:l2] != str2 {
				return ""
			}
			str1 = str1[l2:]
		}
	}
}

func main() {
	var str1, str2 = "ABCABC", "ABC"
	fmt.Println(gcdOfStrings(str1, str2))
}
