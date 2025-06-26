package main

import (
	"fmt"
	"slices"
)

func main() {
	res := threeSum([]int{0, 0, 0})
	fmt.Println(res)
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	set := make(map[[3]int]struct{})
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			t := nums[j] + nums[k] + nums[i]
			if t == 0 {
				set[[3]int{nums[j], nums[k], nums[i]}] = struct{}{}
				j++
				k--
			} else if t < 0 {
				j++
			} else {
				k--
			}
		}
	}

	for k := range set {
		res = append(res, k[:])
	}

	return res
}

/*def threeSum(self, nums: List[int]) -> List[List[int]]:
  res = []
  nums.sort()

  for i in range(len(nums)):
      if i > 0 and nums[i] == nums[i-1]:
          continue

      j = i + 1
      k = len(nums) - 1

      while j < k:
          total = nums[i] + nums[j] + nums[k]

          if total > 0:
              k -= 1
          elif total < 0:
              j += 1
          else:
              res.append([nums[i], nums[j], nums[k]])
              j += 1

              while nums[j] == nums[j-1] and j < k:
                  j += 1

  return res*/
