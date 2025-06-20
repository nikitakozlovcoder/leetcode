package main

import (
	"fmt"
	"slices"
)

func main() {
	res := permuteUnique([]int{3, 3, 0, 3})
	fmt.Println(len(res))
	for _, v := range res {
		fmt.Println(v)
	}
}

func permuteUnique(nums []int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	dfc(nums, 0, []int{}, used, &res)
	return res
}

func dfc(nums []int,
	idx int,
	comb []int,
	used []bool,
	res *[][]int) {
	if idx == len(nums) {
		*res = append(*res, comb)
		return
	}

	for i, n := range nums {
		if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
			continue
		}

		nextComb := append([]int{}, comb...)
		nextComb = append(nextComb, n)
		nextUsed := append([]bool{}, used...)
		nextUsed[i] = true
		dfc(nums, idx+1, nextComb, nextUsed, res)
	}
}
