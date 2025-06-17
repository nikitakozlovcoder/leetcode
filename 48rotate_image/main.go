package main

import "fmt"

func main() {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(m)
	fmt.Println(m)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
