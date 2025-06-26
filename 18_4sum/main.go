package main

import (
	"fmt"
	"slices"
)

func main() {
	res := fourSum([]int{2, 2, 2, 2, 2}, 8)
	fmt.Println(res)
}

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			k := j + 1
			v := len(nums) - 1
			for k < v {
				t := nums[k] + nums[j] + nums[i] + nums[v]
				if t == target {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[v]})
					k++

					for nums[k-1] == nums[k] && k < v {
						k++
					}
				} else if t < target {
					k++
				} else {
					v--
				}
			}
		}
	}

	return res
}

func fourSumSlow(nums []int, target int) [][]int {
	slices.Sort(nums)
	set := make(map[[4]int]struct{})
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			k := j + 1
			v := len(nums) - 1
			for k < v {
				t := nums[k] + nums[j] + nums[i] + nums[v]
				if t == target {
					set[[4]int{nums[i], nums[j], nums[k], nums[v]}] = struct{}{}
					k++
				} else if t < target {
					k++
				} else {
					v--
				}
			}
		}
	}

	res := make([][]int, 0)
	for k := range set {
		res = append(res, k[:])
	}

	return res
}
