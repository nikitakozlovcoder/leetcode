package main

import "fmt"

func main() {
	result := generateParenthesis(3)
	fmt.Println(result)
}

func generateParenthesis(n int) []string {
	res := make([]string, 0, 100)
	generate(0, 0, "", n, &res)
	return res
}

func generate(open int, close int, s string, n int, res *[]string) {
	if open == close && close+open == 2*n {
		*res = append(*res, s)
		return
	}

	if open < n {
		generate(open+1, close, s+"(", n, res)
	}

	if close < open {
		generate(open, close+1, s+")", n, res)
	}
}
