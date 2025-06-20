package main

import (
	"fmt"
	"sort"
)

func main() {
	res := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(res)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Slice(strs, func(a, b int) bool {
		return len(strs[a]) < len(strs[b])
	})

	res := strs[0]
	for _, v := range strs[1:] {
		for i := 0; i < len(res); i++ {
			resPref := res[:i+1]
			vPref := v[:i+1]
			if vPref != resPref {
				res = res[:i]
				break
			}
		}
	}

	return res
}
