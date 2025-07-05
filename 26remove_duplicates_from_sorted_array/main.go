package main

import "fmt"

func main() {
	arr := []int{0}
	res := removeDuplicates(arr)
	fmt.Println(res)
	fmt.Println(arr)
}

func removeDuplicates(nums []int) int {
	uniq := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[uniq] = nums[i+1]
			uniq++
		}
	}

	return uniq
}
