package main

import "fmt"

func main() {
	nums := []int{2, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
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
		Reverse(nums)
		return
	}

	for i := len(nums) - 1; i > 1; i-- {
		if nums[i] > nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
			return
		}
	}

	PushRight(nums, nums[0])
}

func PushRight(nums []int, val int) {
	for i := 0; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}

	nums[len(nums)-1] = val
}

func Reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
