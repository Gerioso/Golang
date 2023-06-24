package main

import (
	"fmt"
)

func removeDuplicates(nums []int) (int, []int) {
	if len(nums) < 2 {
		return 0, nil
	}
	count := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[i-2] {
			nums[count] = nums[i]
			count++
		}
	}
	return count, nums
}

func main() {
	result, nums := removeDuplicates([]int{1, 1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 3})
	fmt.Println(result, "/n", nums)

}
