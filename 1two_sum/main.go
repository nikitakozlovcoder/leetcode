package main

import "fmt"

func main() {
	res := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(res)

}

func twoSum(nums []int, target int) []int {
	h := map[int]int{}

	for i, num := range nums {
		h[num] = i
	}

	for i, num := range nums {
		var diff = target - num
		if idx, ok := h[diff]; ok && idx != i {
			return []int{idx, i}
		}
	}

	return []int{}
}
