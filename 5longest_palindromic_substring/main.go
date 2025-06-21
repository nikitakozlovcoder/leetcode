package main

import "fmt"

func main() {
	res := longestPalindrome("a")
	fmt.Println(res)
}

// Given a string s, return the longest palindromic substring in s.
func longestPalindrome(s string) string {
	// dp[i][j] = i начало j конец
	dp := make([][]bool, len(s))
	for start := range dp {
		dp[start] = make([]bool, len(s))
	}

	best := ""
	for start := len(s) - 1; start >= 0; start-- {
		for end := 0; end <= len(s)-1; end++ {
			// если строка пустая или состоит из 1 символа, то это палиндром
			if start >= end {
				dp[start][end] = true
			}

			if start < end && s[start] == s[end] {
				// если в центре палиндром и по бокам символы равны, то строка палиндром
				middle := dp[start+1][end-1]
				dp[start][end] = middle
			}

			if dp[start][end] {
				if end-start+1 > len(best) {
					best = s[start : end+1]
				}
			}
		}
	}

	return best
}
