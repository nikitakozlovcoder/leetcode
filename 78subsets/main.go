package main

import "fmt"

func main() {
	res := subsets([]int{1, 2, 3})
	for _, v := range res {
		fmt.Println(v)
	}
}

func subsets(nums []int) [][]int {
	res := [][]int{}
	rec(nums, &res, []int{}, 0)
	return res
}

func rec(nums []int, res *[][]int, comb []int, n int) {
	if n == len(nums) {
		*res = append(*res, comb)
		return
	}

	rec(nums, res, comb, n+1)
	rec(nums, res, append([]int{nums[n]}, comb...), n+1)
}
