package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := isPalindrome(11211)
	fmt.Println(res)
}

func isPalindrome(x int) bool {
	t := x
	reversed := 0
	for t > 0 {
		reversed = reversed*10 + t%10
		t = t / 10
	}

	return x == reversed
}

func isPalindrome2(x int) bool {
	s := strconv.Itoa(x)
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
