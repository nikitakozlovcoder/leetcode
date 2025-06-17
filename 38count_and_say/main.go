package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := countAndSay(4)
	fmt.Println(res)
}

type ParsedSeq struct {
	Char byte
	Len  int
}

func countAndSay(n int) string {
	res := "1"
	for range n - 1 {
		t := ""
		count := 1
		for i := 1; i < len(res); i++ {
			if res[i-1] == res[i] {
				count++
			} else {

				t += strconv.Itoa(count) + string(res[i-1])
				count = 1
			}
		}

		t += strconv.Itoa(count) + string(res[len(res)-1])
		res = t
	}

	return res
}

func dfs(n int) string {
	if n == 1 {
		return "1"
	}

	prev := dfs(n - 1)
	parsed := parse(prev)
	str := ""
	for _, c := range parsed {
		str += fmt.Sprintf("%v%c", c.Len, c.Char)
	}

	return str
}

func parse(s string) []*ParsedSeq {
	parsed := []*ParsedSeq{
		{s[0], 1},
	}

	for i := 1; i < len(s); i++ {
		prev := parsed[len(parsed)-1]
		if prev.Char == s[i] {
			prev.Len++
		} else {
			parsed = append(parsed, &ParsedSeq{s[i], 1})
		}
	}

	return parsed
}
