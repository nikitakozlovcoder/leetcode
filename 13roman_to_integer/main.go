package main

import "fmt"

func main() {
	res := romanToInt("MCMXCIV")
	fmt.Println(res)
}

func romanToInt(s string) int {
	conv := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	res := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && conv[string(s[i])] < conv[string(s[i+1])] {
			res -= conv[string(s[i])]
		} else {
			res += conv[string(s[i])]
		}
	}

	return res
}
