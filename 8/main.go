package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func myAtoi(s string) int {
	if s == "" || s == " " {
		return 0
	}
	if utf8.RuneCountInString(s) == 1 {
		r, _ := strconv.Atoi(s)
		return r
	}
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	c := []rune(s)
	if unicode.IsLetter(c[0]) || string(c[0]) == "+" && string(c[1]) == "-" || string(c[1]) == "+" {
		return 0
	}
	var str, res string
	for _, k := range s {
		if unicode.IsDigit(k) || string(k) == "-" || string(k) == "." {
			str += string(k)
		}
	}
	var (
		index int
	)
	if strings.Contains(str, ".") {
		index = strings.Index(str, ".")
		str = str[:index]
		if index == 0 {
			return 0
		}
	}
	str2 := strings.Split(str, "")
	if str2[len(str2)-1] == "-" {
		str2[len(str2)-1] = ""
	}
	for i := range str2 {
		res += str2[i]
	}
	res = strings.TrimSpace(res)
	son, _ := strconv.Atoi(res)
	if son <= -2147483648 {
		return -2147483648
	}
	if son >= 2147483647 {
		return 2147483647
	}
	return son
}

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.TrimSpace(str)
	fmt.Println(myAtoi(str))
}
