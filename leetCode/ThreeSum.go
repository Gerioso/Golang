package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums) - 1
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < n-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := n
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > 0 {
				k--
			} else {
				j++
			}

		}
	}
	return result
}

func main() {
	result := threeSum([]int{-1, 0, 3, 1, -3, 2, -1, -1})
	fmt.Println(result)
}
