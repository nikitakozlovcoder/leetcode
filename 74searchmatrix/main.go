package main

import "fmt"

func main() {
	res := searchMatrix([][]int{{1}}, 2)
	fmt.Println(res)
}

func searchMatrix(matrix [][]int, target int) bool {
	return bSearch(matrix, 0, len(matrix)*len(matrix[0])-1, target)
}

func bSearch(matrix [][]int, left, right, target int) bool {
	if left > right {
		return false
	}

	mid := left + (right-left)/2
	t := get(matrix, mid)
	if t == target {
		return true
	}

	if t > target {
		return bSearch(matrix, left, mid-1, target)
	}

	return bSearch(matrix, mid+1, right, target)
}

func get(matrix [][]int, idx int) int {
	n := len(matrix[0])
	return matrix[idx/n][idx%n]
}
