package main

import "fmt"

func main() {
	res := letterCombinations("23456789")
	fmt.Println(res)
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return make([]string, 0)
	}

	mapping := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	result := make([]string, 0)
	dfs(digits, 0, mapping, "", &result)
	return result
}

func dfs(digits string,
	idx int,
	mapping map[string][]string,
	curStr string,
	res *[]string) {
	if idx == len(digits) {
		*res = append(*res, curStr)
		return
	}

	key := string(digits[idx])
	letters := mapping[key]
	for _, l := range letters {
		dfs(digits, idx+1, mapping, curStr+l, res)
	}
}
