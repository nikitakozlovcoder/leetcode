package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	res := simplifyPath("/home/user/Documents/../Pictures")
	fmt.Println(res)

}

func simplifyPath(path string) string {
	sections := strings.Split(path, "/")
	res := []string{}
	skip := 0
	for _, v := range slices.Backward(sections) {
		if v == "" || v == "." {
			continue
		}

		if v == ".." {
			skip++
			continue
		}

		if skip == 0 {
			res = append(res, v)
		} else {
			skip--
		}
	}

	slices.Reverse(res)
	return "/" + strings.Join(res, "/")
}
