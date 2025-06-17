package main

import "fmt"

func main() {
	res := intToRoman(3749)
	fmt.Println(res)
}

func intToRoman(num int) string {
	conv := []struct {
		num  int
		conv string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""
	for num > 0 {
		for _, c := range conv {
			if num >= c.num {
				num -= c.num
				result += c.conv
				break
			}
		}
	}

	return result
}
