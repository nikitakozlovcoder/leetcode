package main

import "fmt"

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	result := minPathSum(grid)
	fmt.Println(result)
}

func minPathSum(grid [][]int) int {
	memo := [200][200]int{}
	n := len(grid)
	m := len(grid[0])
	memo[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		memo[0][i] = memo[0][i-1] + grid[0][i]
	}

	for i := 1; i < n; i++ {
		memo[i][0] = memo[i-1][0] + grid[i][0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			memo[i][j] = min(memo[i-1][j], memo[i][j-1]) + grid[i][j]
		}
	}

	return memo[n-1][m-1]
}
