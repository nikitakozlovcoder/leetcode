package main

import (
	"fmt"
	"slices"
)

func main() {
	res := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	fmt.Println(res)
}

/*func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	slices.Sort(candidates)
	rec(candidates, target, &res, []int{}, math.MinInt32)
	return res

}

func rec(candidates []int, target int, res *[][]int, picked []int, prevCandiate int) {
	if target == 0 {
		*res = append(*res, picked)
		return
	}

	if len(candidates) == 0 {
		return
	}

	if prevCandiate == candidates[0] {
		return
	}

	newPicked := append([]int{}, picked...)
	newPicked = append(newPicked, candidates[0])

	rec(candidates[1:], target-candidates[0], res, newPicked, candidates[0])
	rec(candidates[1:], target, res, picked, candidates[0])
}*/

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	slices.Sort(candidates)
	rec(candidates, target, 0, &res, []int{})
	return res

}

func rec(candidates []int, target int, start int, res *[][]int, picked []int) {
	if target == 0 {
		*res = append(*res, picked)
		return
	}

	if target < 0 {
		return
	}

	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		newPicked := append([]int{}, picked...)
		newPicked = append(newPicked, candidates[i])
		rec(candidates, target-candidates[i], i+1, res, newPicked)
	}
}
