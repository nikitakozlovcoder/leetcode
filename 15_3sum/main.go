package main

import (
	"fmt"
	"slices"
)

func main() {
	res := threeSum([]int{0, 0, 0})
	fmt.Println(res)
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	set := make(map[[3]int]struct{})
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			t := nums[j] + nums[k] + nums[i]
			if t == 0 {
				set[[3]int{nums[j], nums[k], nums[i]}] = struct{}{}
				j++
				k--
			} else if t < 0 {
				j++
			} else {
				k--
			}
		}
	}

	for k := range set {
		res = append(res, k[:])
	}

	return res
}
