package main

import "fmt"

func main() {
	res := lengthOfLongestSubstring("abba")
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	length := 0
	m := make(map[byte]int)
	for right < len(s) {
		c := s[right]
		if _, ok := m[c]; ok {
			left = max(left, m[c]+1)
		}

		win := right - left + 1
		m[c] = right
		length = max(length, win)
		right++
	}

	return length
}
