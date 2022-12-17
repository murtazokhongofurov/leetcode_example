package main

import "fmt"

// func subsets(nums []int) [][]int {

// }

// func main() {
// 	var nums = []int{1, 2, 3}
// 	fmt.Println(subsets(nums))

// }

func allPathsSourceTarget(graph [][]int) [][]int {
	var paths [][]int
	n := len(graph)

	var backtrack func(int, []int)
	backtrack = func(node int, curPath []int) {
		if node == n-1 {
			paths = append(paths, append(curPath, node))
			return
		}
		for _, child := range graph[node] {
			backtrack(child, append(curPath, node))
		}
	}
	backtrack(0, []int{})
	return paths
}

func main() {
	nums := [][]int{{1},{}}
	fmt.Println(allPathsSourceTarget(nums))
}