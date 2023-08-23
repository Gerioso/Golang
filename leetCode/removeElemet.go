package main

import (
	"fmt"
)

func removeElement(nums []int, val int) (int, []int) {
	top := len(nums) - 1
	for i := top; i >= 0; i-- {
		if nums[i] == val {
			nums[i], nums[top] = nums[top], nums[i]
			top--
		}
	}
	return top + 1, nums
}

func main() {
	result, nums := removeElement([]int{1, 1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 3}, 1)
	fmt.Println(result, "/n", nums)

}
