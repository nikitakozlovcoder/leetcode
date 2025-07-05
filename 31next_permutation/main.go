package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	nums := []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation2(nums []int) {
	if len(nums) < 2 {
		return
	}

	isDescOrdered := true
	for i, num := range nums[1:] {
		if num > nums[i] {
			isDescOrdered = false
		}
	}

	if isDescOrdered {
		slices.Reverse(nums)
		return
	}

	for i := len(nums) - 2; i >= 0; i-- {
		numIdx := -1
		num := math.MaxInt32
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < num && nums[i] < nums[j] {
				numIdx = j
				num = nums[j]
			}
		}

		if numIdx != -1 {
			nums[i], nums[numIdx] = nums[numIdx], nums[i]
			slices.Sort(nums[i+1:])
			return
		}
	}
}

func nextPermutation(nums []int) {
	i := len(nums) - 1
	for i > 0 && nums[i-1] >= nums[i] {
		i--
	}

	if i == 0 {
		slices.Reverse(nums)
		return
	}

	j := len(nums) - 1
	for j >= i-1 && nums[j] <= nums[i-1] {
		j--
	}

	nums[i-1], nums[j] = nums[j], nums[i-1]
	slices.Reverse(nums[i:])
}
