package main

import "fmt"

func main() {
	res := searchRange([]int{2, 2}, 2)
	fmt.Println(res)
}

func searchRange(nums []int, target int) []int {
	left := binSearch2(nums, target, true)
	right := binSearch2(nums, target, false)
	/*idx := binSearch(nums, target)
	if idx == -1 {
		return []int{-1, -1}
	}

	left, right := idx, idx
	for left > 0 {
		if nums[left-1] == target {
			left--
		} else {
			break
		}
	}

	for right < len(nums)-1 {
		if nums[right+1] == target {
			right++
		} else {
			break
		}
	}*/

	return []int{left, right}
}

func binSearch2(nums []int, target int, left bool) int {
	l := 0
	r := len(nums) - 1
	idx := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			idx = mid
			if left {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return idx
}

func binSearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
