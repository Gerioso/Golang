package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}

	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

func main() {
	result := removeDuplicates([]int{1, 1, 2})
	fmt.Println(result, "/n")

}
