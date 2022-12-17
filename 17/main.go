package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	str := []string{}
	if len(digits) == 0 {
		return str
	}
	maps := map[string]string{
		"1": "", "2": "abc", "3": "def",
		"4": "ghi", "5": "jkl", "6": "mno",
		"7": "pqrs", "8": "tuv", "9": "wxyz",
	}
	res := []string{""}
	for i := range digits {
		term := []string{}
		for k := range maps[string(digits[i])] {
			for s := range res {
				term = append(term, res[s]+string(maps[string(digits[i])][k]))
			}
		}
		res = term
	}
	return res

}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(letterCombinations(str))
}
