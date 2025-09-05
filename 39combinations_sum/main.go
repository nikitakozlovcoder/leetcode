package main

import "fmt"

func main() {
	res := combinationSum([]int{2, 3, 6, 7}, 7)
	fmt.Println(res)
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	rec(candidates, target, []int{}, &res, 0)
	return res
}

func rec(candidates []int, target int, combination []int, res *[][]int, n int) {
	if target == 0 {
		*res = append(*res, combination)
		return
	}

	if target < 0 {
		return
	}

	for i := n; i < len(candidates); i++ {
		comb := append([]int{candidates[i]}, combination...)
		rec(candidates, target-candidates[i], comb, res, i)
	}
}
