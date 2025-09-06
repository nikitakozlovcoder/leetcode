package main

func main() {}

func numTrees(n int) int {
	dp := make([]int, n)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		total := 0
		for j := 1; j <= i; j++ {
			total += dp[j-1] * dp[i-j]
		}
		dp[i] = total
	}

	return dp[n]
}
