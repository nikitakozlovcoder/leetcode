package main

import (
	"fmt"
	"math"
)

func main() {
	res := reverse(1463847412)
	fmt.Println(res)
}

// Given a signed 32-bit integer x, return x with its digits reversed.
// If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
func reverse(x int) int {
	res := 0
	for x > 0 {
		t := x % 10
		x = x / 10
		if res > (math.MaxInt32-t)/10 {
			return 0
		}

		res = res*10 + t
	}

	for x < 0 {
		t := x % 10
		x = x / 10
		if res < (math.MinInt32-t)/10 {
			return 0
		}

		res = res*10 + t
	}

	return res
}
