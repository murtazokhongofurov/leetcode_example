package main

import "fmt"

func intToRoman(num int) string {
	d := map[int]string{
		1000: "M", 900: "CM", 500: "D", 400: "CD", 100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I",
	}
	str := ""
	for key, value := range d {
		for num/key != 0 {
			str += value
			num -= key
		}
	}
	return str
}

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(intToRoman(num))
}
