package main

import "fmt"

func rotate(matrix [][]int) {
	fmt.Println("1-> ", matrix)
	for i, k := 0, len(matrix)-1; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[k][j] = matrix[k][j], matrix[i][j]
		}
		k--
	}
	fmt.Println("2-> ", matrix)
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	fmt.Println("3-> ", matrix)
}

func main() {
	var matrix = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
}