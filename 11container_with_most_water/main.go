package main

import "fmt"

func main() {
	res := maxArea([]int{1, 1})
	fmt.Println(res)
}

func maxArea(heights []int) int {
	result := 0
	for i := 0; i < len(heights); i++ {
		for j := i + 1; j < len(heights); j++ {
			left := heights[i]
			right := heights[j]
			height := min(left, right)
			width := j - i
			area := width * height
			if area > result {
				result = area
			}
		}
	}

	return result
}
